package transfer

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/clients"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type Handler interface {
	InitiateTransfer(c *gin.Context)
	GetAllTransfers(c *gin.Context)
	VerifyTransfer(c *gin.Context)
	TransferWebhook(c *gin.Context)
}

type handler struct {
	service *Service
}

func NewHandler(service *Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) InitiateTransfer(c *gin.Context) {
	var req clients.TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	ref, err := h.service.InitiateTransfer(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, ErrDuplicateTransfer) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Transfer Already exists",
				"status":  "failed",
				"data":    nil,
			})
			return
		}
		if errors.Is(err, ErrFailedInitiate) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": ErrFailedInitiate,
				"status":  "failed",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to initiate transfer",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transfer Queued Successfully",
		"status":  "success",
		"data":    ref,
	})
}

func (h *handler) GetAllTransfers(c *gin.Context) {
	filter := utils.Pagination{
		Page:     utils.GetQueryInt(c.Request, "page", 1),
		PageSize: utils.GetQueryInt(c.Request, "pageSize", 5),
	}

	transfers, err := h.service.GetAllTransfers(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve transfers",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	sortParams := utils.ParseSortParams(c.QueryArray("sortBy"))
	transfers.Meta.SortBy = sortParams

	links := utils.BuildLinks(c.Request, filter.Page, int(transfers.Meta.TotalItems), filter.PageSize, sortParams)
	transfers.Links = links

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfers retrieved",
		"data":    transfers,
	})
}

func (h *handler) VerifyTransfer(c *gin.Context) {
	ref := c.Query("ref")

	tf, err := h.service.VerifyTransfer(ref)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": tf.Message,
			"status":  tf.Status,
			"data":    tf.Data,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": tf.Message,
		"status":  tf.Status,
		"data":    tf.Data,
	})
}

func (h *handler) TransferWebhook(c *gin.Context) {
	var req Transfer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	err := h.service.HandleWebhook(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Trandfer not found",
				"status":  "failed",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update transfer",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "status updated",
		"status":  "success",
	})
}
