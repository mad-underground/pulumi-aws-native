// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationAzureBlob.
 */
export function getLocationAzureBlob(args: GetLocationAzureBlobArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationAzureBlobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:datasync:getLocationAzureBlob", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationAzureBlobArgs {
    /**
     * The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
     */
    locationArn: string;
}

export interface GetLocationAzureBlobResult {
    /**
     * The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.
     */
    readonly agentArns?: string[];
    /**
     * Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.
     */
    readonly azureAccessTier?: enums.datasync.LocationAzureBlobAzureAccessTier;
    /**
     * The specific authentication type that you want DataSync to use to access your Azure Blob Container.
     */
    readonly azureBlobAuthenticationType?: enums.datasync.LocationAzureBlobAzureBlobAuthenticationType;
    /**
     * Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.
     */
    readonly azureBlobType?: enums.datasync.LocationAzureBlobAzureBlobType;
    /**
     * The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the Azure Blob Location that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationAzureBlobTag[];
}
/**
 * Resource schema for AWS::DataSync::LocationAzureBlob.
 */
export function getLocationAzureBlobOutput(args: GetLocationAzureBlobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationAzureBlobResult> {
    return pulumi.output(args).apply((a: any) => getLocationAzureBlob(a, opts))
}

export interface GetLocationAzureBlobOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
     */
    locationArn: pulumi.Input<string>;
}
