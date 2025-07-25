---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/packetbeat/current/protocol-metrics-packetbeat.html
applies_to:
  stack: ga
---

# Protocol-Specific Metrics [protocol-metrics-packetbeat]

Packetbeat exposes per-protocol metrics under the [HTTP monitoring endpoint](/reference/packetbeat/http-endpoint.md). These metrics are exposed under the `/inputs/` path. They can be used to observe the activity of Packetbeat for the monitored protocol.


## AF_PACKET Metrics [_af_packet_metrics]

| Metric | Description |
| --- | --- |
| `device` | Name of the device being monitored. |
| `socket_packets` | Number of packets delivered by the kernel to the shared buffer. |
| `socket_drops` | Number of packets dropped by the kernel on the socket. |
| `socket_queue_freezes` | Number of kernel queue freezes on the socket. |
| `packets` | Number of packets handled by Packetbeat. |
| `polls` | Number of blocking syscalls made waiting for packets. |


## TCP Metrics [_tcp_metrics]

| Metric | Description |
| --- | --- |
| `device` | Name of the device being monitored. |
| `received_events_total` | Number of packets processed. |
| `received_bytes_total` | Number of bytes processed. |
| `tcp_overlaps` | Number of packets shrunk due to overlap. |
| `tcp.dropped_because_of_gaps` | Number of packets dropped because of gaps. |
| `arrival_period` | Histogram of the elapsed time between packet arrivals. |
| `processing_time` | Histogram of the elapsed time between packet receipt and publication. |
| `fin_flags_total` | Number of TCP FIN (finish) flags observed. |
| `syn_flags_total` | Number of TCP SYN (synchronization) flags observed. |
| `rst_flags_total` | Number of TCP RST (reset) flags observed. |
| `psh_flags_total` | Number of TCP PSH (push) flags observed. |
| `ack_flags_total` | Number of TCP ACK (acknowledgement) flags observed. |
| `urg_flags_total` | Number of TCP URG (urgent) flags observed. |
| `ece_flags_total` | Number of TCP ECE (ECN echo) flags observed. |
| `cwr_flags_total` | Number of TCP CWR (congestion window reduced) flags observed. |
| `ns_flags_total` | Number of TCP NS (nonce sum) flags observed. |
| `received_headers_total` | Number of headers observed, including unprocessed packets. |


## UDP Metrics [_udp_metrics]

| Metric | Description |
| --- | --- |
| `device` | Name of the device being monitored. |
| `received_events_total` | Number of packets processed. |
| `received_bytes_total` | Number of bytes processed. |
| `arrival_period` | Histogram of the elapsed time between packet arrivals. |
| `processing_time` | Histogram of the elapsed time between packet receipt and publication. |

