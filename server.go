package main

import (
	"github.com/hmsayem/employee-server/controller"
	"github.com/hmsayem/employee-server/repository"
	"github.com/hmsayem/employee-server/router"
	"github.com/hmsayem/employee-server/service"
	"os"
)

var (
	fireRepo           = repository.NewFirestoreRepository()
	employeeService    = service.NewEmployeeService(fireRepo)
	employeeController = controller.NewEmployeeController(employeeService)
	httpRouter         = router.NewMuxRouter()
)

func main() {
	httpRouter.Get("/employees", employeeController.GetEmployees)
	httpRouter.Post("/employees", employeeController.AddEmployee)
	httpRouter.Serve(os.Getenv("SERVER_PORT"))
}
