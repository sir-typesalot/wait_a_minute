package main

import (
	"fmt"
	"github.com/sir-typesalot/wait_a_minute/queries"
)

func main() {
	var s string = "Hello, World!"
  queries.GetQueries(200)
	fmt.Println(s)
}
