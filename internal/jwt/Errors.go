package jwt

type TokenIsValidError struct {}

func (*TokenIsValidError) Error() string {
	return "Your token is still valid"
}