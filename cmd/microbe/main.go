package main

import (
	"fmt"
	"github.com/JakubOboza/microbe/pkg/microbe"
)

func main() {

	// Gin uses PORT env if you wanna force port exampl PORT=6787
	app := microbe.NewMicrobe()

	err := app.Run()

	if err != nil {
		fmt.Println(err)
	}

}
