package main

import (
	"flag"
	"go_shell/conf"
	"go_shell/lib"
	"log"
	"net/http"
)

func main() {
	//初始化配置
	err := Init()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", lib.Index)             //   设置访问路由
	http.HandleFunc("/index", lib.Index)        //   设置访问路由
	http.HandleFunc("/doExecte", lib.DoExecute) //   设置访问路由
	err = http.ListenAndServe(":18888", nil)    //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/**
 * 初始化
 */
func Init() error {
	configpath := flag.String("f", "./conf/cfg.toml", "config file")
	flag.Parse()

	err := conf.InitConfig(*configpath)
	if err != nil {
		panic(err)
	}
	return nil
}
