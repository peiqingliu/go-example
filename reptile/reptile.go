package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/**
	爬虫步骤
	1、明确目标(确定在哪个网站搜索)
	2、爬（爬取内容）
	3、取（筛选内容）
	3、处理数据（按照你的想法去处理）
 */
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)
func main()  {
	GetEmail()
}

func GetEmail()  {

	// 1、去网站拿数据
	resp,err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err,"http.Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr := string(pageBytes)
	fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	results  := re.FindAllStringSubmatch(pageStr,-1)
	fmt.Println(results )

	// 遍历结果
	for _,result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}