package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	mockdb "github.com/amirazad1/simple-store/db/mock"
	db "github.com/amirazad1/simple-store/db/sqlc"
	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetProductAPI(t *testing.T) {
	product := randomProduct()

	testCases := []struct {
		name          string
		productId     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productId: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				//requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name:      "NotFound",
			productId: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productId: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			productId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/products/%d", tc.productId)
			request, err := http.NewRequest("GET", url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomProduct() db.Product {
	num, _ := faker.RandomInt(1, 1000)
	price, _ := faker.RandomInt(1, 10000000)
	now := time.Now()
	return db.Product{
		ID:   int64(num[0]),
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
		BuyDate:       now,
		SalePrice: sql.NullInt64{
			Int64: int64(price[0]),
			Valid: true,
		},
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotProduct db.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}
