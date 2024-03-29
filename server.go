package main

import (
	"github.com/hmsayem/clean-architecture-implementation/cache"
	"github.com/hmsayem/clean-architecture-implementation/controller"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"github.com/hmsayem/clean-architecture-implementation/router"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"log"
	"os"
)

var ()

func main() {
	fireRepo, err := repository.NewFirestoreRepository()
	if err != nil {
		log.Fatal(err)
	}
	employeeService := service.NewEmployeeService(fireRepo)
	redisCache := cache.NewRedisCache()
	employeeController := controller.NewEmployeeController(employeeService, redisCache)
	httpRouter := router.NewChiRouter()

	httpRouter.Get("/employees", employeeController.GetAll)
	httpRouter.Post("/employees", employeeController.Add)
	httpRouter.Get("/employees/{id}", employeeController.Get)
	httpRouter.Put("/employees/{id}", employeeController.Update)
	httpRouter.Delete("/employees/{id}", employeeController.Delete)
	httpRouter.Serve(os.Getenv("SERVER_PORT"))
}
