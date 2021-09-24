// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::ADMChannel
 *
 * @deprecated ADMChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ADMChannel extends pulumi.CustomResource {
    /**
     * Get an existing ADMChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ADMChannel {
        pulumi.log.warn("ADMChannel is deprecated: ADMChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ADMChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:ADMChannel';

    /**
     * Returns true if the given object is an instance of ADMChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ADMChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ADMChannel.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSecret!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ADMChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ADMChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ADMChannelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ADMChannel is deprecated: ADMChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
        } else {
            inputs["applicationId"] = undefined /*out*/;
            inputs["clientId"] = undefined /*out*/;
            inputs["clientSecret"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ADMChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ADMChannel resource.
 */
export interface ADMChannelArgs {
    applicationId: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    clientSecret: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
}