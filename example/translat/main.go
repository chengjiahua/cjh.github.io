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
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please input the translation word !")
		fmt.Println("example: dict hello")
		os.Exit(0)
	}
	word := strings.Join(args, " ")
	requestUrl := fmt.Sprintf("https://dict.youdao.com/fsearch?client=deskdict&keyfrom=chrome.extension&q=%s&pos=-1&doctype=xml&xmlVersion=3.2&dogVersion=1.0&vendor=unknown&appVer=3.1.17.4208&le=eng", url.QueryEscape(word))
	ch := make(chan []byte)
	go HttpGet(requestUrl, ch)

	xmlObject := ParseXmlData{}
	xml.Unmarshal(<-ch, &xmlObject)
	//fmt.Println("您的输入:", strings.TrimSpace(xmlObject.RawInput))
	fmt.Println("\033[1;31m*******英汉翻译:*******\033[0m")
	for _, v := range xmlObject.CustomTrans.Translation {
		fmt.Println(strings.TrimSpace(v.Content))
	}
	fmt.Println("\033[1;31m*******网络释义:*******\033[0m")
	for _, v := range xmlObject.WebTrans.TransNode {
		key := strings.TrimSpace(v.Key)
		for _, vv := range v.Trans {
			value := strings.TrimSpace(vv.Value)
			fmt.Println(key, ":", value)
			break
		}
	}
}
