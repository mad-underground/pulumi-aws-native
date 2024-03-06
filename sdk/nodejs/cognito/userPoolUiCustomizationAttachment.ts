// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
 */
export class UserPoolUiCustomizationAttachment extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolUiCustomizationAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UserPoolUiCustomizationAttachment {
        return new UserPoolUiCustomizationAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cognito:UserPoolUiCustomizationAttachment';

    /**
     * Returns true if the given object is an instance of UserPoolUiCustomizationAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolUiCustomizationAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolUiCustomizationAttachment.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly css!: pulumi.Output<string | undefined>;
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a UserPoolUiCustomizationAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPoolUiCustomizationAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["css"] = args ? args.css : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["css"] = undefined /*out*/;
            resourceInputs["userPoolId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["clientId", "userPoolId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(UserPoolUiCustomizationAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UserPoolUiCustomizationAttachment resource.
 */
export interface UserPoolUiCustomizationAttachmentArgs {
    clientId: pulumi.Input<string>;
    css?: pulumi.Input<string>;
    userPoolId: pulumi.Input<string>;
}
