// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoTAnalytics::Dataset
 */
export function getDataset(args: GetDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotanalytics:getDataset", {
        "datasetName": args.datasetName,
    }, opts);
}

export interface GetDatasetArgs {
    datasetName: string;
}

export interface GetDatasetResult {
    readonly actions?: outputs.iotanalytics.DatasetAction[];
    readonly contentDeliveryRules?: outputs.iotanalytics.DatasetContentDeliveryRule[];
    readonly id?: string;
    readonly lateDataRules?: outputs.iotanalytics.DatasetLateDataRule[];
    readonly retentionPeriod?: outputs.iotanalytics.DatasetRetentionPeriod;
    readonly tags?: outputs.Tag[];
    readonly triggers?: outputs.iotanalytics.DatasetTrigger[];
    readonly versioningConfiguration?: outputs.iotanalytics.DatasetVersioningConfiguration;
}
/**
 * Resource Type definition for AWS::IoTAnalytics::Dataset
 */
export function getDatasetOutput(args: GetDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatasetResult> {
    return pulumi.output(args).apply((a: any) => getDataset(a, opts))
}

export interface GetDatasetOutputArgs {
    datasetName: pulumi.Input<string>;
}
