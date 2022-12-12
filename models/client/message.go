// Copyright 2022 Tigris Data, Inc.
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

package clientmodel

type Message struct {
	// Possible names
	// "inconsistent_cluster_file",
	// "unreachable_cluster_controller",
	// "no_cluster_controller",
	// "status_incomplete_client",
	// "status_incomplete_coordinators",
	// "status_incomplete_error",
	// "status_incomplete_timeout",
	// "status_incomplete_cluster",
	// "quorum_not_reachable"
	Name        string
	Description string
}
