package components

import (
	"crm-glonass/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// OTP
	service_errors.OtpExists:         409,
	service_errors.OtpUsed:           409,
	service_errors.OtpNotValid:       400,
	service_errors.UsernameExists:    409,
	service_errors.EmailExists:       409,
	service_errors.UsernameNotExists: 404,
	service_errors.PermissionDenied:  403,
	service_errors.RecordNotFound:    404,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
