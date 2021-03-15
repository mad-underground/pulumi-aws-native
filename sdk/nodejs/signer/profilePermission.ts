// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html
 */
export class ProfilePermission extends pulumi.CustomResource {
    /**
     * Get an existing ProfilePermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProfilePermission {
        return new ProfilePermission(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Signer:ProfilePermission';

    /**
     * Returns true if the given object is an instance of ProfilePermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfilePermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfilePermission.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-action
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-principal
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profilename
     */
    public readonly profileName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profileversion
     */
    public readonly profileVersion!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-statementid
     */
    public readonly statementId!: pulumi.Output<string>;

    /**
     * Create a ProfilePermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfilePermissionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.action === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.principal === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.profileName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'profileName'");
            }
            if ((!args || args.statementId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'statementId'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["principal"] = args ? args.principal : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["profileVersion"] = args ? args.profileVersion : undefined;
            inputs["statementId"] = args ? args.statementId : undefined;
        } else {
            inputs["action"] = undefined /*out*/;
            inputs["principal"] = undefined /*out*/;
            inputs["profileName"] = undefined /*out*/;
            inputs["profileVersion"] = undefined /*out*/;
            inputs["statementId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfilePermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProfilePermission resource.
 */
export interface ProfilePermissionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-action
     */
    readonly action: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-principal
     */
    readonly principal: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profilename
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profileversion
     */
    readonly profileVersion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-statementid
     */
    readonly statementId: pulumi.Input<string>;
}
