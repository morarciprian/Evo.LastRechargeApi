package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /LastRecharge/api/Mobile?RequestId=1010&Msisdn=99980363&Channel=WEB
	router.GET("/LastRecharge/api/Mobile", mobile)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /LastRecharge/api/Fixed?RequestId=1010&Msisdn=99980363&Channel=WEB
	router.GET("/LastRecharge/api/Fixed", fixed)

	router.Run(":808")
}
