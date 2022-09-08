// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Template Resource Type.
 */
export function getTemplate(args: GetTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:quicksight:getTemplate", {
        "awsAccountId": args.awsAccountId,
        "templateId": args.templateId,
    }, opts);
}

export interface GetTemplateArgs {
    awsAccountId: string;
    templateId: string;
}

export interface GetTemplateResult {
    /**
     * <p>The Amazon Resource Name (ARN) of the template.</p>
     */
    readonly arn?: string;
    /**
     * <p>A display name for the template.</p>
     */
    readonly name?: string;
    /**
     * <p>A list of resource permissions to be set on the template. </p>
     */
    readonly permissions?: outputs.quicksight.TemplateResourcePermission[];
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
     */
    readonly tags?: outputs.quicksight.TemplateTag[];
}

export function getTemplateOutput(args: GetTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTemplateResult> {
    return pulumi.output(args).apply(a => getTemplate(a, opts))
}

export interface GetTemplateOutputArgs {
    awsAccountId: pulumi.Input<string>;
    templateId: pulumi.Input<string>;
}
