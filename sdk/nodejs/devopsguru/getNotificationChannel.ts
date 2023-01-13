// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
 */
export function getNotificationChannel(args: GetNotificationChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationChannelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:devopsguru:getNotificationChannel", {
        "id": args.id,
    }, opts);
}

export interface GetNotificationChannelArgs {
    /**
     * The ID of a notification channel.
     */
    id: string;
}

export interface GetNotificationChannelResult {
    /**
     * The ID of a notification channel.
     */
    readonly id?: string;
}
/**
 * This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
 */
export function getNotificationChannelOutput(args: GetNotificationChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNotificationChannelResult> {
    return pulumi.output(args).apply((a: any) => getNotificationChannel(a, opts))
}

export interface GetNotificationChannelOutputArgs {
    /**
     * The ID of a notification channel.
     */
    id: pulumi.Input<string>;
}
