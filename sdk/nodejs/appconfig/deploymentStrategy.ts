// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppConfig::DeploymentStrategy
 *
 * @deprecated DeploymentStrategy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DeploymentStrategy extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentStrategy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentStrategy {
        pulumi.log.warn("DeploymentStrategy is deprecated: DeploymentStrategy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DeploymentStrategy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appconfig:DeploymentStrategy';

    /**
     * Returns true if the given object is an instance of DeploymentStrategy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentStrategy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentStrategy.__pulumiType;
    }

    public readonly deploymentDurationInMinutes!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly finalBakeTimeInMinutes!: pulumi.Output<number | undefined>;
    public readonly growthFactor!: pulumi.Output<number>;
    public readonly growthType!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly replicateTo!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.appconfig.DeploymentStrategyTags[] | undefined>;

    /**
     * Create a DeploymentStrategy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DeploymentStrategy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DeploymentStrategyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DeploymentStrategy is deprecated: DeploymentStrategy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deploymentDurationInMinutes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentDurationInMinutes'");
            }
            if ((!args || args.growthFactor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'growthFactor'");
            }
            if ((!args || args.replicateTo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicateTo'");
            }
            resourceInputs["deploymentDurationInMinutes"] = args ? args.deploymentDurationInMinutes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["finalBakeTimeInMinutes"] = args ? args.finalBakeTimeInMinutes : undefined;
            resourceInputs["growthFactor"] = args ? args.growthFactor : undefined;
            resourceInputs["growthType"] = args ? args.growthType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["replicateTo"] = args ? args.replicateTo : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["deploymentDurationInMinutes"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["finalBakeTimeInMinutes"] = undefined /*out*/;
            resourceInputs["growthFactor"] = undefined /*out*/;
            resourceInputs["growthType"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["replicateTo"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeploymentStrategy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentStrategy resource.
 */
export interface DeploymentStrategyArgs {
    deploymentDurationInMinutes: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    finalBakeTimeInMinutes?: pulumi.Input<number>;
    growthFactor: pulumi.Input<number>;
    growthType?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    replicateTo: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.appconfig.DeploymentStrategyTagsArgs>[]>;
}
