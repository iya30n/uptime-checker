package User

func (u *User) HasVerified() bool {
	return !u.EmailVerifiedAt.IsZero()
}
