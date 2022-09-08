// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::TaskTemplate.
 */
export function getTaskTemplate(args: GetTaskTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:connect:getTaskTemplate", {
        "arn": args.arn,
    }, opts);
}

export interface GetTaskTemplateArgs {
    /**
     * The identifier (arn) of the task template.
     */
    arn: string;
}

export interface GetTaskTemplateResult {
    /**
     * The identifier (arn) of the task template.
     */
    readonly arn?: string;
    readonly clientToken?: string;
    /**
     * The constraints for the task template
     */
    readonly constraints?: outputs.connect.ConstraintsProperties;
    /**
     * The identifier of the contact flow.
     */
    readonly contactFlowArn?: string;
    readonly defaults?: outputs.connect.TaskTemplateDefaultFieldValue[];
    /**
     * The description of the task template.
     */
    readonly description?: string;
    /**
     * The list of task template's fields
     */
    readonly fields?: outputs.connect.TaskTemplateField[];
    /**
     * The identifier (arn) of the instance.
     */
    readonly instanceArn?: string;
    /**
     * The name of the task template.
     */
    readonly name?: string;
    readonly status?: enums.connect.TaskTemplateStatus;
    /**
     * One or more tags.
     */
    readonly tags?: outputs.connect.TaskTemplateTag[];
}

export function getTaskTemplateOutput(args: GetTaskTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskTemplateResult> {
    return pulumi.output(args).apply(a => getTaskTemplate(a, opts))
}

export interface GetTaskTemplateOutputArgs {
    /**
     * The identifier (arn) of the task template.
     */
    arn: pulumi.Input<string>;
}
