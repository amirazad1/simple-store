package service

import (
	"context"
	"database/sql"
	"github.com/amirazad1/simple-store/db/sqlc"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomProductWitNum(t *testing.T, num int32) db.Product {
	price, _ := faker.RandomInt(1, 10000000)
	arg := db.CreateProductParams{
		Name: faker.Name(),
		Brand: sql.NullString{
			String: faker.LastName(),
			Valid:  true,
		},
		Model: sql.NullString{
			String: faker.FirstName(),
			Valid:  true,
		},
		InitNumber:    num,
		PresentNumber: num,
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

func createRandomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username: faker.Name(),
		Password: faker.Password(),
		FullName: faker.Name(),
		Mobile:   faker.Phonenumber()[0:10],
		PasswordChangedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	_, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Mobile, user.Mobile)
	require.WithinDuration(t, arg.PasswordChangedAt.Time, user.PasswordChangedAt.Time, time.Minute)

	return user
}

func TestSaleTx(t *testing.T) {
	sale := NewStore(testDB)

	n := 5
	product := createRandomProductWitNum(t, int32(n))
	user := createRandomUser(t)

	errs := make(chan error)
	results := make(chan int64)
	prices := make(chan int64)

	for i := 0; i < n; i++ {
		go func() {
			price, _ := faker.RandomInt(1, 1000000)
			result, err := sale.SaleTx(context.Background(), SaleTxParams{
				FactorParam: db.CreateFactorParams{
					CustomerName: sql.NullString{
						String: faker.Name(),
						Valid:  true,
					},
					CustomerMobile: sql.NullString{
						String: faker.Phonenumber()[0:10],
						Valid:  true,
					},
					Seller: user.Username,
				},
				DetailParam: db.CreateFactorDetailParams{
					FactorID:  0,
					ProductID: product.ID,
					SaleCount: 1,
					//SaleCount: int32(saleCount[i]),
					SalePrice: int64(price[0]),
				},
			})

			errs <- err
			results <- result
			prices <- int64(price[0])
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		price := <-prices

		detail, err := testQueries.GetFactorDetail(context.Background(), result)
		require.NoError(t, err)
		require.Equal(t, int32(1), detail.SaleCount)
		require.Equal(t, price, detail.SalePrice)
	}
}
