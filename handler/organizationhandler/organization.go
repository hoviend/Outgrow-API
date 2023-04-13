package organizationhandler

import (
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/enum"
	"outgrow/service/accountservice"
	"outgrow/service/eventservice"
	"outgrow/service/organizationservice"
	"outgrow/service/transactionservice"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	OrganizationService *organizationservice.OrganizationService
	AccountService      *accountservice.AccountService
	TransactionService  *transactionservice.TransactionService
	EventService        *eventservice.EventService
}

func NewOrganizationHandler(
	organizationSvc *organizationservice.OrganizationService,
	accountSvc *accountservice.AccountService,
	transactionSvc *transactionservice.TransactionService,
	eventSvc *eventservice.EventService,
) *OrganizationHandler {
	return &OrganizationHandler{
		OrganizationService: organizationSvc,
		AccountService:      accountSvc,
		TransactionService:  transactionSvc,
		EventService:        eventSvc,
	}
}

func (h *OrganizationHandler) getOrganizationFromURL(c *fiber.Ctx, opt *dto.GetOrganizationOption) (*ent.Organization, error) {
	idString := c.Params("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if opt == nil {
		opt = &dto.GetOrganizationOption{
			ID: id,
		}
	} else {
		opt.ID = id
	}

	organization, err := h.OrganizationService.GetOrganizationByParam(c.UserContext(), *opt)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return organization, nil
}

func (h *OrganizationHandler) GetOrganization(c *fiber.Ctx) error {
	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	data := dto.GetUserOrganizationsResponse{}
	data.BuildResponse(organization)

	resp := dto.SuccessResponse{
		Data: data,
	}

	return c.JSON(resp)
}

func (h *OrganizationHandler) CreateOrganizationAccount(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload dto.CreateOrganizationAccountPayload
	err := c.BodyParser(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.Validate(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	isCategoryValid, err := h.AccountService.CheckOrganizationAccountCategory(ctx, organization.ID, payload.AccountCategoryID)
	if err != nil {
		return err
	}

	if !isCategoryValid {
		return fiber.NewError(fiber.StatusNotFound, "invalid account category")
	}

	newAccount, err := h.AccountService.CreateOrganizationAccount(ctx, payload)
	if err != nil {
		return err
	}

	resp := dto.SuccessResponse{
		Message: "account has successfully created",
		Data: dto.DefaultCreateIDResponse{
			ID: newAccount.ID,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *OrganizationHandler) GetOrganizationAccountCategories(c *fiber.Ctx) error {
	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	var query dto.GetAccountsParam
	query.Page = c.QueryInt("page")
	query.PerPage = c.QueryInt("per_page")
	query.Name = c.Query("name")

	opt := dto.GetAccountCategoriesOption{
		Paginate:           &query.PaginateParam,
		OrganizationID:     organization.ID,
		CategoryNameFilter: query.Name,
	}

	categories, paginate, err := h.AccountService.GetAccountCategories(ctx, opt)
	if err != nil {
		return err
	}

	var respData []dto.GetAccountCategoryResponse
	for _, cat := range categories {
		var data dto.GetAccountCategoryResponse

		data.ID = cat.ID
		data.Name = cat.Name
		data.AccountType.ID = cat.Edges.Type.ID
		data.AccountType.Name = cat.Edges.Type.Name

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Data:       respData,
		Pagination: &paginate,
	}

	return c.JSON(resp)

}

func (h *OrganizationHandler) GetOrganizationChartOfAccounts(c *fiber.Ctx) error {
	var query dto.GetAccountsParam
	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	ctx := c.UserContext()

	query.Page = c.QueryInt("page")
	query.PerPage = c.QueryInt("per_page")
	query.Name = c.Query("name")

	opt := dto.GetOrganizationAccountsOption{
		Paginate:          &query.PaginateParam,
		OrganizationID:    organization.ID,
		AccountNameFilter: query.Name,
	}

	accounts, paginate, err := h.AccountService.GetOrganizationAccounts(
		ctx,
		opt,
	)
	if err != nil {
		return err
	}

	respData := []dto.GetChartofAccountsResponse{}

	for _, acc := range accounts {
		data := dto.GetChartofAccountsResponse{}
		data.BuildResponse(acc)

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Data:       respData,
		Pagination: &paginate,
	}

	return c.JSON(resp)
}

func (h *OrganizationHandler) GetOrganizationEventTypes(c *fiber.Ctx) error {
	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	eventTypes, err := h.OrganizationService.GetOrganizationEventTypes(ctx, organization.ID)
	if err != nil {
		return err
	}

	var respData []dto.OrganizationEventTypeSelectionResponse
	for _, et := range eventTypes {
		var data dto.OrganizationEventTypeSelectionResponse

		data.ID = et.ID
		data.Name = et.Name

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Data: respData,
	}

	return c.JSON(resp)
}

// Create custom event type for the organization.
func (h *OrganizationHandler) CreateOrganizationEventType(c *fiber.Ctx) error {
	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	var payload dto.CreateOrganizationEventTypePayload
	err = c.BodyParser(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.Validate(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	err = h.EventService.ValidateEventRule(ctx, payload.Rules)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payload.OrganizationID = organization.ID
	eventType, err := h.OrganizationService.CreateOrganizationEventType(ctx, payload)
	if err != nil {
		return err
	}

	resp := dto.SuccessResponse{
		Message: "custom event type has successfully created",
		Data: dto.DefaultCreateIDResponse{
			ID: eventType.ID,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *OrganizationHandler) CreateOrganizationEvent(c *fiber.Ctx) error {
	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	var payload dto.CreateOrganizationEventPayload
	err = c.BodyParser(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.Validate(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// validate & get event type
	eventType, err := h.OrganizationService.GetOrganizationEventTypeByID(
		ctx,
		payload.EventTypeID,
		organization.ID,
	)
	if err != nil || eventType == nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// create event
	payload.OrganizationID = organization.ID
	event, err := h.OrganizationService.CreateEvent(ctx, payload)
	if err != nil {
		return err
	}

	amountDetails := h.EventService.CalculateTotalAmounts(eventType.Rules, payload.Amount)

	transactions := []*ent.Transaction{}
	for _, er := range eventType.Rules {
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

		transaction := &ent.Transaction{
			AccountID:       er.AccountID,
			Amount:          transactionAmount,
			TransactionDate: time.Now(),
			TransactionType: er.TransactionType,
			EventID:         event.ID,
			Notes:           payload.Notes,
		}

		transactions = append(transactions, transaction)
	}

	err = h.TransactionService.CreateTransactions(ctx, transactions)
	if err != nil {
		return err
	}

	resp := dto.SuccessResponse{
		Message: "event has successfully created",
		Data: dto.DefaultCreateUUIDResponse{
			ID: event.ID,
		},
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *OrganizationHandler) GetOrganizationTransactions(c *fiber.Ctx) error {
	var query dto.PaginateParam

	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	query.Page = c.QueryInt("page")
	query.PerPage = c.QueryInt("per_page")

	opt := dto.GetTransactionsOption{
		Paginate:       &query,
		OrganizationID: organization.ID,
		JoinAccount:    true,
		JoinEvent:      true,
	}

	transactions, paginate, err := h.TransactionService.GetTransactions(ctx, opt)
	if err != nil {
		return err
	}

	respData := []dto.GetOrganizationTransactionResponse{}

	for _, trx := range transactions {
		data := dto.GetOrganizationTransactionResponse{}
		data.BuildResponse(trx)

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Data:       respData,
		Pagination: &paginate,
	}

	return c.JSON(resp)
}
