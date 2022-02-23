// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A Kendra index
 */
export function getIndex(args: GetIndexArgs, opts?: pulumi.InvokeOptions): Promise<GetIndexResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:kendra:getIndex", {
        "id": args.id,
    }, opts);
}

export interface GetIndexArgs {
    id: string;
}

export interface GetIndexResult {
    readonly arn?: string;
    /**
     * Capacity units
     */
    readonly capacityUnits?: outputs.kendra.IndexCapacityUnitsConfiguration;
    /**
     * A description for the index
     */
    readonly description?: string;
    /**
     * Document metadata configurations
     */
    readonly documentMetadataConfigurations?: outputs.kendra.IndexDocumentMetadataConfiguration[];
    readonly id?: string;
    readonly name?: string;
    readonly roleArn?: string;
    /**
     * Tags for labeling the index
     */
    readonly tags?: outputs.kendra.IndexTag[];
    readonly userContextPolicy?: enums.kendra.IndexUserContextPolicy;
    readonly userTokenConfigurations?: outputs.kendra.IndexUserTokenConfiguration[];
}

export function getIndexOutput(args: GetIndexOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIndexResult> {
    return pulumi.output(args).apply(a => getIndex(a, opts))
}

export interface GetIndexOutputArgs {
    id: pulumi.Input<string>;
}