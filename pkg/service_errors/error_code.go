package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"
	// OTP
	OtpExists   = "Otp exists"
	OtpUsed     = "Otp used"
	OtpNotValid = "Otp invalid"

	// User
	EmailExists       = "Email exists"
	EmailNotExists    = "Member not exists"
	UsernameExists    = "Username exists"
	UsernameNotExists = "Username not exists"
	PermissionDenied  = "Permission denied"

	// DB
	RecordNotFound = "record not found"

	// Role
	RoleExists = "Role with such name already exists"
)
