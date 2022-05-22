package topicModel

import (
	"wait_a_minute/util"
	"wait_a_minute/models/categoryModel"
)

type Topic struct {
	ID int
	Name string
	Desc string
	Tags string
	CategoryID int
	RequestID int
}

func GetAllTopics() ([]Topic, error) {

	var data []Topic

    db := util.GetConnection()

	result, err := db.Query("SELECT * FROM category")
	if err != nil { return data, err }

	defer result.Close()

	for result.Next() {
        var topic Topic
		// Check for error and return
        if err := result.Scan(&topic.ID, &topic.Name, &topic.Desc,
			&topic.Tags, &topic.CategoryID, &topic.RequestID); 
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
	result, _ := db.Exec(query, category[0].ID, name, desc, tags)
	
	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		return 200
	} else {
		return 500
	}
}


