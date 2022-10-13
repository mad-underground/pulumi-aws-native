// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::CodeDeploy::Application resource creates an AWS CodeDeploy application
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:codedeploy:getApplication", {
        "applicationName": args.applicationName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * A name for the application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
     */
    applicationName: string;
}

export interface GetApplicationResult {
    /**
     * The metadata that you apply to CodeDeploy applications to help you organize and categorize them. Each tag consists of a key and an optional value, both of which you define. 
     */
    readonly tags?: outputs.codedeploy.ApplicationTag[];
}

export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    /**
     * A name for the application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
     */
    applicationName: pulumi.Input<string>;
}
