#!/usr/bin/env bats
#
# Copyright 2019 HAProxy Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

load '../../libs/dataplaneapi'
load "../../libs/get_json_path"
load '../../libs/haproxy_config_setup'
load '../../libs/resource_client'
load '../../libs/version'

load 'utils/_helpers'

@test "servers: Return an array of runtime servers' settings" {
  resource_get "$_RUNTIME_SERVER_BASE_PATH" "backend=test_backend"
	assert_equal "$SC" 200

	for name in "server_01" "server_02" "server_01"; do
  	assert_equal "$(get_json_path "$BODY" ".[] | select(.name | contains(\"$name\") ).name")" "$name"
  done
}

@test "servers: Replace server transient settings" {
  resource_put "$_RUNTIME_SERVER_BASE_PATH/server_01" "data/transient.json" "backend=test_backend"
	assert_equal "$SC" 200
  assert_equal "$(get_json_path "$BODY" '.name')" "server_01"
}

@test "servers: Return one server runtime settings" {
  resource_get "$_RUNTIME_SERVER_BASE_PATH/server_01" "backend=test_backend"
	assert_equal "$SC" 200
  assert_equal "$(get_json_path "$BODY" '.name')" "server_01"
}
