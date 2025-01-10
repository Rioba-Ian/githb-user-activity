package main

import (
	"flag"
)

func main() {
	ghUserName := flag.String("username", "Rioba-Ian", "github username that will be used")
	flag.Parse()

	HandleResponse(*ghUserName)
}
