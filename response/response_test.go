// Copyright 2018 MH Rohman Masyhar
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

package response_test

import (
	"testing"

	"github.com/rohmanhm/brover/response"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("good response", func(t *testing.T) {
		resp := response.New()
		resp.Data([]string{
			"banana",
			"apple",
		})
		expected := `{"data":["banana","apple"]}`
		assert.Equal(t, expected, resp.String())
	})
	t.Run("bad response", func(t *testing.T) {
		resp := response.New()
		resp.Errors(&response.Errors{
			Message: "bad response",
		})
		expected := `{"errors":{"message":"bad response"}}`
		assert.Equal(t, expected, resp.String())
	})
}
