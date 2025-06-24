package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
		var person Person
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
