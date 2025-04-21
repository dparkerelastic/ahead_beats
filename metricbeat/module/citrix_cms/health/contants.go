package health

const Count_API = "$count=true"

const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"
const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries"
const LoadIndexSummaries_API_PATH = "$filter=Granularity eq 60 and CreatedDate gt 2024-01-01T00:00:00.000Z&$top=1000&$count=true"
const MachineMetric_API = "/monitorodata/MachineMetric"
const SessionActivitySummaries_API = "/monitorodata/SessionActivitySummaries"
const SessionActivitySummaries_API_PATH = "$apply=filter(SummaryDate gt 2025-02-01T00:00:00Z)/groupby((DesktopGroupId),aggregate(ConnectedSessionCount with average as AverageConnectedSessionCount,DisconnectedSessionCount with average as AverageDisconnectedSessionCount,ConcurrentSessionCount with average as AverageConcurrentSessionCount,TotalLogOnDuration with average as AverageTotalLogOnDuration,TotalLogOnCount with sum as SumTotalLogOnCount))&$count=true&$top=1000"

// These are the details API endpoints for the metrics
const LogonMetric_Details_API = "/monitorodata/LogOnMetrics"
const LogonMetric_Details_API_PATH = "$expand=Session($expand=CurrentConnection),Session($expand=Machine),Session($expand=User),Session($expand=SessionMetricsLatest)&$count=true&%$top=1000"

const MachineMetric_Details_API = "/monitorodata/MachineMetric"
const MachineMetric_Details_API_PATH = "$expand=Machine($expand=CurrentLoadIndex),Machine($expand=Catalog),Machine($expand=DesktopGroup),Machine($expand=Hypervisor),Machine($expand=MachineCost)&$count=true&%$top=1000"

const SessionMetrics_Details_API = "/monitorodata/SessionMetrics"
const SessionMetrics_Details_API_PATH = "$expand=Session($expand=Failure),Session($expand=CurrentConnection),Session($expand=User),Session($expand=Machine),Session($expand=SessionMetricsLatest)&$count=true&%$top=1000"
