// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Permission
 */
export function getPermission(args: GetPermissionArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lambda:getPermission", {
        "id": args.id,
    }, opts);
}

export interface GetPermissionArgs {
    id: string;
}

export interface GetPermissionResult {
    readonly id?: string;
}

export function getPermissionOutput(args: GetPermissionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPermissionResult> {
    return pulumi.output(args).apply(a => getPermission(a, opts))
}

export interface GetPermissionOutputArgs {
    id: pulumi.Input<string>;
}