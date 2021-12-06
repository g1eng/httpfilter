package main

import (
	"os"
)

func main() {
	if len(os.Args) == 0 {
		Serve()
	} else if len(os.Args) != 0 && os.Args[0] == "--router" {
		ServeRouter()
	}
}
