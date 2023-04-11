package userhandler

import (
	"outgrow/dto"
	"outgrow/ent"
	"outgrow/service/masterservice"
	"outgrow/service/organizationservice"
	"outgrow/service/userservice"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService         *userservice.UserService
	OrganizationService *organizationservice.OrganizationService
	MasterService       *masterservice.MasterService
}

func NewUserHandler(
	userSvc *userservice.UserService,
	organizationSvc *organizationservice.OrganizationService,
	masterSvc *masterservice.MasterService,
) *UserHandler {
	return &UserHandler{
		UserService:         userSvc,
		OrganizationService: organizationSvc,
		MasterService:       masterSvc,
	}
}

func (h *UserHandler) getUserFromURL(c *fiber.Ctx) (*ent.User, error) {
	idString := c.Params("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := h.UserService.GetUserByID(c.UserContext(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return user, nil
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.getUserFromURL(c)
	if err != nil {
		return err
	}

	userResp := dto.GetUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
	resp := dto.SuccessResponse{
		Data: userResp,
	}
	return c.JSON(resp)
}

func (h *UserHandler) GetUserOrganizations(c *fiber.Ctx) error {
	var query dto.PaginateParam

	user, err := h.getUserFromURL(c)
	if err != nil {
		return err
	}

	query.Page = c.QueryInt("page")
	query.PerPage = c.QueryInt("per_page")

	opt := dto.GetUserOrganizationsOption{
		Paginate: &query,
	}
	organizations, paginate, err := h.UserService.GetUserOrganizations(
		c.UserContext(),
		user.ID,
		opt,
	)
	if err != nil {
		return err
	}

	respData := []dto.GetUserOrganizationsResponse{}

	for _, org := range organizations {
		data := dto.GetUserOrganizationsResponse{}
		data.BuildResponse(org)

		respData = append(respData, data)
	}

	resp := dto.SuccessResponse{
		Data:       respData,
		Pagination: &paginate,
	}

	return c.JSON(resp)
}

func (h *UserHandler) CreateOrganizationByUser(c *fiber.Ctx) error {
	var payload dto.UserCreateOrganizationPayload
	ctx := c.UserContext()
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

	user, err := h.getUserFromURL(c)
	if err != nil {
		return err
	}

	org, err := h.UserService.CreateOrganizationByUser(ctx, payload.OrganizationName, user)
	if err != nil {
		return err
	}

	// check masters data
	masterEventTypes, masterAccountTypes, err := h.MasterService.GetMasterEventAndAccountData(ctx)
	if err != nil {
		return err
	}

	// if exists then copy all of the data
	if len(masterEventTypes) > 0 && len(masterAccountTypes) > 0 {
		err = h.OrganizationService.CopyMasters(
			ctx,
			org.ID,
			masterEventTypes,
			masterAccountTypes,
		)
		if err != nil {
			return err
		}
	}

	resp := dto.SuccessResponse{
		Message: "organization successfully created",
		Data:    dto.UserCreateOrganizationResponse{OrganizationID: org.ID},
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UserHandler) UserJoinOrganization(c *fiber.Ctx) error {
	var payload dto.UserJoinOrganizationPayload

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

	user, err := h.getUserFromURL(c)
	if err != nil {
		return err
	}

	err = h.OrganizationService.JoinOrganization(
		c.UserContext(),
		payload.PublicID,
		user,
	)
	if err != nil {
		return err
	}

	resp := dto.SuccessResponse{
		Message: "successfully join the organization",
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}
