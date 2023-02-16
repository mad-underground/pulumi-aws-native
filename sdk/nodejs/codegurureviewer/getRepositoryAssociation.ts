// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.
 */
export function getRepositoryAssociation(args: GetRepositoryAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:codegurureviewer:getRepositoryAssociation", {
        "associationArn": args.associationArn,
    }, opts);
}

export interface GetRepositoryAssociationArgs {
    /**
     * The Amazon Resource Name (ARN) of the repository association.
     */
    associationArn: string;
}

export interface GetRepositoryAssociationResult {
    /**
     * The Amazon Resource Name (ARN) of the repository association.
     */
    readonly associationArn?: string;
}
/**
 * This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.
 */
export function getRepositoryAssociationOutput(args: GetRepositoryAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryAssociationResult> {
    return pulumi.output(args).apply((a: any) => getRepositoryAssociation(a, opts))
}

export interface GetRepositoryAssociationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the repository association.
     */
    associationArn: pulumi.Input<string>;
}
