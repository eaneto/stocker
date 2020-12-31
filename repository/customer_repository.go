package repository

import "github.com/eaneto/stocker/domain"

type CustomerRepository interface {
	Save(customer domain.CustomerEntity) error
}
