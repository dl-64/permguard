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

package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	aziclients "github.com/permguard/permguard/internal/agents/clients"
	azmodels "github.com/permguard/permguard/pkg/agents/models"
	azconfigs "github.com/permguard/permguard/pkg/configs"
)

const (
	commandNameForTenantsDelete = "tenants.delete"
)

// runECommandForDeleteTenant runs the command for creating a tenant.
func runECommandForDeleteTenant(cmd *cobra.Command, v *viper.Viper) error {
	ctx, printer, err := createContextAndPrinter(cmd, v)
	if err != nil {
		color.Red("invalid inputs")
		return ErrCommandSilent
	}
	aapTarget := ctx.GetAAPTarget()
	client, err := aziclients.NewGrpcAAPClient(aapTarget)
	if err != nil {
		printer.Error(fmt.Errorf("invalid aap target %s", aapTarget))
		return ErrCommandSilent
	}
	accountID := v.GetInt64(azconfigs.FlagName(commandNameForTenant, flagCommonAccountID))
	tenantID := v.GetString(azconfigs.FlagName(commandNameForTenantsDelete, flagTenantID))
	tenant, err := client.DeleteTenant(accountID, tenantID)
	if err != nil {
		printer.Error(err)
		return ErrCommandSilent
	}
	output := map[string]any{}
	if ctx.IsTerminalOutput() {
		tenantID := tenant.TenantID
		tenantName := tenant.Name
		output[tenantID] = tenantName
	} else if ctx.IsJSONOutput() {
		output["tenant"] = []*azmodels.Tenant{tenant}
	}
	printer.Print(output)
	return nil
}

// createCommandForTenantDelete creates a command for managing tenantdelete.
func createCommandForTenantDelete(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete a tenant",
		Long: `This command delete a tenant.

Examples:
  # delete a tenant with id 19159d69-e902-418e-966a-148c4d5169a4 and account 301990992055
  permguard authn tenants delete --account 301990992055 --tenantid 19159d69-e902-418e-966a-148c4d5169a4
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForDeleteTenant(cmd, v)
		},
	}
	command.Flags().String(flagTenantID, "", "tenant id")
	v.BindPFlag(azconfigs.FlagName(commandNameForTenantsDelete, flagTenantID), command.Flags().Lookup(flagTenantID))
	return command
}
