// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.
 */
export function getDimension(args: GetDimensionArgs, opts?: pulumi.InvokeOptions): Promise<GetDimensionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iot:getDimension", {
        "name": args.name,
    }, opts);
}

export interface GetDimensionArgs {
    /**
     * A unique identifier for the dimension.
     */
    name: string;
}

export interface GetDimensionResult {
    /**
     * The ARN (Amazon resource name) of the created dimension.
     */
    readonly arn?: string;
    /**
     * Specifies the value or list of values for the dimension.
     */
    readonly stringValues?: string[];
    /**
     * Metadata that can be used to manage the dimension.
     */
    readonly tags?: outputs.iot.DimensionTag[];
}

export function getDimensionOutput(args: GetDimensionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDimensionResult> {
    return pulumi.output(args).apply(a => getDimension(a, opts))
}

export interface GetDimensionOutputArgs {
    /**
     * A unique identifier for the dimension.
     */
    name: pulumi.Input<string>;
}
