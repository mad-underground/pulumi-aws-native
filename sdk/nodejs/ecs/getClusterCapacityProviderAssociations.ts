// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Associate a set of ECS Capacity Providers with a specified ECS Cluster
 */
export function getClusterCapacityProviderAssociations(args: GetClusterCapacityProviderAssociationsArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterCapacityProviderAssociationsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecs:getClusterCapacityProviderAssociations", {
        "cluster": args.cluster,
    }, opts);
}

export interface GetClusterCapacityProviderAssociationsArgs {
    cluster: string;
}

export interface GetClusterCapacityProviderAssociationsResult {
    readonly capacityProviders?: (enums.ecs.ClusterCapacityProviderAssociationsCapacityProvider | string)[];
    readonly defaultCapacityProviderStrategy?: outputs.ecs.ClusterCapacityProviderAssociationsCapacityProviderStrategy[];
}

export function getClusterCapacityProviderAssociationsOutput(args: GetClusterCapacityProviderAssociationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterCapacityProviderAssociationsResult> {
    return pulumi.output(args).apply(a => getClusterCapacityProviderAssociations(a, opts))
}

export interface GetClusterCapacityProviderAssociationsOutputArgs {
    cluster: pulumi.Input<string>;
}