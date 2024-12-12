package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/dto"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/domain"
	digitaltaxi "github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/usecases/payperday"
)

// AcceptedContentTypes is a list of all the accepted content types
var AcceptedContentTypes = []string{"application/json", "application/x-www-form-urlencoded"}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	usecase digitaltaxi.PayPerDay
}

// NewPresentationHandlers initializes a new rest handlers usecase
func NewPresentationHandlers(usecase digitaltaxi.PayPerDay) *PresentationHandlersImpl {
	return &PresentationHandlersImpl{
		usecase: usecase,
	}
}

// CreateCoverType handles the POST request to create a new cover type
func (h *PresentationHandlersImpl) CreateCoverType(c *gin.Context) {
	input := &dto.CoverTypeInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	output, err := h.usecase.CreateCoverType(c.Request.Context(), input)
	if err != nil {
		jsonErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// ListCoverTypes handles the GET request to list cover types
func (h *PresentationHandlersImpl) ListCoverTypes(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()

	pageSizeStr := queryParams.Get("page_size")
	page := queryParams.Get("page")

	if page == "" || pageSizeStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": "List may be very large. Please provide pagination information"})

		return
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		jsonErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	pageNumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		jsonErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if pageNumber <= 0 {
		pageNumber = 1
	}

	pagination := &domain.Pagination{
		PageSize: pageSize,
		Page:     pageNumber,
	}

	output, err := h.usecase.ListCoverTypes(ctx.Request.Context(), pagination)
	if err != nil {
		jsonErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}

// CreateProductRate handles the POST request to create a new product rate
func (h *PresentationHandlersImpl) CreateProductRate(c *gin.Context) {
	input := &dto.ProductRateInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	output, err := h.usecase.CreateProductRate(c.Request.Context(), input)
	if err != nil {
		jsonErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// ListProductRates handles the GET request to list rates
func (h *PresentationHandlersImpl) ListProductRates(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	pageSizeStr := queryParams.Get("page_size")
	page := queryParams.Get("page")

	if page == "" || pageSizeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": "List may be very large. Please provide pagination information"})

		return
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	pageNumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if pageNumber <= 0 {
		pageNumber = 1
	}

	pagination := &domain.Pagination{
		PageSize: pageSize,
		Page:     pageNumber,
	}

	output, err := h.usecase.ListProductRates(c.Request.Context(), pagination)
	if err != nil {
		jsonErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func jsonErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
