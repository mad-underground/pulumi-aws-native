// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource represents an individual schema version of a schema defined in Glue Schema Registry.
 */
export function getSchemaVersion(args: GetSchemaVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:glue:getSchemaVersion", {
        "versionId": args.versionId,
    }, opts);
}

export interface GetSchemaVersionArgs {
    /**
     * Represents the version ID associated with the schema version.
     */
    versionId: string;
}

export interface GetSchemaVersionResult {
    /**
     * Represents the version ID associated with the schema version.
     */
    readonly versionId?: string;
}

export function getSchemaVersionOutput(args: GetSchemaVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemaVersionResult> {
    return pulumi.output(args).apply(a => getSchemaVersion(a, opts))
}

export interface GetSchemaVersionOutputArgs {
    /**
     * Represents the version ID associated with the schema version.
     */
    versionId: pulumi.Input<string>;
}
