package controller

import (
	"fmt"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/NewChakrit/golang_gin_vuejs_full-stack/services"
)

type Controller struct {
	TransactionService services.TransactionService
}

type Config struct {
	R                  *gin.Engine
	TransactionService services.TransactionService
}

func NewController(c *Config) {
	controller := &Controller{
		TransactionService: c.TransactionService,
	}

	apiRouters := c.R.Group("/api") // start path with /api

	{
		apiRouters.POST("/txn/add", controller.AddTransactions)
	}
}

func (c *Controller) FindAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func (c *Controller) AddTransactions(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.TransactionService.Add(ctx.Request.Context(), txn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Create transaction Successfully!"})
}

func (c *Controller) EditTransactions(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.TransactionService.Edit(ctx, txn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Update Transaction ID: %v, Successfully!", txn.ID)})
}

func (c *Controller) DeleteTransactions(ctx *gin.Context) {
	//var request struct {
	//	ID int `json:"id"`
	//}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err = ctx.ShouldBindJSON(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = c.TransactionService.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Transaction ID: %v, Has been delete", id)})
}
