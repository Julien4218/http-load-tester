# http-load-tester

This tool allows to generate an http load traffic at the specified rate (Request Per Minute) and defined minimum parallelism. The pacing will adjust based on the response time of each batch.

See [local.yaml](./tests/local.yaml) for more info on how to define the http configuration and batch specifications.

Optionally, Metrics for each http attempt and results can be sent to New Relic. To do so use the following environment variable `NEW_RELIC_LICENSE_KEY=<my-license-key>`

A `--dryrun` option is available to skip any http calls.

```bash
make test-unit compile-linux-x86 && NEW_RELIC_LICENSE_KEY=<my-license-key> MY_KEY=abc ./bin/linux/http-load-tester --config tests/local.yaml --dryrun
```
