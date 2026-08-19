package caller

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ResponseModel struct {
	Successful bool   `json:"successful"`
	ErrorCode  string `json:"error_code"`
	Data       any    `json:"data"`
}

func Success(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusOK, ResponseModel{
		Successful: true,
		ErrorCode:  "",
		Data:       data,
	})
}

func BadRequest(c *fiber.Ctx, errs any) error {
	errorCode := ""
	if errs != nil {
		errorCode = fmt.Sprint(errs)
	}
	return Response(c, fiber.StatusBadRequest, ResponseModel{
		Successful: false,
		ErrorCode:  errorCode,
		Data:       nil,
	})
}

func NotFound(c *fiber.Ctx, errs any) error {
	errorCode := ""
	if errs != nil {
		errorCode = fmt.Sprint(errs)
	}
	return Response(c, fiber.StatusNotFound, ResponseModel{
		Successful: false,
		ErrorCode:  errorCode,
		Data:       nil,
	})
}

func InternalServerError(c *fiber.Ctx, err error) error {
	if err != nil {
		log.Printf("internalError: %v", err)
	}
	return Response(c, fiber.StatusInternalServerError, ResponseModel{
		Successful: false,
		ErrorCode:  "internal server error",
		Data:       nil,
	})
}

func Response(c *fiber.Ctx, httpCode int, data any) error {
	return c.Status(httpCode).JSON(data)
}
