package main

import (
    "net/http"
    "net/http/pprof"
)

func main() {
    // 主函数中添加
    go func() {
        http.HandleFunc("/debug/pprof/block", pprof.Index)
        http.HandleFunc("/debug/pprof/goroutine", pprof.Index)
        http.HandleFunc("/debug/pprof/heap", pprof.Index)
        http.HandleFunc("/debug/pprof/threadcreate", pprof.Index)

        http.ListenAndServe("0.0.0.0:8888", nil)
    }()

    var finishWaiter chan int
    <-finishWaiter
}