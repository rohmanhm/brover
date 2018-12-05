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

// Package response provides response utilities
package response

import "net/http"

// Response JSON
type Response struct {
	DataValue       interface{} `json:"data,omitempty"`
	ErrorsValue     *Errors     `json:"errors,omitempty"`
	StatusCodeValue int         `json:"-"`
	Context         interface{} `json:"-"`
}

// Errors for JSON errors
type Errors struct {
	Message string `json:"message"`
}

// Data Response
func (r *Response) Data(data interface{}) *Response {
	if r.StatusCodeValue == 0 {
		r.StatusCode(http.StatusOK)
	}

	r.DataValue = data
	return r
}

// StatusCode Response
func (r *Response) StatusCode(code int) *Response {
	r.StatusCodeValue = code
	return r
}

// Errors Response value
func (r *Response) Errors(err *Errors) *Response {
	if r.StatusCodeValue == 0 {
		r.StatusCode(http.StatusBadRequest)
	}

	r.ErrorsValue = err
	return r
}

// ErrorsMessage Response value
func (r *Response) ErrorsMessage(msg string) *Response {
	if r.ErrorsValue != nil {
		r.ErrorsValue.Message = msg
	}

	return r
}

// Render let user custom their own renders
// func (r *Response) Render() *Response {
// 	return nil
// }
