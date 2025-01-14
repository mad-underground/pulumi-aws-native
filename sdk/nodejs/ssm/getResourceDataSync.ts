// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::ResourceDataSync
 */
export function getResourceDataSync(args: GetResourceDataSyncArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceDataSyncResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ssm:getResourceDataSync", {
        "syncName": args.syncName,
    }, opts);
}

export interface GetResourceDataSyncArgs {
    syncName: string;
}

export interface GetResourceDataSyncResult {
    readonly syncSource?: outputs.ssm.ResourceDataSyncSyncSource;
}
/**
 * Resource Type definition for AWS::SSM::ResourceDataSync
 */
export function getResourceDataSyncOutput(args: GetResourceDataSyncOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceDataSyncResult> {
    return pulumi.output(args).apply((a: any) => getResourceDataSync(a, opts))
}

export interface GetResourceDataSyncOutputArgs {
    syncName: pulumi.Input<string>;
}
