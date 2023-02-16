// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelCard.
 */
export function getModelCard(args: GetModelCardArgs, opts?: pulumi.InvokeOptions): Promise<GetModelCardResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getModelCard", {
        "modelCardName": args.modelCardName,
    }, opts);
}

export interface GetModelCardArgs {
    /**
     * The unique name of the model card.
     */
    modelCardName: string;
}

export interface GetModelCardResult {
    readonly content?: outputs.sagemaker.ModelCardContent;
    /**
     * Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
     */
    readonly createdBy?: outputs.sagemaker.ModelCardUserContext;
    /**
     * The date and time the model card was created.
     */
    readonly creationTime?: string;
    /**
     * Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
     */
    readonly lastModifiedBy?: outputs.sagemaker.ModelCardUserContext;
    /**
     * The date and time the model card was last modified.
     */
    readonly lastModifiedTime?: string;
    /**
     * The Amazon Resource Name (ARN) of the successfully created model card.
     */
    readonly modelCardArn?: string;
    /**
     * The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
     */
    readonly modelCardProcessingStatus?: enums.sagemaker.ModelCardProcessingStatus;
    /**
     * The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
     */
    readonly modelCardStatus?: enums.sagemaker.ModelCardStatus;
    /**
     * A version of the model card.
     */
    readonly modelCardVersion?: number;
    /**
     * Key-value pairs used to manage metadata for model cards.
     */
    readonly tags?: outputs.sagemaker.ModelCardTag[];
}
/**
 * Resource Type definition for AWS::SageMaker::ModelCard.
 */
export function getModelCardOutput(args: GetModelCardOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelCardResult> {
    return pulumi.output(args).apply((a: any) => getModelCard(a, opts))
}

export interface GetModelCardOutputArgs {
    /**
     * The unique name of the model card.
     */
    modelCardName: pulumi.Input<string>;
}
