// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::InstanceStorageConfig
 */
export function getInstanceStorageConfig(args: GetInstanceStorageConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceStorageConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:connect:getInstanceStorageConfig", {
        "associationId": args.associationId,
        "instanceArn": args.instanceArn,
        "resourceType": args.resourceType,
    }, opts);
}

export interface GetInstanceStorageConfigArgs {
    associationId: string;
    /**
     * Connect Instance ID with which the storage config will be associated
     */
    instanceArn: string;
    resourceType: enums.connect.InstanceStorageConfigInstanceStorageResourceType;
}

export interface GetInstanceStorageConfigResult {
    readonly associationId?: string;
    readonly kinesisFirehoseConfig?: outputs.connect.InstanceStorageConfigKinesisFirehoseConfig;
    readonly kinesisStreamConfig?: outputs.connect.InstanceStorageConfigKinesisStreamConfig;
    readonly kinesisVideoStreamConfig?: outputs.connect.InstanceStorageConfigKinesisVideoStreamConfig;
    readonly s3Config?: outputs.connect.InstanceStorageConfigS3Config;
    readonly storageType?: enums.connect.InstanceStorageConfigStorageType;
}

export function getInstanceStorageConfigOutput(args: GetInstanceStorageConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceStorageConfigResult> {
    return pulumi.output(args).apply(a => getInstanceStorageConfig(a, opts))
}

export interface GetInstanceStorageConfigOutputArgs {
    associationId: pulumi.Input<string>;
    /**
     * Connect Instance ID with which the storage config will be associated
     */
    instanceArn: pulumi.Input<string>;
    resourceType: pulumi.Input<enums.connect.InstanceStorageConfigInstanceStorageResourceType>;
}