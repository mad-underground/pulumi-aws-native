// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getExtensionAssociation(args: GetExtensionAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetExtensionAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appconfig:getExtensionAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetExtensionAssociationArgs {
    id: string;
}

export interface GetExtensionAssociationResult {
    readonly arn?: string;
    readonly extensionArn?: string;
    readonly id?: string;
    readonly parameters?: {[key: string]: string};
    readonly resourceArn?: string;
}
/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export function getExtensionAssociationOutput(args: GetExtensionAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExtensionAssociationResult> {
    return pulumi.output(args).apply((a: any) => getExtensionAssociation(a, opts))
}

export interface GetExtensionAssociationOutputArgs {
    id: pulumi.Input<string>;
}
