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

package testutils

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	azmocks "github.com/permguard/permguard/internal/cli/testutils/mocks"
	azcli "github.com/permguard/permguard/pkg/cli"
)

func BaseCommandTest(t *testing.T, cmdFunc func(azcli.CLIDependenciesProvider, *viper.Viper)(*cobra.Command), outputs []string) () {
	assert := assert.New(t)
	v := viper.New()
	depsMocks := azmocks.NewCLIDependenciesMock()
	cmd := cmdFunc(depsMocks, v)
	assert.NotNil(cmd, "The command should not be nil")

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"-h"})

	err := cmd.Execute()

	assert.NoError(err)
	output := buf.String()
	for _, out := range outputs {
		assert.Contains(output, out)
	}
}