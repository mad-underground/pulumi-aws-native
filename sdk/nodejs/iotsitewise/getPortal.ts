// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Portal
 */
export function getPortal(args: GetPortalArgs, opts?: pulumi.InvokeOptions): Promise<GetPortalResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotsitewise:getPortal", {
        "portalId": args.portalId,
    }, opts);
}

export interface GetPortalArgs {
    /**
     * The ID of the portal.
     */
    portalId: string;
}

export interface GetPortalResult {
    /**
     * Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.
     */
    readonly alarms?: outputs.iotsitewise.AlarmsProperties;
    /**
     * The email address that sends alarm notifications.
     */
    readonly notificationSenderEmail?: string;
    /**
     * The ARN of the portal, which has the following format.
     */
    readonly portalArn?: string;
    /**
     * The AWS SSO application generated client ID (used with AWS SSO APIs).
     */
    readonly portalClientId?: string;
    /**
     * The AWS administrator's contact email address.
     */
    readonly portalContactEmail?: string;
    /**
     * A description for the portal.
     */
    readonly portalDescription?: string;
    /**
     * The ID of the portal.
     */
    readonly portalId?: string;
    /**
     * A friendly name for the portal.
     */
    readonly portalName?: string;
    /**
     * The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.
     */
    readonly portalStartUrl?: string;
    /**
     * The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.
     */
    readonly roleArn?: string;
}

export function getPortalOutput(args: GetPortalOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortalResult> {
    return pulumi.output(args).apply(a => getPortal(a, opts))
}

export interface GetPortalOutputArgs {
    /**
     * The ID of the portal.
     */
    portalId: pulumi.Input<string>;
}
