package main

import (
	"fmt"
	"github.com/MR5356/go-console-loading/loading"
	"os"
	"time"
)

func main() {
	ld := loading.New(os.Args[1])

	ld.Start("Processing...")
	time.Sleep(time.Second * 10)
	fmt.Println()
	ld.Stop()
}
