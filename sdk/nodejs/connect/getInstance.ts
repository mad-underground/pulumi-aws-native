// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::Instance
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:connect:getInstance", {
        "arn": args.arn,
    }, opts);
}

export interface GetInstanceArgs {
    /**
     * An instanceArn is automatically generated on creation based on instanceId.
     */
    arn: string;
}

export interface GetInstanceResult {
    /**
     * An instanceArn is automatically generated on creation based on instanceId.
     */
    readonly arn?: string;
    /**
     * The attributes for the instance.
     */
    readonly attributes?: outputs.connect.InstanceAttributes;
    /**
     * Timestamp of instance creation logged as part of instance creation.
     */
    readonly createdTime?: string;
    /**
     * An instanceId is automatically generated on creation and assigned as the unique identifier.
     */
    readonly id?: string;
    /**
     * Specifies the creation status of new instance.
     */
    readonly instanceStatus?: enums.connect.InstanceStatus;
    /**
     * Service linked role created as part of instance creation.
     */
    readonly serviceRole?: string;
}

export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply(a => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    /**
     * An instanceArn is automatically generated on creation based on instanceId.
     */
    arn: pulumi.Input<string>;
}
