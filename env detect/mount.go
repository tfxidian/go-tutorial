package main

import (
	"bufio"
	"fmt"
	"github.com/cdk-team/CDK/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
)

type Mount struct {
	Device     string
	Path       string
	Filesystem string
	Flags      string
}

// The checkClose function calls close on a Closer and panics with a
// runtime error if the Closer returns an error
func checkClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(&errors.CDKRuntimeError{Err: err})
	}
}

func GetMounts() ([]Mount, error) {
	readPath := "/proc/self/mounts"
	file, err := os.Open(readPath)
	if err != nil {
		log.Printf("[Err] Open %s failed.", readPath)
		return nil, err
	}
	fmt.Println("test")
	defer checkClose(file)
	mounts := []Mount(nil)
	reader := bufio.NewReaderSize(file, 64*1024)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return mounts, nil
			}
			return nil, err
		}
		if isPrefix {
			return nil, syscall.EIO
		}
		parts := strings.SplitN(string(line), " ", 5)
		if len(parts) != 5 {
			return nil, syscall.EIO
		}
		mounts = append(mounts, Mount{parts[0], parts[1], parts[2], parts[3]})
	}
}

func main() {
	mounts, _ := GetMounts()
	for _, m := range mounts {
		fmt.Println(m)
	}
}
