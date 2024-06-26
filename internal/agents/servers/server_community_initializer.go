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

package servers

import (
	"fmt"

	aziaap "github.com/permguard/permguard/internal/agents/services/aap"
	azipap "github.com/permguard/permguard/internal/agents/services/pap"
	azipdp "github.com/permguard/permguard/internal/agents/services/pdp"
	azservers "github.com/permguard/permguard/pkg/agents/servers"
	azservices "github.com/permguard/permguard/pkg/agents/services"
	azstorage "github.com/permguard/permguard/pkg/agents/storage"
	azcopier "github.com/permguard/permguard/pkg/extensions/copier"
	azipostgres "github.com/permguard/permguard/plugin/storage/postgres"
)

// CommunityServerInitializer is the community service factory initializer.
type CommunityServerInitializer struct {
	host      azservices.HostKind
	hostInfos map[azservices.HostKind]*azservices.HostInfo
	storages  []azstorage.StorageKind
	services  []azservices.ServiceKind
}

// NewCommunityServerInitializer creates a new community server initializer.
func NewCommunityServerInitializer(host azservices.HostKind) (azservers.ServerInitializer, error) {
	serviceKindsDescriptions := map[azservices.HostKind]*azservices.HostInfo{
		azservices.HostAllInOne: {Name: "AllInOne", Use: "all-in-one", Short: "Start all the services", Long: "Using this option all the services are started."},
		azservices.HostAAP:      {Name: "AAP (Account Administration Point)", Use: "pdp", Short: "Start the AAP service", Long: "Using this option the Account Administration Point (AAP) service is started."},
		azservices.HostPAP:      {Name: "PAP (Policy Administration Point)", Use: "pap", Short: "Start the PAP service", Long: "Using this option the Policy Administration Point (PAP) service is started."},
		azservices.HostPIP:      {Name: "PIP (Policy Information Point)", Use: "pip", Short: "Start the PIP service", Long: "Using this option the Policy Information Point service (PIP) is started."},
		azservices.HostPRP:      {Name: "PRP (Policy Retrieval Point)", Use: "pip", Short: "Start the PRP service", Long: "Using this option the Policy Retrieval Point (PRP) service is started."},
		azservices.HostIDP:      {Name: "IDP (Identity Provider)", Use: "idp", Short: "Start the IDP service", Long: "Using this option the Identity Provider (IDP) service is started."},
		azservices.HostPDP:      {Name: "PDP (Policy Decision Point)", Use: "pdp", Short: "Start the PDP service", Long: "Using this option the Policy Decision Point (PDP) service is started."},
	}
	hosts := []azservices.HostKind{azservices.HostAllInOne, azservices.HostAAP, azservices.HostPAP, azservices.HostPIP, azservices.HostPRP, azservices.HostIDP, azservices.HostPDP}
	storages := []azstorage.StorageKind{azstorage.StoragePostgres, azstorage.StorageBadger}
	services := []azservices.ServiceKind{azservices.ServiceAAP, azservices.ServicePAP, azservices.ServicePIP, azservices.ServicePRP, azservices.ServiceIDP, azservices.ServicePDP}

	if !host.IsValid(hosts) {
		panic(fmt.Sprintf("server: invalid server kind: %s", host))
	}
	return &CommunityServerInitializer{
		host:      host,
		hostInfos: serviceKindsDescriptions,
		storages:  storages,
		services:  host.GetServices(hosts, services),
	}, nil
}

// HasCentralStorage returns true if a central storage is required.
func (c *CommunityServerInitializer) HasCentralStorage() bool {
	return true
}

// HasProximityStorage returns true if a proximity storage is required.
func (c *CommunityServerInitializer) HasProximityStorage() bool {
	return azservices.HostAllInOne.Equal(c.host) || azservices.HostPDP.Equal(c.host)
}

// GetHost returns the service kind set as host.
func (c *CommunityServerInitializer) GetHost() azservices.HostKind {
	return c.host
}

// GetHostInfo returns the infos of the service kind set as host.
func (c *CommunityServerInitializer) GetHostInfo() *azservices.HostInfo {
	return c.hostInfos[c.host]
}

// GetStorages returns the active storage kinds.
func (c *CommunityServerInitializer) GetStorages(centralStorageEngine azstorage.StorageKind, proximityStorageEngine azstorage.StorageKind) []azstorage.StorageKind {
	storages := []azstorage.StorageKind{}
	for _, storageKind := range c.storages {
		if azstorage.StorageNone.Equal(storageKind) {
			continue
		}
		if centralStorageEngine == storageKind || proximityStorageEngine == storageKind {
			storages = append(storages, storageKind)
		}
	}
	return storages
}

// GetStoragesFactories returns the storage factories providers.
func (c *CommunityServerInitializer) GetStoragesFactories(centralStorageEngine azstorage.StorageKind, proximityStorageEngine azstorage.StorageKind) (map[azstorage.StorageKind]azstorage.StorageFactoryProvider, error) {
	factories := map[azstorage.StorageKind]azstorage.StorageFactoryProvider{}
	for _, storageKind := range c.GetStorages(centralStorageEngine, proximityStorageEngine) {
		switch storageKind {
		case azstorage.StoragePostgres:
			fFactCfg := func() (azstorage.StorageFactoryConfig, error) { return azipostgres.NewPostgresStorageFactoryConfig() }
			fFact := func(config azstorage.StorageFactoryConfig) (azstorage.StorageFactory, error) {
				return azipostgres.NewPostgresStorageFactory(config.(*azipostgres.PostgresStorageFactoryConfig))
			}
			fcty, err := azstorage.NewStorageFactoryProvider(fFactCfg, fFact)
			if err != nil {
				return nil, err
			}
			factories[storageKind] = *fcty
			continue
		case azstorage.StorageBadger:
			// TODO: To be implemented.
			continue
		}
	}
	return factories, nil
}

// GetServices returns the active service kinds.
func (c *CommunityServerInitializer) GetServices() []azservices.ServiceKind {
	return azcopier.CopySlice(c.services)
}

// GetServicesFactories returns the service factories providers.
func (c *CommunityServerInitializer) GetServicesFactories() (map[azservices.ServiceKind]azservices.ServiceFactoryProvider, error) {
	factories := map[azservices.ServiceKind]azservices.ServiceFactoryProvider{}
	for _, serviceKind := range c.services {
		switch serviceKind {
		case azservices.ServiceAAP:
			fFactCfg := func() (azservices.ServiceFactoryConfig, error) { return aziaap.NewAAPServiceFactoryConfig() }
			fFact := func(config azservices.ServiceFactoryConfig) (azservices.ServiceFactory, error) {
				return aziaap.NewAAPServiceFactory(config.(*aziaap.AAPServiceFactoryConfig))
			}
			fcty, err := azservices.NewServiceFactoryProvider(fFactCfg, fFact)
			if err != nil {
				return nil, err
			}
			factories[serviceKind] = *fcty
			continue
		case azservices.ServicePAP:
			fFactCfg := func() (azservices.ServiceFactoryConfig, error) { return azipap.NewPAPServiceFactoryConfig() }
			fFact := func(config azservices.ServiceFactoryConfig) (azservices.ServiceFactory, error) {
				return azipap.NewPAPServiceFactory(config.(*azipap.PAPServiceFactoryConfig))
			}
			fcty, err := azservices.NewServiceFactoryProvider(fFactCfg, fFact)
			if err != nil {
				return nil, err
			}
			factories[serviceKind] = *fcty
			continue
		case azservices.ServicePIP:
			continue
		case azservices.ServicePRP:
			continue
		case azservices.ServiceIDP:
			continue
		case azservices.ServicePDP:
			fFactCfg := func() (azservices.ServiceFactoryConfig, error) { return azipdp.NewPDPServiceFactoryConfig() }
			fFact := func(config azservices.ServiceFactoryConfig) (azservices.ServiceFactory, error) {
				return azipdp.NewPDPServiceFactory(config.(*azipdp.PDPServiceFactoryConfig))
			}
			fcty, err := azservices.NewServiceFactoryProvider(fFactCfg, fFact)
			if err != nil {
				return nil, err
			}
			factories[serviceKind] = *fcty
		}
	}
	return factories, nil
}
