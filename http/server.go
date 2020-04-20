package http

import (
	"encoding/json"
	"github.com/aracoool/face/logs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//Server server's structure
type Server struct {
	router  *mux.Router
	command Command
	config  Config
}

//registerRoutes registers server routes
func (s *Server) registerRoutes() {
	s.router.HandleFunc("/logs", authenticationMiddleware(s.config.JwtSecret, s.handlerLogCreate)).Methods("POST")
	s.router.HandleFunc("/logs", authenticationMiddleware(s.config.JwtSecret, s.handlerLogsList)).Methods("GET")
}

//run runs the server
func (s *Server) run() {
	log.Info("Running a server on a port " + s.config.ServerHost)
	log.Fatal(http.ListenAndServe(s.config.ServerHost, s.router))
}

//respondWithJson responds with JSON
func (s *Server) respondWithJson(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ServerError{
			Status:  "error",
			Message: err.Error(),
		})
	}
}

//handlerLogCreate creates a record of a log
func (s *Server) handlerLogCreate(w http.ResponseWriter, r *http.Request) {
	var record logs.Record

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		s.respondWithJson(w, r, ServerError{
			Status:  "error",
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	_ = json.Unmarshal(reqBody, &record)

	record, err = s.command.Create(record)
	if err != nil {
		log.Error(err)
		s.respondWithJson(w, r, ServerError{
			Status:  "error",
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	s.respondWithJson(w, r, record, http.StatusCreated)
}

//handlerLogsList returns a list of logs based on the search criteria
func (s *Server) handlerLogsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	records, err := s.command.List(r.URL.Query())
	if err != nil {
		log.Error(err)
		s.respondWithJson(w, r, ServerError{
			Status:  "error",
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	s.respondWithJson(w, r, records, http.StatusOK)
}

//RunNewServer creates a new server and runs it
func RunNewServer(config Config) {
	router := mux.NewRouter().StrictSlash(true)

	server := Server{
		config: config,
		router: router,
		command: Command{
			Repository: &logs.MysqlRepository{Credentials: config.Dsn},
		},
	}
	server.registerRoutes()
	server.run()
}
