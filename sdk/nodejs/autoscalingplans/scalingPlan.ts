// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AutoScalingPlans::ScalingPlan
 *
 * @deprecated ScalingPlan is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ScalingPlan extends pulumi.CustomResource {
    /**
     * Get an existing ScalingPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScalingPlan {
        pulumi.log.warn("ScalingPlan is deprecated: ScalingPlan is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ScalingPlan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:autoscalingplans:ScalingPlan';

    /**
     * Returns true if the given object is an instance of ScalingPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingPlan.__pulumiType;
    }

    public readonly applicationSource!: pulumi.Output<outputs.autoscalingplans.ScalingPlanApplicationSource>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly scalingInstructions!: pulumi.Output<outputs.autoscalingplans.ScalingPlanScalingInstruction[]>;
    public /*out*/ readonly scalingPlanName!: pulumi.Output<string>;
    public /*out*/ readonly scalingPlanVersion!: pulumi.Output<string>;

    /**
     * Create a ScalingPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ScalingPlan is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ScalingPlanArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ScalingPlan is deprecated: ScalingPlan is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationSource'");
            }
            if ((!args || args.scalingInstructions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingInstructions'");
            }
            resourceInputs["applicationSource"] = args ? args.applicationSource : undefined;
            resourceInputs["scalingInstructions"] = args ? args.scalingInstructions : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["scalingPlanName"] = undefined /*out*/;
            resourceInputs["scalingPlanVersion"] = undefined /*out*/;
        } else {
            resourceInputs["applicationSource"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["scalingInstructions"] = undefined /*out*/;
            resourceInputs["scalingPlanName"] = undefined /*out*/;
            resourceInputs["scalingPlanVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingPlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScalingPlan resource.
 */
export interface ScalingPlanArgs {
    applicationSource: pulumi.Input<inputs.autoscalingplans.ScalingPlanApplicationSourceArgs>;
    scalingInstructions: pulumi.Input<pulumi.Input<inputs.autoscalingplans.ScalingPlanScalingInstructionArgs>[]>;
}
