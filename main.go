package main

import (
	"flag"
	"github.com/aracoool/face/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	file, err := os.OpenFile("var/logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)
	log.SetLevel(log.InfoLevel)
}

func main() {
	dsn := flag.String("dsn", "", "A data source name (DSN)")
	jwtSecret := flag.String("jk", "", "JWT secret key")
	serverHost := flag.String("h", "", "Server host")
	flag.Parse()


	http.RunNewServer(http.Config{
		Dsn:        *dsn,
		JwtSecret:  *jwtSecret,
		ServerHost: *serverHost,
	})
}
