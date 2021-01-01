package controller

import (
	"net/http"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/service"
)

type BaseCustomerController interface {
	Create(name string) int
	FindAll() ([]domain.Customer, int)
}

type CustomerController struct {
	CustomerService service.BaseCustomerService
}

func NewCustomerController() BaseCustomerController {
	return CustomerController{
		CustomerService: service.NewCustomerService(),
	}
}

func (controller CustomerController) Create(name string) int {
	controller.CustomerService.Create(name)
	return http.StatusCreated
}

func (controller CustomerController) FindAll() ([]domain.Customer, int) {
	return controller.CustomerService.FindAll(), http.StatusOK
}
