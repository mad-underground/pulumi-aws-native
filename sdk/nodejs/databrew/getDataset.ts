// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Dataset.
 */
export function getDataset(args: GetDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:databrew:getDataset", {
        "name": args.name,
    }, opts);
}

export interface GetDatasetArgs {
    /**
     * Dataset name
     */
    name: string;
}

export interface GetDatasetResult {
    /**
     * Dataset format
     */
    readonly format?: enums.databrew.DatasetFormat;
    /**
     * Format options for dataset
     */
    readonly formatOptions?: outputs.databrew.DatasetFormatOptions;
    /**
     * Input
     */
    readonly input?: outputs.databrew.DatasetInput;
    /**
     * PathOptions
     */
    readonly pathOptions?: outputs.databrew.DatasetPathOptions;
}
/**
 * Resource schema for AWS::DataBrew::Dataset.
 */
export function getDatasetOutput(args: GetDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatasetResult> {
    return pulumi.output(args).apply((a: any) => getDataset(a, opts))
}

export interface GetDatasetOutputArgs {
    /**
     * Dataset name
     */
    name: pulumi.Input<string>;
}
