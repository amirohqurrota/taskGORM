package database

import (
	"GORM-2/config"
	"GORM-2/models"
)

// func GetArticles() (*[]models.Article, error) {
// 	var articles []models.Article

// 	if err := config.DB.Find(&articles).Error; err != nil {
// 		return &[]models.Article{}, err
// 	}
// 	return &articles, nil
// }

// func GetUsers() (*[]models.User, error) {
// 	var users []models.User
// 	if err := config.DB.Find(&users).Error; err != nil {
// 		return &[]models.Use{}, err
// 	}
// 	return *users, nil
// }

// func GetbyIDArticles(id int) (*models.Article, error) {
// 	var article models.Article

// 	if err := config.DB.Where("id=?", id).First(&article).Error; err != nil {
// 		return &models.Article{}, err
// 	}
// 	return &article, nil
// }

func ReqUserBook() (*[]models.User, error) {
	var user []models.User

	if err := config.DB.Find(&user).Error; err != nil {
		return &[]models.User{}, err
	}
	return &user, nil
}

func ReqAddNewUser(newUser models.User) (*models.User, error) {
	// var newUser = models.User{
	// 	Model:     gorm.Model{},
	// 	Name:      name,
	// 	Age:       age,
	// 	Sex:       sex,
	// 	Client_id: client_id,
	// }

	if err := config.DB.Save(&newUser).Error; err != nil {
		return &models.User{}, err
	}
	return &newUser, nil

}

func ReqDeleteUser(id int) (*models.User, error) {
	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return &models.User{}, err
	}
	return &models.User{}, nil

}
