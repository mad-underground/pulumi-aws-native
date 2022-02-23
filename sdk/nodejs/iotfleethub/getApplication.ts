// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTFleetHub::Application
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotfleethub:getApplication", {
        "applicationId": args.applicationId,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The ID of the application.
     */
    applicationId: string;
}

export interface GetApplicationResult {
    /**
     * The ARN of the application.
     */
    readonly applicationArn?: string;
    /**
     * When the Application was created
     */
    readonly applicationCreationDate?: number;
    /**
     * Application Description, should be between 1 and 2048 characters.
     */
    readonly applicationDescription?: string;
    /**
     * The ID of the application.
     */
    readonly applicationId?: string;
    /**
     * When the Application was last updated
     */
    readonly applicationLastUpdateDate?: number;
    /**
     * Application Name, should be between 1 and 256 characters.
     */
    readonly applicationName?: string;
    /**
     * The current state of the application.
     */
    readonly applicationState?: string;
    /**
     * The URL of the application.
     */
    readonly applicationUrl?: string;
    /**
     * A message indicating why Create or Delete Application failed.
     */
    readonly errorMessage?: string;
    /**
     * The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
     */
    readonly roleArn?: string;
    /**
     * The AWS SSO application generated client ID (used with AWS SSO APIs).
     */
    readonly ssoClientId?: string;
    /**
     * A list of key-value pairs that contain metadata for the application.
     */
    readonly tags?: outputs.iotfleethub.ApplicationTag[];
}

export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply(a => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    /**
     * The ID of the application.
     */
    applicationId: pulumi.Input<string>;
}