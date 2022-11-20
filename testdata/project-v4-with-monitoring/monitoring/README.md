
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
