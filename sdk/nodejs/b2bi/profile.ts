// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::B2BI::Profile Resource Type
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:b2bi:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    public readonly businessName!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly email!: pulumi.Output<string | undefined>;
    public /*out*/ readonly logGroupName!: pulumi.Output<string>;
    public readonly logging!: pulumi.Output<enums.b2bi.ProfileLogging>;
    public /*out*/ readonly modifiedAt!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly phone!: pulumi.Output<string>;
    public /*out*/ readonly profileArn!: pulumi.Output<string>;
    public /*out*/ readonly profileId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.businessName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'businessName'");
            }
            if ((!args || args.logging === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logging'");
            }
            if ((!args || args.phone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phone'");
            }
            resourceInputs["businessName"] = args ? args.businessName : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["logGroupName"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["profileArn"] = undefined /*out*/;
            resourceInputs["profileId"] = undefined /*out*/;
        } else {
            resourceInputs["businessName"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["logGroupName"] = undefined /*out*/;
            resourceInputs["logging"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["phone"] = undefined /*out*/;
            resourceInputs["profileArn"] = undefined /*out*/;
            resourceInputs["profileId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["logging"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Profile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    businessName: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    logging: pulumi.Input<enums.b2bi.ProfileLogging>;
    name?: pulumi.Input<string>;
    phone: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
