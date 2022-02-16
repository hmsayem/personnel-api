package main

import (
	"github.com/hmsayem/clean-architecture-implementation/controller"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"github.com/hmsayem/clean-architecture-implementation/router"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"os"
)

var (
	fireRepo           = repository.NewFirestoreRepository()
	employeeService    = service.NewEmployeeService(fireRepo)
	employeeController = controller.NewEmployeeController(employeeService)
	httpRouter         = router.NewChiRouter()
)

func main() {
	httpRouter.Get("/employees", employeeController.GetEmployees)
	httpRouter.Post("/employees", employeeController.AddEmployee)
	httpRouter.Serve(os.Getenv("SERVER_PORT"))
}
