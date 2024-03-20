package httpcontroller

import (
	"darkness8129/validation-service/app/service"
	"darkness8129/validation-service/packages/logging"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Router   *gin.Engine
	Logger   logging.Logger
	Services service.Services
}

type controllerOptions struct {
	RouterGroup *gin.RouterGroup
	Logger      logging.Logger
	Services    service.Services
}

func New(opt Options) {
	opt.Router.Use(gin.Logger(), gin.Recovery(), corsMiddleware)

	controllerOpt := controllerOptions{
		RouterGroup: opt.Router.Group("/api/v1"),
		Logger:      opt.Logger.Named("httpController"),
		Services:    opt.Services,
	}

	newCreditCardController(controllerOpt)
}

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

// httpErr provides a base error type for all http controller errors.
type httpErr struct {
	Type    httpErrType `json:"-"`
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
}

type httpErrType string

const (
	httpErrTypeServer httpErrType = "server"
	httpErrTypeClient httpErrType = "client"
)

// httpResponseBody provides a base response type for all http controllers.
type httpResponseBody struct {
	Response interface{} `json:"response"`
	Err      *httpErr    `json:"err,omitempty"`
}

// errorDecorator provides unified error handling for all http controllers
func errorDecorator(logger logging.Logger, handler func(c *gin.Context) (interface{}, *httpErr)) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logger.Named("errorHandler")

		// handle panics
		defer func() {
			err := recover()
			if err != nil {
				err := c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%v", err))
				if err != nil {
					logger.Error("failed to abort with error", "err", err)
				}
			}
		}()

		body, err := handler(c)
		if err != nil {
			if err.Type == httpErrTypeServer {
				logger.Error("internal server error", "err", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponseBody{Response: body, Err: err})
			} else {
				logger.Info("expected client error", "err", err)
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, httpResponseBody{Response: body, Err: err})
			}

			return
		}

		logger.Info("successfully handled request")
		c.JSON(http.StatusOK, httpResponseBody{Response: body})
	}
}
