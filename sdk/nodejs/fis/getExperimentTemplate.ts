// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export function getExperimentTemplate(args: GetExperimentTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetExperimentTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:fis:getExperimentTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetExperimentTemplateArgs {
    id: string;
}

export interface GetExperimentTemplateResult {
    readonly actions?: {[key: string]: outputs.fis.ExperimentTemplateAction};
    readonly description?: string;
    readonly experimentOptions?: outputs.fis.ExperimentTemplateExperimentOptions;
    readonly id?: string;
    readonly logConfiguration?: outputs.fis.ExperimentTemplateLogConfiguration;
    readonly roleArn?: string;
    readonly stopConditions?: outputs.fis.ExperimentTemplateStopCondition[];
    readonly targets?: {[key: string]: outputs.fis.ExperimentTemplateTarget};
}
/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export function getExperimentTemplateOutput(args: GetExperimentTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExperimentTemplateResult> {
    return pulumi.output(args).apply((a: any) => getExperimentTemplate(a, opts))
}

export interface GetExperimentTemplateOutputArgs {
    id: pulumi.Input<string>;
}
