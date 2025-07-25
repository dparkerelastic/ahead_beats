---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-metricset-syncgateway-memory.html
---

% This file is generated! See scripts/docs_collector.py

# SyncGateway memory metricset [metricbeat-metricset-syncgateway-memory]

::::{warning}
This functionality is in beta and is subject to change. The design and code is less mature than official GA features and is being provided as-is with no warranties. Beta features are not subject to the support SLA of official GA features.
::::


SyncGateway `memory` metriset contains detailed information about the memory usage of the SyncGateway Go’s process.

## Fields [_fields]

For a description of each field in the metricset, see the [exported fields](/reference/metricbeat/exported-fields-syncgateway.md) section.

Here is an example document generated by this metricset:

```json
{
    "@timestamp": "2017-10-12T08:05:34.853Z",
    "event": {
        "dataset": "syncgateway.memory",
        "duration": 115000,
        "module": "syncgateway"
    },
    "metricset": {
        "name": "memory",
        "period": 10000
    },
    "service": {
        "address": "127.0.0.1:39195",
        "type": "syncgateway"
    },
    "syncgateway": {
        "memory": {
            "Alloc": 159398816,
            "BuckHashSys": 1759627,
            "DebugGC": false,
            "EnableGC": true,
            "Frees": 55790048,
            "GCCPUFraction": 0.008328526565497294,
            "GCSys": 13555712,
            "HeapAlloc": 159398816,
            "HeapIdle": 67780608,
            "HeapInuse": 264388608,
            "HeapObjects": 1877567,
            "HeapReleased": 26927104,
            "HeapSys": 332169216,
            "LastGC": 1620310398040906500,
            "Lookups": 0,
            "MCacheInuse": 20832,
            "MCacheSys": 32768,
            "MSpanInuse": 4030632,
            "MSpanSys": 4358144,
            "Mallocs": 57667615,
            "NextGC": 265304208,
            "NumForcedGC": 0,
            "NumGC": 51,
            "OtherSys": 2457453,
            "PauseTotalNs": 88693194,
            "StackInuse": 3375104,
            "StackSys": 3375104,
            "Sys": 357708024,
            "TotalAlloc": 2719416432
        }
    }
}
```
