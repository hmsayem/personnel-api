package main

import (
	"github.com/hmsayem/employee-server/controller"
	"github.com/hmsayem/employee-server/repository"
	"github.com/hmsayem/employee-server/router"
	"github.com/hmsayem/employee-server/service"
)

var (
	fireRepo           = repository.NewFirestoreRepository()
	employeeService    = service.NewEmployeeService(fireRepo)
	employeeController = controller.NewEmployeeController(employeeService)
	httpRouter         = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"

	httpRouter.GET("/employees", employeeController.GetEmployees)
	httpRouter.POST("/employees", employeeController.AddEmployee)
	httpRouter.SERVE(port)
}
