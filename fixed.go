package main

import (
	"database/sql"
	"log"
	"encoding/json"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/gin-gonic/gin.v1"
)

func fixed(c *gin.Context) {
	requestid := c.Query("RequestId") // shortcut for c.Request.URL.Query().Get("RequestId")
	msisdn := c.Query("Msisdn")
	channel := c.Query("Channel")

	var rrequestid string
	var rmsisdn string
	var rrecharge string
	var rdate string

	conn, err := sql.Open("mssql", "server=192.168.144.241;user id=reporter;password=blogic;")
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("exec [db_RTPC_LR].[dbo].[stp_api_get_postpaid] ?, ?, ?")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(requestid, msisdn, channel)

	err = row.Scan(&rrequestid, &rmsisdn, &rrecharge, &rdate)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	mobilerechargeresponse := &MobileLastRechargeResponse{RequestId: rrequestid, Msisdn: rmsisdn, Recharge: rrecharge, Date: rdate}
	mobilerechargeresponsejson, err := json.Marshal(mobilerechargeresponse)
	if err != nil {
		log.Fatal("Marshal failed: %s", err.Error())
	}
	//c.String(http.StatusOK, "API called with: %s %s %s %s", rrequestid, rmsisdn, rrecharge, rdate)
	c.String(http.StatusOK, "%s", mobilerechargeresponsejson)
}
