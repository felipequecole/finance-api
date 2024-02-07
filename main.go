package main

import (
	"finances-api/configuration"
	"finances-api/controller"
	"finances-api/repository"
	"finances-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func loadYaml(fileName string, model interface{}) {
	validate := validator.New()
	f, fsError := os.Open(fileName)
	if fsError != nil {
		log.Fatalln(fmt.Sprintf("Failed to read configuration file with name %s", fileName))
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(fmt.Sprintf("Failed to close file %s", fileName))
		}
	}(f)

	decoder := yaml.NewDecoder(f)

	if decodeError := decoder.Decode(model); decodeError != nil {
		log.Printf("Failed to close file %s", fileName)
		panic(decodeError)
	}

	if validationError := validate.Struct(model); validationError != nil {
		log.Printf("Invalid configuration file %s", fileName)
		panic(validationError)
	}
}

func main() {

	// Load configs
	var appConfig configuration.AppConfig
	loadYaml("config.yaml", &appConfig)

	var secrets configuration.Secrets
	loadYaml("secrets.yaml", &secrets)

	// Connect to database
	db := configuration.DatabaseConnection(appConfig, secrets)

	// Init Expense related
	expenseRepository := repository.NewExpenseRepositoryImpl(db)
	expenseService := service.NewExpenseService(expenseRepository)
	expenseController := controller.NewExpenseController(expenseService)

	// Init Gin
	r := gin.Default()

	// Expense Routes
	r.GET("/expenses/:id", expenseController.FindById)
	r.GET("/expenses", expenseController.FindAll)

	// Start Server
	if err := r.Run(fmt.Sprintf("%s:%d", appConfig.App.Host, appConfig.App.Port)); err != nil {
		panic(err)
	}
}
