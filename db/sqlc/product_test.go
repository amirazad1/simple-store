package db

import (
	"context"
	"database/sql"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomProduct(t *testing.T) Product {
	num, _ := faker.RandomInt(1, 1000)
	price, _ := faker.RandomInt(1, 10000000)
	arg := CreateProductParams{
		Name: faker.Name(),
		Brand: sql.NullString{
			String: faker.LastName(),
			Valid:  true,
		},
		Model: sql.NullString{
			String: faker.FirstName(),
			Valid:  true,
		},
		InitNumber:    int32(num[0]),
		PresentNumber: int32(num[0]),
		BuyPrice:      int64(price[0]),
		BuyDate:       time.Now(),
		SalePrice: sql.NullInt64{
			Int64: int64(price[0]),
			Valid: true,
		},
	}

	result, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	id, err := result.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, id)

	product, err := testQueries.GetProduct(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Model, product.Model)
	require.Equal(t, arg.InitNumber, product.InitNumber)
	require.Equal(t, arg.PresentNumber, product.PresentNumber)
	require.Equal(t, arg.BuyPrice, product.BuyPrice)
	require.WithinDuration(t, arg.BuyDate, product.BuyDate, time.Hour*24)
	require.Equal(t, arg.SalePrice, product.SalePrice)

	return product
}

func TestQueries_CreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestQueries_GetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Brand, product2.Brand)
	require.Equal(t, product1.Model, product2.Model)
	require.Equal(t, product1.InitNumber, product2.InitNumber)
	require.Equal(t, product1.PresentNumber, product2.PresentNumber)
	require.Equal(t, product1.BuyPrice, product2.BuyPrice)
	require.Equal(t, product1.BuyDate, product2.BuyDate)
	require.Equal(t, product1.SalePrice, product2.SalePrice)
}

func TestQueries_UpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	num, _ := faker.RandomInt(1, 1000)
	price, _ := faker.RandomInt(1, 10000000)
	arg := UpdateProductParams{
		ID:   product1.ID,
		Name: faker.Name(),
		Brand: sql.NullString{
			String: faker.LastName(),
			Valid:  true,
		},
		Model: sql.NullString{
			String: faker.FirstName(),
			Valid:  true,
		},
		InitNumber:    int32(num[0]),
		PresentNumber: int32(num[0]),
		BuyPrice:      int64(price[0]),
		BuyDate:       time.Now(),
		SalePrice: sql.NullInt64{
			Int64: int64(price[0]),
			Valid: true,
		},
	}

	result, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	affected, err := result.RowsAffected()
	require.NoError(t, err)
	require.NotEmpty(t, affected)

	product2, err := testQueries.GetProduct(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)
	require.Equal(t, arg.Name, product2.Name)
	require.Equal(t, arg.Brand, product2.Brand)
	require.Equal(t, arg.Model, product2.Model)
	require.Equal(t, arg.InitNumber, product2.InitNumber)
	require.Equal(t, arg.PresentNumber, product2.PresentNumber)
	require.Equal(t, arg.BuyPrice, product2.BuyPrice)
	require.WithinDuration(t, arg.BuyDate, product2.BuyDate, time.Hour*24)
	require.Equal(t, arg.SalePrice, product2.SalePrice)
}

func TestQueries_UpdateProductPresent(t *testing.T) {
	product1 := createRandomProduct(t)

	num, _ := faker.RandomInt(1, 1000)
	arg := UpdateProductPresentParams{
		ID:            product1.ID,
		PresentNumber: int32(num[0]),
	}

	result, err := testQueries.UpdateProductPresent(context.Background(), arg)
	require.NoError(t, err)
	affected, err := result.RowsAffected()
	require.NoError(t, err)
	require.NotEmpty(t, affected)

	product2, err := testQueries.GetProduct(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)
	require.Equal(t, arg.PresentNumber, product2.PresentNumber)
}

func TestQueries_DeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestQueries_ListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
