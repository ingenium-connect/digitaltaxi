package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/dto"
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

func jsonErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}
