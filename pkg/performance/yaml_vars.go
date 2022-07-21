// Copyright 2021 Red Hat, Inc.
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

package performance

import (
	"github.com/maistra/maistra-test-tool/pkg/util"
)

var (
	smcpName         string            = util.Getenv("SMCPNAME", "basic")
	meshNamespace    string            = util.Getenv("MESHNAMESPACE", "istio-system")
	appNSPrefix      string            = util.Getenv("APPNSPREFIX", "app-perf-test")
	nsCountBundle    string            = util.Getenv("NSCOUNTBUNDLE", "10,50,100,500")
	nsAcceptanceTime string            = util.Getenv("NSACCEPTANCETIME", "5")
	prometheusAPIMap map[string]string = map[string]string{
		"xds_ppctb": "histogram_quantile(0.99, sum(rate(pilot_proxy_convergence_time_bucket[1m])) by (le))",
	}
)
