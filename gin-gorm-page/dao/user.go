package dao

import (
	"gin-gorm-page/config"
	models "gin-gorm-page/model"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, error) {
	var users []models.User
	offset := (pagination.Page - 1) * pagination.Limit
	// queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := config.DB.Model(&models.User{}).Where(user).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
