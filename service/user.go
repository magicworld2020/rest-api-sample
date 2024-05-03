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

// func (BookService) GetBookList() []model.Book {
// 	tests := make([]model.Book, 0)
// 	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return tests
// }

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
