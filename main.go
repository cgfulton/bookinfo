package main

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/cgfulton/bookinfo-prompt/internal/debug"
	"github.com/cgfulton/bookinfo-prompt/istio"
)

func main() {
	if c, err := istio.NewCompleter(); err != nil {
		debug.LogError("Completer error", err)
		os.Exit(0)
	} else {
		cmd := getNamespace(c)
		fmt.Println("-------")
		fmt.Println("Executing: " + cmd)
		fmt.Println("-------")
		cmd = "kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml"
		fmt.Println("Executing: " + cmd)
		fmt.Println("-------")
	}
}

func getNamespace(c *istio.Completer) string {
	fmt.Println("Select the namespace to host your application")
	if ns := prompt.Input("> ", c.Complete); ns != "" {
		fmt.Println("You selected namespace: " + ns)
		return fmt.Sprintf("kubectl label namespace %s istio-injection=enabled", ns)
	}
	return getNamespace(c)
}
