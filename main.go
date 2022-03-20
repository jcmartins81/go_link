package main

import (
	"fmt"
	"go_link/app"
	"log"
	"os"
)

func main() {
	fmt.Println("hello")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
