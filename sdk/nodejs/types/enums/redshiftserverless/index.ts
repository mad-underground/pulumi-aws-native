// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const NamespaceLogExport = {
    Useractivitylog: "useractivitylog",
    Userlog: "userlog",
    Connectionlog: "connectionlog",
} as const;

export type NamespaceLogExport = (typeof NamespaceLogExport)[keyof typeof NamespaceLogExport];

export const NamespaceStatus = {
    Available: "AVAILABLE",
    Modifying: "MODIFYING",
    Deleting: "DELETING",
} as const;

export type NamespaceStatus = (typeof NamespaceStatus)[keyof typeof NamespaceStatus];
