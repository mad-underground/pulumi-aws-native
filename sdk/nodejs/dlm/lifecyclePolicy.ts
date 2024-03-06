// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DLM::LifecyclePolicy
 *
 * @deprecated LifecyclePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class LifecyclePolicy extends pulumi.CustomResource {
    /**
     * Get an existing LifecyclePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LifecyclePolicy {
        pulumi.log.warn("LifecyclePolicy is deprecated: LifecyclePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new LifecyclePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dlm:LifecyclePolicy';

    /**
     * Returns true if the given object is an instance of LifecyclePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecyclePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecyclePolicy.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly copyTags!: pulumi.Output<boolean | undefined>;
    public readonly createInterval!: pulumi.Output<number | undefined>;
    public readonly crossRegionCopyTargets!: pulumi.Output<outputs.dlm.LifecyclePolicyCrossRegionCopyTargets | undefined>;
    public readonly defaultPolicy!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly exclusions!: pulumi.Output<outputs.dlm.LifecyclePolicyExclusions | undefined>;
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    public readonly extendDeletion!: pulumi.Output<boolean | undefined>;
    public readonly policyDetails!: pulumi.Output<outputs.dlm.LifecyclePolicyPolicyDetails | undefined>;
    public readonly retainInterval!: pulumi.Output<number | undefined>;
    public readonly state!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a LifecyclePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated LifecyclePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: LifecyclePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LifecyclePolicy is deprecated: LifecyclePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["copyTags"] = args ? args.copyTags : undefined;
            resourceInputs["createInterval"] = args ? args.createInterval : undefined;
            resourceInputs["crossRegionCopyTargets"] = args ? args.crossRegionCopyTargets : undefined;
            resourceInputs["defaultPolicy"] = args ? args.defaultPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["exclusions"] = args ? args.exclusions : undefined;
            resourceInputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            resourceInputs["extendDeletion"] = args ? args.extendDeletion : undefined;
            resourceInputs["policyDetails"] = args ? args.policyDetails : undefined;
            resourceInputs["retainInterval"] = args ? args.retainInterval : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["copyTags"] = undefined /*out*/;
            resourceInputs["createInterval"] = undefined /*out*/;
            resourceInputs["crossRegionCopyTargets"] = undefined /*out*/;
            resourceInputs["defaultPolicy"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["exclusions"] = undefined /*out*/;
            resourceInputs["executionRoleArn"] = undefined /*out*/;
            resourceInputs["extendDeletion"] = undefined /*out*/;
            resourceInputs["policyDetails"] = undefined /*out*/;
            resourceInputs["retainInterval"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LifecyclePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LifecyclePolicy resource.
 */
export interface LifecyclePolicyArgs {
    copyTags?: pulumi.Input<boolean>;
    createInterval?: pulumi.Input<number>;
    crossRegionCopyTargets?: pulumi.Input<inputs.dlm.LifecyclePolicyCrossRegionCopyTargetsArgs>;
    defaultPolicy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    exclusions?: pulumi.Input<inputs.dlm.LifecyclePolicyExclusionsArgs>;
    executionRoleArn?: pulumi.Input<string>;
    extendDeletion?: pulumi.Input<boolean>;
    policyDetails?: pulumi.Input<inputs.dlm.LifecyclePolicyPolicyDetailsArgs>;
    retainInterval?: pulumi.Input<number>;
    state?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
