// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EnvironmentEngineType = {
    Microfocus: "microfocus",
    Bluage: "bluage",
} as const;

/**
 * The target platform for the environment.
 */
export type EnvironmentEngineType = (typeof EnvironmentEngineType)[keyof typeof EnvironmentEngineType];
