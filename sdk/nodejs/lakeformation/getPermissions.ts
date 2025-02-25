// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::LakeFormation::Permissions
 */
export function getPermissions(args: GetPermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lakeformation:getPermissions", {
        "id": args.id,
    }, opts);
}

export interface GetPermissionsArgs {
    id: string;
}

export interface GetPermissionsResult {
    readonly id?: string;
    readonly permissions?: string[];
    readonly permissionsWithGrantOption?: string[];
}
/**
 * Resource Type definition for AWS::LakeFormation::Permissions
 */
export function getPermissionsOutput(args: GetPermissionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPermissionsResult> {
    return pulumi.output(args).apply((a: any) => getPermissions(a, opts))
}

export interface GetPermissionsOutputArgs {
    id: pulumi.Input<string>;
}
