package fiber

import (
	"backend/type/response"
	"errors"
	"strings"

	uu "github.com/bsthun/goutils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Case of *fiber.Error
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(response.ErrorResponse{
			Success: false,
			Code:    "FIBER_ERROR",
			Message: fiberErr.Error(),
		})
	}

	// Case of ErrorInstance
	var respErr *uu.ErrorInstance
	if errors.As(err, &respErr) {
		block := respErr.Errors[len(respErr.Errors)-1]
		if block.Err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
				Success: false,
				Code:    block.Code,
				Message: block.Message,
				Error:   block.Err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Message: block.Message,
			Code:    block.Code,
		})
	}

	// Case of validator.ValidationErrors
	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		var lists []string
		for _, err := range valErr {
			lists = append(lists, err.Field()+" ("+err.Tag()+")")
		}

		message := strings.Join(lists[:], ", ")

		return c.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "VALIDATION failed on field " + message,
			Error:   valErr.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(&response.ErrorResponse{
		Success: false,
		Code:    "SERVER_SIDE_ERROR",
		Message: "Unknown server side error",
		Error:   err.Error(),
	})
}
