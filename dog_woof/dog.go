package dogwoof

import (
	"fmt"
	"strings"
)

func WhenDogGrow(s string) string {
	fmt.Println("in dog.go")
	return "Grown " + strings.ToUpper(s)
}