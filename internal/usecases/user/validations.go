package user

import "social/internal/domain"

func ValidatePagination(limit, offset int) error {
	if limit <= 0 || offset < 0 {
		return domain.ErrInvalidPagination
	}
	return nil
}

func ValidateUserID(id string) error {
	if id == "" {
		return domain.ErrInvalidUserID
	}
	return nil
}
