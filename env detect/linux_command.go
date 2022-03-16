
package main

import (
	"github.com/cdk-team/CDK/conf"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ans := []string{}
	for _, cmd := range conf.LinuxCommandChecklist {
		_, err := exec.LookPath(cmd)
		if err == nil {
			ans = append(ans, cmd)
		}
	}
	log.Printf("available commands:\n\t%s\n", strings.Join(ans, ","))
}
