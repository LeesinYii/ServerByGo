package main

import(
	"fmt"
	"net/http"
	"strings"
	"log"
)

func main() {
	//设置访问路由
	http.HandleFunc("/",sayHelloName) 

	// 设置监听端口
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}

/**
 * 请求处理函数
 */
func sayHelloName(w http.ResponseWriter,r *http.Request) {
	// 解析参数，默认不解析
	r.ParseForm() 

	// 信息输出到服务器输出
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}

	// 写入到 w 是 输出到客户端
	fmt.Fprintf(w,"Hello  astaxiel")
}
