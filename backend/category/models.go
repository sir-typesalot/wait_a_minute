package categoryModel

import (
	"fmt"
	"wait_a_minute/backend/util"
)

type Category struct {
	Category_ID int
	Name string
	Description string
	Tags string
	Request_ID int
}

func GetAllCategories() ([]Category, error) {

	var data []Category

    db := util.GetConnection()

	result, err := db.Queryx("SELECT * FROM category")
	if err != nil { return data, err }

	defer result.Close()

	for result.Next() {
        var category Category
		// Check for error and return
        if err := result.StructScan(&category); err != nil { return data, err }
		// Add data to slice
        data = append(data, category)
    }
	return data, nil
}

func GetCategoryByName(name string, strict bool) (Category, error) {

	var category Category

    db := util.GetConnection()

	query := ""
	if !strict {
		query = "SELECT * FROM category WHERE name LIKE %?%"
	} else {
		query = "SELECT * FROM category WHERE name = ?"
	}
	result := db.QueryRowx(query, name)

	if err := result.StructScan(&category);
	err != nil {
		fmt.Println(err.Error())
	}
	return category, nil
}

func CreateCategory(name string, desc string, tags string) int {

	db := util.GetConnection()
	// TODO : Check if request already exists 
	query := `INSERT INTO requests_log 
		(type, mapping_id, title, description, tags, status, change_datetime) 
		VALUES ('topic', 0, ?, ?, ?, 'created', NOW())`
	result, _ := db.Exec(query, name, desc, tags)
	
	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		return 200
	} else {
		return 500
	}
}
