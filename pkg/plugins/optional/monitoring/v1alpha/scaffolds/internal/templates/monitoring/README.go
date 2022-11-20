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

package readme

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &ReadMeManifest{}

// ReadMeManifest scaffolds a file that defines the monitoring README file
type ReadMeManifest struct {
	machinery.TemplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file template
func (f *ReadMeManifest) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("monitoring", "README.md")
	}

	f.TemplateBody = ReadMeTemplate
	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

// nolint: lll
const ReadMeTemplate = `
## Operator Monitoring

In this folder we will outline what operators require in order to meet the best practices and examples for creating metrics, [recording rules](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/#recording-rules) and alerts. It is based on the general guidelines in [Operator Capability Levels](https://sdk.operatorframework.io/docs/overview/operator-capabilities/).

**Note:** For more technical documentation of how to add metrics to your operator, please read the [Metrics](https://book.kubebuilder.io/reference/metrics.html) section of the Kubebuilder documentation.

### Monitoring Guidelines

The following are recommended guidelines for monitoring an operator:

1. **Health and Performance metrics** - for all of the operator components
    1.1. Metrics should be implemented based on the best practice guidelines.
    1.2. **Metrics Documentation** - All metrics should have documentation.
    1.3. **Metrics Tests** - Metrics should include tests that verify that they exist and that their value is correct.
2. **Alerts** for when things are not working as expected for each of the operator's components
    2.1  Alerts should be implemented based on the guidelines below.
    2.2. **Alerts Runbooks** - An alert can include a \"runbook_url\" annotation and an alert runbook, See /docs/monitoring/runbooks/ (Optional).
    2.3. **Alerts Tests** - Alerts should include E2E Testing and unit tests.
3. **Events** - Custom Resources should emit custom events for the operations taking place.
`
