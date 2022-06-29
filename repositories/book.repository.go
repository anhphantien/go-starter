package repositories

// import (
// 	"go-starter/entities"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/jinzhu/copier"
// 	"gorm.io/gorm/clause"
// )

// type BookRepository struct{}

// func (r BookRepository) FindOneByID(id any) (book entities.Book, err error) {
// 	err = CreateSqlBuilder(book).
// 		Joins("User").
// 		Where("book.id = ?", utils.ConvertToID(id)).
// 		Take(&book).Error
// 	return book, err
// }

// func (r BookRepository) Create(body dto.CreateBookBody) (book entities.Book, err error) {
// 	copier.Copy(&book, body)
// 	err = CreateSqlBuilder(book).Create(&book).Error
// 	return book, err
// }

// func (r BookRepository) Update(c *fiber.Ctx, body dto.UpdateBookBody) (book entities.Book, err error) {
// 	id := c.Params("id")

// 	book, err = BookRepository{}.FindOneByID(id)
// 	if err != nil {
// 		return book, err
// 	}

// 	copier.Copy(&book, body)
// 	err = CreateSqlBuilder(book).
// 		Omit(clause.Associations). // skip auto create/update
// 		Updates(utils.FilterRequestBody(c, body)).Error
// 	return book, err
// }

// func (r BookRepository) Delete(id any) (err error) {
// 	book, err := BookRepository{}.FindOneByID(id)
// 	if err != nil {
// 		return err
// 	}

// 	err = CreateSqlBuilder(book).Delete(&book).Error
// 	return err
// }
