package lib

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err != nil {
		Response(w, fmt.Sprint("ParseForm error....err:",err))
		return
	}

	username, pass, err := AuthUser(r)
	if err != nil {
		Response(w, "账号 bad ...")
		return
	}

	res := ""
	for _, act := range Shells {
		res +=  fmt.Sprint("<a href='/doExecte?username=",username,"&pass=",pass,"&sh=",act,"'>", act, "<a/><br/>")
	}
	Response(w, res)
	return
}


func DoExecute(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err != nil {
		fmt.Println("ParseForm error....err:",err)
		return
	}

	username, pass, err := AuthUser(r)
	if err != nil {
		Response(w, err.Error())
		return
	}

	key := r.FormValue("sh")
	err = Execute(key)
	if err != nil {
		Response(w, err.Error())
		return
	}

	out, err := QuickExecute(key)
	if err != nil {
		Response(w, err.Error())
		return
	}

	res := fmt.Sprint("<a href='/index?username=",username,"&pass=",pass,"'>返回<a/><br/>")
	res += out
	Response(w, res)
	return
}