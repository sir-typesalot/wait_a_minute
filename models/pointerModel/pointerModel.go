package pointerModel

import (
	"wait_a_minute/util"
)

type Pointer struct {
	ID int
	Name string
	Desc string
	Tags string
	TopicID int
	RequestID int
}
// id, topic_id, title, description, tags, request_id
// Need to convert to these to enable sqlx

func GetPointers(topicID int) ([]Pointer, error) {

	var data []Pointer

    db := util.GetConnection()

	result, err := db.Query("SELECT * FROM topic_pointer WHERE topic_id = ?", topicID)
	if err != nil { return data, err }

	defer result.Close()

	for result.Next() {
        var pointer Pointer
		// Check for error and return
        if err := result.Scan(&pointer.ID, &pointer.Name, &pointer.Desc,
			&pointer.Tags, &pointer.TopicID, &pointer.RequestID); 
			err != nil { return data, err }
		// Add data to slice
        data = append(data, pointer)
    }
	return data, nil
}

func CreateNewPointer(name string, desc string, tags string, topicID int) int {
	// TODO: Consider double checking topic ID here?
	db := util.GetConnection()
	
	query := `INSERT INTO requests_log
		(type, mapping_id, title, description, tags, status, change_datetime) 
		VALUES ('pointer', ?, ?, ?, ?, 'created', NOW())`
	result, _ := db.Exec(query, topicID, name, desc, tags)
	
	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		return 200
	} else {
		return 500
	}
}


