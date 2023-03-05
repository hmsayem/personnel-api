package main

import (
	"github.com/hmsayem/clean-architecture-implementation/cache"
	"github.com/hmsayem/clean-architecture-implementation/controller"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"github.com/hmsayem/clean-architecture-implementation/router"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"os"
)

var (
	fireRepo           = repository.NewFirestoreRepository()
	employeeService    = service.NewEmployeeService(fireRepo)
	redisCache         = cache.NewRedisCache(os.Getenv("REDIS_SERVER_HOST"), 0, 0)
	employeeController = controller.NewEmployeeController(employeeService, redisCache)
	httpRouter         = router.NewChiRouter()
)

func main() {
	httpRouter.Get("/employees", employeeController.GetAll)
	httpRouter.Get("/employees/{id}", employeeController.Get)
	httpRouter.Put("/employees/{id}", employeeController.Update)
	httpRouter.Post("/employees", employeeController.Add)
	httpRouter.Serve(os.Getenv("SERVER_PORT"))
}
