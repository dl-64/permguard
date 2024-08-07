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
	commandNameForAccountsDelete = "accounts.delete"
)

// runECommandForDeleteAccount runs the command for creating an account.
func runECommandForDeleteAccount(cmd *cobra.Command, v *viper.Viper) error {
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
	accountID := v.GetInt64(azconfigs.FlagName(commandNameForAccountsDelete, flagCommonAccountID))
	account, err := client.DeleteAccount(accountID)
	if err != nil {
		printer.Error(err)
		return ErrCommandSilent
	}
	output := map[string]any{}
	if ctx.IsTerminalOutput() {
		accountID := fmt.Sprintf("%d", account.AccountID)
		output[accountID] = account.Name
	} else if ctx.IsJSONOutput() {
		output["accounts"] = []*azmodels.Account{account}
	}
	printer.Print(output)
	return nil
}

// createCommandForAccountDelete creates a command for managing accountdelete.
func createCommandForAccountDelete(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete an account",
		Long: `This command delete an account.

Examples:
  # delete the account with the account id 301990992055
  permguard accounts delete --account 301990992055
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForDeleteAccount(cmd, v)
		},
	}
	command.Flags().Int64(flagCommonAccountID, 0, "account id")
	v.BindPFlag(azconfigs.FlagName(commandNameForAccountsDelete, flagCommonAccountID), command.Flags().Lookup(flagCommonAccountID))
	return command
}
