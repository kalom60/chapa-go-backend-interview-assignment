package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	api := r.Group("/api")

	tx := api.Group("/transactions")
	{
		tx.POST("")
		tx.GET("/verify/:tx_ref")
	}

	tf := api.Group("/transfers")
	{
		tf.POST("", s.transfer.InitiateTransfer)
		tf.GET("/verify/:ref", s.transfer.VerifyTransfer)
		tf.GET("", s.transfer.GetAllTransfers)
	}

	bk := api.Group("/banks")
	{
		bk.GET("", s.bank.GetAllBanks)
		bk.GET("/:id", s.bank.GetBankByBankID)
	}

	api.POST("/webhooks/transactions")
	api.POST("/webhooks/transfers")

	return r
}
