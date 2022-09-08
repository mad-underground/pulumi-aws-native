// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::RedshiftServerless::Workgroup Resource Type
 */
export function getWorkgroup(args: GetWorkgroupArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkgroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:redshiftserverless:getWorkgroup", {
        "workgroupName": args.workgroupName,
    }, opts);
}

export interface GetWorkgroupArgs {
    workgroupName: string;
}

export interface GetWorkgroupResult {
    readonly workgroup?: outputs.redshiftserverless.Workgroup;
}

export function getWorkgroupOutput(args: GetWorkgroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkgroupResult> {
    return pulumi.output(args).apply(a => getWorkgroup(a, opts))
}

export interface GetWorkgroupOutputArgs {
    workgroupName: pulumi.Input<string>;
}
