package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func main() {
	slog.Info("hello, world")
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}
	path, err := os.Getwd()
	slog.Info(path, err)

}
