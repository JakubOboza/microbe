package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/JakubOboza/microbe/pkg/microscope"
)

/*
THIS IS A SIMPLE MICROBE CLI CLIENT


examples: (λ = returned value in shell)
./bin/microscope -operation base64enc -data "this is a client that kinda works"

λ dGhpcyBpcyBhIGNsaWVudCB0aGF0IGtpbmRhIHdvcmtz

./bin/microscope -operation base64dec -data dGhpcyBpcyBhIGNsaWVudCB0aGF0IGtpbmRhIHdvcmtz

λ this is a client that kinda works

*/

func dataIsrequired(data string) {
	if strings.TrimSpace(data) == "" {
		fmt.Println("data param is required and can't be empty")
		flag.PrintDefaults()
		panic(-1)
	}
}

func main() {

	endpointPtr := flag.String("endpoint", "http://localhost:8080", "microbe endpoint")
	operationPtr := flag.String("operation", "base64enc", "api call to make")
	dataPtr := flag.String("data", "", "data to be passed to api call")

	flag.Parse()

	client := microscope.NewClient(*endpointPtr)

	switch *operationPtr {
	case "ping":
		result, err := client.Ping()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(*result)
		}
	case "base64enc":
		dataIsrequired(*dataPtr)
		result, err := client.Base64Encode(*dataPtr)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(*result)
		}
	case "base64dec":
		dataIsrequired(*dataPtr)
		result, err := client.Base64Decode(*dataPtr)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(*result)
		}
	}

}
