# http-load-tester

This tool allows to generate an http load traffic at the specified rate (Request Per Minute) and defined minimum parallelism. The pacing will adjust based on the response time of each batch.

See [local.yaml](./tests/local.yaml) for more info on how to define the http configuration and batch specifications.

A `--dryrun` option is available to skip any http calls.

```bash
make test-unit compile-linux-x86 && ./bin/linux/http-load-tester --config tests/local.yaml
```

## Observability

Optionally, Metrics for each http attempt and results can be sent to New Relic. To do so use the following environment variable `NEW_RELIC_LICENSE_KEY=<my-license-key>` with the application.
If using an alternate environment than US, specify the metric endpoint by using an environment variable like this `NEW_RELIC_METRIC_API=https://metric-api.eu.newrelic.com/metric/v1`.

The dashboard [dashboard.json](.doc/dashboard.json) can be imported in New Relic to visualize the metrics. Update with your `ACCOUNT_ID` in the json file before importing.

Otherwise the metrics can be queried with:

Success/Fail Count
```
SELECT sum(httploadtester_results_test_count) as 'Total Count', sum(httploadtester_results_success_count) as 'Success Count', sum(httploadtester_results_fail_count) as 'Fail Count'
FROM Metric SINCE 10 MINUTES AGO TIMESERIES
```

Average Response Time
```
SELECT average(`Custom/httploadtester_results_elapsed_time_ms`) as 'Average Response Time (ms)'
FROM Metric SINCE 10 MINUTES AGO TIMESERIES
```
