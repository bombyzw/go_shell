package lib

import (
	"errors"
	"fmt"
	"go_shell/conf"
	"io"
	"log"
	"os/exec"
	"strings"
)

func asyncLog(reader io.ReadCloser) error {
	cache := "" //缓存不足一行的日志信息
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志

			fmt.Printf("%s%s\n", cache, line)
			cache = s[len(s)-1]
		}
	}
	return nil
}

func Execute(key string) error {
	shCmd, ok := conf.Cfg.ShellMap[key]
	if !ok {
		return errors.New("no key .... key:" + key)
	}
	cmd := exec.Command("sh", "-c", shCmd.Cmd)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
		return err
	}

	return nil
}

func QuickExecute(key string) (string, error) {
	shCmd, ok := conf.Cfg.ShellMap[key]
	if !ok {
		return "", errors.New("no key .... key:" + key)
	}

	cmd := exec.Command("sh", "-c", shCmd.Cmd)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	outStr := string(out)
	outArr := strings.Split(outStr, "\n")

	return strings.Join(outArr, "<br/>"), nil
}
