package health

// const MachineDetails_API = "/monitorodata/Machines?%24expand=CurrentLoadIndex%2CCatalog%2CHypervisor%2CMachineCost%2CLoadIndex%2CResourceUtilization%2CMachineFailures%2CMachineMetric%2CMachineMetricSummary%2CMachineHotFixLogs&%24count=true&%24top=1"
// const MachineDetails_API = "/monitorodata/Machines?%24expand=CurrentLoadIndex%2CCatalog%2CHypervisor%2CMachineCost%2CLoadIndex%2CResourceUtilization%2CMachineFailures%2CMachineMetric%2CMachineMetricSummary%2CMachineHotFixLogs&%24count=true"
const MachineDetails_API = "/monitorodata/Machines?$expand=CurrentLoadIndex,Catalog,Hypervisor,MachineCost,LoadIndex,ResourceUtilization,MachineFailures,MachineMetric,MachineMetricSummary,MachineHotFixLogs&$count=true&$top=1000"

// const MachineDetails_API = "/monitorodata/Machines?$count=true&top=1000"
const MachineLoadIndex_API = "/monitorodata/Machines?$expand=CurrentLoadIndex"
const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries"

// const ServerOSDesktopSummaries_API = "/monitorodata/ServerOSDesktopSummaries?$filter=SummaryDate eq datetime'2023-10-01T00:00:00Z' and SummaryDate lt datetime'2023-10-02T00:00:00Z'&$orderby=SummaryDate desc&$top=1&$format=json"
const LoadIndexSummaries_API = "/monitorodata/LoadIndexSummaries?%24filter=Granularity%20eq%2060%20and%20CreatedDate%20gt%202024-01-01T00%3A00%3A00.000Z&%24top=1000&%24count=true"
const MachineMetric_API = "/monitorodata/MachineMetric?%24count=true"

//const MachineLoadIndex_API_V2 = "/monitorodata/Machines?$expand=CurrentLoadIndex&$filter=CurrentLoadIndex/EffectiveLoadIndex ne 0"
