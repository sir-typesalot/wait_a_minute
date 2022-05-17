package models

import (
	"fmt"
	"log"
	"wait_a_minute/util"
)

func TestIt() []string {

    db := util.GetConnection()
	fmt.Println("Connected to Database")

	result, err := db.Query("Select * from category")
	defer result.Close()

	data := make([]string, 3)
	if err != nil {
		log.Fatal(err)
	}
	print("Success")
	// for result.Next() {
	// 	data = append(data, value)
	// }
	return data
}
