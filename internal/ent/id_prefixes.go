// Copyright 2023 The Infratographer Authors
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

package entc

import (
	"context"
	"fmt"

	"go.infratographer.com/instance-api/internal/ent/generated/instance"
	"go.infratographer.com/instance-api/internal/ent/generated/instancemetadata"
	"go.infratographer.com/instance-api/internal/ent/generated/instanceprovider"
	"go.infratographer.com/x/gidx"
)

// prefixMap maps prefixes to table names.
var prefixMap = map[string]string{
	"instanc": instance.Table,
	"instinm": instancemetadata.Table,
	"instpvd": instanceprovider.Table,
}

// NodeTypeFromPrefixedID maps a PrefixedID to the underlying data table.
func NodeTypeFromPrefixedID(ctx context.Context, id gidx.PrefixedID) (string, error) {
	p := id.Prefix()
	if p == "" {
		return "", fmt.Errorf("prefix is missing")
	}
	typ := prefixMap[p]
	if typ == "" {
		return "", fmt.Errorf("prefix '%s' is unknown", p)
	}
	return typ, nil
}
