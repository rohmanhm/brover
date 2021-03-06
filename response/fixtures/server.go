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

import (
	"net/http"

	"github.com/rohmanhm/brover/response"
)

// Mux custom server
type Mux struct {
	http.ServeMux
}

func (c *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	current.ServeHTTP(w, r)
}

// SetupServer default
func SetupServer() *http.Server {
	mux := &Mux{}
	mux.HandleFunc(RouteText(RouteTestResponse), handleTestResponse)

	server := &http.Server{}
	server.Handler = mux

	return server
}

func handleTestResponse(w http.ResponseWriter, r *http.Request) {
	result := response.New()
	result.Data([]string{
		"banana",
		"apple",
	}).
		Render(w)
}
