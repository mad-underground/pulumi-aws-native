// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Asset
 */
export function getAsset(args: GetAssetArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotsitewise:getAsset", {
        "assetId": args.assetId,
    }, opts);
}

export interface GetAssetArgs {
    /**
     * The ID of the asset
     */
    assetId: string;
}

export interface GetAssetResult {
    /**
     * The ARN of the asset
     */
    readonly assetArn?: string;
    /**
     * A description for the asset
     */
    readonly assetDescription?: string;
    readonly assetHierarchies?: outputs.iotsitewise.AssetHierarchy[];
    /**
     * The ID of the asset
     */
    readonly assetId?: string;
    /**
     * The ID of the asset model from which to create the asset.
     */
    readonly assetModelId?: string;
    /**
     * A unique, friendly name for the asset.
     */
    readonly assetName?: string;
    readonly assetProperties?: outputs.iotsitewise.AssetProperty[];
    /**
     * A list of key-value pairs that contain metadata for the asset.
     */
    readonly tags?: outputs.iotsitewise.AssetTag[];
}
/**
 * Resource schema for AWS::IoTSiteWise::Asset
 */
export function getAssetOutput(args: GetAssetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssetResult> {
    return pulumi.output(args).apply((a: any) => getAsset(a, opts))
}

export interface GetAssetOutputArgs {
    /**
     * The ID of the asset
     */
    assetId: pulumi.Input<string>;
}
