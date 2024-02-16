// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A security profile defines a set of expected behaviors for devices in your account.
 */
export function getSecurityProfile(args: GetSecurityProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityProfileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getSecurityProfile", {
        "securityProfileName": args.securityProfileName,
    }, opts);
}

export interface GetSecurityProfileArgs {
    /**
     * A unique identifier for the security profile.
     */
    securityProfileName: string;
}

export interface GetSecurityProfileResult {
    /**
     * A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.
     */
    readonly additionalMetricsToRetainV2?: outputs.iot.SecurityProfileMetricToRetain[];
    /**
     * Specifies the destinations to which alerts are sent.
     */
    readonly alertTargets?: {[key: string]: outputs.iot.SecurityProfileAlertTarget};
    /**
     * Specifies the behaviors that, when violated by a device (thing), cause an alert.
     */
    readonly behaviors?: outputs.iot.SecurityProfileBehavior[];
    /**
     * A structure containing the mqtt topic for metrics export.
     */
    readonly metricsExportConfig?: outputs.iot.MetricsExportConfigProperties;
    /**
     * The ARN (Amazon resource name) of the created security profile.
     */
    readonly securityProfileArn?: string;
    /**
     * A description of the security profile.
     */
    readonly securityProfileDescription?: string;
    /**
     * Metadata that can be used to manage the security profile.
     */
    readonly tags?: outputs.iot.SecurityProfileTag[];
    /**
     * A set of target ARNs that the security profile is attached to.
     */
    readonly targetArns?: string[];
}
/**
 * A security profile defines a set of expected behaviors for devices in your account.
 */
export function getSecurityProfileOutput(args: GetSecurityProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityProfileResult> {
    return pulumi.output(args).apply((a: any) => getSecurityProfile(a, opts))
}

export interface GetSecurityProfileOutputArgs {
    /**
     * A unique identifier for the security profile.
     */
    securityProfileName: pulumi.Input<string>;
}
