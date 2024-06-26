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
	commandNameForIdentitySourcesUpdate = "identitysources.update"
)

// runECommandForUpdateIdentitySource runs the command for creating an identity source.
func runECommandForUpdateIdentitySource(cmd *cobra.Command, v *viper.Viper) error {
	return runECommandForUpsertIdentitySource(cmd, v, commandNameForIdentitySourcesUpdate, false)
}

// createCommandForIdentitySourceUpdate creates a command for managing identity source update.
func createCommandForIdentitySourceUpdate(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "update",
		Short: "Update an identity source",
		Long: `This command update an identity source.

Examples:
  # update an identity source with name permguard, id 19159d69-e902-418e-966a-148c4d5169a4 and account 301990992055
  permguard authn identitysources update --account 301990992055 --identitysourceid 19159d69-e902-418e-966a-148c4d5169a4 --name permguard
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForUpdateIdentitySource(cmd, v)
		},
	}
	command.Flags().String(flagIdentitySourceID, "", "identity source id")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitySourcesUpdate, flagIdentitySourceID), command.Flags().Lookup(flagIdentitySourceID))
	command.Flags().String(flagCommonName, "", "identity source name")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitySourcesUpdate, flagCommonName), command.Flags().Lookup(flagCommonName))
	return command
}
