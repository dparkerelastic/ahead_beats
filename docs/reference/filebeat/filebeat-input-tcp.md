---
navigation_title: "TCP"
mapped_pages:
  - https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-tcp.html
applies_to:
  stack: ga
---

# TCP input [filebeat-input-tcp]


Use the `TCP` input to read events over TCP.

Example configuration:

```yaml
filebeat.inputs:
- type: tcp
  max_message_size: 10MiB
  host: "localhost:9000"
```

## Configuration options [_configuration_options_19]

The `tcp` input supports the following configuration options plus the [Common options](#filebeat-input-tcp-common-options) described later.


### `max_message_size` [filebeat-input-tcp-tcp-max-message-size]

The maximum size of the message received over TCP. The default is `20MiB`.


### `host` [filebeat-input-tcp-tcp-host]

The host and TCP port to listen on for event streams.


### `network` [filebeat-input-tcp-tcp-network]

The network type. Acceptable values are: "tcp" (default), "tcp4", "tcp6"


### `framing` [filebeat-input-tcp-tcp-framing]

Specify the framing used to split incoming events.  Can be one of `delimiter` or `rfc6587`.  `delimiter` uses the characters specified in `line_delimiter` to split the incoming events.  `rfc6587` supports octet counting and non-transparent framing as described in [RFC6587](https://tools.ietf.org/html/rfc6587).  `line_delimiter` is used to split the events in non-transparent framing.  The default is `delimiter`.


### `line_delimiter` [filebeat-input-tcp-tcp-line-delimiter]

Specify the characters used to split the incoming events. The default is *\n*.


### `max_connections` [filebeat-input-tcp-tcp-max-connections]

The at most number of connections to accept at any given point in time.


### `timeout` [filebeat-input-tcp-tcp-timeout]

The number of seconds of inactivity before a remote connection is closed. The default is `300s`.


#### `ssl` [filebeat-input-tcp-tcp-ssl]

Configuration options for SSL parameters like the certificate, key and the certificate authorities to use.

See [SSL](/reference/filebeat/configuration-ssl.md) for more information.


## Metrics [_metrics_15]

This input exposes metrics under the [HTTP monitoring endpoint](/reference/filebeat/http-endpoint.md). These metrics are exposed under the `/inputs` path. They can be used to observe the activity of the input.

| Metric | Description |
| --- | --- |
| `device` | Host/port of the TCP stream. |
| `received_events_total` | Total number of packets (events) that have been received. |
| `received_bytes_total` | Total number of bytes received. |
| `receive_queue_length` | Aggregated size of the system receive queues (IPv4 and IPv6) (linux only) (gauge). |
| `arrival_period` | Histogram of the time between successive packets in nanoseconds. |
| `processing_time` | Histogram of the time taken to process packets in nanoseconds. |


## Common options [filebeat-input-tcp-common-options]

The following configuration options are supported by all inputs.


#### `enabled` [_enabled_26]

Use the `enabled` option to enable and disable inputs. By default, enabled is set to true.


#### `tags` [_tags_25]

A list of tags that Filebeat includes in the `tags` field of each published event. Tags make it easy to select specific events in Kibana or apply conditional filtering in Logstash. These tags will be appended to the list of tags specified in the general configuration.

Example:

```yaml
filebeat.inputs:
- type: tcp
  . . .
  tags: ["json"]
```


#### `fields` [filebeat-input-tcp-fields]

Optional fields that you can specify to add additional information to the output. For example, you might add fields that you can use for filtering log data. Fields can be scalar values, arrays, dictionaries, or any nested combination of these. By default, the fields that you specify here will be grouped under a `fields` sub-dictionary in the output document. To store the custom fields as top-level fields, set the `fields_under_root` option to true. If a duplicate field is declared in the general configuration, then its value will be overwritten by the value declared here.

```yaml
filebeat.inputs:
- type: tcp
  . . .
  fields:
    app_id: query_engine_12
```


#### `fields_under_root` [fields-under-root-tcp]

If this option is set to true, the custom [fields](#filebeat-input-tcp-fields) are stored as top-level fields in the output document instead of being grouped under a `fields` sub-dictionary. If the custom field names conflict with other field names added by Filebeat, then the custom fields overwrite the other fields.


#### `processors` [_processors_25]

A list of processors to apply to the input data.

See [Processors](/reference/filebeat/filtering-enhancing-data.md) for information about specifying processors in your config.


#### `pipeline` [_pipeline_25]

The ingest pipeline ID to set for the events generated by this input.

::::{note}
The pipeline ID can also be configured in the Elasticsearch output, but this option usually results in simpler configuration files. If the pipeline is configured both in the input and output, the option from the input is used.
::::


::::{important}
The `pipeline` is always lowercased. If `pipeline: Foo-Bar`, then the pipeline name in {{es}} needs to be defined as `foo-bar`.
::::



#### `keep_null` [_keep_null_25]

If this option is set to true, fields with `null` values will be published in the output document. By default, `keep_null` is set to `false`.


#### `index` [_index_25]

If present, this formatted string overrides the index for events from this input (for elasticsearch outputs), or sets the `raw_index` field of the event’s metadata (for other outputs). This string can only refer to the agent name and version and the event timestamp; for access to dynamic fields, use `output.elasticsearch.index` or a processor.

Example value: `"%{[agent.name]}-myindex-%{+yyyy.MM.dd}"` might expand to `"filebeat-myindex-2019.11.01"`.


#### `publisher_pipeline.disable_host` [_publisher_pipeline_disable_host_25]

By default, all events contain `host.name`. This option can be set to `true` to disable the addition of this field to all events. The default value is `false`.


