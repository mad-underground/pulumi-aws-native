// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html
 */
export class InfrastructureConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing InfrastructureConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InfrastructureConfiguration {
        return new InfrastructureConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ImageBuilder:InfrastructureConfiguration';

    /**
     * Returns true if the given object is an instance of InfrastructureConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InfrastructureConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InfrastructureConfiguration.__pulumiType;
    }

    public /*out*/ readonly Arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-description
     */
    public readonly Description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instanceprofilename
     */
    public readonly InstanceProfileName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancetypes
     */
    public readonly InstanceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-keypair
     */
    public readonly KeyPair!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-logging
     */
    public readonly Logging!: pulumi.Output<any | string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-name
     */
    public readonly Name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-resourcetags
     */
    public readonly ResourceTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-securitygroupids
     */
    public readonly SecurityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-snstopicarn
     */
    public readonly SnsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-subnetid
     */
    public readonly SubnetId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-tags
     */
    public readonly Tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-terminateinstanceonfailure
     */
    public readonly TerminateInstanceOnFailure!: pulumi.Output<boolean | undefined>;

    /**
     * Create a InfrastructureConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InfrastructureConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.InstanceProfileName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'InstanceProfileName'");
            }
            if ((!args || args.Name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'Name'");
            }
            inputs["Description"] = args ? args.Description : undefined;
            inputs["InstanceProfileName"] = args ? args.InstanceProfileName : undefined;
            inputs["InstanceTypes"] = args ? args.InstanceTypes : undefined;
            inputs["KeyPair"] = args ? args.KeyPair : undefined;
            inputs["Logging"] = args ? args.Logging : undefined;
            inputs["Name"] = args ? args.Name : undefined;
            inputs["ResourceTags"] = args ? args.ResourceTags : undefined;
            inputs["SecurityGroupIds"] = args ? args.SecurityGroupIds : undefined;
            inputs["SnsTopicArn"] = args ? args.SnsTopicArn : undefined;
            inputs["SubnetId"] = args ? args.SubnetId : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["TerminateInstanceOnFailure"] = args ? args.TerminateInstanceOnFailure : undefined;
            inputs["Arn"] = undefined /*out*/;
        } else {
            inputs["Arn"] = undefined /*out*/;
            inputs["Description"] = undefined /*out*/;
            inputs["InstanceProfileName"] = undefined /*out*/;
            inputs["InstanceTypes"] = undefined /*out*/;
            inputs["KeyPair"] = undefined /*out*/;
            inputs["Logging"] = undefined /*out*/;
            inputs["Name"] = undefined /*out*/;
            inputs["ResourceTags"] = undefined /*out*/;
            inputs["SecurityGroupIds"] = undefined /*out*/;
            inputs["SnsTopicArn"] = undefined /*out*/;
            inputs["SubnetId"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
            inputs["TerminateInstanceOnFailure"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InfrastructureConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InfrastructureConfiguration resource.
 */
export interface InfrastructureConfigurationArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-description
     */
    readonly Description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instanceprofilename
     */
    readonly InstanceProfileName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancetypes
     */
    readonly InstanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-keypair
     */
    readonly KeyPair?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-logging
     */
    readonly Logging?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-name
     */
    readonly Name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-resourcetags
     */
    readonly ResourceTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-securitygroupids
     */
    readonly SecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-snstopicarn
     */
    readonly SnsTopicArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-subnetid
     */
    readonly SubnetId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-tags
     */
    readonly Tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-terminateinstanceonfailure
     */
    readonly TerminateInstanceOnFailure?: pulumi.Input<boolean>;
}
