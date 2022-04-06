// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Personalize::Dataset.
 */
export function getDataset(args: GetDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:personalize:getDataset", {
        "datasetArn": args.datasetArn,
    }, opts);
}

export interface GetDatasetArgs {
    /**
     * The ARN of the dataset
     */
    datasetArn: string;
}

export interface GetDatasetResult {
    /**
     * The ARN of the dataset
     */
    readonly datasetArn?: string;
    readonly datasetImportJob?: outputs.personalize.DatasetImportJob;
}

export function getDatasetOutput(args: GetDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatasetResult> {
    return pulumi.output(args).apply(a => getDataset(a, opts))
}

export interface GetDatasetOutputArgs {
    /**
     * The ARN of the dataset
     */
    datasetArn: pulumi.Input<string>;
}