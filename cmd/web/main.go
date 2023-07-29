package main

import (
	"fmt"
	"os"

	"github.com/iamanders/tdr3/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
