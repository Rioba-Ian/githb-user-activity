package main

import (
	"flag"
	"fmt"
)

func main() {
	ghUserName := flag.String("username", "Rioba-Ian", "github username that will be used")
	flag.Parse()

	ghResults := HandleResponse(*ghUserName)

	fmt.Println(ghResults)

}
