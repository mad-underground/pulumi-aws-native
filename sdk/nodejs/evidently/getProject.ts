// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Project
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:evidently:getProject", {
        "arn": args.arn,
    }, opts);
}

export interface GetProjectArgs {
    arn: string;
}

export interface GetProjectResult {
    readonly arn?: string;
    readonly dataDelivery?: outputs.evidently.ProjectDataDeliveryObject;
    readonly description?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.evidently.ProjectTag[];
}

export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

export interface GetProjectOutputArgs {
    arn: pulumi.Input<string>;
}
