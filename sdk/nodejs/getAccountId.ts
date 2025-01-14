// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAccountId(opts?: pulumi.InvokeOptions): Promise<GetAccountIdResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:index:getAccountId", {
    }, opts);
}

export interface GetAccountIdResult {
    readonly accountId: string;
}
export function getAccountIdOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountIdResult> {
    return pulumi.output(getAccountId(opts))
}
