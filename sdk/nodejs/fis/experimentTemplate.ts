// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export class ExperimentTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ExperimentTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExperimentTemplate {
        return new ExperimentTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:fis:ExperimentTemplate';

    /**
     * Returns true if the given object is an instance of ExperimentTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExperimentTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExperimentTemplate.__pulumiType;
    }

    public readonly actions!: pulumi.Output<{[key: string]: outputs.fis.ExperimentTemplateAction} | undefined>;
    public readonly description!: pulumi.Output<string>;
    public readonly experimentOptions!: pulumi.Output<outputs.fis.ExperimentTemplateExperimentOptions | undefined>;
    public readonly logConfiguration!: pulumi.Output<outputs.fis.ExperimentTemplateLogConfiguration | undefined>;
    public readonly roleArn!: pulumi.Output<string>;
    public readonly stopConditions!: pulumi.Output<outputs.fis.ExperimentTemplateStopCondition[]>;
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    public readonly targets!: pulumi.Output<{[key: string]: outputs.fis.ExperimentTemplateTarget}>;

    /**
     * Create a ExperimentTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExperimentTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.stopConditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stopConditions'");
            }
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["experimentOptions"] = args ? args.experimentOptions : undefined;
            resourceInputs["logConfiguration"] = args ? args.logConfiguration : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stopConditions"] = args ? args.stopConditions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
        } else {
            resourceInputs["actions"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["experimentOptions"] = undefined /*out*/;
            resourceInputs["logConfiguration"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["stopConditions"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["tags.*"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ExperimentTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExperimentTemplate resource.
 */
export interface ExperimentTemplateArgs {
    actions?: pulumi.Input<{[key: string]: pulumi.Input<inputs.fis.ExperimentTemplateActionArgs>}>;
    description: pulumi.Input<string>;
    experimentOptions?: pulumi.Input<inputs.fis.ExperimentTemplateExperimentOptionsArgs>;
    logConfiguration?: pulumi.Input<inputs.fis.ExperimentTemplateLogConfigurationArgs>;
    roleArn: pulumi.Input<string>;
    stopConditions: pulumi.Input<pulumi.Input<inputs.fis.ExperimentTemplateStopConditionArgs>[]>;
    tags: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    targets: pulumi.Input<{[key: string]: pulumi.Input<inputs.fis.ExperimentTemplateTargetArgs>}>;
}
