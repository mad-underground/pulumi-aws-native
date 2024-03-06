// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::VoiceChannel
 *
 * @deprecated VoiceChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class VoiceChannel extends pulumi.CustomResource {
    /**
     * Get an existing VoiceChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VoiceChannel {
        pulumi.log.warn("VoiceChannel is deprecated: VoiceChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new VoiceChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:VoiceChannel';

    /**
     * Returns true if the given object is an instance of VoiceChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VoiceChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VoiceChannel.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VoiceChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated VoiceChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: VoiceChannelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("VoiceChannel is deprecated: VoiceChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(VoiceChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VoiceChannel resource.
 */
export interface VoiceChannelArgs {
    applicationId: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
}
