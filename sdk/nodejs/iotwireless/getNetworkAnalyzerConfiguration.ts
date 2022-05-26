// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Create and manage NetworkAnalyzerConfiguration resource.
 */
export function getNetworkAnalyzerConfiguration(args: GetNetworkAnalyzerConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkAnalyzerConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotwireless:getNetworkAnalyzerConfiguration", {
        "name": args.name,
    }, opts);
}

export interface GetNetworkAnalyzerConfigurationArgs {
    /**
     * Name of the network analyzer configuration
     */
    name: string;
}

export interface GetNetworkAnalyzerConfigurationResult {
    /**
     * Arn for network analyzer configuration, Returned upon successful create.
     */
    readonly arn?: string;
    /**
     * The description of the new resource
     */
    readonly description?: string;
    /**
     * Trace content for your wireless gateway and wireless device resources
     */
    readonly traceContent?: outputs.iotwireless.TraceContentProperties;
    /**
     * List of wireless gateway resources that have been added to the network analyzer configuration
     */
    readonly wirelessDevices?: string[];
    /**
     * List of wireless gateway resources that have been added to the network analyzer configuration
     */
    readonly wirelessGateways?: string[];
}

export function getNetworkAnalyzerConfigurationOutput(args: GetNetworkAnalyzerConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkAnalyzerConfigurationResult> {
    return pulumi.output(args).apply(a => getNetworkAnalyzerConfiguration(a, opts))
}

export interface GetNetworkAnalyzerConfigurationOutputArgs {
    /**
     * Name of the network analyzer configuration
     */
    name: pulumi.Input<string>;
}
