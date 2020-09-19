package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Response(w http.ResponseWriter, resp string)  {
	n, err := io.WriteString(w, resp)
	if err != nil {
		fmt.Println("WriteString error....err:",err)
		fmt.Println("WriteString error....n:",n)
		return
	}
	return
}

func AuthUser(r *http.Request) (string,  string,  error) {
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	//query := r.URL.Query()
	var username string // 初始化定义变量
	var pass string // 初始化定义变量
	if r.Method == "GET" {
		username = r.FormValue("username")
		pass = r.FormValue("pass")
	} else if r.Method == "POST" {
		username = r.PostFormValue("username")
		pass = r.PostFormValue("pass")
	}

	if username != "root" || pass != "123456" {
		return username, pass, errors.New(fmt.Sprint("账号不对....:",username, pass))
	}

	return username, pass, nil
}