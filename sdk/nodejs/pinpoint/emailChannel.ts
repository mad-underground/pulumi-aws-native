// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::EmailChannel
 *
 * @deprecated EmailChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EmailChannel extends pulumi.CustomResource {
    /**
     * Get an existing EmailChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EmailChannel {
        pulumi.log.warn("EmailChannel is deprecated: EmailChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EmailChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pinpoint:EmailChannel';

    /**
     * Returns true if the given object is an instance of EmailChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailChannel.__pulumiType;
    }

    public readonly applicationId!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly configurationSet!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly fromAddress!: pulumi.Output<string>;
    public readonly identity!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EmailChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: EmailChannelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EmailChannel is deprecated: EmailChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.fromAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fromAddress'");
            }
            if ((!args || args.identity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identity'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["configurationSet"] = args ? args.configurationSet : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["fromAddress"] = args ? args.fromAddress : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["configurationSet"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["fromAddress"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EmailChannel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EmailChannel resource.
 */
export interface EmailChannelArgs {
    applicationId: pulumi.Input<string>;
    configurationSet?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    fromAddress: pulumi.Input<string>;
    identity: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
}
