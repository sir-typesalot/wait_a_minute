package topicModel

import (
	"fmt"
	"wait_a_minute/backend/models/categoryModel"
	"wait_a_minute/backend/util"
)

type Topic struct {
	Topic_ID int
	Name string
	Description string
	Tags string
	Category_ID int
	Request_ID int
}

func GetAllTopics(categoryID int) ([]Topic, error) {

	var data []Topic

	var query string
	if categoryID != 0 {
		query = fmt.Sprintf("SELECT * FROM topic WHERE category_id = %d", categoryID)
	} else {
		query = "SELECT * FROM topic"
	}
	db := util.GetConnection()
	result, err := db.Queryx(query)
	if err != nil { return data, err }

	defer result.Close()

	for result.Next() {
        var topic Topic
		// Check for error and return
        if err := result.StructScan(&topic); 
			err != nil { return data, err }
		// Add data to slice
        data = append(data, topic)
    }
	return data, nil
}

func CreateNewTopic(name string, desc string, tags string, categoryName string) int {
	category, _ := categoryModel.GetCategoryByID(categoryName, true)

	db := util.GetConnection()
	
	query := `INSERT INTO requests_log 
		(type, mapping_id, title, description, tags, status, change_datetime) 
		VALUES ('topic', ?, ?, ?, ?, 'created', NOW())`
	result, _ := db.Exec(query, category[0].Category_ID, name, desc, tags)
	
	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		return 200
	} else {
		return 500
	}
}


