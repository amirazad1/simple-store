package service

import (
	"context"
	"database/sql"
	"errors"
	dbc "github.com/amirazad1/simple-store/db/sqlc"
)

type Sale struct {
	db *sql.DB
	*dbc.Queries
}

func NewSale(dbs *sql.DB) *Sale {
	return &Sale{
		db:      dbs,
		Queries: dbc.New(dbs),
	}
}

type SaleTxParams struct {
	FactorParam dbc.CreateFactorParams
	DetailParam dbc.CreateFactorDetailParams
}

func (sale *Sale) SaleTx(ctx context.Context, arg SaleTxParams) (int64, error) {
	tx, err := sale.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	qtx := sale.Queries.WithTx(tx)

	product, err := qtx.GetProduct(ctx, arg.DetailParam.ProductID)
	if err != nil {
		return -1, err
	}

	if product.PresentNumber < arg.DetailParam.SaleCount {
		return -1, errors.New("not enough product for sale")
	}

	//create factor if id is zero
	if arg.DetailParam.FactorID == 0 {
		result, err := sale.Queries.CreateFactor(ctx, arg.FactorParam)
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
