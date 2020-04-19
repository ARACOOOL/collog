package http

type Config struct {
	// ServerHost is the host where server will be launched ex. ":8080"
	ServerHost string
	// Dsn is a database connection source
	Dsn        string
	// JWT secret key
	JwtSecret  string
}
