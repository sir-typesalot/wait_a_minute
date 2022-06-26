package immigrationModel

import (
	"fmt"
	"wait_a_minute/backend/util"
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

// For now, just use raw admin ID to complete transaction
// Will NEED to change this to use the Admin Model and auth
func ApproveRequest(requestID int, adminID int) int {

	request := getRequestbyID(requestID)
	var query string

	switch request.Type {
	case "topic":
		query = `INSERT INTO topic
		(name, description, tags, request_id, category_id) 
		VALUES (?, ?, ?, ?, ?)`
	case "category":
		query = `INSERT INTO category
		(name, description, tags, request_id) 
		VALUES (?, ?, ?, ?)`
	case "pointer":
		query = `INSERT INTO topic_pointer
		(title, description, tags, request_id, topic_id) 
		VALUES (?, ?, ?, ?, ?)`
	default:
		query = "topic"
	}

	db := util.GetConnection()
	result, err := db.Exec(query, request.Title, request.Description, request.Tags, request.ID, request.Mapping_ID)
	if err != nil {
		fmt.Println(err.Error())
		return 500
	}
	
	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		
		update, _ := db.Exec(`UPDATE requests_log 
			SET status = 'approved', admin_id = ?, state_reason = 'valid'
			WHERE id = ?`, adminID, request.ID)
		rowsAff, _ := update.RowsAffected()
		if rowsAff == 1 {
			return 200
		} else {
			return 300
		}

	} else {
		return 500
	}
}

func ChangeRequestStatus(requestID int, adminID int) int {
	return 0
}

func getRequestbyID(ID int) Requests {
	var request Requests
	db := util.GetConnection()

	result := db.QueryRowx("SELECT * FROM requests_log WHERE id = ?", ID)

	if err := result.StructScan(&request);
	err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(request.Title)
	}
	
	return request
}


