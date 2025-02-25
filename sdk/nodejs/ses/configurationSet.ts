// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::SES::ConfigurationSet.
 */
export class ConfigurationSet extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigurationSet {
        return new ConfigurationSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ses:ConfigurationSet';

    /**
     * Returns true if the given object is an instance of ConfigurationSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationSet.__pulumiType;
    }

    public readonly deliveryOptions!: pulumi.Output<outputs.ses.ConfigurationSetDeliveryOptions | undefined>;
    /**
     * The name of the configuration set.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly reputationOptions!: pulumi.Output<outputs.ses.ConfigurationSetReputationOptions | undefined>;
    public readonly sendingOptions!: pulumi.Output<outputs.ses.ConfigurationSetSendingOptions | undefined>;
    public readonly suppressionOptions!: pulumi.Output<outputs.ses.ConfigurationSetSuppressionOptions | undefined>;
    public readonly trackingOptions!: pulumi.Output<outputs.ses.ConfigurationSetTrackingOptions | undefined>;
    public readonly vdmOptions!: pulumi.Output<outputs.ses.ConfigurationSetVdmOptions | undefined>;

    /**
     * Create a ConfigurationSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigurationSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["deliveryOptions"] = args ? args.deliveryOptions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["reputationOptions"] = args ? args.reputationOptions : undefined;
            resourceInputs["sendingOptions"] = args ? args.sendingOptions : undefined;
            resourceInputs["suppressionOptions"] = args ? args.suppressionOptions : undefined;
            resourceInputs["trackingOptions"] = args ? args.trackingOptions : undefined;
            resourceInputs["vdmOptions"] = args ? args.vdmOptions : undefined;
        } else {
            resourceInputs["deliveryOptions"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["reputationOptions"] = undefined /*out*/;
            resourceInputs["sendingOptions"] = undefined /*out*/;
            resourceInputs["suppressionOptions"] = undefined /*out*/;
            resourceInputs["trackingOptions"] = undefined /*out*/;
            resourceInputs["vdmOptions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ConfigurationSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigurationSet resource.
 */
export interface ConfigurationSetArgs {
    deliveryOptions?: pulumi.Input<inputs.ses.ConfigurationSetDeliveryOptionsArgs>;
    /**
     * The name of the configuration set.
     */
    name?: pulumi.Input<string>;
    reputationOptions?: pulumi.Input<inputs.ses.ConfigurationSetReputationOptionsArgs>;
    sendingOptions?: pulumi.Input<inputs.ses.ConfigurationSetSendingOptionsArgs>;
    suppressionOptions?: pulumi.Input<inputs.ses.ConfigurationSetSuppressionOptionsArgs>;
    trackingOptions?: pulumi.Input<inputs.ses.ConfigurationSetTrackingOptionsArgs>;
    vdmOptions?: pulumi.Input<inputs.ses.ConfigurationSetVdmOptionsArgs>;
}
