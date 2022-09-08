// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::Application
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appstream:getApplication", {
        "arn": args.arn,
    }, opts);
}

export interface GetApplicationArgs {
    arn: string;
}

export interface GetApplicationResult {
    readonly appBlockArn?: string;
    readonly arn?: string;
    readonly attributesToDelete?: string[];
    readonly createdTime?: string;
    readonly description?: string;
    readonly displayName?: string;
    readonly iconS3Location?: outputs.appstream.ApplicationS3Location;
    readonly launchParameters?: string;
    readonly launchPath?: string;
    readonly workingDirectory?: string;
}

export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    arn: pulumi.Input<string>;
}
