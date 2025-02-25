// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents an AWS account that is a part of a collaboration
 */
export class Membership extends pulumi.CustomResource {
    /**
     * Get an existing Membership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Membership {
        return new Membership(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cleanrooms:Membership';

    /**
     * Returns true if the given object is an instance of Membership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Membership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Membership.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly collaborationArn!: pulumi.Output<string>;
    public /*out*/ readonly collaborationCreatorAccountId!: pulumi.Output<string>;
    public readonly collaborationIdentifier!: pulumi.Output<string>;
    public readonly defaultResultConfiguration!: pulumi.Output<outputs.cleanrooms.MembershipProtectedQueryResultConfiguration | undefined>;
    public /*out*/ readonly membershipIdentifier!: pulumi.Output<string>;
    public readonly paymentConfiguration!: pulumi.Output<outputs.cleanrooms.MembershipPaymentConfiguration | undefined>;
    public readonly queryLogStatus!: pulumi.Output<enums.cleanrooms.MembershipQueryLogStatus>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms membership.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Membership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembershipArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.collaborationIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collaborationIdentifier'");
            }
            if ((!args || args.queryLogStatus === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queryLogStatus'");
            }
            resourceInputs["collaborationIdentifier"] = args ? args.collaborationIdentifier : undefined;
            resourceInputs["defaultResultConfiguration"] = args ? args.defaultResultConfiguration : undefined;
            resourceInputs["paymentConfiguration"] = args ? args.paymentConfiguration : undefined;
            resourceInputs["queryLogStatus"] = args ? args.queryLogStatus : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["collaborationArn"] = undefined /*out*/;
            resourceInputs["collaborationCreatorAccountId"] = undefined /*out*/;
            resourceInputs["membershipIdentifier"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["collaborationArn"] = undefined /*out*/;
            resourceInputs["collaborationCreatorAccountId"] = undefined /*out*/;
            resourceInputs["collaborationIdentifier"] = undefined /*out*/;
            resourceInputs["defaultResultConfiguration"] = undefined /*out*/;
            resourceInputs["membershipIdentifier"] = undefined /*out*/;
            resourceInputs["paymentConfiguration"] = undefined /*out*/;
            resourceInputs["queryLogStatus"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["collaborationIdentifier"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Membership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Membership resource.
 */
export interface MembershipArgs {
    collaborationIdentifier: pulumi.Input<string>;
    defaultResultConfiguration?: pulumi.Input<inputs.cleanrooms.MembershipProtectedQueryResultConfigurationArgs>;
    paymentConfiguration?: pulumi.Input<inputs.cleanrooms.MembershipPaymentConfigurationArgs>;
    queryLogStatus: pulumi.Input<enums.cleanrooms.MembershipQueryLogStatus>;
    /**
     * An arbitrary set of tags (key-value pairs) for this cleanrooms membership.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
