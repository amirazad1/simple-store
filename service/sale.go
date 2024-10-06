package service

import (
	"context"
	"errors"
	dbc "github.com/amirazad1/simple-store/db/sqlc"
	"github.com/amirazad1/simple-store/util/e"
)

type SaleTxParams struct {
	FactorParam dbc.CreateFactorParams
	DetailParam dbc.CreateFactorDetailParams
}

func (store *SQLStore) SaleTx(ctx context.Context, arg SaleTxParams) (int64, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	qtx := store.Queries.WithTx(tx)

	product, err := qtx.GetProduct(ctx, arg.DetailParam.ProductID)
	if err != nil {
		return -1, errors.New(e.PRODUCT_NOT_EXISTS)
	}

	if product.PresentNumber < arg.DetailParam.SaleCount {
		return -1, errors.New("not enough product for sale")
	}

	//create factor if id is zero
	if arg.DetailParam.FactorID == 0 {
		result, err := store.Queries.CreateFactor(ctx, arg.FactorParam)
		if err != nil {
			return -1, err
		}

		arg.DetailParam.FactorID, err = result.LastInsertId()
		if err != nil {
			return -1, err
		}
	}

	detailResult, err := qtx.CreateFactorDetail(ctx, arg.DetailParam)
	if err != nil {
		return -1, err
	}

	detailID, err := detailResult.LastInsertId()
	if err != nil {
		return -1, err
	}

	productArg := dbc.UpdateProductPresentParams{
		PresentNumber: product.PresentNumber - arg.DetailParam.SaleCount,
		ID:            arg.DetailParam.ProductID,
	}

	_, err = qtx.UpdateProductPresent(ctx, productArg)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return detailID, nil
}
