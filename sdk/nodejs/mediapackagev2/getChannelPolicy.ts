// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * <p>Represents a resource-based policy that allows or denies access to a channel.</p>
 */
export function getChannelPolicy(args: GetChannelPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:mediapackagev2:getChannelPolicy", {
        "channelGroupName": args.channelGroupName,
        "channelName": args.channelName,
    }, opts);
}

export interface GetChannelPolicyArgs {
    channelGroupName: string;
    channelName: string;
}

export interface GetChannelPolicyResult {
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MediaPackageV2::ChannelPolicy` for more information about the expected schema for this property.
     */
    readonly policy?: any;
}
/**
 * <p>Represents a resource-based policy that allows or denies access to a channel.</p>
 */
export function getChannelPolicyOutput(args: GetChannelPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChannelPolicyResult> {
    return pulumi.output(args).apply((a: any) => getChannelPolicy(a, opts))
}

export interface GetChannelPolicyOutputArgs {
    channelGroupName: pulumi.Input<string>;
    channelName: pulumi.Input<string>;
}
