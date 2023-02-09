// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApplicationRuntimeEnvironmentType = {
    WineStaging: "WINE-STAGING",
} as const;

/**
 * Type of the runtime environment used to run games. For initial launch it only includes wine staging branch but Motif
 * will follow up with Proton support.
 */
export type ApplicationRuntimeEnvironmentType = (typeof ApplicationRuntimeEnvironmentType)[keyof typeof ApplicationRuntimeEnvironmentType];

export const StreamGroupStreamClass = {
    Mini: "MINI",
    Low: "LOW",
    Mid: "MID",
    High: "HIGH",
    Ultra: "ULTRA",
} as const;

/**
 * These are pre-packed Motif virtual instance type used to run customer games where mini spec maps to high-end
 * integrated graphics and ultra maps to high-end gaming pc. These have different cost implications.
 */
export type StreamGroupStreamClass = (typeof StreamGroupStreamClass)[keyof typeof StreamGroupStreamClass];