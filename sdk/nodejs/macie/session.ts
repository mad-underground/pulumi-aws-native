// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html
 */
export class Session extends pulumi.CustomResource {
    /**
     * Get an existing Session resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Session {
        return new Session(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Macie:Session';

    /**
     * Returns true if the given object is an instance of Session.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Session {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Session.__pulumiType;
    }

    public /*out*/ readonly AwsAccountId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
     */
    public readonly FindingPublishingFrequency!: pulumi.Output<string | undefined>;
    public /*out*/ readonly ServiceRole!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
     */
    public readonly Status!: pulumi.Output<string | undefined>;

    /**
     * Create a Session resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SessionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            inputs["FindingPublishingFrequency"] = args ? args.FindingPublishingFrequency : undefined;
            inputs["Status"] = args ? args.Status : undefined;
            inputs["AwsAccountId"] = undefined /*out*/;
            inputs["ServiceRole"] = undefined /*out*/;
        } else {
            inputs["AwsAccountId"] = undefined /*out*/;
            inputs["FindingPublishingFrequency"] = undefined /*out*/;
            inputs["ServiceRole"] = undefined /*out*/;
            inputs["Status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Session.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Session resource.
 */
export interface SessionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
     */
    readonly FindingPublishingFrequency?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
     */
    readonly Status?: pulumi.Input<string>;
}
