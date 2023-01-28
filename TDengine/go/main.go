/*
 * In this test program, we'll create a database and insert 4 records then select out.
 */
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/taosdata/driver-go/v3/taosSql"
)

type config struct {
	hostName   string
	serverPort string
	user       string
	password   string
}

var configPara config
var taosDriverName = "taosSql"
var url string

func init() {
	flag.StringVar(&configPara.hostName, "h", "", "The host to connect to TDengine server.")
	flag.StringVar(&configPara.serverPort, "p", "", "The TCP/IP port number to use for the connection to TDengine server.")
	flag.StringVar(&configPara.user, "u", "root", "The TDengine user name to use when connecting to the server.")
	flag.StringVar(&configPara.password, "P", "taosdata", "The password to use when connecting to the server.")
	flag.Parse()
}

func printAllArgs() {
	fmt.Printf("============= args parse result: =============\n")
	fmt.Printf("hostName:             %v\n", configPara.hostName)
	fmt.Printf("serverPort:           %v\n", configPara.serverPort)
	fmt.Printf("usr:                  %v\n", configPara.user)
	fmt.Printf("password:             %v\n", configPara.password)
	fmt.Printf("================================================\n")
}

func main() {
	printAllArgs()

	url = "root:taosdata@/tcp(" + configPara.hostName + ":" + configPara.serverPort + ")/"

	taos, err := sql.Open(taosDriverName, url)
	checkErr(err, "open database error")
	defer taos.Close()

	taos.Exec("create database if not exists test")
	taos.Exec("use test")
	taos.Exec("create table if not exists tb1 (ts timestamp, a int)")
	_, err = taos.Exec("insert into tb1 values(now, 0)(now+1s,1)(now+2s,2)(now+3s,3)")
	checkErr(err, "failed to insert")
	rows, err := taos.Query("select * from tb1")
	checkErr(err, "failed to select")

	defer rows.Close()
	for rows.Next() {
		var r struct {
			ts time.Time
			a  int
		}
		err := rows.Scan(&r.ts, &r.a)
		if err != nil {
			fmt.Println("scan error:\n", err)
			return
		}
		fmt.Println(r.ts, r.a)
	}
}

func checkErr(err error, prompt string) {
	if err != nil {
		fmt.Println("ERROR: %s\n", prompt)
		panic(err)
	}
}
