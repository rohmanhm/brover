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

	"github.com/kataras/iris/httptest"
	"github.com/rohmanhm/brover/response/fixtures"
)

func TestIrisResponse(t *testing.T) {
	app := fixtures.SetupIrisServer()
	e := httptest.New(t, app)

	e.GET(fixtures.RouteText(fixtures.RouteTestResponse)).
		Expect().Status(httptest.StatusOK).Body().NotEmpty()

	e.GET(fixtures.RouteText(fixtures.RouteTestResponseError)).
		Expect().Status(httptest.StatusBadRequest).Body().NotEmpty()

	e.GET(fixtures.RouteText(fixtures.RouteTestResponseNotFound)).
		Expect().Status(httptest.StatusNotFound).Body().NotEmpty()
}
