package bank

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kalom60/chapa-go-backend-interview-assignment/pkg/utils"
)

type Handler interface {
	GetAllBanks(c *gin.Context)
	GetBankByBankID(c *gin.Context)
}

type handler struct {
	service *Service
}

func NewHandler(service *Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllBanks(c *gin.Context) {
	filter := utils.Pagination{
		Page:     utils.GetQueryInt(c.Request, "page", 1),
		PageSize: utils.GetQueryInt(c.Request, "pageSize", 5),
	}

	banks, err := h.service.GetAllBanks(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve banks",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	sortParams := utils.ParseSortParams(c.QueryArray("sortBy"))
	banks.Meta.SortBy = sortParams

	links := utils.BuildLinks(c.Request, filter.Page, int(banks.Meta.TotalItems), filter.PageSize, sortParams)
	banks.Links = links

	c.JSON(http.StatusOK, gin.H{
		"message": "Banks retrieved",
		"data":    banks,
	})
}

func (h *handler) GetBankByBankID(c *gin.Context) {
	id := c.Param("id")
	bankID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid bank id",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	bank, err := h.service.GetBankByBankID(c.Request.Context(), bankID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve bank",
			"status":  "failed",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bank retrieved",
		"data":    bank,
	})
}
