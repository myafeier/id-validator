package main

import (
	"fmt"

	idvalidator "github.com/guanguans/id-validator"
)

func main() {
	valid := idvalidator.IsValid("430419198710033703", false)
	fmt.Println(valid)
}
