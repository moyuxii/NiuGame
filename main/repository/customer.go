package repository

import (
	"NiuGame/main/model"
	"log"
)

type CustomerRepo struct {
	DB model.DataBase
}

func (r *CustomerRepo) CheckCustomerPasswd(customer model.Customer) bool {
	var count int64
	if err := r.DB.SqlLite.Model(model.Customer{}).Where(&customer).Count(&count).Error; err != nil {
		log.Panicln(err)
	}
	return count == 1
}
