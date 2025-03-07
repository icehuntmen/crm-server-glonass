package services

import (
	"crm-glonass/api/dto"
	"crm-glonass/config"
	"crm-glonass/constants"
	"crm-glonass/pkg/logging"
	"crm-glonass/pkg/service_errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenService struct {
	Logger logging.Logger
	Cfg    *config.Config
}

type tokenDto struct {
	Id           string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) TokenInterface {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		Logger: logger,
		Cfg:    cfg,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.Cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.Cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	act := jwt.MapClaims{}

	act[constants.UserIdKey] = token.Id
	act[constants.EmailKey] = token.Email
	act[constants.MobileNumberKey] = token.MobileNumber
	act[constants.RolesKey] = token.Roles
	act[constants.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, act)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.Cfg.Jwt.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constants.UserIdKey] = token.Id
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(s.Cfg.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.Cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
