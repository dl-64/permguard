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

package persistence

import (
	"fmt"
	"os"
	"path/filepath"

	azerrors "github.com/permguard/permguard/pkg/extensions/errors"
	aziclicommon "github.com/permguard/permguard/internal/cli/common"
)

// PersistenceManager implements the internal manager for the persistence file.
type PersistenceManager struct {
	rootDir string
	ctx	*aziclicommon.CliCommandContext
}

// NewPersistenceManager creates a new persistenceuration manager.
func NewPersistenceManager(rootDir string, ctx *aziclicommon.CliCommandContext) *PersistenceManager {
	return &PersistenceManager{
		rootDir: rootDir,
		ctx:     ctx,
	}
}

// CheckFileIfExists checks if a file exists.
func (p *PersistenceManager) CheckFileIfExists(relative bool, name string) (bool, error) {
	if relative {
		name = filepath.Join(p.rootDir, name)
	}
	if _, err := os.Stat(name); err == nil {
        return false, nil
	}
	return true, nil
}

// CreateFileIfNotExists creates a file if it does not exist.
func (p *PersistenceManager) CreateFileIfNotExists(relative bool, name string) (bool, error) {
	if relative {
		name = filepath.Join(p.rootDir, name)
	}
    if _, err := os.Stat(name); err == nil {
        return false, nil
    } else if os.IsNotExist(err) {
        dir := filepath.Dir(name)
        err := os.MkdirAll(dir, 0755)
        if err != nil {
			return false, azerrors.WrapSystemError(azerrors.ErrCliDirectoryOperation, fmt.Sprintf("cli: failed to create directory %s", dir))
        }
        file, err := os.Create(name)
        if err != nil {
            return false, azerrors.WrapSystemError(azerrors.ErrCliFileOperation, fmt.Sprintf("cli: failed to create file %s", name))
        }
        defer file.Close()
    } else if os.IsExist(err) {
		return false, nil
	} else {
		return false, azerrors.WrapSystemError(azerrors.ErrCliDirectoryOperation, fmt.Sprintf("cli: failed to stat directory %s", name))
    }
    return true, nil
}

// CreateDirIfNotExists creates a directory if it does not exist.
func (p *PersistenceManager) CreateDirIfNotExists(relative bool, name string) (bool, error) {
	if relative {
		name = filepath.Join(p.rootDir, name)
	}
	if _, err := os.Stat(name); err == nil {
		return false, nil
	} else if os.IsNotExist(err) {
		err := os.MkdirAll(name, 0755)
		if err != nil {
			return false, azerrors.WrapSystemError(azerrors.ErrCliDirectoryOperation, fmt.Sprintf("cli: failed to create directory %s", name))
		}
	} else {
		return false, azerrors.WrapSystemError(azerrors.ErrCliDirectoryOperation, fmt.Sprintf("cli: failed to stat directory %s", name))
	}
	return true, nil
}

// WriteFileIfNotExists writes a file if it does not exist.
func (p *PersistenceManager) WriteFileIfNotExists(relative bool, name string, data []byte, perm os.FileMode) (bool, error) {
	if relative {
		name = filepath.Join(p.rootDir, name)
	}
    if _, err := os.Stat(name); err == nil {
        return false, nil
	} else if os.IsExist(err) {
		return false, nil
	} else if os.IsNotExist(err) {
		return p.WriteFile(false, name, data, perm)
	} else {
		return false, azerrors.WrapSystemError(azerrors.ErrCliDirectoryOperation, fmt.Sprintf("cli: failed to stat directory %s", name))
	}
}

// WriteFile writes a file.
func (p *PersistenceManager) WriteFile(relative bool, name string, data []byte, perm os.FileMode) (bool, error)  {
	if relative {
		name = filepath.Join(p.rootDir, name)
	}
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		return false, azerrors.WrapSystemError(azerrors.ErrCliFileOperation, fmt.Sprintf("cli: failed to write file %s", name))
	}
	return true, nil
}