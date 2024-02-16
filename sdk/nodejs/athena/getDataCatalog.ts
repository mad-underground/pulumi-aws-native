// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Athena::DataCatalog
 */
export function getDataCatalog(args: GetDataCatalogArgs, opts?: pulumi.InvokeOptions): Promise<GetDataCatalogResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:athena:getDataCatalog", {
        "name": args.name,
    }, opts);
}

export interface GetDataCatalogArgs {
    /**
     * The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. 
     */
    name: string;
}

export interface GetDataCatalogResult {
    /**
     * A description of the data catalog to be created. 
     */
    readonly description?: string;
    /**
     * Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type. 
     */
    readonly parameters?: {[key: string]: string};
    /**
     * A list of comma separated tags to add to the data catalog that is created. 
     */
    readonly tags?: outputs.athena.DataCatalogTag[];
    /**
     * The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. 
     */
    readonly type?: enums.athena.DataCatalogType;
}
/**
 * Resource schema for AWS::Athena::DataCatalog
 */
export function getDataCatalogOutput(args: GetDataCatalogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataCatalogResult> {
    return pulumi.output(args).apply((a: any) => getDataCatalog(a, opts))
}

export interface GetDataCatalogOutputArgs {
    /**
     * The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. 
     */
    name: pulumi.Input<string>;
}
