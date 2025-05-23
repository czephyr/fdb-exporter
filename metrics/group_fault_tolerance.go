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

import (
	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/models"
)

type DbClusterFaultTolerance struct {
	metricGroup
}

func NewDbClusterFaultTolerance(info *MetricReporter) *DbClusterFaultTolerance {
	return &DbClusterFaultTolerance{*newMetricGroup("fault_tolerance", info.GetScopeOrExit("cluster"), info)}
}

func (d *DbClusterFaultTolerance) GetMetrics(status *models.FullStatus) {
	scope := d.GetScopeOrExit("default")
	if !isValidClusterFaultTolerance(status) {
		log.Error().Msg("failed to get database fault tolerance")
		return
	}
	SetGauge(scope, "zone_failures_availability", GetBaseTags(), status.Cluster.FaultTolerance.MaxZoneFailuresWithoutLosingAvailability)
	SetGauge(scope, "zone_failures_data", GetBaseTags(), status.Cluster.FaultTolerance.MaxZoneFailuresWithoutLosingData)
}
