// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Logs::LogGroup
 */
export function getLogGroup(args: GetLogGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetLogGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:logs:getLogGroup", {
        "logGroupName": args.logGroupName,
    }, opts);
}

export interface GetLogGroupArgs {
    /**
     * The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
     */
    logGroupName: string;
}

export interface GetLogGroupResult {
    /**
     * The CloudWatch log group ARN.
     */
    readonly arn?: string;
    /**
     * The body of the policy document you want to use for this topic.
     *
     * You can only add one policy per topic.
     *
     * The policy must be in JSON string format.
     *
     * Length Constraints: Maximum length of 30720
     */
    readonly dataProtectionPolicy?: any;
    /**
     * The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
     */
    readonly kmsKeyId?: string;
    /**
     * The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.
     */
    readonly retentionInDays?: number;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.logs.LogGroupTag[];
}

export function getLogGroupOutput(args: GetLogGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogGroupResult> {
    return pulumi.output(args).apply(a => getLogGroup(a, opts))
}

export interface GetLogGroupOutputArgs {
    /**
     * The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
     */
    logGroupName: pulumi.Input<string>;
}
