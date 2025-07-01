package handlers

import (
	"CalculatorAppBackend/internal/calculationService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(service calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: service}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculation(c echo.Context) error {
	var request calculationService.CalculationRequest
	if err := c.Bind(&request); err != nil { // расшифровываем сообщение
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(request.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc) // 204 status

}

func (h *CalculationHandler) PatchCalculation(c echo.Context) error {
	id := c.Param("id")
	var request calculationService.CalculationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedCalculation, err := h.service.UpdateCalculation(id, request.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, updatedCalculation)
}

func (h *CalculationHandler) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent) // 204 status

}
