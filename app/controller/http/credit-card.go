package httpcontroller

import (
	"darkness8129/validation-service/packages/logging"

	"github.com/gin-gonic/gin"
)

type creditCardController struct {
	logger logging.Logger
}

func newCreditCardController(opt controllerOptions) {
	logger := opt.Logger.Named("creditCardController")

	c := creditCardController{
		logger: logger,
	}

	group := opt.RouterGroup.Group("/credit-cards")
	group.POST("/validate", errorDecorator(logger, c.validate))
}

type validateRequestBody struct {
	CCNumber   string `json:"ccNumber" binding:"required"`
	CCExpMonth string `json:"ccExpMonth" binding:"required"`
	CCExpYear  string `json:"ccExpYear" binding:"required"`
}

type validateResponseBody struct {
	Valid bool `json:"valid"`
}

func (ctrl *creditCardController) validate(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("validate")

	var body validateRequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		logger.Info("invalid request body", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}
	logger.Debug("parsed request body", "body", body)

	// TODO: validation

	logger.Info("credit card successfully validated")
	return validateResponseBody{true}, nil
}
