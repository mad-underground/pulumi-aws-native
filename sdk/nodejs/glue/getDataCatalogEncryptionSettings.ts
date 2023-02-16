// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::DataCatalogEncryptionSettings
 */
export function getDataCatalogEncryptionSettings(args: GetDataCatalogEncryptionSettingsArgs, opts?: pulumi.InvokeOptions): Promise<GetDataCatalogEncryptionSettingsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:glue:getDataCatalogEncryptionSettings", {
        "id": args.id,
    }, opts);
}

export interface GetDataCatalogEncryptionSettingsArgs {
    id: string;
}

export interface GetDataCatalogEncryptionSettingsResult {
    readonly dataCatalogEncryptionSettings?: outputs.glue.DataCatalogEncryptionSettings;
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::Glue::DataCatalogEncryptionSettings
 */
export function getDataCatalogEncryptionSettingsOutput(args: GetDataCatalogEncryptionSettingsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataCatalogEncryptionSettingsResult> {
    return pulumi.output(args).apply((a: any) => getDataCatalogEncryptionSettings(a, opts))
}

export interface GetDataCatalogEncryptionSettingsOutputArgs {
    id: pulumi.Input<string>;
}
