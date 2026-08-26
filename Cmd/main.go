package main

import (
	"fmt"
     "GitHub/management-/API"
	"GitHub/management-/Controller"

)


func main() {

	api := api.APIRouts{}
	api.StartAPI(controller.Controller{})
	fmt.Printf("main server = %v\n", api)
}

