// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cdaxbeat

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "cdaxbeat", asset.ModuleFieldsPri, AssetCdaxbeat); err != nil {
		panic(err)
	}
}

// AssetCdaxbeat returns asset data.
// This is the base64 encoded zlib format compressed contents of module/cdaxbeat.
func AssetCdaxbeat() string {
	return "eJyMkk1uwyAQhfec4in79AAseoReIcIwdqfhxzJYjW9fYdcpYEfqrKw3+OPNPK640yKhjXp0pJIAEidLEpddughgIksqkkRHSQnAUNQTj4mDl3gXAJ4EuGBmSwLomayJcu1e4ZWj6p5caRlJYpjCPP4qJ+SaVNJ6tnTLn3i2duadlu8wmUI/JW/1kRGhR/okjHYe2EMHa0kn9kNWI8FRmljHt+LHzYRTj5shu+6l9hC6L9KpkDfhtnVt8MPpRC74A6pcUa42kH8MeUS3ay2NJHa0jYWqXxlqW9XtkXTwJiKy1wSr4vo0uGet8oF939lWgznz1S5pjf1w4nX6e73OrCbkdNrx/kr8BAAA//87p9nW"
}
