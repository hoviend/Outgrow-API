package enum

// currently still unused yet
type UserStatus string

const (
	UserStatusAwaitingVerification UserStatus = "AWAITING_VERIFICATION"
	UserStatusVerified             UserStatus = "VERIFIED"
)

func (UserStatus) Values() (kinds []string) {
	for _, s := range []UserStatus{UserStatusAwaitingVerification, UserStatusVerified} {
		kinds = append(kinds, string(s))
	}

	return
}
