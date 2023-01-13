// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Amazon OpenSearchServerless collection resource
 */
export function getCollection(args: GetCollectionArgs, opts?: pulumi.InvokeOptions): Promise<GetCollectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:opensearchserverless:getCollection", {
        "id": args.id,
    }, opts);
}

export interface GetCollectionArgs {
    /**
     * The identifier of the collection
     */
    id: string;
}

export interface GetCollectionResult {
    /**
     * The Amazon Resource Name (ARN) of the collection.
     */
    readonly arn?: string;
    /**
     * The endpoint for the collection.
     */
    readonly collectionEndpoint?: string;
    /**
     * The OpenSearch Dashboards endpoint for the collection.
     */
    readonly dashboardEndpoint?: string;
    /**
     * The description of the collection
     */
    readonly description?: string;
    /**
     * The identifier of the collection
     */
    readonly id?: string;
}
/**
 * Amazon OpenSearchServerless collection resource
 */
export function getCollectionOutput(args: GetCollectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCollectionResult> {
    return pulumi.output(args).apply((a: any) => getCollection(a, opts))
}

export interface GetCollectionOutputArgs {
    /**
     * The identifier of the collection
     */
    id: pulumi.Input<string>;
}
