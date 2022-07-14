package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wait_a_minute/backend/immigration"

	"github.com/gin-gonic/gin"
)

type RequestArgs struct {
	RequestID int `json:"requestID"`
	AdminID int `json:"adminID"`
}

func GetRequests(c *gin.Context) {

	data, err := immigrationModel.GetRequests("created")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "requests.html", gin.H{
			"Content": data,
			"ContentName": "Requests",
			"AddItemURL": "/topic/create",
		})
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
