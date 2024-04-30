package repository

import (
	"fmt"
	"omc_go/internal/domain"
)

type repository struct{}

func Repository() domain.Repository {
	return &repository{}
}

func (r *repository) Hello() error {
	fmt.Println("Hello, world")
	return nil
}
