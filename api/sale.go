package api

import (
	"database/sql"
	db "github.com/amirazad1/simple-store/db/sqlc"
	"github.com/amirazad1/simple-store/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSaleRequest struct {
	CustomerName   string `json:"customer_name" `
	CustomerMobile string `json:"customer_mobile"`
	Seller         string `json:"seller" binding:"required"`
	FactorID       int64  `json:"factor_id" binding:"min=0"`
	ProductID      int64  `json:"product_id" binding:"required,min=1"`
	SaleCount      int32  `json:"sale_count" binding:"required,min=1"`
	SalePrice      int64  `json:"sale_price" binding:"required,min=1"`
}

func (server *Server) createSale(ctx *gin.Context) {
	var req CreateSaleRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := service.SaleTxParams{
		FactorParam: db.CreateFactorParams{
			CustomerName: sql.NullString{
				String: req.CustomerName,
				Valid:  true,
			},
			CustomerMobile: sql.NullString{
				String: req.CustomerMobile,
				Valid:  true,
			},
			Seller: req.Seller,
		},
		DetailParam: db.CreateFactorDetailParams{
			FactorID:  req.FactorID,
			ProductID: req.ProductID,
			SaleCount: req.SaleCount,
			SalePrice: req.SalePrice,
		},
	}

	result, err := server.store.SaleTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
