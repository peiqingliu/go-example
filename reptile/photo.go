package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//并发爬取图片

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)
func main()  {

	//1、初始化管道 带有缓存
	chanImageUrls = make(chan string,10000)
	chanTask = make(chan string,10)
	// 2.爬虫协程 20个
	for i := 0;i<20;i++ {
		waitGroup.Add(1)  //计数
		// url := "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html"
		url := "https://www.58pic.com/piccate/2-0-0-ty1-se4-p"+ strconv.Itoa(i) +".html"
		go getImgUrls(url)
	}
	//3、任务统计协程，统计20个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i :=0;i<5;i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string)  {
	//获取所有的图片链接切片
	urls :=getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _,url := range urls{
		chanImageUrls <- url
	}

	chanTask <- url  //用于监控协程已经完成了多少任务
	waitGroup.Done()  //用于表示当前协程完成任务
}

//获取当前页图片链接
func getImgs(url string) (urls []string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg) //传入正则表达式，得到正则表达式对象
	results  := re.FindAllStringSubmatch(pageStr,-1)  //用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
	fmt.Printf("共找到%d条结果\n", len(results))
	for _,result := range results{
		fmt.Println("准备存取结果result:",result)
		url := result[0]
		urls = append(urls,url)
	}
	return
}

// 抽取根据url获取内容
func GetPageStr(url string)(pageStr string)  {
	resp,err :=http.Get(url)
	HandleError1(err,"http.Get url")
	defer resp.Body.Close()
	//2、读取页面内容,转成字节
	pageBytes,err := ioutil.ReadAll(resp.Body)
	HandleError1(err, "ioutil.ReadAll")
	//3、字节转成字符串
	pageStr = string(pageBytes)
	return pageStr
}

//下载图片
func DownloadImg()  {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

//下载保存
func DownloadFile(url string,filename string) (ok bool)  {
	resp,err := http.Get(url)
	HandleError1(err,"http.get.url")
	HandleError1(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError1(err, "resp.body")
	filename = "E:/img/" + filename
	// 写出数据
	err = ioutil.WriteFile(filename,bytes,0666)
	if err != nil {
		return false
	}else {
		return true
	}
}

// 任务统计协程
func CheckOK()  {
	var count int
	for {
		url := <- chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count ++
		if count == 20 {
			//关闭通道
			close(chanImageUrls)
			break
		}
	}
}
// 处理异常
func HandleError1(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}