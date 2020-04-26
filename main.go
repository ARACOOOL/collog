package main

import (
	"flag"
	"github.com/aracoool/face/http"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	dsn := flag.String("dsn", "", "A data source name (DSN)")
	jwtSecret := flag.String("jk", "", "JWT secret key")
	serverHost := flag.String("h", "", "Server host")
	serverLogsDir := flag.String("ld", "var/logs", "Directory for the server logs")
	flag.Parse()

	initLogs(*serverLogsDir)

	http.RunNewServer(http.Config{
		Dsn:        *dsn,
		JwtSecret:  *jwtSecret,
		ServerHost: *serverHost,
	})
}

func initLogs(logsDir string) {
	file, err := os.OpenFile(logsDir + "/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error: problem with the creation of a logs file ("+logsDir+"), check you logs directory\n")
		os.Exit(2)
	}

	logrus.SetOutput(file)
	logrus.SetLevel(logrus.InfoLevel)
}