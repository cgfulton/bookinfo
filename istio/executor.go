package istio

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cgfulton/bookinfo-prompt/internal/debug"
)

func Executor(s string) {
	ExecuteAndGetResult(s)
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	return out.String()
}
