// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::IoTFleetWise::Campaign Resource Type
 */
export function getCampaign(args: GetCampaignArgs, opts?: pulumi.InvokeOptions): Promise<GetCampaignResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotfleetwise:getCampaign", {
        "name": args.name,
    }, opts);
}

export interface GetCampaignArgs {
    name: string;
}

export interface GetCampaignResult {
    readonly arn?: string;
    readonly creationTime?: string;
    readonly dataExtraDimensions?: string[];
    readonly description?: string;
    readonly lastModificationTime?: string;
    readonly signalsToCollect?: outputs.iotfleetwise.CampaignSignalInformation[];
    readonly status?: enums.iotfleetwise.CampaignStatus;
    readonly tags?: outputs.iotfleetwise.CampaignTag[];
}

export function getCampaignOutput(args: GetCampaignOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCampaignResult> {
    return pulumi.output(args).apply(a => getCampaign(a, opts))
}

export interface GetCampaignOutputArgs {
    name: pulumi.Input<string>;
}