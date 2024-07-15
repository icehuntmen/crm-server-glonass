package constants

const (
	// User
	AdminRoleName      string = "admin"
	DefaultRoleName    string = "default"
	DefaultUserName    string = "admin"
	RedisOtpDefaultKey string = "otp"

	// Claims
	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserId"
	FirstNameKey           string = "FirstName"
	LastNameKey            string = "LastName"
	UserNameKey            string = "UserName"
	EmailKey               string = "Email"
	MobileNumberKey        string = "MobileNumber"
	RolesKey               string = "Roles"
	ExpireTimeKey          string = "ExpireTime"
)

const (
	// Roles
	RoleAdminName   string = "admin"
	RoleManagerName string = "manager"
	RoleAgentName   string = "agent"
	RoleUserName    string = "user"
)
