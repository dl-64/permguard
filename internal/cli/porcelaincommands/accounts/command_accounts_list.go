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

package accounts

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	aziclicommon "github.com/permguard/permguard/internal/cli/common"
	azcli "github.com/permguard/permguard/pkg/cli"
	azoptions "github.com/permguard/permguard/pkg/cli/options"
)

const (
	// commandNameForAccountsList is the command name for accounts list.
	commandNameForAccountsList = "accounts.list"
)

// runECommandForListAccounts runs the command for creating an account.
func runECommandForListAccounts(deps azcli.CliDependenciesProvider, cmd *cobra.Command, v *viper.Viper) error {
	ctx, printer, err := aziclicommon.CreateContextAndPrinter(deps, cmd, v)
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
		return aziclicommon.ErrCommandSilent
	}
	aapTarget := ctx.GetAAPTarget()
	client, err := deps.CreateGrpcAAPClient(aapTarget)
	if err != nil {
		printer.Error(fmt.Errorf("invalid aap target %s", aapTarget))
		return aziclicommon.ErrCommandSilent
	}

	page := v.GetInt32(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonPage))
	pageSize := v.GetInt32(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonPageSize))
	accountID := v.GetInt64(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonAccountID))
	name := v.GetString(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonName))

	accounts, err := client.FetchAccountsBy(page, pageSize, accountID, name)
	if err != nil {
		if ctx.IsTerminalOutput() {
			printer.Println("Failed list the accounts.")
			if ctx.IsVerboseTerminalOutput() {
				printer.Error(err)
			}
		}
		return aziclicommon.ErrCommandSilent
	}
	output := map[string]any{}
	if ctx.IsTerminalOutput() {
		for _, account := range accounts {
			accountID := fmt.Sprintf("%d", account.AccountID)
			output[accountID] = account.Name
		}
	} else if ctx.IsJSONOutput() {
		output["accounts"] = accounts
	}
	printer.PrintlnMap(output)
	return nil
}

// createCommandForAccountList creates a command for managing accountlist.
func createCommandForAccountList(deps azcli.CliDependenciesProvider, v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List remote accounts",
		Long: aziclicommon.BuildCliLongTemplate(`This command lists all remote accounts.

Examples:
  # list all accounts and output the result in json format
  permguard accounts list --output json
  # list all accounts for page 1 and page size 100
  permguard accounts list --page 1 --size 100
  # list accounts and filter by account
  permguard accounts list --account 301
  # list accounts and filter by account and name
  permguard accounts list --account 301--name dev
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForListAccounts(deps, cmd, v)
		},
	}
	command.Flags().Int32P(aziclicommon.FlagCommonPage, aziclicommon.FlagCommonPageShort, 1, "page number")
	v.BindPFlag(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonPage), command.Flags().Lookup(aziclicommon.FlagCommonPage))
	command.Flags().Int32P(aziclicommon.FlagCommonPageSize, aziclicommon.FlagCommonPageSizeShort, 1000, "page size")
	v.BindPFlag(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonPageSize), command.Flags().Lookup(aziclicommon.FlagCommonPageSize))
	command.Flags().Int64(aziclicommon.FlagCommonAccountID, 0, "account id filter")
	v.BindPFlag(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonAccountID), command.Flags().Lookup(aziclicommon.FlagCommonAccountID))
	command.Flags().String(aziclicommon.FlagCommonName, "", "account name filter")
	v.BindPFlag(azoptions.FlagName(commandNameForAccountsList, aziclicommon.FlagCommonName), command.Flags().Lookup(aziclicommon.FlagCommonName))
	return command
}
