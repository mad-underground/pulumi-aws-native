// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::OpsWorks::Layer
 */
export function getLayer(args: GetLayerArgs, opts?: pulumi.InvokeOptions): Promise<GetLayerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:opsworks:getLayer", {
        "id": args.id,
    }, opts);
}

export interface GetLayerArgs {
    id: string;
}

export interface GetLayerResult {
    readonly attributes?: {[key: string]: string};
    readonly autoAssignElasticIps?: boolean;
    readonly autoAssignPublicIps?: boolean;
    readonly customInstanceProfileArn?: string;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::OpsWorks::Layer` for more information about the expected schema for this property.
     */
    readonly customJson?: any;
    readonly customRecipes?: outputs.opsworks.LayerRecipes;
    readonly customSecurityGroupIds?: string[];
    readonly enableAutoHealing?: boolean;
    readonly id?: string;
    readonly installUpdatesOnBoot?: boolean;
    readonly lifecycleEventConfiguration?: outputs.opsworks.LayerLifecycleEventConfiguration;
    readonly loadBasedAutoScaling?: outputs.opsworks.LayerLoadBasedAutoScaling;
    readonly name?: string;
    readonly packages?: string[];
    readonly shortname?: string;
    readonly tags?: outputs.Tag[];
    readonly useEbsOptimizedInstances?: boolean;
    readonly volumeConfigurations?: outputs.opsworks.LayerVolumeConfiguration[];
}
/**
 * Resource Type definition for AWS::OpsWorks::Layer
 */
export function getLayerOutput(args: GetLayerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLayerResult> {
    return pulumi.output(args).apply((a: any) => getLayer(a, opts))
}

export interface GetLayerOutputArgs {
    id: pulumi.Input<string>;
}
