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

package metrics

import "os"

const UnknownValue = "unknown"

func GetBaseTags() map[string]string {
	return map[string]string{
		"env":     getEnv(),
		"service": getService(),
		"version": getVersion(),
		"cluster": getClusterName(),
	}
}

func mergeTags(tagSets ...map[string]string) map[string]string {
	res := make(map[string]string)
	for _, tagSet := range tagSets {
		for k, v := range tagSet {
			if _, ok := res[k]; !ok {
				res[k] = v
			} else if res[k] == "unknown" {
				res[k] = v
			}
		}
	}
	return res
}

func getEnv() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "default_env"
	}
	return env
}

func getService() string {
	service := os.Getenv("SERVICE")
	if service == "" {
		service = "default_service"
	}
	return service
}

func getVersion() string {
	// TODO this should be the actual foundationdb version
	version := os.Getenv("FDB_VERSION")
	if version == "" {
		version = "default_version"
	}
	return version
}

func getClusterName() string {
	clusterName := os.Getenv("FDB_CLUSTER_NAME")
	if clusterName == "" {
		clusterName = "default_cluster_name"
	}
	return clusterName
}

func getDefaultValue(tagKey string) string {
	switch tagKey {
	case "env":
		return getEnv()
	case "service":
		return getService()
	case "version":
		return getVersion()
	case "cluster_name":
		return getClusterName()
	default:
		return UnknownValue
	}
}

func StandardizeTags(tags map[string]string, stdKeys []string) map[string]string {
	res := tags
	for _, tagKey := range stdKeys {
		if _, ok := tags[tagKey]; !ok {
			// tag is missing, need to add it
			res[tagKey] = getDefaultValue(tagKey)
		} else if res[tagKey] == "" {
			res[tagKey] = getDefaultValue(tagKey)
		}
	}
	for k := range res {
		extraTag := true
		// result has an extra tag that should not be there
		for _, stdKey := range stdKeys {
			if stdKey == k {
				extraTag = false
			}
		}
		if extraTag {
			delete(res, k)
		}
	}
	return res
}
func convertBool(value bool) float64 {
	if value {
		return 1
	} else {
		return 0
	}
}
