package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func IndexHandlerName(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello world " + name,
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandlerXml(c *gin.Context) {

	c.XML(200, Person{
		FirstName: "Mohamed",
		LastName:  "Labouardy",
	})

}

func main() {

	router := gin.Default()

	router.GET("/", IndexHandler)
	router.GET("/:name", IndexHandlerName)
	router.GET("/:name/xml", IndexHandlerXml)

	router.Run()

}
