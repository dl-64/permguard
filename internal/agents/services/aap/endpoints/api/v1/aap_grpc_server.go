// Copyright 2024 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"

	azmodels "github.com/permguard/permguard/pkg/agents/models"
	azservices "github.com/permguard/permguard/pkg/agents/services"
)

// AAPService is the service for the AAP.
type AAPService interface {
	Setup() error

	CreateAccount(account *azmodels.Account) (*azmodels.Account, error)
	UpdateAccount(account *azmodels.Account) (*azmodels.Account, error)
	DeleteAccount(accountID int64) (*azmodels.Account, error)
	GetAllAccounts(filter map[string]any) ([]azmodels.Account, error)

	CreateIdentitySource(identitySource *azmodels.IdentitySource) (*azmodels.IdentitySource, error)
	UpdateIdentitySource(identitySource *azmodels.IdentitySource) (*azmodels.IdentitySource, error)
	DeleteIdentitySource(accountID int64, identitySourceID string) (*azmodels.IdentitySource, error)
	GetAllIdentitySources(accountID int64, fields map[string]any) ([]azmodels.IdentitySource, error)

	CreateIdentity(identity *azmodels.Identity) (*azmodels.Identity, error)
	UpdateIdentity(identity *azmodels.Identity) (*azmodels.Identity, error)
	DeleteIdentity(accountID int64, identityID string) (*azmodels.Identity, error)
	GetAllIdentities(accountID int64, fields map[string]any) ([]azmodels.Identity, error)

	CreateTenant(tenant *azmodels.Tenant) (*azmodels.Tenant, error)
	UpdateTenant(tenant *azmodels.Tenant) (*azmodels.Tenant, error)
	DeleteTenant(accountID int64, tenantID string) (*azmodels.Tenant, error)
	GetAllTenants(accountID int64, fields map[string]any) ([]azmodels.Tenant, error)
}

// NewV1AAPServer creates a new AAP server.
func NewV1AAPServer(endpointCtx *azservices.EndpointContext, Service AAPService) (*V1AAPServer, error) {
	return &V1AAPServer{
		ctx:     endpointCtx,
		service: Service,
	}, nil
}

// V1AAPServer is the gRPC server for the AAP.
type V1AAPServer struct {
	UnimplementedV1AAPServiceServer
	ctx     *azservices.EndpointContext
	service AAPService
}

// CreateAccount creates a new account.
func (s V1AAPServer) CreateAccount(ctx context.Context, accountRequest *AccountCreateRequest) (*AccountResponse, error) {
	account, err := s.service.CreateAccount(&azmodels.Account{Name: accountRequest.Name})
	if err != nil {
		return nil, err
	}
	return MapAgentAccountToGrpcAccountResponse(account)
}

// UpdateAccount updates an account.
func (s V1AAPServer) UpdateAccount(ctx context.Context, accountRequest *AccountUpdateRequest) (*AccountResponse, error) {
	account, err := s.service.UpdateAccount((&azmodels.Account{AccountID: accountRequest.AccountID, Name: accountRequest.Name}))
	if err != nil {
		return nil, err
	}
	return MapAgentAccountToGrpcAccountResponse(account)
}

// DeleteAccount deletes an account.
func (s V1AAPServer) DeleteAccount(ctx context.Context, accountRequest *AccountDeleteRequest) (*AccountResponse, error) {
	account, err := s.service.DeleteAccount(accountRequest.AccountID)
	if err != nil {
		return nil, err
	}
	return MapAgentAccountToGrpcAccountResponse(account)
}

// GetAllAccounts returns all the accounts.
func (s V1AAPServer) GetAllAccounts(ctx context.Context, accountRequest *AccountGetRequest) (*AccountListResponse, error) {
	fields := map[string]any{}
	if accountRequest.AccountID != nil {
		fields[azmodels.FieldAccountAccountID] = *accountRequest.AccountID
	}
	if accountRequest.Name != nil {
		fields[azmodels.FieldAccountName] = *accountRequest.Name

	}
	accounts, err := s.service.GetAllAccounts(fields)
	if err != nil {
		return nil, err
	}
	accountList := &AccountListResponse{
		Accounts: make([]*AccountResponse, len(accounts)),
	}
	for i, account := range accounts {
		cvtedAccount, err := MapAgentAccountToGrpcAccountResponse(&account)
		if err != nil {
			return nil, err
		}
		accountList.Accounts[i] = cvtedAccount
	}
	return accountList, nil
}

// CreateIdentitySource creates a new identity source.
func (s V1AAPServer) CreateIdentitySource(ctx context.Context, identitySourceRequest *IdentitySourceCreateRequest) (*IdentitySourceResponse, error) {
	identitySource, err := s.service.CreateIdentitySource(&azmodels.IdentitySource{AccountID: identitySourceRequest.AccountID, Name: identitySourceRequest.Name})
	if err != nil {
		return nil, err
	}
	return MapAgentIdentitySourceToGrpcIdentitySourceResponse(identitySource)
}

// UpdateIdentitySource updates an identity source.
func (s V1AAPServer) UpdateIdentitySource(ctx context.Context, identitySourceRequest *IdentitySourceUpdateRequest) (*IdentitySourceResponse, error) {
	identitySource, err := s.service.UpdateIdentitySource((&azmodels.IdentitySource{IdentitySourceID: identitySourceRequest.IdentitySourceID, AccountID: identitySourceRequest.AccountID, Name: identitySourceRequest.Name}))
	if err != nil {
		return nil, err
	}
	return MapAgentIdentitySourceToGrpcIdentitySourceResponse(identitySource)
}

// DeleteIdentitySource deletes an identity source.
func (s V1AAPServer) DeleteIdentitySource(ctx context.Context, identitySourceRequest *IdentitySourceDeleteRequest) (*IdentitySourceResponse, error) {
	identitySource, err := s.service.DeleteIdentitySource(identitySourceRequest.AccountID, identitySourceRequest.IdentitySourceID)
	if err != nil {
		return nil, err
	}
	return MapAgentIdentitySourceToGrpcIdentitySourceResponse(identitySource)
}

// GetAllIdentitySources returns all the identity sources.
func (s V1AAPServer) GetAllIdentitySources(ctx context.Context, identitySourceRequest *IdentitySourceGetRequest) (*IdentitySourceListResponse, error) {
	fields := map[string]any{}
	fields[azmodels.FieldIdentitySourceAccountID] = identitySourceRequest.AccountID
	if identitySourceRequest.Name != nil {
		fields[azmodels.FieldIdentitySourceName] = *identitySourceRequest.Name
	}
	if identitySourceRequest.IdentitySourceID != nil {
		fields[azmodels.FieldIdentitySourceIdentitySourceID] = *identitySourceRequest.IdentitySourceID
	}
	identitySources, err := s.service.GetAllIdentitySources(identitySourceRequest.AccountID, fields)
	if err != nil {
		return nil, err
	}
	identitySourceList := &IdentitySourceListResponse{
		IdentitySources: make([]*IdentitySourceResponse, len(identitySources)),
	}
	for i, identitySource := range identitySources {
		cvtedIdentitySource, err := MapAgentIdentitySourceToGrpcIdentitySourceResponse(&identitySource)
		if err != nil {
			return nil, err
		}
		identitySourceList.IdentitySources[i] = cvtedIdentitySource
	}
	return identitySourceList, nil
}

// CreateIdentity creates a new identity.
func (s V1AAPServer) CreateIdentity(ctx context.Context, identityRequest *IdentityCreateRequest) (*IdentityResponse, error) {
	identity, err := s.service.CreateIdentity(&azmodels.Identity{AccountID: identityRequest.AccountID, IdentitySourceID: identityRequest.IdentitySourceID, Kind: identityRequest.Kind, Name: identityRequest.Name})
	if err != nil {
		return nil, err
	}
	return MapAgentIdentityToGrpcIdentityResponse(identity)
}

// UpdateIdentity updates an identity.
func (s V1AAPServer) UpdateIdentity(ctx context.Context, identityRequest *IdentityUpdateRequest) (*IdentityResponse, error) {
	identity, err := s.service.UpdateIdentity((&azmodels.Identity{IdentityID: identityRequest.IdentityID, AccountID: identityRequest.AccountID, Kind: identityRequest.Kind, Name: identityRequest.Name}))
	if err != nil {
		return nil, err
	}
	return MapAgentIdentityToGrpcIdentityResponse(identity)
}

// DeleteIdentity deletes an identity.
func (s V1AAPServer) DeleteIdentity(ctx context.Context, identityRequest *IdentityDeleteRequest) (*IdentityResponse, error) {
	identity, err := s.service.DeleteIdentity(identityRequest.AccountID, identityRequest.IdentityID)
	if err != nil {
		return nil, err
	}
	return MapAgentIdentityToGrpcIdentityResponse(identity)
}

// GetAllIdentities returns all the identities.
func (s V1AAPServer) GetAllIdentities(ctx context.Context, identityRequest *IdentityGetRequest) (*IdentityListResponse, error) {
	fields := map[string]any{}
	fields[azmodels.FieldIdentityAccountID] = identityRequest.AccountID
	if identityRequest.IdentitySourceID != nil {
		fields[azmodels.FieldIdentityIdentitySourceID] = *identityRequest.IdentitySourceID
	}
	if identityRequest.IdentityID != nil {
		fields[azmodels.FieldIdentityIdentityID] = *identityRequest.IdentityID
	}
	if identityRequest.Kind != nil {
		fields[azmodels.FieldIdentityKind] = *identityRequest.Kind
	}
	if identityRequest.Name != nil {
		fields[azmodels.FieldIdentityName] = *identityRequest.Name
	}
	identities, err := s.service.GetAllIdentities(identityRequest.AccountID, fields)
	if err != nil {
		return nil, err
	}
	identityList := &IdentityListResponse{
		Identities: make([]*IdentityResponse, len(identities)),
	}
	for i, identity := range identities {
		cvtedIdentity, err := MapAgentIdentityToGrpcIdentityResponse(&identity)
		if err != nil {
			return nil, err
		}
		identityList.Identities[i] = cvtedIdentity
	}
	return identityList, nil
}

// CreateTenant creates a new tenant.
func (s V1AAPServer) CreateTenant(ctx context.Context, tenantRequest *TenantCreateRequest) (*TenantResponse, error) {
	tenant, err := s.service.CreateTenant(&azmodels.Tenant{AccountID: tenantRequest.AccountID, Name: tenantRequest.Name})
	if err != nil {
		return nil, err
	}
	return MapAgentTenantToGrpcTenantResponse(tenant)
}

// UpdateTenant updates an tenant.
func (s V1AAPServer) UpdateTenant(ctx context.Context, tenantRequest *TenantUpdateRequest) (*TenantResponse, error) {
	tenant, err := s.service.UpdateTenant((&azmodels.Tenant{TenantID: tenantRequest.TenantID, AccountID: tenantRequest.AccountID, Name: tenantRequest.Name}))
	if err != nil {
		return nil, err
	}
	return MapAgentTenantToGrpcTenantResponse(tenant)
}

// DeleteTenant deletes an tenant.
func (s V1AAPServer) DeleteTenant(ctx context.Context, tenantRequest *TenantDeleteRequest) (*TenantResponse, error) {
	tenant, err := s.service.DeleteTenant(tenantRequest.AccountID, tenantRequest.TenantID)
	if err != nil {
		return nil, err
	}
	return MapAgentTenantToGrpcTenantResponse(tenant)
}

// GetAllTenants returns all the tenants.
func (s V1AAPServer) GetAllTenants(ctx context.Context, tenantRequest *TenantGetRequest) (*TenantListResponse, error) {
	fields := map[string]any{}
	fields[azmodels.FieldTenantAccountID] = tenantRequest.AccountID
	if tenantRequest.Name != nil {
		fields[azmodels.FieldTenantName] = *tenantRequest.Name
	}
	if tenantRequest.TenantID != nil {
		fields[azmodels.FieldTenantTenantID] = *tenantRequest.TenantID
	}
	tenants, err := s.service.GetAllTenants(tenantRequest.AccountID, fields)
	if err != nil {
		return nil, err
	}
	tenantList := &TenantListResponse{
		Tenants: make([]*TenantResponse, len(tenants)),
	}
	for i, tenant := range tenants {
		cvtedTenant, err := MapAgentTenantToGrpcTenantResponse(&tenant)
		if err != nil {
			return nil, err
		}
		tenantList.Tenants[i] = cvtedTenant
	}
	return tenantList, nil
}
