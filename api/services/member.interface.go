package services

import "crm-glonass/api/dto"

type MemberInterface interface {
	Register(createDTO *dto.MemberCreate) error
	Login(*dto.MemberAuth) (*dto.TokenDetail, error)
	Update(updateDTO *dto.MemberUpdate) (*dto.MemberResponse, error)
}
