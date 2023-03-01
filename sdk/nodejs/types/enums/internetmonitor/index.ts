// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const MonitorConfigState = {
    Pending: "PENDING",
    Active: "ACTIVE",
    Inactive: "INACTIVE",
    Error: "ERROR",
} as const;

export type MonitorConfigState = (typeof MonitorConfigState)[keyof typeof MonitorConfigState];

export const MonitorProcessingStatusCode = {
    Ok: "OK",
    Inactive: "INACTIVE",
    CollectingData: "COLLECTING_DATA",
    InsufficientData: "INSUFFICIENT_DATA",
    FaultService: "FAULT_SERVICE",
    FaultAccessCloudwatch: "FAULT_ACCESS_CLOUDWATCH",
} as const;

export type MonitorProcessingStatusCode = (typeof MonitorProcessingStatusCode)[keyof typeof MonitorProcessingStatusCode];
