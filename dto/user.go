package dto

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"outgrow/ent"
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type GetUserOrganizationsParam struct {
	PaginateParam
	OrganizationID string
}

type GetUserOrganizationsResponse struct {
	Name     string `json:"name"`
	PublicID string `json:"public_id"`
}

type UserCreateOrganizationPayload struct {
	OrganizationName string `json:"organization_name" validate:"required"`
}

type UserCreateOrganizationResponse struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

type GetUserOrganizationsOption struct {
	Paginate *PaginateParam
}

type UserJoinOrganizationPayload struct {
	PublicID string `json:"public_id" validate:"required"`
}

func (r *GetUserOrganizationsResponse) BuildResponse(o *ent.Organization) {
	r.Name = o.Name
	r.PublicID = o.PublicID
}

// generate unique 8-letters string
func GeneratePublicID() string {
	source := rand.NewSource(time.Now().UnixNano())

	random := rand.New(source)

	randomBytes := make([]byte, 32)
	random.Read(randomBytes)

	hash := sha256.Sum256(randomBytes)

	res := hex.EncodeToString(hash[:])[:8]

	return res
}
