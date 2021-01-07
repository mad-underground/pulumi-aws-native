// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getSsmParameterString(args: GetSsmParameterStringArgs, opts?: pulumi.InvokeOptions): Promise<GetSsmParameterStringResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("cloudformation:index:getSsmParameterString", {
        "name": args.name,
    }, opts);
}

export interface GetSsmParameterStringArgs {
    readonly name: string;
}

export interface GetSsmParameterStringResult {
    readonly value: string;
}