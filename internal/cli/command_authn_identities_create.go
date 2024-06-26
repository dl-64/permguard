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
	commandNameForIdentitiesCreate = "identities.create"
)

// runECommandForCreateIdentity runs the command for creating an identity.
func runECommandForCreateIdentity(cmd *cobra.Command, v *viper.Viper) error {
	return runECommandForUpsertIdentity(cmd, v, commandNameForIdentitiesCreate, true)
}

// createCommandForIdentityCreate creates a command for managing identitycreate.
func createCommandForIdentityCreate(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create an identity",
		Long: `This command create an identity.

Examples:
  # create an identity with name identity1 and account 301990992055
  permguard authn identities create --account 301990992055 --name identity1
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForCreateIdentity(cmd, v)
		},
	}
	command.Flags().String(flagIdentitySourceID, "", "identity source id")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitiesCreate, flagIdentitySourceID), command.Flags().Lookup(flagIdentitySourceID))
	command.Flags().String(flagIdentityKind, "", "identity kind")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitiesCreate, flagIdentityKind), command.Flags().Lookup(flagIdentityKind))
	command.Flags().String(flagCommonName, "", "identity name")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitiesCreate, flagCommonName), command.Flags().Lookup(flagCommonName))
	return command
}
