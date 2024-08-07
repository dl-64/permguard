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
	azconfigs "github.com/permguard/permguard/pkg/configs"
)

const (
	commandNameForAccountsList = "accounts.list"
)

// runECommandForListAccounts runs the command for creating an account.
func runECommandForListAccounts(cmd *cobra.Command, v *viper.Viper) error {
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

	accountID := v.GetInt64(azconfigs.FlagName(commandNameForAccountsList, flagCommonAccountID))
	name := v.GetString(azconfigs.FlagName(commandNameForAccountsList, flagCommonName))

	accounts, err := client.GetAccountsBy(accountID, name)
	if err != nil {
		printer.Error(err)
		return ErrCommandSilent
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
	printer.Print(output)
	return nil
}

// createCommandForAccountList creates a command for managing accountlist.
func createCommandForAccountList(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List accounts",
		Long: `This command lists all the accounts.

Examples:
  # list all accounts
  permguard accounts list
  # list accounts and filter by account
  permguard accounts list --account 301
  # list accounts and filter by account and name
  permguard accounts list --account 301--name dev
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForListAccounts(cmd, v)
		},
	}
	command.Flags().Int64(flagCommonAccountID, 0, "account id filter")
	v.BindPFlag(azconfigs.FlagName(commandNameForAccountsList, flagCommonAccountID), command.Flags().Lookup(flagCommonAccountID))
	command.Flags().String(flagCommonName, "", "account name filter")
	v.BindPFlag(azconfigs.FlagName(commandNameForAccountsList, flagCommonName), command.Flags().Lookup(flagCommonName))
	return command
}
