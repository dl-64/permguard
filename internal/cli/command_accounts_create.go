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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	azconfigs "github.com/permguard/permguard/pkg/configs"
)

const (
	commandNameForAccountsCreate = "accounts.create"
)

// runECommandForCreateAccount runs the command for creating an account.
func runECommandForCreateAccount(cmd *cobra.Command, v *viper.Viper) error {
	return runECommandForUpsertAccount(cmd, v, commandNameForAccountsCreate, true)
}

// createCommandForAccountCreate creates a command for managing accountcreate.
func createCommandForAccountCreate(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create an account",
		Long: `This command create an account.

Examples:
  # create an account with name mycorporate
  permguard accounts create --name mycorporate
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForCreateAccount(cmd, v)
		},
	}
	command.Flags().String(flagCommonName, "", "account name")
	v.BindPFlag(azconfigs.FlagName(commandNameForAccountsCreate, flagCommonName), command.Flags().Lookup(flagCommonName))
	return command
}
