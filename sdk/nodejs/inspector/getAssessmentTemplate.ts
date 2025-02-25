// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Inspector::AssessmentTemplate
 */
export function getAssessmentTemplate(args: GetAssessmentTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:inspector:getAssessmentTemplate", {
        "arn": args.arn,
    }, opts);
}

export interface GetAssessmentTemplateArgs {
    arn: string;
}

export interface GetAssessmentTemplateResult {
    readonly arn?: string;
}
/**
 * Resource Type definition for AWS::Inspector::AssessmentTemplate
 */
export function getAssessmentTemplateOutput(args: GetAssessmentTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssessmentTemplateResult> {
    return pulumi.output(args).apply((a: any) => getAssessmentTemplate(a, opts))
}

export interface GetAssessmentTemplateOutputArgs {
    arn: pulumi.Input<string>;
}
