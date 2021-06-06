package main

/**
 * Created by vscode
 * User: chengjiahua
 * Date: 2021/6/6
 * Time: 16:16
 */

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ParseXmlData struct {
	XMLName     xml.Name     `xml:"yodaodict"`
	RawInput    string       `xml:"return-phrase"`
	CustomTrans CustomNode   `xml:"custom-translation"`
	WebTrans    WebTransList `xml:"yodao-web-dict"`
}

type CustomNode struct {
	Type        string        `xml:"type"`
	Translation []Translation `xml:"translation"`
}

type WebTransList struct {
	TransNode []WebTransNode `xml:"web-translation"`
}

type WebTransNode struct {
	Key   string      `xml:"key"`
	Trans []TransNode `xml:"trans"`
}

type TransNode struct {
	Value string `xml:"value,CDATA"`
}

type Translation struct {
	Content string `xml:"content,CDATA"`
}

func HttpGet(url string, ch chan []byte) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		fmt.Println("please enter ctrl+c,http request err:", errReq.Error())
		return
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Header.Add("Connection", "close")
	resp, errDo := client.Do(request)
	if errDo == nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if errDo != nil {
		fmt.Println("please enter ctrl+c ,request error:", errDo.Error())
		return
	}
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println("please enter ctrl+c ,io read error:", errRead.Error())
		return
	}
	ch <- body
}
