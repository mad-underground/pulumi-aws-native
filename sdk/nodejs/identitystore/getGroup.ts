// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IdentityStore::Group
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:identitystore:getGroup", {
        "groupId": args.groupId,
        "identityStoreId": args.identityStoreId,
    }, opts);
}

export interface GetGroupArgs {
    /**
     * The unique identifier for a group in the identity store.
     */
    groupId: string;
    /**
     * The globally unique identifier for the identity store.
     */
    identityStoreId: string;
}

export interface GetGroupResult {
    /**
     * A string containing the description of the group.
     */
    readonly description?: string;
    /**
     * A string containing the name of the group. This value is commonly displayed when the group is referenced.
     */
    readonly displayName?: string;
    /**
     * The unique identifier for a group in the identity store.
     */
    readonly groupId?: string;
}
/**
 * Resource Type definition for AWS::IdentityStore::Group
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    /**
     * The unique identifier for a group in the identity store.
     */
    groupId: pulumi.Input<string>;
    /**
     * The globally unique identifier for the identity store.
     */
    identityStoreId: pulumi.Input<string>;
}
