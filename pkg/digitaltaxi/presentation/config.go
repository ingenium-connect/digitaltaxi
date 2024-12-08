package presentation

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/presentation/rest"
)

const serverTimeoutSeconds = 120

var allowedOriginPatterns = []string{
	`^(https?://)?(.+)-?.ibima\.co.ke$`,
}

// PayPerDayAllowedHeaders is a list of CORS allowed headers for the pay-per-day
// service
var PayPerDayAllowedHeaders = []string{
	"Accept",
	"Accept-Charset",
	"Accept-Language",
	"Accept-Encoding",
	"Origin",
	"Host",
	"User-Agent",
	"Content-Length",
	"Content-Type",
	"Authorization",
	"X-Authorization",
}

// PrepareServer sets up the HTTP server
func PrepareServer(ctx context.Context, port int) {
	// start up the router
	ginEngine := gin.Default()

	err := StartGinRouter(ctx, ginEngine)
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	addr := fmt.Sprintf(":%v", port)

	if err := ginEngine.Run(addr); err != nil {
		helpers.LogStartupError(ctx, err)
	}
}

func isAllowedOrigin(origin string, compiledPatterns []*regexp.Regexp) bool {
	for _, pattern := range compiledPatterns {
		if pattern.MatchString(origin) {
			return true
		}
	}

	return false
}

// Compile the regex patterns into a slice of *regexp.Regexp
func compilePatterns(patterns []string) []*regexp.Regexp {
	var compiledPatterns []*regexp.Regexp

	for _, pattern := range patterns {
		compiledPattern := regexp.MustCompile(pattern)
		compiledPatterns = append(compiledPatterns, compiledPattern)
	}

	return compiledPatterns
}

// StartGinRouter sets up the GIN router
func StartGinRouter(ctx context.Context, engine *gin.Engine) error {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Set allowed origins
	compiledPatterns := compilePatterns(allowedOriginPatterns)

	engine.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowMethods:     []string{http.MethodPut, http.MethodPatch, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     PayPerDayAllowedHeaders,
		ExposeHeaders:    []string{"Content-Length", "Link"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// localhost to allow
			if origin == "http://localhost:8000" || origin == "http://localhost:8080" {
				return true
			}

			allowed := isAllowedOrigin(origin, compiledPatterns)
			return allowed
		},
		MaxAge:          12 * time.Hour,
		AllowWebSockets: true,
	}))

	datastore := datastore.NewDbService()

	infrastructure := infrastructure.NewInfrastructureInteractor(datastore)
	_ = rest.NewPresentationHandlers(infrastructure)

	_ = engine.Group("/api/v1")
	{

	}

	return nil
}
