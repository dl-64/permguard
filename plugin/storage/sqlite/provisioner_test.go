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

package sqlite

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSQLiteStorageFactory tests the SQLiteStorageFactory.
func TestSQLiteStorageProvisioner(t *testing.T) {
	assert := assert.New(t)

	storageProvisioner, err := NewSQLiteStorageProvisioner()
	assert.NotNil(storageProvisioner, "storage provisioner should not be nil")
	assert.Nil(err, "error should be nil")
	assert.Nil(storageProvisioner.AddFlags(&flag.FlagSet{}), "error should be nil")
	// //TODO: Complete this part of the test
	// assert.Nil(storageProvisioner.InitFromViper(viper.New()), "error should be nil")
}