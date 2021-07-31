// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getAzs(args?: GetAzsArgs, opts?: pulumi.InvokeOptions): Promise<GetAzsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws-native:index:getAzs", {
        "region": args.region,
    }, opts);
}

export interface GetAzsArgs {
    region?: string;
}

export interface GetAzsResult {
    readonly azs: string[];
}
