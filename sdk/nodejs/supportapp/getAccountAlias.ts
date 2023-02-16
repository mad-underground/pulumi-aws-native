// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.
 */
export function getAccountAlias(args: GetAccountAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountAliasResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:supportapp:getAccountAlias", {
        "accountAliasResourceId": args.accountAliasResourceId,
    }, opts);
}

export interface GetAccountAliasArgs {
    /**
     * Unique identifier representing an alias tied to an account
     */
    accountAliasResourceId: string;
}

export interface GetAccountAliasResult {
    /**
     * An account alias associated with a customer's account.
     */
    readonly accountAlias?: string;
    /**
     * Unique identifier representing an alias tied to an account
     */
    readonly accountAliasResourceId?: string;
}
/**
 * An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.
 */
export function getAccountAliasOutput(args: GetAccountAliasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountAliasResult> {
    return pulumi.output(args).apply((a: any) => getAccountAlias(a, opts))
}

export interface GetAccountAliasOutputArgs {
    /**
     * Unique identifier representing an alias tied to an account
     */
    accountAliasResourceId: pulumi.Input<string>;
}
