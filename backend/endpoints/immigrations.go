package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wait_a_minute/backend/models/immigrationModel"

	"github.com/gin-gonic/gin"
)

type RequestArgs struct {
	RequestID int `json:"requestID"`
	AdminID int `json:"adminID"`
}

func GetRequests(c *gin.Context) {

	data, err := immigrationModel.GetRequests("created")
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func ApproveReq(c *gin.Context) {
	var vars RequestArgs
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&vars)

	status := immigrationModel.ApproveRequest(vars.RequestID, vars.AdminID)
	if status == 200 {
		fmt.Println("Request Approved")
		c.IndentedJSON(http.StatusOK, "Request Approved")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error in Approval process")
	}
}
