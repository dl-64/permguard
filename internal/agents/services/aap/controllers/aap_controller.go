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

package controllers

import (
	azmodels "github.com/permguard/permguard/pkg/agents/models"
	azservices "github.com/permguard/permguard/pkg/agents/services"
	azStorage "github.com/permguard/permguard/pkg/agents/storage"
)

// AAPController is the controller for the AAP service.
type AAPController struct {
	ctx     *azservices.ServiceContext
	storage azStorage.AAPCentralStorage
}

// Setup initializes the service.
func (s AAPController) Setup() error {
	return nil
}

// NewAAPController creates a new AAP controller.
func NewAAPController(serviceContext *azservices.ServiceContext, aapCentralStorage azStorage.AAPCentralStorage) (*AAPController, error) {
	service := AAPController{
		ctx:     serviceContext,
		storage: aapCentralStorage,
	}
	return &service, nil
}

// CreateApplication creates a new application.
func (s AAPController) CreateApplication(application *azmodels.Application) (*azmodels.Application, error) {
	return s.storage.CreateApplication(application)
}

// UpdateApplication updates an application.
func (s AAPController) UpdateApplication(application *azmodels.Application) (*azmodels.Application, error) {
	return s.storage.UpdateApplication(application)
}

// DeleteApplication delete an application.
func (s AAPController) DeleteApplication(applicationID int64) (*azmodels.Application, error) {
	return s.storage.DeleteApplication(applicationID)
}

// FetchApplications returns all applications filtering by search criteria.
func (s AAPController) FetchApplications(page int32, pageSize int32, fields map[string]any) ([]azmodels.Application, error) {
	return s.storage.FetchApplications(page, pageSize, fields)
}

// CreateIdentitySource creates a new identity source.
func (s AAPController) CreateIdentitySource(identitySource *azmodels.IdentitySource) (*azmodels.IdentitySource, error) {
	return s.storage.CreateIdentitySource(identitySource)
}

// UpdateIdentitySource updates an identity source.
func (s AAPController) UpdateIdentitySource(identitySource *azmodels.IdentitySource) (*azmodels.IdentitySource, error) {
	return s.storage.UpdateIdentitySource(identitySource)
}

// DeleteIdentitySource delete an identity source.
func (s AAPController) DeleteIdentitySource(applicationID int64, identitySourceID string) (*azmodels.IdentitySource, error) {
	return s.storage.DeleteIdentitySource(applicationID, identitySourceID)
}

// FetchIdentitySources returns all identity sources filtering by search criteria.
func (s AAPController) FetchIdentitySources(page int32, pageSize int32, applicationID int64, fields map[string]any) ([]azmodels.IdentitySource, error) {
	return s.storage.FetchIdentitySources(page, pageSize, applicationID, fields)
}

// CreateIdentity creates a new identity.
func (s AAPController) CreateIdentity(identity *azmodels.Identity) (*azmodels.Identity, error) {
	return s.storage.CreateIdentity(identity)
}

// UpdateIdentity updates an identity.
func (s AAPController) UpdateIdentity(identity *azmodels.Identity) (*azmodels.Identity, error) {
	return s.storage.UpdateIdentity(identity)
}

// DeleteIdentity delete an identity.
func (s AAPController) DeleteIdentity(applicationID int64, identityID string) (*azmodels.Identity, error) {
	return s.storage.DeleteIdentity(applicationID, identityID)
}

// FetchIdentities returns all identities filtering by search criteria.
func (s AAPController) FetchIdentities(page int32, pageSize int32, applicationID int64, fields map[string]any) ([]azmodels.Identity, error) {
	return s.storage.FetchIdentities(page, pageSize, applicationID, fields)
}

// CreateTenant creates a new tenant.
func (s AAPController) CreateTenant(tenant *azmodels.Tenant) (*azmodels.Tenant, error) {
	return s.storage.CreateTenant(tenant)
}

// UpdateTenant updates a tenant.
func (s AAPController) UpdateTenant(tenant *azmodels.Tenant) (*azmodels.Tenant, error) {
	return s.storage.UpdateTenant(tenant)
}

// DeleteTenant delete a tenant.
func (s AAPController) DeleteTenant(applicationID int64, tenantID string) (*azmodels.Tenant, error) {
	return s.storage.DeleteTenant(applicationID, tenantID)
}

// FetchTenants returns all tenants filtering by search criteria.
func (s AAPController) FetchTenants(page int32, pageSize int32, applicationID int64, fields map[string]any) ([]azmodels.Tenant, error) {
	return s.storage.FetchTenants(page, pageSize, applicationID, fields)
}
