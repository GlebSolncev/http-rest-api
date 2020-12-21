package apiServer

import (
	"github.com/GlebSolncev/http-rest-api/app"
	"github.com/GlebSolncev/http-rest-api/app/models"
	"github.com/GlebSolncev/http-rest-api/configs"
	"github.com/GlebSolncev/http-rest-api/routers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//APIServer ...
type APIServer struct {
	config         *configs.Config
	logger         *logrus.Logger
	router         *mux.Router
	DatabaseConfig configs.DatabaseConfig
}

// New ..
func New(config *configs.Config) *APIServer {
	return &APIServer{
		config:         config,
		logger:         logrus.New(),
		router:         mux.NewRouter(),
		DatabaseConfig: config.DatabaseConfig,
	}
}

// Start ...
func (s *APIServer) Start() error {
	// Init logger
	err := s.configureLogger()
	app.Check(err)
	s.logger.Info("Starting API server")

	// Build connection to DB
	logrus.Info("Build connection to DB")

	DBUrl := s.config.DatabaseConfig
	configs.DB, err = gorm.Open("mysql", configs.DbURL(DBUrl)) //configs.BuildDatabaseConfig()
	app.Check(err)

	defer configs.DB.Close()
	configs.DB.AutoMigrate(&models.Image{})

	// Route init
	logrus.Info("Init Router")
	res := routers.RouteList()

	err = res.Run(s.config.BindAddr)
	app.Check(err)

	return nil
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	app.Check(err)
	s.logger.SetLevel(level)

	return nil
}

func NewConfig() *configs.Config {
	return configs.NewConfig()
}
