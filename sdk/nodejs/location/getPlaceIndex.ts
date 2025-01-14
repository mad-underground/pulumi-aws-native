// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Location::PlaceIndex Resource Type
 */
export function getPlaceIndex(args: GetPlaceIndexArgs, opts?: pulumi.InvokeOptions): Promise<GetPlaceIndexResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:location:getPlaceIndex", {
        "indexName": args.indexName,
    }, opts);
}

export interface GetPlaceIndexArgs {
    indexName: string;
}

export interface GetPlaceIndexResult {
    readonly arn?: string;
    readonly createTime?: string;
    readonly dataSourceConfiguration?: outputs.location.PlaceIndexDataSourceConfiguration;
    readonly description?: string;
    readonly indexArn?: string;
    readonly pricingPlan?: enums.location.PlaceIndexPricingPlan;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
    readonly updateTime?: string;
}
/**
 * Definition of AWS::Location::PlaceIndex Resource Type
 */
export function getPlaceIndexOutput(args: GetPlaceIndexOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlaceIndexResult> {
    return pulumi.output(args).apply((a: any) => getPlaceIndex(a, opts))
}

export interface GetPlaceIndexOutputArgs {
    indexName: pulumi.Input<string>;
}
