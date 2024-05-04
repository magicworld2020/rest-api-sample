package service

import (
	"errors"

	"github.com/magicworld2020/rest-api-sample/model"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) AddUser(user *model.User) error {
	result := DbEngine.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userService *UserService) GetUserByUserID(userID string) (*model.User, error) {
	user := new(model.User)
	result := DbEngine.Where("user_id = ?", userID).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return user, nil
}
func (userService *UserService) UpdateUser(user *model.User) error {
	result := DbEngine.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func (BookService) UpdateBook(newBook *model.Book) error {
// 	_, err := DbEngine.Id(newBook.Id).Update(newBook)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (BookService) DeleteBook(id int) error {
// 	book := new(model.Book)
// 	_, err := DbEngine.Id(id).Delete(book)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
