package main

import (
	"fmt"
	"os"
	"os/exec"
	_ "plugin"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: hangme <iterations>  # 0 means exit immediately\n")
		os.Exit(99)
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse argument as number: %v\n", err)
		os.Exit(1)
	}
	if count == 0 {
		os.Exit(0)
	}
	fmt.Printf("%s\n", runtime.Version())
	for i := 0; i < count; i++ {
		cmd := exec.Command("./hangme", "0")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	}
	return
}
