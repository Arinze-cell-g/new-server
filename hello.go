package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    // handler is block of code that is executed for the specific endpoint
	//.GET .POST  is the route,
	
	//  create a new git router
	 router := gin.Default()
     
	 // "/" is the endpoint ,hello handler is the handler
	 router.GET("/helloworld", hellohandler)
     // anotherhandler
	 router.POST("/another", arise)
     
     
     //run the server on port 3000
	_  = router.Run(":3000")
}

func hellohandler(c *gin.Context) {
	c.JSON(200, gin.H{
	"message": "say ahmed",
	"arinze" : "ait",

	})

	
}

func arise(c *gin.Context) {
	fmt.Println("arise at your service")
	c.JSON(200, gin.H{
		"message": "learn go for now",
	    "arinze" : "multi talented",

	})
}




