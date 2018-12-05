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

package fixtures

// RouteInt type
type RouteInt int

const (
	// RouteTestResponse enum
	RouteTestResponse RouteInt = iota
	// RouteTestResponseError enum
	RouteTestResponseError
	// RouteTestResponseNotFound enum
	RouteTestResponseNotFound
)

var routeTestResponseText = map[RouteInt]string{
	RouteTestResponse:         "/test-response",
	RouteTestResponseError:    "/test-response-error",
	RouteTestResponseNotFound: "/test-response-notfound",
}

// RouteText value
func RouteText(code RouteInt) string {
	return routeTestResponseText[code]
}
