// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IVS::Channel
 */
export class Channel extends pulumi.CustomResource {
    /**
     * Get an existing Channel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Channel {
        return new Channel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ivs:Channel';

    /**
     * Returns true if the given object is an instance of Channel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Channel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Channel.__pulumiType;
    }

    /**
     * Channel ARN is automatically generated on creation and assigned as the unique identifier.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether the channel is authorized.
     */
    public readonly authorized!: pulumi.Output<boolean | undefined>;
    /**
     * Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.
     */
    public /*out*/ readonly ingestEndpoint!: pulumi.Output<string>;
    /**
     * Channel latency mode.
     */
    public readonly latencyMode!: pulumi.Output<enums.ivs.ChannelLatencyMode | undefined>;
    /**
     * Channel
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Channel Playback URL.
     */
    public /*out*/ readonly playbackUrl!: pulumi.Output<string>;
    /**
     * Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: “” (recording is disabled).
     */
    public readonly recordingConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    public readonly tags!: pulumi.Output<outputs.ivs.ChannelTag[] | undefined>;
    /**
     * Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
     */
    public readonly type!: pulumi.Output<enums.ivs.ChannelType | undefined>;

    /**
     * Create a Channel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["latencyMode"] = args ? args.latencyMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recordingConfigurationArn"] = args ? args.recordingConfigurationArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["ingestEndpoint"] = undefined /*out*/;
            resourceInputs["playbackUrl"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["authorized"] = undefined /*out*/;
            resourceInputs["ingestEndpoint"] = undefined /*out*/;
            resourceInputs["latencyMode"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["playbackUrl"] = undefined /*out*/;
            resourceInputs["recordingConfigurationArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Channel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Channel resource.
 */
export interface ChannelArgs {
    /**
     * Whether the channel is authorized.
     */
    authorized?: pulumi.Input<boolean>;
    /**
     * Channel latency mode.
     */
    latencyMode?: pulumi.Input<enums.ivs.ChannelLatencyMode>;
    /**
     * Channel
     */
    name?: pulumi.Input<string>;
    /**
     * Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: “” (recording is disabled).
     */
    recordingConfigurationArn?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the asset model.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ivs.ChannelTagArgs>[]>;
    /**
     * Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
     */
    type?: pulumi.Input<enums.ivs.ChannelType>;
}
