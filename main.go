package main

import (
	"flag"
	"github.com/aracoool/face/http"
)

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
