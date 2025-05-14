package product_repository_test

import (
	fixture "getfund-api-v2/test/internal/domain/product/adapter/repository/product_repository_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetProductById_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetProductById(0)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetProductById_WhenNotFound_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetProductById(0)

	// Assert
	verify.Should(t, err.Error()).Be("product not found")
}
