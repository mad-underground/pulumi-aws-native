// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaPackage::OriginEndpoint
 */
export class OriginEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing OriginEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OriginEndpoint {
        return new OriginEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mediapackage:OriginEndpoint';

    /**
     * Returns true if the given object is an instance of OriginEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OriginEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OriginEndpoint.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<outputs.mediapackage.OriginEndpointAuthorization | undefined>;
    /**
     * The ID of the Channel the OriginEndpoint is associated with.
     */
    public readonly channelId!: pulumi.Output<string>;
    public readonly cmafPackage!: pulumi.Output<outputs.mediapackage.OriginEndpointCmafPackage | undefined>;
    public readonly dashPackage!: pulumi.Output<outputs.mediapackage.OriginEndpointDashPackage | undefined>;
    /**
     * A short text description of the OriginEndpoint.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly hlsPackage!: pulumi.Output<outputs.mediapackage.OriginEndpointHlsPackage | undefined>;
    /**
     * A short string appended to the end of the OriginEndpoint URL.
     */
    public readonly manifestName!: pulumi.Output<string | undefined>;
    public readonly mssPackage!: pulumi.Output<outputs.mediapackage.OriginEndpointMssPackage | undefined>;
    /**
     * Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination
     */
    public readonly origination!: pulumi.Output<enums.mediapackage.OriginEndpointOrigination | undefined>;
    /**
     * Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint.
     */
    public readonly startoverWindowSeconds!: pulumi.Output<number | undefined>;
    /**
     * A collection of tags associated with a resource
     */
    public readonly tags!: pulumi.Output<outputs.mediapackage.OriginEndpointTag[] | undefined>;
    /**
     * Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint.
     */
    public readonly timeDelaySeconds!: pulumi.Output<number | undefined>;
    /**
     * The URL of the packaged OriginEndpoint for consumption.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.
     */
    public readonly whitelist!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OriginEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OriginEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.channelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelId'");
            }
            resourceInputs["authorization"] = args ? args.authorization : undefined;
            resourceInputs["channelId"] = args ? args.channelId : undefined;
            resourceInputs["cmafPackage"] = args ? args.cmafPackage : undefined;
            resourceInputs["dashPackage"] = args ? args.dashPackage : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hlsPackage"] = args ? args.hlsPackage : undefined;
            resourceInputs["manifestName"] = args ? args.manifestName : undefined;
            resourceInputs["mssPackage"] = args ? args.mssPackage : undefined;
            resourceInputs["origination"] = args ? args.origination : undefined;
            resourceInputs["startoverWindowSeconds"] = args ? args.startoverWindowSeconds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeDelaySeconds"] = args ? args.timeDelaySeconds : undefined;
            resourceInputs["whitelist"] = args ? args.whitelist : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["authorization"] = undefined /*out*/;
            resourceInputs["channelId"] = undefined /*out*/;
            resourceInputs["cmafPackage"] = undefined /*out*/;
            resourceInputs["dashPackage"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hlsPackage"] = undefined /*out*/;
            resourceInputs["manifestName"] = undefined /*out*/;
            resourceInputs["mssPackage"] = undefined /*out*/;
            resourceInputs["origination"] = undefined /*out*/;
            resourceInputs["startoverWindowSeconds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["timeDelaySeconds"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["whitelist"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OriginEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OriginEndpoint resource.
 */
export interface OriginEndpointArgs {
    authorization?: pulumi.Input<inputs.mediapackage.OriginEndpointAuthorizationArgs>;
    /**
     * The ID of the Channel the OriginEndpoint is associated with.
     */
    channelId: pulumi.Input<string>;
    cmafPackage?: pulumi.Input<inputs.mediapackage.OriginEndpointCmafPackageArgs>;
    dashPackage?: pulumi.Input<inputs.mediapackage.OriginEndpointDashPackageArgs>;
    /**
     * A short text description of the OriginEndpoint.
     */
    description?: pulumi.Input<string>;
    hlsPackage?: pulumi.Input<inputs.mediapackage.OriginEndpointHlsPackageArgs>;
    /**
     * A short string appended to the end of the OriginEndpoint URL.
     */
    manifestName?: pulumi.Input<string>;
    mssPackage?: pulumi.Input<inputs.mediapackage.OriginEndpointMssPackageArgs>;
    /**
     * Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination
     */
    origination?: pulumi.Input<enums.mediapackage.OriginEndpointOrigination>;
    /**
     * Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint.
     */
    startoverWindowSeconds?: pulumi.Input<number>;
    /**
     * A collection of tags associated with a resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.mediapackage.OriginEndpointTagArgs>[]>;
    /**
     * Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint.
     */
    timeDelaySeconds?: pulumi.Input<number>;
    /**
     * A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.
     */
    whitelist?: pulumi.Input<pulumi.Input<string>[]>;
}
