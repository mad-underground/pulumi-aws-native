// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::UserProfile
 */
export function getUserProfile(args: GetUserProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetUserProfileResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getUserProfile", {
        "domainId": args.domainId,
        "userProfileName": args.userProfileName,
    }, opts);
}

export interface GetUserProfileArgs {
    /**
     * The ID of the associated Domain.
     */
    domainId: string;
    /**
     * A name for the UserProfile.
     */
    userProfileName: string;
}

export interface GetUserProfileResult {
    /**
     * The user profile Amazon Resource Name (ARN).
     */
    readonly userProfileArn?: string;
    /**
     * A collection of settings.
     */
    readonly userSettings?: outputs.sagemaker.UserProfileUserSettings;
}

export function getUserProfileOutput(args: GetUserProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserProfileResult> {
    return pulumi.output(args).apply(a => getUserProfile(a, opts))
}

export interface GetUserProfileOutputArgs {
    /**
     * The ID of the associated Domain.
     */
    domainId: pulumi.Input<string>;
    /**
     * A name for the UserProfile.
     */
    userProfileName: pulumi.Input<string>;
}
