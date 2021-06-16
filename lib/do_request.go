package lib

import (
	"fmt"
	"go_shell/conf"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err != nil {
		Response(w, fmt.Sprint("ParseForm error....err:", err))
		return
	}

	username, pass, err := AuthUser(r)
	if err != nil {
		Response(w, "账号 bad ...")
		return
	}

	res := "<html><h2>测试脚本：</h2>"
	for _, act := range conf.Cfg.Shells {
		res += fmt.Sprint("<a href='/doExecte?username=", username, "&pass=", pass, "&sh=", act, "'>", conf.Cfg.ShellMap[act].Name, "<a/><br/>")
	}
	res += "<br /><br />"

	res += "<h2>另外脚本：</h2>"
	for _, act := range conf.Cfg.SyncShells {
		res += fmt.Sprint("<a href='/doExecte?username=", username, "&pass=", pass, "&sh=", act, "'>", conf.Cfg.ShellMap[act].Name, "<a/><br/>")
	}

	res += "</html>"
	Response(w, res)
	return
}

func DoExecute(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err != nil {
		fmt.Println("ParseForm error....err:", err)
		return
	}

	username, pass, err := AuthUser(r)
	if err != nil {
		Response(w, err.Error())
		return
	}

	res := fmt.Sprint("<a href='/index?username=", username, "&pass=", pass, "'>返回<a/><br/>")

	key := r.FormValue("sh")
	out, err := QuickExecute(key)
	if err != nil {
		res += err.Error()
		Response(w, res)
		return
	}

	res += out
	Response(w, res)
	return
}
