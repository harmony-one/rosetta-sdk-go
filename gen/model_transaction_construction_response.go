// Copyright 2020 Coinbase, Inc.
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

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

// TransactionConstructionResponse The transaction construction response should return the network
// fee to use in transaction construction. This network fee should be the gas price or cost per byte
// not some calculate value based on the method in the transaction construction request. Any
// calculations (like the one previously described) should be included in the metadata.
type TransactionConstructionResponse struct {
	NetworkFee *Amount                 `json:"network_fee"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
}
