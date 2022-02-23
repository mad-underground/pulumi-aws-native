// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Aws Route53 Recovery Readiness Check Schema and API specification.
 */
export function getReadinessCheck(args: GetReadinessCheckArgs, opts?: pulumi.InvokeOptions): Promise<GetReadinessCheckResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53recoveryreadiness:getReadinessCheck", {
        "readinessCheckName": args.readinessCheckName,
    }, opts);
}

export interface GetReadinessCheckArgs {
    /**
     * Name of the ReadinessCheck to create.
     */
    readinessCheckName: string;
}

export interface GetReadinessCheckResult {
    /**
     * The Amazon Resource Name (ARN) of the readiness check.
     */
    readonly readinessCheckArn?: string;
    /**
     * The name of the resource set to check.
     */
    readonly resourceSetName?: string;
    /**
     * A collection of tags associated with a resource.
     */
    readonly tags?: outputs.route53recoveryreadiness.ReadinessCheckTag[];
}

export function getReadinessCheckOutput(args: GetReadinessCheckOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReadinessCheckResult> {
    return pulumi.output(args).apply(a => getReadinessCheck(a, opts))
}

export interface GetReadinessCheckOutputArgs {
    /**
     * Name of the ReadinessCheck to create.
     */
    readinessCheckName: pulumi.Input<string>;
}