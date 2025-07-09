package transaction

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Chapa-Et/chapa-go"
	"github.com/gin-gonic/gin"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

func computeHMAC(body []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil))
}

type Handler interface {
	InitiateTransaction(c *gin.Context)
	GetTransactions(c *gin.Context)
	GetAllTransactions(c *gin.Context)
	VerifyTransaction(c *gin.Context)
	TransactionWebhook(c *gin.Context)
}

type handler struct {
	webhookSecret string
	service       *Service
}

func NewHandler(webhookSecret string, service *Service) Handler {
	return &handler{
		webhookSecret: webhookSecret,
		service:       service,
	}
}

func (h *handler) InitiateTransaction(c *gin.Context) {
	var req chapa.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	ref, err := h.service.InitiateTransaction(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, ErrDuplicateTransaction) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Transaction Already exists",
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
			"message": "Failed to initiate transaction",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": ref,
	})
}

func (h *handler) GetTransactions(c *gin.Context) {
	transactions, err := h.service.GetTransactions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve transactions",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transactions retrieved",
		"data":    transactions,
	})
}

func (h *handler) GetAllTransactions(c *gin.Context) {
	filter := utils.Pagination{
		Page:     utils.GetQueryInt(c.Request, "page", 1),
		PageSize: utils.GetQueryInt(c.Request, "pageSize", 5),
	}

	transactions, err := h.service.GetAllTransactions(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve transactions",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	sortParams := utils.ParseSortParams(c.QueryArray("sortBy"))
	transactions.Meta.SortBy = sortParams

	links := utils.BuildLinks(c.Request, filter.Page, int(transactions.Meta.TotalItems), filter.PageSize, sortParams)
	transactions.Links = links

	c.JSON(http.StatusOK, gin.H{
		"message": "Transactions retrieved",
		"data":    transactions,
	})
}

func (h *handler) VerifyTransaction(c *gin.Context) {
	ref := c.Param("tx_ref")

	tf, err := h.service.VerifyTransaction(ref)
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

func (h *handler) TransactionWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	sig := c.GetHeader("X-Chapa-Signature")
	if sig == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Missing signature",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	expectedSig := computeHMAC(body, h.webhookSecret)

	if !hmac.Equal([]byte(sig), []byte(expectedSig)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid signature",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	var payload Transaction
	if err := json.Unmarshal(body, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	err = h.service.HandleWebhook(c.Request.Context(), payload)
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
			"message": "Failed to update transaction",
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
