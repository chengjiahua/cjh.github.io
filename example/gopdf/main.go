package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

//Breakdown .
func Breakdown(pdf *gofpdf.Fpdf) {
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.CellFormat(
			0, 10,
			fmt.Sprintf("第 %d 页,共 {nb} 页", pdf.PageNo()),
			"", 0, "C", false, 0, "",
		)
	})
	pdf.AliasNbPages("")

	// pdf.SetFont("Arial", "", 14)
	pdf.AddUTF8Font("notosanssc", "", "./simfang.ttf")
	pdf.SetFont("notosanssc", "", 11)

	pdf.AddPage()

	type countryType struct {
		name, ip, types, count, percent, start, last, timePercent string
	}
	countryList := make([]countryType, 0, 8)

	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err != nil {
			panic(err)
		}

		defer fl.Close()

		scanner := bufio.NewScanner(fl)
		var c countryType
		for scanner.Scan() {
			lineStr := scanner.Text()
			list := strings.Split(lineStr, ";")

			if len(list) == 8 {
				c.name = list[0]
				c.ip = list[1]
				c.types = list[2]
				c.count = list[3]
				c.percent = list[4]
				c.start = list[5]
				c.last = list[6]
				c.timePercent = list[7]
				countryList = append(countryList, c)
			} else {
				fmt.Printf("error tokenizing %s\n", lineStr)
			}
		}

		if len(countryList) == 0 {
			fmt.Printf("error loading data from %s\n", fileStr)
		}
	}
	loadData("/root/test/break.txt")

	left := (210.0 - 8*25 - 5) / 2
	// left := 0.0
	pdf.SetX(left)
	pdf.CellFormat(205, 13, "故障报表        时间周期：DAY   时间范围：2021-01-12 09:01:16 至 2021-01-13 09:01:16",
		"1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(left)
	// "设施名称", "IP(域名)", "设施类型", "报警次数", "百分比", "异常开始时间", "累计持续时间", "百分比"
	pdf.CellFormat(25, 8, "设施名称", "1", 0, "", false, 0, "")
	pdf.CellFormat(25, 8, "IP(域名)", "1", 0, "", false, 0, "")
	pdf.CellFormat(20, 8, "设施类型", "1", 0, "", false, 0, "")
	pdf.CellFormat(20, 8, "报警次数", "1", 0, "", false, 0, "")
	pdf.CellFormat(25, 8, "次数比例", "1", 0, "", false, 0, "")
	pdf.CellFormat(45, 8, "异常开始", "1", 0, "", false, 0, "")
	pdf.CellFormat(20, 8, "持续时间", "1", 0, "", false, 0, "")
	pdf.CellFormat(25, 8, "时间比例", "1", 0, "", false, 0, "")

	pdf.Ln(-1)

	for _, c := range countryList {
		pdf.SetX(left)
		pdf.CellFormat(25, 8, c.name, "1", 0, "", false, 0, "")
		pdf.CellFormat(25, 8, c.ip, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 8, c.types, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 8, c.count, "1", 0, "", false, 0, "")
		pdf.CellFormat(25, 8, c.percent, "1", 0, "", false, 0, "")
		pdf.CellFormat(45, 8, c.start, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 8, c.last, "1", 0, "", false, 0, "")
		pdf.CellFormat(25, 8, c.timePercent, "1", 0, "", false, 0, "")
		pdf.Ln(-1)

	}

}

// Statis .
func Statis(pdf *gofpdf.Fpdf) {
	// pdf.SetFont("Arial", "", 14)
	// pdf.AddUTF8Font("notosanssc", "", "./simfang.ttf")

	buf, err := ioutil.ReadFile("/root/test/simfang.ttf")
	if err != nil {
		fmt.Println(err)
	}

	pdf.AddUTF8FontFromBytes("notosanssc", "", buf)
	pdf.SetFont("notosanssc", "", 8)

	pdf.AddPage()

	type countryType struct {
		a, b, c, d, e, f, g, h, i, j, k, l, m string
	}
	countryList := make([]countryType, 0, 8)

	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err != nil {
			panic(err)
		}

		defer fl.Close()

		scanner := bufio.NewScanner(fl)
		var c countryType
		for scanner.Scan() {
			//Austria;Vienna;83859;8075
			lineStr := scanner.Text()
			list := strings.Split(lineStr, ";")

			if len(list) == 13 {
				c.a = list[0]
				c.b = list[1]
				c.c = list[2]
				c.d = list[3]
				c.e = list[4]
				c.f = list[5]
				c.g = list[6]
				c.h = list[7]
				c.i = list[8]
				c.j = list[9]
				c.k = list[10]
				c.l = list[11]
				c.m = list[12]
				countryList = append(countryList, c)
			} else {
				fmt.Printf("error tokenizing %s\n", lineStr)
			}
		}

		if len(countryList) == 0 {
			fmt.Printf("error loading data from %s\n", fileStr)
		}
	}
	loadData("/root/test/static.txt")
	left := (210.0 - 13*15) / 2
	// left := 0.0
	pdf.SetX(left)
	pdf.CellFormat(195, 10, " 统计报表        时间周期：DAY   时间范围：2021-01-12 09:01:16 至 2021-01-13 09:01:16",
		"1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(left)
	// "设施名称", "IP(域名)", "设施类型", "报警次数", "百分比", "异常开始时间", "累计持续时间", "百分比"
	pdf.CellFormat(18, 8, "设施名称", "1", 0, "", false, 0, "")
	pdf.CellFormat(18, 8, "IP(域名)", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "设施类型", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "报警次数", "1", 0, "", false, 0, "")
	pdf.CellFormat(12, 8, "CPU大小", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "CPU平均值", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "CPU最大值", "1", 0, "", false, 0, "")
	pdf.CellFormat(12, 8, "内存大小", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "内存平均值", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "内存最大值", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "磁盘空间", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "磁盘使用率", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 8, "综合利用率", "1", 0, "", false, 0, "")

	pdf.Ln(-1)

	for _, c := range countryList {
		pdf.SetX(left)
		pdf.CellFormat(18, 8, c.a, "1", 0, "", false, 0, "")
		pdf.CellFormat(18, 8, c.b, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.c, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.d, "1", 0, "", false, 0, "")
		pdf.CellFormat(12, 8, c.e, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.f, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.g, "1", 0, "", false, 0, "")
		pdf.CellFormat(12, 8, c.h, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.i, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.j, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.k, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.l, "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 8, c.m, "1", 0, "", false, 0, "")
		pdf.Ln(-1)

	}
}

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "") // 210 x 297
	Breakdown(pdf)
	// Statis(pdf)
	if err := pdf.OutputFileAndClose("Breakdown.pdf"); err != nil {
		panic(err.Error())
	}

	fmt.Println(fmt.Sprintf("%d%%", 5))
}
