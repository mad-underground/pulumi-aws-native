// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Kendra DataSource
 */
export function getDataSource(args: GetDataSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:kendra:getDataSource", {
        "id": args.id,
        "indexId": args.indexId,
    }, opts);
}

export interface GetDataSourceArgs {
    id: string;
    indexId: string;
}

export interface GetDataSourceResult {
    readonly arn?: string;
    readonly customDocumentEnrichmentConfiguration?: outputs.kendra.DataSourceCustomDocumentEnrichmentConfiguration;
    readonly dataSourceConfiguration?: outputs.kendra.DataSourceConfiguration;
    readonly description?: string;
    readonly id?: string;
    readonly indexId?: string;
    readonly name?: string;
    readonly roleArn?: string;
    readonly schedule?: string;
    /**
     * Tags for labeling the data source
     */
    readonly tags?: outputs.kendra.DataSourceTag[];
}

export function getDataSourceOutput(args: GetDataSourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceResult> {
    return pulumi.output(args).apply(a => getDataSource(a, opts))
}

export interface GetDataSourceOutputArgs {
    id: pulumi.Input<string>;
    indexId: pulumi.Input<string>;
}