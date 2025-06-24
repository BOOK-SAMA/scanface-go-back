package handle

import (
	"CRUD_DOCKER_HOMEPAGE/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Type         string `json:"type"`
	Scanface_num string `json:"scanface_num"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	File_date    string `json:"file_date"`
	File_time    string `json:"file_time"`
	Emp_code     string `json:"emp_code"`
	New_emp_code string `json:"new_emp_code"`
}

func Cutstring(c string) Info {
	raw := c
	result := []string{
		raw[0:3],   // Type
		raw[3:7],   // Scanface_num
		raw[7:15],  // Date
		raw[15:21], // Time
		raw[21:29], // File_date
		raw[29:35], // File_time
		raw[35:43], // Emp_code
		raw[43:51], // New_emp_code
	}
	var info Info
	info.Type = result[0]
	info.Scanface_num = result[1]
	info.Date = result[2]
	info.Time = result[3]
	info.File_date = result[4]
	info.File_time = result[5]
	info.Emp_code = result[6]
	info.New_emp_code = result[7]

	return info
}

func Receiveinfo_useformdata(c *gin.Context) {

	if c.Request.ContentLength == 0 {
		c.JSON(400, gin.H{"error": "ไม่มีข้อมูล request ใน body"})
		return
	}
	// fmt.Println(c.Request)
	info := Info{
		Type:         c.Request.FormValue("type"),
		Scanface_num: c.Request.FormValue("scanface_num"),
		Date:         c.Request.FormValue("date"),
		Time:         c.Request.FormValue("time"),
		File_date:    c.Request.FormValue("file_date"),
		File_time:    c.Request.FormValue("file_time"),
		Emp_code:     c.Request.FormValue("emp_code"),
		New_emp_code: c.Request.FormValue("new_emp_code"),
	}
	// fmt.Println(info)
	if err := insertInfo(db.DB, &info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed", "detail": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status": http.StatusOK,
			"Text":   "สามารถนำข้อมูลเข้า database ได้แล้ว",
		})
		return
	}

}

func insertInfo(conn *sql.DB, info *Info) error {

	if conn == nil {
		return fmt.Errorf("DB connection is nil")
	}

	query := `
		INSERT INTO scanme
		(type, scan_num, date, time, file_date, file_time, empcode, empcode_main)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := conn.Exec(query,
		info.Type,
		info.Scanface_num,
		info.Date,
		info.Time,
		info.File_date,
		info.File_time,
		info.Emp_code,
		info.New_emp_code,
	)
	return err
}

func Receiveinfo_userawdata(c *gin.Context) {
	var json Info
	c.Bind(&json)
	if err := insertInfo(db.DB, &json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed", "detail": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status": http.StatusOK,
			"Text":   "สามารถนำข้อมูลเข้า database ได้แล้ว",
		})
		return
	}
}
