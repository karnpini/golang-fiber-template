package interceptors

import (
	e "github.com/corp-ais/golang-fiber-template/pkg/errors"
	"github.com/corp-ais/golang-fiber-template/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ResponseStatus struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
	ResponseData   interface{}    `json:"responseData"`
}

func SendJSONResponse(ctx *fiber.Ctx, statusCode int, responseCode string, data interface{}) {
	clientName, _ := ctx.Locals("clientName").(string)

	logger.Info("Sending JSON response",
		zap.String("method", ctx.Method()),
		zap.String("url", ctx.OriginalURL()),
		zap.Int("statusCode", statusCode),
		zap.String("clientName", clientName),
	)
	response := Response{
		ResponseStatus: ResponseStatus{
			ResponseCode:    responseCode,
			ResponseMessage: e.GetMsg(responseCode),
		},
		ResponseData: data,
	}
	ctx.Status(statusCode).JSON(response)
}
