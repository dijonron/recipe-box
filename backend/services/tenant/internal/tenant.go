package tenant

import (
	"errors"
)

const (
	INVALID_CREDENTIALS string = "invalid credentials"
	FAILED_TO_GEN_JWT   string = "failed to generate jwt"
)

var (
	errInvalidCredentials = errors.New(INVALID_CREDENTIALS)
	errFailedToGenJwt     = errors.New(FAILED_TO_GEN_JWT)
)

type tenantManager struct {
}

var _ TenantManager = (*tenantManager)(nil)

func NewTenantManager() TenantManager {
	return &tenantManager{}
}
