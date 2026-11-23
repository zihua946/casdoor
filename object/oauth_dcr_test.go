// Copyright 2026 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicMCPClientAdmission(t *testing.T) {
	for _, tc := range []struct {
		name    string
		app     *Application
		allowed bool
	}{
		{"deleted", nil, false},
		{"dynamic", &Application{RegistrationAccessToken: "registration", Organization: "mcp-local", Type: "MCP", GrantTypes: []string{"authorization_code"}}, true},
		{"static", &Application{Organization: "mcp-local", Type: "MCP", GrantTypes: []string{"authorization_code"}}, false},
		{"disabled", &Application{RegistrationAccessToken: "registration", Organization: "mcp-local", Type: "MCP", GrantTypes: []string{"authorization_code"}, DisableSignin: true}, false},
		{"wrong type", &Application{RegistrationAccessToken: "registration", Organization: "mcp-local", Type: "Web", GrantTypes: []string{"authorization_code"}}, false},
		{"no authorization code", &Application{RegistrationAccessToken: "registration", Organization: "mcp-local", Type: "MCP", GrantTypes: []string{"client_credentials"}}, false},
		{"no organization", &Application{RegistrationAccessToken: "registration", Type: "MCP", GrantTypes: []string{"authorization_code"}}, false},
	} {
		t.Run(tc.name, func(t *testing.T) { assert.Equal(t, tc.allowed, tc.app.IsDynamicMCPClient()) })
	}
}
