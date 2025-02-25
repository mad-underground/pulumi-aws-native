// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::FSx::StorageVirtualMachine
 */
export function getStorageVirtualMachine(args: GetStorageVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageVirtualMachineResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:fsx:getStorageVirtualMachine", {
        "storageVirtualMachineId": args.storageVirtualMachineId,
    }, opts);
}

export interface GetStorageVirtualMachineArgs {
    storageVirtualMachineId: string;
}

export interface GetStorageVirtualMachineResult {
    readonly activeDirectoryConfiguration?: outputs.fsx.StorageVirtualMachineActiveDirectoryConfiguration;
    readonly resourceArn?: string;
    readonly storageVirtualMachineId?: string;
    readonly svmAdminPassword?: string;
    readonly tags?: outputs.Tag[];
    readonly uuid?: string;
}
/**
 * Resource Type definition for AWS::FSx::StorageVirtualMachine
 */
export function getStorageVirtualMachineOutput(args: GetStorageVirtualMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStorageVirtualMachineResult> {
    return pulumi.output(args).apply((a: any) => getStorageVirtualMachine(a, opts))
}

export interface GetStorageVirtualMachineOutputArgs {
    storageVirtualMachineId: pulumi.Input<string>;
}
