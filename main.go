package main

import (
	"fmt"
	"wait_a_minute/queries"
)

func main() {
	var s string = "Hello, World!"
  	queries.GetQueries(200)
	fmt.Println(s)
}
