package api

import (
	"database/sql"
	db "github.com/amirazad1/simple-store/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateProductRequest struct {
	Name       string         `json:"name" binding:"required"`
	Brand      sql.NullString `json:"brand"`
	Model      sql.NullString `json:"model"`
	InitNumber int32          `json:"init_number" binding:"required,min=1"`
	BuyPrice   int64          `json:"buy_price" binding:"required,min=1"`
	BuyDate    time.Time      `json:"buy_date" binding:"required"`
	SalePrice  sql.NullInt64  `json:"sale_price"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req CreateProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		Name:          req.Name,
		Brand:         req.Brand,
		Model:         req.Model,
		InitNumber:    req.InitNumber,
		PresentNumber: req.InitNumber,
		BuyPrice:      req.BuyPrice,
		BuyDate:       req.BuyDate,
		SalePrice:     req.SalePrice,
	}

	result, err := server.store.Queries.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, id)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.Queries.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type listProductRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

func (server *Server) listProduct(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.Queries.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}
