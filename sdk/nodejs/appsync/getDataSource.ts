// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::DataSource
 */
export function getDataSource(args: GetDataSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appsync:getDataSource", {
        "id": args.id,
    }, opts);
}

export interface GetDataSourceArgs {
    id: string;
}

export interface GetDataSourceResult {
    readonly dataSourceArn?: string;
    readonly description?: string;
    readonly dynamoDBConfig?: outputs.appsync.DataSourceDynamoDBConfig;
    readonly elasticsearchConfig?: outputs.appsync.DataSourceElasticsearchConfig;
    readonly httpConfig?: outputs.appsync.DataSourceHttpConfig;
    readonly id?: string;
    readonly lambdaConfig?: outputs.appsync.DataSourceLambdaConfig;
    readonly openSearchServiceConfig?: outputs.appsync.DataSourceOpenSearchServiceConfig;
    readonly relationalDatabaseConfig?: outputs.appsync.DataSourceRelationalDatabaseConfig;
    readonly serviceRoleArn?: string;
    readonly type?: string;
}

export function getDataSourceOutput(args: GetDataSourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceResult> {
    return pulumi.output(args).apply(a => getDataSource(a, opts))
}

export interface GetDataSourceOutputArgs {
    id: pulumi.Input<string>;
}
