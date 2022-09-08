// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodePipeline::Webhook
 *
 * @deprecated Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Webhook {
        pulumi.log.warn("Webhook is deprecated: Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Webhook(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:codepipeline:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    public readonly authentication!: pulumi.Output<string>;
    public readonly authenticationConfiguration!: pulumi.Output<outputs.codepipeline.WebhookAuthConfiguration>;
    public readonly filters!: pulumi.Output<outputs.codepipeline.WebhookFilterRule[]>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly registerWithThirdParty!: pulumi.Output<boolean | undefined>;
    public readonly targetAction!: pulumi.Output<string>;
    public readonly targetPipeline!: pulumi.Output<string>;
    public readonly targetPipelineVersion!: pulumi.Output<number>;
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Webhook is deprecated: Webhook is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authentication === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authentication'");
            }
            if ((!args || args.authenticationConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationConfiguration'");
            }
            if ((!args || args.filters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filters'");
            }
            if ((!args || args.targetAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetAction'");
            }
            if ((!args || args.targetPipeline === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetPipeline'");
            }
            if ((!args || args.targetPipelineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetPipelineVersion'");
            }
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["authenticationConfiguration"] = args ? args.authenticationConfiguration : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["registerWithThirdParty"] = args ? args.registerWithThirdParty : undefined;
            resourceInputs["targetAction"] = args ? args.targetAction : undefined;
            resourceInputs["targetPipeline"] = args ? args.targetPipeline : undefined;
            resourceInputs["targetPipelineVersion"] = args ? args.targetPipelineVersion : undefined;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["authentication"] = undefined /*out*/;
            resourceInputs["authenticationConfiguration"] = undefined /*out*/;
            resourceInputs["filters"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["registerWithThirdParty"] = undefined /*out*/;
            resourceInputs["targetAction"] = undefined /*out*/;
            resourceInputs["targetPipeline"] = undefined /*out*/;
            resourceInputs["targetPipelineVersion"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Webhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    authentication: pulumi.Input<string>;
    authenticationConfiguration: pulumi.Input<inputs.codepipeline.WebhookAuthConfigurationArgs>;
    filters: pulumi.Input<pulumi.Input<inputs.codepipeline.WebhookFilterRuleArgs>[]>;
    name?: pulumi.Input<string>;
    registerWithThirdParty?: pulumi.Input<boolean>;
    targetAction: pulumi.Input<string>;
    targetPipeline: pulumi.Input<string>;
    targetPipelineVersion: pulumi.Input<number>;
}
