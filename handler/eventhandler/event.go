package eventhandler

import (
	"outgrow/dto"
	"outgrow/enum"
	"outgrow/service/accountservice"
	"outgrow/service/eventservice"

	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	EventService   *eventservice.EventService
	AccountService *accountservice.AccountService
}

func NewEventHandler(
	eventSvc *eventservice.EventService,
	accountSvc *accountservice.AccountService,
) *EventHandler {
	return &EventHandler{
		EventService:   eventSvc,
		AccountService: accountSvc,
	}
}

func (h *EventHandler) GenerateEventSimulations(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload dto.EventSimulationPayload

	err := c.BodyParser(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.Validate(payload)
	if err != nil {
		errData := dto.Err{
			Message: "error validation",
			Data:    err,
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errData)
	}

	eventType, err := h.EventService.GetOrganizationEventTypeByID(ctx, payload.EventTypeID)
	if err != nil {
		return err
	}

	if len(eventType.Rules) == 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "event rules are not found")
	}

	amountDetails := h.EventService.CalculateTotalAmounts(eventType.Rules, payload.Amount)

	isValidAmount := amountDetails.TotalDebitPercentage == amountDetails.TotalCreditPercentage
	if !isValidAmount {
		return fiber.NewError(fiber.StatusBadRequest, "mismatch total amount")
	}

	respData := []dto.EventSimulationResponse{}
	for _, er := range eventType.Rules {
		account, err := h.AccountService.GetOrganizationAccountByID(ctx, er.AccountID, false)
		if err != nil {
			return err
		}

		var transactionAmount float64
		if er.RuleType == enum.EventRulesFlat {
			transactionAmount = er.Amount
		} else {
			if er.TransactionType == enum.TransactionTypeDebit {
				transactionAmount = (er.Amount / 100) * amountDetails.TotalDebitAmount
			} else {
				transactionAmount = (er.Amount / 100) * amountDetails.TotalCreditAmount
			}
		}

		data := dto.EventSimulationResponse{
			AccountName:     account.Name,
			TransactionType: er.TransactionType,
			Amount:          transactionAmount,
		}

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Message: "success",
		Data:    respData,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
