// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::LayerVersion
 */
export function getLayerVersion(args: GetLayerVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetLayerVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lambda:getLayerVersion", {
        "id": args.id,
    }, opts);
}

export interface GetLayerVersionArgs {
    /**
     * The ARN of the version.
     */
    id: string;
}

export interface GetLayerVersionResult {
    /**
     * The ARN of the version.
     */
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::Lambda::LayerVersion
 */
export function getLayerVersionOutput(args: GetLayerVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLayerVersionResult> {
    return pulumi.output(args).apply((a: any) => getLayerVersion(a, opts))
}

export interface GetLayerVersionOutputArgs {
    /**
     * The ARN of the version.
     */
    id: pulumi.Input<string>;
}
