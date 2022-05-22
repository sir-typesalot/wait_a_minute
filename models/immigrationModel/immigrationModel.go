package immigrationModel

import (
	"wait_a_minute/util"
)

type Requests struct {
	ID int
	Type string
	Mapping_ID int
	Title string
	Description string
	Tags string
	Status string
	Change_Datetime string
	Admin_ID *int
	State_Reason *string
}

func GetRequests(status string) ([]Requests, error) {

	var data []Requests

	var query string
	if status != "" {
		query = "SELECT * FROM requests_log WHERE status = ?"
	} else {
		query = "SELECT * FROM requests_log"
	}

    db := util.GetConnection()
	result, err := db.Queryx(query, status)
	if err != nil { return data, err }

	defer result.Close()

	for result.Next() {
        var request Requests
		// Check for error and return
        if err := result.StructScan(&request); 
			err != nil { return data, err }
		// Add data to slice
        data = append(data, request)
    }
	return data, nil
}

func ApproveRequest() {

}


