/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &MetricsManifest{}

// MetricsManifest scaffolds a file that defines the prometheus metrics
type MetricsManifest struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file template
func (f *MetricsManifest) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("monitoring", "metrics", "metrics.go")
	}

	f.TemplateBody = MetricsTemplate
	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

// nolint: lll
const MetricsTemplate = `
/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import "{{ .Repo }}/monitoring/metrics/util"

var metrics = [][]util.Metric{
	OperatorMetrics,
}

// Register OperatorMetrics metrics defined in example_metrics.go
func RegisterMetrics() {
	util.RegisterMetrics(metrics)
}

func ListMetrics() []util.Metric {
	return util.ListMetrics(metrics)
}
`
