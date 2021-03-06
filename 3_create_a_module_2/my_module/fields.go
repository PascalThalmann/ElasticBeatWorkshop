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

package my_module

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "my_module", asset.ModuleFieldsPri, AssetMyModule); err != nil {
		panic(err)
	}
}

// AssetMyModule returns asset data.
// This is the base64 encoded zlib format compressed contents of module/my_module.
func AssetMyModule() string {
	return "eJy0lE+PmzAQxe98iqe9r9T2UolDP0KvPaJZPMAUYyN7KOXbV4YQNoT8UbXrSxKb/N57eGZe0fKUo5uKzpvBcgaoqOUcL+e9lwwIbJki53hjpQwwHMsgvYp3OX5kADYGTh8S0VNQ+AraMDrWIOUbk2L0oY2N7/EtAypha2I+M17hqONLP2np1HOOOvihP+0cGNhQp58brxLLRefd+eAImdY+57oO5VbRK/b7SHsrhisarBYd/S0M2wuVzZd/+82l7o6WzWJ5wnpX3xWISkELlY6fV7gTczme0ZjRSOh0uZ13oj6IqzGKM34ERVAINGEUbfAlb/wQQM7ga96JG5TjXefszOf4Zmc+z/UGLEbmtjA0xY9LQFNcGynRMTYceOmrRRdkOWgElSp/+CpMHJyhaY7zPY+kQzA0nWUOg6XC/rgEyWoiJnZcXFWDtehJmzXZrAht0ogQa9dobHa0o/7aOy/SV1w9skZoeRp92IOfyJHWT1pqaPUM9ahZT9dhpJKS0v9xUMXvhtyNCfDgPePhJLgQWavy9jx4qHfFqm+16P+gbjXLE6/hiav6xdzOzSMO4jRm/wIAAP//4lXlDw=="
}
