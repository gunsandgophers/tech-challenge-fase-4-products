package repositories

import (
	// "errors"
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	coreerror "tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertProduct(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	repo := NewProductRepositoryDB(conn)

	product := entities.RestoreProduct(
		uuid.NewString(), "Um produto", entities.PRODUCT_CATEGORY_DRINKS,
		13.37, "Uma descrição", "url",
	)

	conn.On("Exec", mock.Anything, product.GetId(),
		product.GetName(),
		product.GetCategory().String(),
		product.GetPrice(),
		product.GetDescription(),
		product.GetImage(),
	).Return(nil).Once()

	err := repo.Insert(product)

	assert.Nil(t, err)
}

func TestUpdateProduct(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	repo := NewProductRepositoryDB(conn)

	product := entities.RestoreProduct(
		uuid.NewString(), "Um produto", entities.PRODUCT_CATEGORY_DRINKS,
		13.37, "Uma descrição", "url",
	)

	conn.On("Exec", mock.Anything,
		product.GetName(),
		product.GetCategory().String(),
		product.GetPrice(),
		product.GetDescription(),
		product.GetImage(),
		product.GetId(),
	).Return(nil).Once()

	err := repo.Update(product)

	assert.Nil(t, err)
}

func TestDeleteProduct(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("Exec", mock.Anything, productID).Return(nil).Once()

	err := repo.Delete(productID)

	assert.Nil(t, err)
}

func TestDeleteProductWithErrProductNotFound(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("Exec", mock.Anything, productID).Return(errors.New("no rows in result set")).Once()

	err := repo.Delete(productID)

	assert.EqualError(t, err, coreerror.ErrProductNotFound.Error())
}

func TestDeleteProductWithError(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("Exec", mock.Anything, productID).Return(errors.New("error")).Once()

	err := repo.Delete(productID)

	assert.EqualError(t, err, "error")
}

func TestFindProductByID(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)
	row := mocks.NewMockRowDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("QueryRow", mock.Anything, productID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return(nil).Once()

	_, err := repo.FindProductByID(productID)

	assert.Nil(t, err)
}

func TestFindProductByIDWithErrProductNotFound(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)
	row := mocks.NewMockRowDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("QueryRow", mock.Anything, productID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return(errors.New("no rows in result set")).Once()

	_, err := repo.FindProductByID(productID)

	assert.EqualError(t, err, coreerror.ErrProductNotFound.Error())
}

func TestFindProductByIDWithError(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)
	row := mocks.NewMockRowDB(t)

	repo := NewProductRepositoryDB(conn)

	productID := uuid.NewString()

	conn.On("QueryRow", mock.Anything, productID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return(errors.New("error")).Once()

	_, err := repo.FindProductByID(productID)

	assert.EqualError(t, err, "error")
}

func TestFindProductByCategory(t *testing.T) {

	conn := mocks.NewMockConnectionDB(t)
	rows := mocks.NewMockRowsDB(t)

	repo := NewProductRepositoryDB(conn)

	productCategory := entities.PRODUCT_CATEGORY_DESSERTS

	conn.On("Query", mock.Anything, productCategory.String(),
		mock.Anything, mock.Anything,
	).Return(rows, nil).Once()

	rows.On("Next").Return(true).Once()
	rows.On("Next").Return(false).Once()


	rows.On("Scan", mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return(nil).Once()

	_, err := repo.FindProductByCategory(productCategory, 1, 2)

	assert.Nil(t, err)
}
