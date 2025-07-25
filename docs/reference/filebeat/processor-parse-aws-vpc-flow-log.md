---
navigation_title: "parse_aws_vpc_flow_log"
mapped_pages:
  - https://www.elastic.co/guide/en/beats/filebeat/current/processor-parse-aws-vpc-flow-log.html
applies_to:
  stack: ga
---

# Parse AWS VPC Flow Log [processor-parse-aws-vpc-flow-log]


The `parse_aws_vpc_flow_log` processor decodes AWS VPC Flow log messages.

Below is an example configuration that decodes the `message` field using the default version 2 VPC flow log format.

```yaml
processors:
  - parse_aws_vpc_flow_log:
      format: version account-id interface-id srcaddr dstaddr srcport dstport protocol packets bytes start end action log-status
      field: message
```

The `parse_aws_vpc_flow_log` processor has the following configuration settings.

| Name | Required | Default | Description |
| --- | --- | --- | --- |
| `field` | no | `message` | Source field containing the VPC flow log message. |
| `target_field` | no | `aws.vpcflow` | Target field for the VPC flow log object. This applies only to the original VPC flow log fields. ECS fields are written to the standard location. |
| `format` | yes |  | VPC flow log format. This supports VPC flow log fields from versions 2 through 5. It will accept a string or a list of strings. Each format must have a unique number of fields to enable matching it to a flow log message. |
| `mode` | no | `ecs` | Mode controls what fields are generated. The available options are `original`, `ecs`, and `ecs_and_original`. `original` generates the fields specified in the format string. `ecs` maps the original fields to ECS and removes the original fields that are mapped to ECS. `ecs_and_original` maps the original fields to ECS and retains all the original fields. |
| `ignore_missing` | no | false | Ignore missing source field. |
| `ignore_failure` | no | false | Ignore failures while parsing and transforming the flow log message. |
| `id` | no |  | Instance ID for debugging purposes. |


## Modes [_modes]


### Original [_original]

This mode returns the same fields found in the `format` string. It will drop any fields whose value a dash (`-`). It converts the strings into the appropriate data types. These are the known field names and their data types.

::::{note}
The AWS VPC flow field names use underscores instead of dashes within Filebeat. You may configure the `format` using field names that contain either.
::::


| VPC Flow Log Field | Data Type |
| --- | --- |
| account_id | string |
| action | string |
| az_id | string |
| bytes | long |
| dstaddr | ip |
| dstport | integer |
| end | timestamp |
| flow_direction | string |
| instance_id | string |
| interface_id | string |
| log_status | string |
| packets | long |
| pkt_dst_aws_service | string |
| pkt_dstaddr | ip |
| pkt_src_aws_service | string |
| pkt_srcaddr | ip |
| protocol | integer |
| region | string |
| srcaddr | ip |
| srcport | integer |
| start | timestamp |
| sublocation_id | string |
| sublocation_type | string |
| subnet_id | string |
| tcp_flags | integer |
| tcp_flags_array* | integer |
| traffic_path | integer |
| type | string |
| version | integer |
| vpc_id | string |


### ECS [_ecs]

This mode maps the original VPC flow log fields into their associated Elastic Common Schema (ECS) fields. It removes the original fields that were mapped to ECS to reduced duplication. These are the field associations. There may be some transformations applied to derive the ECS field.

| VPC Flow Log Field | ECS Field |
| --- | --- |
| account_id | cloud.account.id |
| action | event.outcome |
| action | event.action |
| action | event.type |
| az_id | cloud.availability_zone |
| bytes | network.bytes |
| bytes | source.bytes |
| dstaddr | destination.address |
| dstaddr | destination.ip |
| dstport | destination.port |
| end | @timestamp |
| end | event.end |
| flow_direction | network.direction |
| instance_id | cloud.instance.id |
| packets | network.packets |
| packets | source.packets |
| protocol | network.iana_number |
| protocol | network.transport |
| region | cloud.region |
| srcaddr | network.type |
| srcaddr | source.address |
| srcaddr | source.ip |
| srcport | source.port |
| start | event.start |


### ECS and Original [_ecs_and_original]

This mode maps the fields into ECS and retains all the original fields. Below is an example document produced using `ecs_and_orignal` mode.

```json
{
  "@timestamp": "2021-03-26T03:29:09Z",
  "aws": {
    "vpcflow": {
      "account_id": "64111117617",
      "action": "REJECT",
      "az_id": "use1-az5",
      "bytes": 1,
      "dstaddr": "10.200.0.0",
      "dstport": 33004,
      "end": "2021-03-26T03:29:09Z",
      "flow_direction": "ingress",
      "instance_id": "i-0axxxxxx1ad77",
      "interface_id": "eni-069xxxxxb7a490",
      "log_status": "OK",
      "packets": 52,
      "pkt_dst_aws_service": "CLOUDFRONT",
      "pkt_dstaddr": "10.200.0.80",
      "pkt_src_aws_service": "AMAZON",
      "pkt_srcaddr": "89.160.20.156",
      "protocol": 17,
      "region": "us-east-1",
      "srcaddr": "89.160.20.156",
      "srcport": 50041,
      "start": "2021-03-26T03:28:12Z",
      "sublocation_id": "fake-id",
      "sublocation_type": "wavelength",
      "subnet_id": "subnet-02d645xxxxxxxdbc0",
      "tcp_flags": 1,
      "tcp_flags_array": [
        "fin"
      ],
      "traffic_path": 1,
      "type": "IPv4",
      "version": 5,
      "vpc_id": "vpc-09676f97xxxxxb8a7"
    }
  },
  "cloud": {
    "account": {
      "id": "64111117617"
    },
    "availability_zone": "use1-az5",
    "instance": {
      "id": "i-0axxxxxx1ad77"
    },
    "region": "us-east-1"
  },
  "destination": {
    "address": "10.200.0.0",
    "ip": "10.200.0.0",
    "port": 33004
  },
  "event": {
    "action": "reject",
    "end": "2021-03-26T03:29:09Z",
    "outcome": "failure",
    "start": "2021-03-26T03:28:12Z",
    "type": [
      "connection",
      "denied"
    ]
  },
  "message": "5 64111117617 eni-069xxxxxb7a490 89.160.20.156 10.200.0.0 50041 33004 17 52 1 1616729292 1616729349 REJECT OK vpc-09676f97xxxxxb8a7 subnet-02d645xxxxxxxdbc0 i-0axxxxxx1ad77 1 IPv4 89.160.20.156 10.200.0.80 us-east-1 use1-az5 wavelength fake-id AMAZON CLOUDFRONT ingress 1",
  "network": {
    "bytes": 1,
    "direction": "ingress",
    "iana_number": "17",
    "packets": 52,
    "transport": "udp",
    "type": "ipv4"
  },
  "related": {
    "ip": [
      "89.160.20.156",
      "10.200.0.0",
      "10.200.0.80"
    ]
  },
  "source": {
    "address": "89.160.20.156",
    "bytes": 1,
    "ip": "89.160.20.156",
    "packets": 52,
    "port": 50041
  }
}
```

