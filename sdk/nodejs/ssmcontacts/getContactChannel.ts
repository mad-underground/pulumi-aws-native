// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSMContacts::ContactChannel
 */
export function getContactChannel(args: GetContactChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetContactChannelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ssmcontacts:getContactChannel", {
        "arn": args.arn,
    }, opts);
}

export interface GetContactChannelArgs {
    /**
     * The Amazon Resource Name (ARN) of the engagement to a contact channel.
     */
    arn: string;
}

export interface GetContactChannelResult {
    /**
     * The Amazon Resource Name (ARN) of the engagement to a contact channel.
     */
    readonly arn?: string;
    /**
     * The details that SSM Incident Manager uses when trying to engage the contact channel.
     */
    readonly channelAddress?: string;
    /**
     * The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
     */
    readonly channelName?: string;
    /**
     * If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
     */
    readonly deferActivation?: boolean;
}
/**
 * Resource Type definition for AWS::SSMContacts::ContactChannel
 */
export function getContactChannelOutput(args: GetContactChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContactChannelResult> {
    return pulumi.output(args).apply((a: any) => getContactChannel(a, opts))
}

export interface GetContactChannelOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the engagement to a contact channel.
     */
    arn: pulumi.Input<string>;
}
