package main

import (
	"github.com/hmsayem/clean-architecture-implementation/cache"
	"github.com/hmsayem/clean-architecture-implementation/controller"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"github.com/hmsayem/clean-architecture-implementation/router"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"os"
)

const (
	REFIS_SERVER_HOST = "locahost:6379"
	REDIS_DB_ID       = 0
	REDIS_EXPIRE      = 100
)

var (
	fireRepo           = repository.NewFirestoreRepository()
	employeeService    = service.NewEmployeeService(fireRepo)
	redisCache         = cache.NewRedisCache(REFIS_SERVER_HOST, REDIS_DB_ID, REDIS_EXPIRE)
	employeeController = controller.NewEmployeeController(employeeService, redisCache)
	httpRouter         = router.NewChiRouter()
)

func main() {
	httpRouter.Get("/employees", employeeController.GetEmployees)
	httpRouter.Get("/employees/{id}", employeeController.GetEmployeeByID)
	httpRouter.Post("/employees", employeeController.AddEmployee)
	httpRouter.Serve(os.Getenv("SERVER_PORT"))
}
