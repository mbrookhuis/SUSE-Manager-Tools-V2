package main

import (
	"fmt"
	"log/slog"
	"runtime"
)

func main() {
	slog.Info("hello, world")
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}

}
