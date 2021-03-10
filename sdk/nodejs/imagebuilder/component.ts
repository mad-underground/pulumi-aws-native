// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html
 */
export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ImageBuilder:Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    public /*out*/ readonly Arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-changedescription
     */
    public readonly ChangeDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-data
     */
    public readonly Data!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-description
     */
    public readonly Description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly Encrypted!: pulumi.Output<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-kmskeyid
     */
    public readonly KmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-name
     */
    public readonly Name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-platform
     */
    public readonly Platform!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-supportedosversions
     */
    public readonly SupportedOsVersions!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-tags
     */
    public readonly Tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly Type!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-uri
     */
    public readonly Uri!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-version
     */
    public readonly Version!: pulumi.Output<string>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.Name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Name'");
            }
            if ((!args || args.Platform === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Platform'");
            }
            if ((!args || args.Version === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Version'");
            }
            inputs["ChangeDescription"] = args ? args.ChangeDescription : undefined;
            inputs["Data"] = args ? args.Data : undefined;
            inputs["Description"] = args ? args.Description : undefined;
            inputs["KmsKeyId"] = args ? args.KmsKeyId : undefined;
            inputs["Name"] = args ? args.Name : undefined;
            inputs["Platform"] = args ? args.Platform : undefined;
            inputs["SupportedOsVersions"] = args ? args.SupportedOsVersions : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["Uri"] = args ? args.Uri : undefined;
            inputs["Version"] = args ? args.Version : undefined;
            inputs["Arn"] = undefined /*out*/;
            inputs["Encrypted"] = undefined /*out*/;
            inputs["Type"] = undefined /*out*/;
        } else {
            inputs["Arn"] = undefined /*out*/;
            inputs["ChangeDescription"] = undefined /*out*/;
            inputs["Data"] = undefined /*out*/;
            inputs["Description"] = undefined /*out*/;
            inputs["Encrypted"] = undefined /*out*/;
            inputs["KmsKeyId"] = undefined /*out*/;
            inputs["Name"] = undefined /*out*/;
            inputs["Platform"] = undefined /*out*/;
            inputs["SupportedOsVersions"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
            inputs["Type"] = undefined /*out*/;
            inputs["Uri"] = undefined /*out*/;
            inputs["Version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Component.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-changedescription
     */
    readonly ChangeDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-data
     */
    readonly Data?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-description
     */
    readonly Description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-kmskeyid
     */
    readonly KmsKeyId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-name
     */
    readonly Name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-platform
     */
    readonly Platform: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-supportedosversions
     */
    readonly SupportedOsVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-tags
     */
    readonly Tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-uri
     */
    readonly Uri?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-version
     */
    readonly Version: pulumi.Input<string>;
}
