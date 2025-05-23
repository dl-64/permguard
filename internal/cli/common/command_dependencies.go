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

package common

import (
	aziclients "github.com/permguard/permguard/internal/transport/clients"
	azlang "github.com/permguard/permguard/pkg/authz/languages"
	azcli "github.com/permguard/permguard/pkg/cli"
	azclients "github.com/permguard/permguard/pkg/transport/clients"
)

// cliDependencies implements the Cli dependencies.
type cliDependencies struct {
	langFactory azlang.LanguageFactory
}

// CreatePrinter creates a new printer.
func (c *cliDependencies) CreatePrinter(verbose bool, output string) (azcli.CliPrinter, error) {
	printer, err := azcli.NewCliPrinterTerminal(verbose, output)
	return printer, err
}

// CreateGrpcZAPClient creates a new gRPC client for the ZAP service.
func (c *cliDependencies) CreateGrpcZAPClient(zapTarget string) (azclients.GrpcZAPClient, error) {
	return aziclients.NewGrpcZAPClient(zapTarget)
}

// CreateGrpcPAPClient creates a new gRPC client for the PAP service.
func (c *cliDependencies) CreateGrpcPAPClient(zapTarget string) (azclients.GrpcPAPClient, error) {
	return aziclients.NewGrpcPAPClient(zapTarget)
}

// CreateGrpcPDPClient creates a new gRPC client for the PDP service.
func (c *cliDependencies) CreateGrpcPDPClient(zapTarget string) (azclients.GrpcPDPClient, error) {
	return aziclients.NewGrpcPDPClient(zapTarget)
}

// CreateGrpcPAPClient creates a new gRPC client for the PAP service.
func (c *cliDependencies) GetLanguageFactory() (azlang.LanguageFactory, error) {
	return c.langFactory, nil
}

// NewCliDependenciesProvider creates a new CliDependenciesProvider.
func NewCliDependenciesProvider(langFactory azlang.LanguageFactory) (azcli.CliDependenciesProvider, error) {
	return &cliDependencies{
		langFactory: langFactory,
	}, nil
}
