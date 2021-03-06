package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------")
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

//处理upload逻辑
func upload1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if (r.Method == "GET") {
		crutime := time.Now().Unix()
		h := md5.new()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t,_ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploasfile")
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := Os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) //此处假设当前目录已存在test目录
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

//处理uploas逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%f", h.Sum(nil))

		t,_ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) //此处假设当前目录瞎已存在test目录
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer f.Close()
		io.Copy(f, file)
	}
} 


func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/upload", upload)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
