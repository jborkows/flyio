package main

import (
	"fmt"
	"os"
	"time"
)

func PrintInfo(msg string) {
	fmt.Fprintf(os.Stderr, "%v %s\n", time.Now(), msg)
}
