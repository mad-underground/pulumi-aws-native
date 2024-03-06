// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TrafficMirrorTarget
 *
 * @deprecated TrafficMirrorTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TrafficMirrorTarget extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TrafficMirrorTarget {
        pulumi.log.warn("TrafficMirrorTarget is deprecated: TrafficMirrorTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TrafficMirrorTarget(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TrafficMirrorTarget';

    /**
     * Returns true if the given object is an instance of TrafficMirrorTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorTarget.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly gatewayLoadBalancerEndpointId!: pulumi.Output<string | undefined>;
    public readonly networkInterfaceId!: pulumi.Output<string | undefined>;
    public readonly networkLoadBalancerArn!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a TrafficMirrorTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TrafficMirrorTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: TrafficMirrorTargetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TrafficMirrorTarget is deprecated: TrafficMirrorTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["gatewayLoadBalancerEndpointId"] = args ? args.gatewayLoadBalancerEndpointId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["networkLoadBalancerArn"] = args ? args.networkLoadBalancerArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["gatewayLoadBalancerEndpointId"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["networkLoadBalancerArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "gatewayLoadBalancerEndpointId", "networkInterfaceId", "networkLoadBalancerArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TrafficMirrorTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TrafficMirrorTarget resource.
 */
export interface TrafficMirrorTargetArgs {
    description?: pulumi.Input<string>;
    gatewayLoadBalancerEndpointId?: pulumi.Input<string>;
    networkInterfaceId?: pulumi.Input<string>;
    networkLoadBalancerArn?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
