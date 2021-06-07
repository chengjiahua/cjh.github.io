package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)


// 定义一个下载结构体，包括需要下载的url，本地保存路径，向服务器发起的连接数，感觉可以换一个变量名
// 默认向服务器发起10个连接
type Download struct {
	Url           string
	TargetPath    string
	TotalSections int
}

var (
	URL           = flag.String("url", "https://github.com/chengjiahua/cjh.github.io/archive/refs/heads/main.zip", "Provide the URL to download")
	Targetfile    = flag.String("target_file", "cjh.github.io-main.zip", "Provide the target file path with extension, example: cjh.github.io-main.zip")
	MaxConcurrent = flag.Int("max_concurrent", 10, "Number of sections/connections to make to the server")
)

func main() {
	flag.Parse()
	startTime := time.Now()
	d := Download{
		// Provide the URL to download,
		//	example: https://www.dropbox.com/s/lgvhj/sample.mp4?dl=1
		Url: *URL,
		// Provide the target file path with extension, example: sample.mp4
		TargetPath: *Targetfile,
		// Number of sections/connections to make to the server
		TotalSections: *MaxConcurrent,
	}
	err := d.Do()
	if err != nil {
		log.Printf("An error occured while downloading the file: %s\n", err)
	}
	// 优化时间显示
	fmt.Printf("Download completed in %v seconds\n", time.Since(startTime).Seconds())

}

// 表示这是一个结构体的方法
// Start the download
func (d Download) Do() error {
	fmt.Println("Checking URL")
	r, err := d.getNewRequest("HEAD")
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	fmt.Printf("Got %v\n", resp.StatusCode)

	if resp.StatusCode > 299 {
		return fmt.Errorf("can't process, response is %v", resp.StatusCode)
	}
	// 等价于 parseInt()，Content-length的单位是字节个数，http HEAD 方法获取了要下载资源的大小
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	fmt.Printf("Size is %v bytes\n", size)

	var sections = make([][2]int, d.TotalSections)
	// 每个协程下载固定大小的文件块
	eachSize := size / d.TotalSections
	fmt.Printf("Each size is %v bytes\n", eachSize)

	// 统计每个协程需要下载的字节起始编号，比如第一个协程下载文件的0~10字节，第二个协程下载11~20字节
	// example: if file size is 100 bytes, our section should like:
	// [[0 10] [11 21] [22 32] [33 43] [44 54] [55 65] [66 76] [77 87] [88 98] [99 99]]
	for i := range sections {
		if i == 0 {
			// starting byte of first section
			sections[i][0] = 0
		} else {
			// starting byte of other sections
			sections[i][0] = sections[i-1][1] + 1
		}

		if i < d.TotalSections-1 {
			// ending byte of other sections
			sections[i][1] = sections[i][0] + eachSize
		} else {
			// ending byte of other sections
			sections[i][1] = size - 1
		}
	}

	// 源码注释里说waitGroup用于等待一组GoRoutine执行完成
	// wg.Wait() 会阻塞住，直到所有的goroutine执行完成
	// 感觉有点像 Java 里的 CountDownLatch
	log.Println(sections)
	var wg sync.WaitGroup
	// download each section concurrently
	for i, s := range sections {
		// 1 是传入的delta
		wg.Add(1)
		go func(i int, s [2]int) {
			// wg.Done()的作用是计数器减1，感觉和java里countDownLatch.countDown()是一样的效果
			defer wg.Done()
			err = d.downloadSection(i, s)
			if err != nil {
				panic(err)
			}
		}(i, s)
	}
	wg.Wait()
	// 等待所有协程都下载完成之后，把分散的文件块merge
	return d.mergeFiles(sections)
}

// 向服务器请求对应资源的[start:end]字节块的数据
// Download a single section and save content to a tmp file
func (d Download) downloadSection(i int, c [2]int) error {
	r, err := d.getNewRequest("GET")
	if err != nil {
		return err
	}
	// Sprintf()主要是拼接一个字符串
	// Get请求，设置Range头部信息
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", c[0], c[1]))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("can't process, response is %v", resp.StatusCode)

	}
	fmt.Printf("Downloaded %v bytes for section %v\n", resp.Header.Get("Content-Length"), i)
	// 返回读取的字节数 b
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 从内存导出到硬盘上
	err = ioutil.WriteFile(fmt.Sprintf("section-%v.tmp", i), b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 表示这是结构体的方法
// Get a new http request
func (d Download) getNewRequest(method string) (*http.Request, error) {
	// 发起一个http请求，由用户传入http请求方法
	r, err := http.NewRequest(
		method,
		d.Url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "Silly Download Manager v001")
	return r, nil
}

// 第一个参数表示是一个结构体方法
// Merge tmp files to single file and delete tmp files
func (d Download) mergeFiles(sections [][2]int) error {
	f, err := os.OpenFile(d.TargetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := range sections {
		tmpFileName := fmt.Sprintf("section-%v.tmp", i)
		b, err := ioutil.ReadFile(tmpFileName)
		if err != nil {
			return err
		}
		// 拼接每一个片段
		n, err := f.Write(b)
		if err != nil {
			return err
		}
		// 拼接完之后，删除原来的文件
		err = os.Remove(tmpFileName)
		if err != nil {
			return err
		}
		fmt.Printf("%v bytes merged\n", n)
	}
	return nil
}
