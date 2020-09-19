package main

import (
	"go_shell/lib"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", lib.Index)      //   设置访问路由
	http.HandleFunc("/index", lib.Index)      //   设置访问路由
	http.HandleFunc("/doExecte", lib.DoExecute)      //   设置访问路由
	err := http.ListenAndServe(":18888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

