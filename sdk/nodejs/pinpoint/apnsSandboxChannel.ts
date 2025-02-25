// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
 *
 * @deprecated ApnsSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ApnsSandboxChannel extends pulumi.CustomResource {
    /**
     * Get an existing ApnsSandboxChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApnsSandboxChannel {
        pulumi.log.warn("ApnsSandboxChannel is deprecated: ApnsSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ApnsSandboxChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:ApnsSandboxChannel';

    /**
     * Returns true if the given object is an instance of ApnsSandboxChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApnsSandboxChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApnsSandboxChannel.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly bundleId!: pulumi.Output<string | undefined>;
    public readonly certificate!: pulumi.Output<string | undefined>;
    public readonly defaultAuthenticationMethod!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly privateKey!: pulumi.Output<string | undefined>;
    public readonly teamId!: pulumi.Output<string | undefined>;
    public readonly tokenKey!: pulumi.Output<string | undefined>;
    public readonly tokenKeyId!: pulumi.Output<string | undefined>;

    /**
     * Create a ApnsSandboxChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ApnsSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ApnsSandboxChannelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ApnsSandboxChannel is deprecated: ApnsSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["bundleId"] = args ? args.bundleId : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["defaultAuthenticationMethod"] = args ? args.defaultAuthenticationMethod : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["tokenKey"] = args ? args.tokenKey : undefined;
            resourceInputs["tokenKeyId"] = args ? args.tokenKeyId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["bundleId"] = undefined /*out*/;
            resourceInputs["certificate"] = undefined /*out*/;
            resourceInputs["defaultAuthenticationMethod"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["privateKey"] = undefined /*out*/;
            resourceInputs["teamId"] = undefined /*out*/;
            resourceInputs["tokenKey"] = undefined /*out*/;
            resourceInputs["tokenKeyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ApnsSandboxChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApnsSandboxChannel resource.
 */
export interface ApnsSandboxChannelArgs {
    applicationId: pulumi.Input<string>;
    bundleId?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    defaultAuthenticationMethod?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    privateKey?: pulumi.Input<string>;
    teamId?: pulumi.Input<string>;
    tokenKey?: pulumi.Input<string>;
    tokenKeyId?: pulumi.Input<string>;
}
