package main

import (
	"fmt"
)

type Team struct {
	Uniqie_Id string `json:"unique_id,omitempty"`
	Name      string `json:"name"`
	Account   string `json:"account,omitempty"`
}

func main() {
	fmt.Println("Hello, 世界")
}
