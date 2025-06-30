package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string ` json:"password"  binding:"required"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NewPersonData struct {
	AccountID  string `json:"accountid"`
	EmpName    string `json:"empName"`
	AppRole    string `json:"approle"`
	Permission string `json:"permission"`
	Extra      string `json:"extra"`
	Position   string `json:"position"`
	OrgCode    string `json:"orgcode"`
	Org1       string `json:"org1"`
	Org2       string `json:"org2"`
	Org3       string `json:"org3"`
	Org4       string `json:"org4"`
}

func Login(c *gin.Context) {
	var user User
	c.Bind(&user)
	fmt.Println(user)

	var testpassword = "testme"
	if user.Password == testpassword {
		data, err := os.ReadFile("info.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// decode JSON จากไฟล์โดยตรง
		var person NewPersonData
		if err := json.Unmarshal(data, &person); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"tokenstring": Gentoken(person),
			"Text":        "you logined",
		})
		return
	}

}

func TestToken(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)
	info := ReciveTokenNew(tokenString)
	c.JSON(http.StatusOK, gin.H{
		"info": info,
		"Text": "you logined",
	})
	return
}
