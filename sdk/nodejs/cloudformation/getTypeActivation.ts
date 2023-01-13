// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Enable a resource that has been published in the CloudFormation Registry.
 */
export function getTypeActivation(args: GetTypeActivationArgs, opts?: pulumi.InvokeOptions): Promise<GetTypeActivationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudformation:getTypeActivation", {
        "arn": args.arn,
    }, opts);
}

export interface GetTypeActivationArgs {
    /**
     * The Amazon Resource Name (ARN) of the extension.
     */
    arn: string;
}

export interface GetTypeActivationResult {
    /**
     * The Amazon Resource Name (ARN) of the extension.
     */
    readonly arn?: string;
    /**
     * Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.
     */
    readonly autoUpdate?: boolean;
    /**
     * The Major Version of the type you want to enable
     */
    readonly majorVersion?: string;
    /**
     * Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
     */
    readonly versionBump?: enums.cloudformation.TypeActivationVersionBump;
}
/**
 * Enable a resource that has been published in the CloudFormation Registry.
 */
export function getTypeActivationOutput(args: GetTypeActivationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTypeActivationResult> {
    return pulumi.output(args).apply((a: any) => getTypeActivation(a, opts))
}

export interface GetTypeActivationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the extension.
     */
    arn: pulumi.Input<string>;
}
