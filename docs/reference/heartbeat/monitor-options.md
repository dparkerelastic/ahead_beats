---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/heartbeat/current/monitor-options.html
applies_to:
  stack: ga
---

# Common monitor options [monitor-options]

You can specify the following options when defining a Heartbeat monitor in any location. These options are the same for all monitors. Each monitor type has additional configuration options that are specific to that monitor type.


### `type` [monitor-type]

The type of monitor to run. See [Monitor types](/reference/heartbeat/configuration-heartbeat-options.md#monitor-types).


### `id` [monitor-id]

A unique identifier for this configuration. This should not change with edits to the monitor configuration regardless of changes to any config fields. Examples: `uploader-service`, `http://example.net`, `us-west-loadbalancer`. Note that this uniqueness is only within a given beat instance. If you want to monitor the same endpoint from multiple locations it is recommended that those heartbeat instances use the same IDs so that their results can be correlated. You can use the `host.geo.name` property to disambiguate them.

When querying against indexed monitor data this is the field you will be aggregating with. Appears in the [exported fields](/reference/heartbeat/exported-fields.md) as `monitor.id`.

If you do not set this explicitly the monitor’s config will be hashed and a generated value used. This value will change with any options change to this monitor making aggregations over time between changes impossible. For this reason it is recommended that you set this manually.


### `name` [monitor-name]

Optional human readable name for this monitor. This value appears in the [exported fields](/reference/heartbeat/exported-fields.md) as `monitor.name`.


### `service.name` [service-name]

Optional APM service name for this monitor. Corresponds to the `service.name` ECS field. Set this when monitoring an app that is also using APM to enable integrations between Uptime and APM data in Kibana.


### `enabled` [monitor-enabled]

A Boolean value that specifies whether the module is enabled. If the `enabled` option is missing from the configuration block, the module is enabled by default.


### `schedule` [monitor-schedule]

A cron-like expression that specifies the task schedule. For example:

* `*/5 * * * * * *` runs the task every 5 seconds (for example, at 10:00:00, 10:00:05, and so on).
* `@every 5s` runs the task every 5 seconds from the time when Heartbeat was started.

The `schedule` option uses a cron-like syntax based on [this `cronexpr` implementation](https://github.com/gorhill/cronexpr#implementation), but adds the `@every` keyword.

For stats on the execution of scheduled tasks you can enable the HTTP stats server with `http.enabled: true` in heartbeat.yml, then run `curl http://localhost:5066/stats | jq .heartbeat.scheduler` to view the scheduler’s stats. Stats are provided for both jobs and tasks. Each time a monitor is scheduled is considered to be a single job, while portions of the work a job does, like DNS lookups and executing network requests are defined as tasks. The stats provided are:

* **jobs.active:** The number of actively running jobs/monitors.
* **jobs.missed_deadline:** The number of jobs that executed after their scheduled time. This can be caused either by overlong long timeouts from the previous job or high load preventing heartbeat from keeping up with work.
* **tasks.active:** The number of tasks currently running.
* **tasks.waiting:** If the global `schedule.limit` option is set, this number will reflect the number of tasks that are ready to execute, but have not been started in order to prevent exceeding `schedule.limit`.

Also see the [task scheduler](/reference/heartbeat/monitors-scheduler.md) settings.


### `ipv4` [monitor-ipv4]

A Boolean value that specifies whether to ping using the ipv4 protocol if hostnames are configured. The default is `true`.


### `ipv6` [monitor-ipv6]

A Boolean value that specifies whether to ping using the ipv6 protocol if hostnames are configured. The default is `true`.


### `mode` [monitor-mode]

If `mode` is `any`, the monitor pings only one IP address for a hostname. If `mode` is `all`, the monitor pings all resolvable IPs for a hostname. The `mode: all` setting is useful if you are using a DNS-load balancer and want to ping every IP address for the specified hostname. The default is `any`.


### `timeout` [monitor-timeout]

The total running time for each ping test. This is the total time allowed for testing the connection and exchanging data. The default is 16 seconds (16s).

If the timeout is exceeded, Heartbeat publishes a `service-down` event. If the value specified for `timeout` is greater than `schedule`, intermediate checks will not be executed by the scheduler.


## `run_from` [monitor-run-from]

Use the `run_from` option to set the geographic location fields relevant to a given heartbeat monitor.

The `run_from` option is used to label the geographic location where the monitor is running. Note, you can also set the `run_from` option on all monitors via the `heartbeat.run_from` option.

The `run_from` option takes two top-level fields:

* `id`: A string used to uniquely identify the geographic location. It is indexed as the `observer.name` field.
* `geo`: A map conforming to [ECS geo fields](ecs://reference/ecs-geo.md). It is indexed under `observer.geo`.

Example:

```yaml
- type: http
  # Set enabled to true (or delete the following line) to enable this example monitor
  enabled: true
  # ID used to uniquely identify this monitor in elasticsearch even if the config changes
  id: geo-test
  # Human readable display name for this service in Uptime UI and elsewhere
  name: Geo Test
  # List or urls to query
  urls: ["http://example.net"]
  # Configure task schedule
  schedule: '@every 10s'
  run_from:
    id: my-custom-geo
    geo:
      name: nyc-dc1-rack1
      location: 40.7128, -74.0060
      continent_name: North America
      country_iso_code: US
      region_name: New York
      region_iso_code: NY
      city_name: New York
```


### `fields` [monitor-fields]

Optional fields that you can specify to add additional information to the output. For example, you might add fields that you can use for filtering log data. Fields can be scalar values, arrays, dictionaries, or any nested combination of these. By default, the fields that you specify here will be grouped under a `fields` sub-dictionary in the output document. To store the custom fields as top-level fields, set the `fields_under_root` option to true. If a duplicate field is declared in the general configuration, then its value will be overwritten by the value declared here.


### `fields_under_root` [monitor-fields-under-root]

If this option is set to true, the custom [fields](#monitor-fields) are stored as top-level fields in the output document instead of being grouped under a `fields` sub-dictionary. If the custom field names conflict with other field names added by Heartbeat, then the custom fields overwrite the other fields.


### `tags` [monitor-tags]

A list of tags that will be sent with the monitor event. This setting is optional.


### `processors` [monitor-processors]

A list of processors to apply to the data generated by the monitor.

See [Processors](/reference/heartbeat/filtering-enhancing-data.md) for information about specifying processors in your config.


### `data_stream` [monitor-data-stream]

Contains options pertaining to data stream naming, following the conventions followed by [Fleet Data Streams](docs-content://reference/fleet/data-streams.md). By default Heartbeat will write to a datastream named `heartbeat-VERSION`.

```yaml
# To enable data streams with the default namespace
data_stream.namespace: default
```


#### `pipeline` [monitor-pipeline]

The {{es}} ingest pipeline ID to set for the events generated by this input.

::::{note}
The pipeline ID can also be configured in the Elasticsearch output, but this option usually results in simpler configuration files. If the pipeline is configured both in the input and output, the option from the input is used.
::::



#### `index` (deprecated) [monitor-index]

This setting is now deprecated in favor of the `data_stream` option. If present, this formatted string overrides the index for events from this input (for elasticsearch outputs), or sets the `raw_index` field of the event’s metadata (for other outputs). This string can only refer to the agent name and version and the event timestamp; for access to dynamic fields, use `output.elasticsearch.index` or a processor.

Example value: `"%{[agent.name]}-myindex-%{+yyyy.MM.dd}"` might expand to `"heartbeat-myindex-2019.11.01"`.


### `keep_null` [monitor-keep-null]

If this option is set to true, fields with `null` values will be published in the output document. By default, `keep_null` is set to `false`.
