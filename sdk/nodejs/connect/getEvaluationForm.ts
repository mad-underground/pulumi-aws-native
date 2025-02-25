// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::EvaluationForm
 */
export function getEvaluationForm(args: GetEvaluationFormArgs, opts?: pulumi.InvokeOptions): Promise<GetEvaluationFormResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:connect:getEvaluationForm", {
        "evaluationFormArn": args.evaluationFormArn,
    }, opts);
}

export interface GetEvaluationFormArgs {
    /**
     * The Amazon Resource Name (ARN) for the evaluation form.
     */
    evaluationFormArn: string;
}

export interface GetEvaluationFormResult {
    /**
     * The description of the evaluation form.
     */
    readonly description?: string;
    /**
     * The Amazon Resource Name (ARN) for the evaluation form.
     */
    readonly evaluationFormArn?: string;
    /**
     * The Amazon Resource Name (ARN) of the instance.
     */
    readonly instanceArn?: string;
    /**
     * The list of evaluation form items.
     */
    readonly items?: outputs.connect.EvaluationFormBaseItem[];
    /**
     * The scoring strategy.
     */
    readonly scoringStrategy?: outputs.connect.EvaluationFormScoringStrategy;
    /**
     * The status of the evaluation form.
     */
    readonly status?: enums.connect.EvaluationFormStatus;
    /**
     * One or more tags.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The title of the evaluation form.
     */
    readonly title?: string;
}
/**
 * Resource Type definition for AWS::Connect::EvaluationForm
 */
export function getEvaluationFormOutput(args: GetEvaluationFormOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEvaluationFormResult> {
    return pulumi.output(args).apply((a: any) => getEvaluationForm(a, opts))
}

export interface GetEvaluationFormOutputArgs {
    /**
     * The Amazon Resource Name (ARN) for the evaluation form.
     */
    evaluationFormArn: pulumi.Input<string>;
}
