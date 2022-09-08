// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Version
 */
export function getVersion(args: GetVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lambda:getVersion", {
        "id": args.id,
    }, opts);
}

export interface GetVersionArgs {
    id: string;
}

export interface GetVersionResult {
    readonly codeSha256?: string;
    readonly description?: string;
    readonly id?: string;
    readonly provisionedConcurrencyConfig?: outputs.lambda.VersionProvisionedConcurrencyConfiguration;
    readonly version?: string;
}

export function getVersionOutput(args: GetVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVersionResult> {
    return pulumi.output(args).apply(a => getVersion(a, opts))
}

export interface GetVersionOutputArgs {
    id: pulumi.Input<string>;
}
