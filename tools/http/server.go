package http

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", helloHandle)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `无障碍功能链接
跳到主要内容无障碍功能帮助
无障碍功能反馈
Google
go的benchmark是多线程吗

找到约 60,200 条结果 （用时 0.44 秒） 
搜索结果
网络搜索结果
Golang单线程&多线程Benchmark - 程序园
www.voidcn.com/code/p-xoqqrccx-q.html
2017年12月31日 - Golang单线程&多线程Benchmark ... 注意：golang有竞速探针build或者run时候请加上-race 参数 ... 什么是多任务，单线程，多线程，超线程; 9.
Go 语言测试（Test）、性能测试（Benchmark） 学习笔记- 郭立东的个人 ...
https://blog.csdn.net/cchd0001/article/details/48181239
2015年9月2日 - are considered benchmarks, and are executed by the “go test” ..... 的不是很规范，因为我并不需要知道每个线程何时结束，所以没有用到channal，主要 .... 当前规模是8K+C代码和2K+lua代码，实现了一个多线程高并发的在线游戏 ...
关于Benchmark 的几个思考- arthurkiller - CSDN博客
https://blog.csdn.net/arthur_killer/article/details/50506338
2016年1月29日 - 其实Go的工具go-tools 里面是有go test 的，可以测benchmark ，并且是 ... 现代PC几乎都是多道系统，cpu 大多是多核多线程，程序其实是在不停地 ...
‎１\ 如何设计并发 · ‎２＼qps峰值与压力的关系 · ‎３＼开始设想的测试思路以及 ...
Go语言性能测试- Go语言中文网- Golang中文社区
https://studygolang.com/articles/11557
Benchmark Go做Benchmar只要在目录下创建一个_test.go后缀的文件，然后添加 ... -4表示4个CPU线程执行；300000表示总共执行了30万次；4531ns/op，表示每次 ...
为何在这个bench下golang开启多核反而慢很多？ - 知乎
https://www.zhihu.com/question/48651236
2016年7月20日 - 这是自己从dhrystone改的一个go的多核bench，bench本身做了什么运算倒不重要，从代码可以看到是开了若干go程，每个go程执行同样的bench，而且在这个代码中 ... 首先做这种运算实体并没有并行但想看看并行性能的benchmark, 请使用testing ...
为什么Go 语言的性能还不如Java？	2017年5月6日
多线程情况下很多变量频繁访问难道每个都要加锁访问吗？	2017年1月8日
为什么Java的并发备受推崇？	2016年11月14日
java的多线程在golang中是如何体现的？	2015年12月9日
www.zhihu.com站内的其它相关信息`)

}
