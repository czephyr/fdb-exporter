// Copyright 2022-2023 Tigris Data, Inc.
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

package metrics

import "github.com/tigrisdata/fdb-exporter/models"

func isValidClusterLockState(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.DatabaseLockState == nil {
		return false
	}
	return true
}

func isValidClusterQos(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.Qos == nil {
		return false
	}
	return true
}

func isValidClusterFaultTolerance(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.FaultTolerance == nil {
		return false
	}
	return true
}

func isValidClusterData(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.Data == nil {
		return false
	}
	return true
}

func isValidClusterLatencyProbe(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.LatencyProbe == nil {
		return false
	}
	return true
}

func isValidClient(status *models.FullStatus) bool {
	if status == nil || status.Client == nil {
		return false
	}
	return true
}

func isValidWorkload(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.Workload == nil {
		return false
	}
	return true
}

func isValidProcesses(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.Processes == nil {
		return false
	}
	return true
}

func isValidBackup(status *models.FullStatus) bool {
	if status == nil || status.Cluster == nil || status.Cluster.Layers == nil || status.Cluster.Layers.Backup == nil {
		return false
	}
	return true
}
