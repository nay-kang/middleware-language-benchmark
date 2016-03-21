package main

import (
    "fmt"
    "net/http"
    //"log"
    "runtime"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}

func main() {
    runtime.GOMAXPROCS(2)
    http.HandleFunc("/", handleRequest) //设置访问的路由
    http.ListenAndServe(":18083", nil) //设置监听的端口
    //if err != nil {
    //    log.Fatal("ListenAndServe: ", err)
    //}
}
