// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for Greengrass V2 deployment.
 */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:greengrassv2:getDeployment", {
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetDeploymentArgs {
    deploymentId: string;
}

export interface GetDeploymentResult {
    readonly deploymentId?: string;
    readonly tags?: any;
}

export function getDeploymentOutput(args: GetDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentResult> {
    return pulumi.output(args).apply(a => getDeployment(a, opts))
}

export interface GetDeploymentOutputArgs {
    deploymentId: pulumi.Input<string>;
}
