package queries
import (
	"fmt"
)

func GetQueries(status int) {
  // Case-Switch for status
  switch status {
  case 200:
    fmt.Println("Success")
  case 404:
    fmt.Println("Not Found")
  default:
    fmt.Println("Undefined")
  }
}