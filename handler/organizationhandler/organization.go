package organizationhandler

import (
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/service/accountservice"
	"outgrow/service/organizationservice"
	"outgrow/service/transactionservice"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	OrganizationService *organizationservice.OrganizationService
	AccountService      *accountservice.AccountService
	TransactionService  *transactionservice.TransactionService
}

func NewOrganizationHandler(
	organizationSvc *organizationservice.OrganizationService,
	accountSvc *accountservice.AccountService,
	transactionSvc *transactionservice.TransactionService,
) *OrganizationHandler {
	return &OrganizationHandler{
		OrganizationService: organizationSvc,
		AccountService:      accountSvc,
		TransactionService:  transactionSvc,
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

func (h *OrganizationHandler) GetOrganizationAccountsSelection(c *fiber.Ctx) error {
	ctx := c.UserContext()

	organization, err := h.getOrganizationFromURL(c, nil)
	if err != nil {
		return err
	}

	accTypes, err := h.AccountService.GetOrganizationAccountTypes(ctx, organization.ID, dto.GetOrganizationAccountTypesOption{WithCategoriesAndAccounts: true})
	if err != nil {
		return err
	}

	var resp []dto.GetAccountSelectionResponse
	for _, at := range accTypes {
		var data dto.GetAccountSelectionResponse

		data.AccountType.ID = at.ID
		data.AccountType.Name = at.Name

		for _, cat := range at.Edges.Categories {
			d := dto.DefaultAccountResponse{
				ID:   cat.ID,
				Name: cat.Name,
			}
			data.Categories = append(data.Categories, d)
		}

		resp = append(resp, data)
	}

	return c.JSON(resp)

}

func (h *OrganizationHandler) GetOrganizationChartOfAccounts(c *fiber.Ctx) error {
	var query dto.GetChartofAccountsParam
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

	var resp []dto.OrganizationEventTypeSelectionResponse
	for _, et := range eventTypes {
		var data dto.OrganizationEventTypeSelectionResponse

		data.ID = et.ID
		data.Name = et.Name

		resp = append(resp, data)
	}

	return c.JSON(resp)
}

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

// func (h *OrganizationHandler) CreateOrganizationEvent(c *fiber.Ctx) error {
// 	ctx := c.UserContext()

// 	organization, err := h.getOrganizationFromURL(c, nil)
// 	if err != nil {
// 		return err
// 	}

// 	var payload dto.CreateOrganizationEventPayload
// 	err = c.BodyParser(&payload)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	err = dto.Validate(payload)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	// validate & get event type
// 	eventType, err := h.OrganizationService.GetOrganizationEventTypeByID(
// 		ctx,
// 		payload.EventTypeID,
// 		organization.ID,
// 	)
// 	if err != nil || eventType == nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	// create event
// 	payload.OrganizationID = organization.ID
// 	_, err = h.OrganizationService.CreateEvent(ctx, payload)
// 	if err != nil {
// 		return err
// 	}

// 	// create transactions based on event rules
// 	eventRules := eventType.Rules
// }

// func ParseEventRules(eventRules []schema.EventRules, eventID uuid.UUID) (transactions []*ent.Transaction, err error) {
// 	// validation
// 	var (
// 		totalDebitPercentage,
// 		totalCreditPercentage float64
// 	)

// 	for _, er := range eventRules {
// 		var transactionAmount float64

// 		if er.RuleType == enum.EventRulesPercentage {
// 			if er.TransactionType == enum.TransactionTypeCredit {
// 				totalCreditPercentage += er.Amount
// 			} else {
// 				totalDebitPercentage += er.Amount
// 			}
// 		} else {
// 			transactionAmount
// 		}

// 		trx := &ent.Transaction{
// 			EventID:         eventID,
// 			AccountID:       er.AccountID,
// 			TransactionType: er.TransactionType,
// 			TransactionDate: time.Now().UTC(),
// 			Amount: ,
// 		}

// 	}
// 	// + total debit and total credit must be equal
// 	isValidAmount := totalDebitPercentage == totalCreditPercentage
// 	if !isValidAmount {
// 		return nil, fmt.Errorf("mismatch total amount")
// 	}
// 	// kalau ada flat, amountnya nya dikurang dulu
// }

// func (h *OrganizationHandler) GetOrganizationTransactions(c *fiber.Ctx) error {
// 	var query dto.PaginateParam

// 	ctx := c.UserContext()

// 	organization, err := h.getOrganizationFromURL(c, nil)
// 	if err != nil {
// 		return err
// 	}

// 	query.Page = c.QueryInt("page")
// 	query.PerPage = c.QueryInt("per_page")

// 	opt := dto.GetTransactionsOption{
// 		Paginate:       &query,
// 		OrganizationID: organization.ID,
// 	}

// 	transactions, paginate, err := h.TransactionService.GetTransactions()

// }
