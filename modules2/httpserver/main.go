package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200
func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/header",writeHeader)
	mux.HandleFunc("/env",getEnv)
	mux.HandleFunc("/logIp",logIp)
	mux.HandleFunc("/healthz",healthz)
	if err := http.ListenAndServe(":8080",mux);err!=nil{
		log.Fatalln("server start error")
	}
}

//接收客户端 request，并将 request 中带的 header 写入 response header
func writeHeader(w http.ResponseWriter,r *http.Request)  {
	for name,sliceValue:=range r.Header{
		var headerValue = ""
		for _,value:=range sliceValue{
			headerValue=headerValue+value+","
		}
		//[a:b]左闭右开
		w.Header().Set(name,headerValue[0:len(headerValue)-1])
	}
}

//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func getEnv(w http.ResponseWriter,r *http.Request)  {
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
}

func logIp(w http.ResponseWriter,r *http.Request)  {
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	log.Printf("{%s}:Request Success! Response code : %d", ip,http.StatusOK)
}

func healthz(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w, "health")
}
