package httpcontroller

import (
	"darkness8129/validation-service/app/service"
	"darkness8129/validation-service/packages/errs"
	"darkness8129/validation-service/packages/logging"

	"github.com/gin-gonic/gin"
)

type creditCardController struct {
	logger   logging.Logger
	services service.Services
}

func newCreditCardController(opt controllerOptions) {
	logger := opt.Logger.Named("creditCardController")

	c := creditCardController{
		logger:   logger,
		services: opt.Services,
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
		return validateResponseBody{false}, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}
	logger.Debug("parsed request body", "body", body)

	err = ctrl.services.CreditCard.Validate(c, service.ValidateOpt{
		CCNumber:   body.CCNumber,
		CCExpMonth: body.CCExpMonth,
		CCExpYear:  body.CCExpYear,
	})
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return validateResponseBody{false}, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to validate credit card", "err", err)
		return validateResponseBody{false}, &httpErr{Type: httpErrTypeServer, Message: "failed to validate credit card"}
	}

	logger.Info("credit card successfully validated")
	return validateResponseBody{true}, nil
}
