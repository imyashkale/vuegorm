package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imyashkale/gormpractice/database"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Skill     string `json:"skill"`
	Age       uint16 `json:"age"`
}

func main() {

	db := database.InitDB().Table("records")

	//migrating
	db.AutoMigrate(&Record{})

	router := gin.Default()

	//post route
	router.Use(cors.Default())
	router.POST("/record", postHandler)
	// start the gin server
	router.Run()
}

//post request handler
func postHandler(c *gin.Context) {
	db := database.InitDB().Table("records")
	var rec Record
	c.BindJSON(&rec)
	db.Create(&rec)
	c.JSON(http.StatusOK, gin.H{
		"first_name": rec.FirstName,
	})
}
