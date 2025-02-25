// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Location::GeofenceCollection Resource Type
 */
export function getGeofenceCollection(args: GetGeofenceCollectionArgs, opts?: pulumi.InvokeOptions): Promise<GetGeofenceCollectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:location:getGeofenceCollection", {
        "collectionName": args.collectionName,
    }, opts);
}

export interface GetGeofenceCollectionArgs {
    collectionName: string;
}

export interface GetGeofenceCollectionResult {
    readonly arn?: string;
    readonly collectionArn?: string;
    readonly createTime?: string;
    readonly description?: string;
    readonly pricingPlan?: enums.location.GeofenceCollectionPricingPlan;
    readonly pricingPlanDataSource?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    readonly updateTime?: string;
}
/**
 * Definition of AWS::Location::GeofenceCollection Resource Type
 */
export function getGeofenceCollectionOutput(args: GetGeofenceCollectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGeofenceCollectionResult> {
    return pulumi.output(args).apply((a: any) => getGeofenceCollection(a, opts))
}

export interface GetGeofenceCollectionOutputArgs {
    collectionName: pulumi.Input<string>;
}
