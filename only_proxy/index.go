package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "runtime"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	resp,err := http.Get("http://localhost:18082/product_list.json")
    if err!=nil{
        fmt.Println(err.Error())
        return
    }
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    bodyStr := string(body)
    fmt.Fprintf(w, bodyStr)
}

func main() {
    runtime.GOMAXPROCS(2)
    http.HandleFunc("/", handleRequest) //设置访问的路由
    err := http.ListenAndServe(":18083", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
