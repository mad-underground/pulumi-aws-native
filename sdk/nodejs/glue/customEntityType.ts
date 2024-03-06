// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::CustomEntityType
 *
 * @deprecated CustomEntityType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class CustomEntityType extends pulumi.CustomResource {
    /**
     * Get an existing CustomEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomEntityType {
        pulumi.log.warn("CustomEntityType is deprecated: CustomEntityType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new CustomEntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:glue:CustomEntityType';

    /**
     * Returns true if the given object is an instance of CustomEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomEntityType.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly contextWords!: pulumi.Output<string[] | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly regexString!: pulumi.Output<string | undefined>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
     */
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a CustomEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated CustomEntityType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: CustomEntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("CustomEntityType is deprecated: CustomEntityType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["contextWords"] = args ? args.contextWords : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regexString"] = args ? args.regexString : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["contextWords"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["regexString"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomEntityType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomEntityType resource.
 */
export interface CustomEntityTypeArgs {
    contextWords?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    regexString?: pulumi.Input<string>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
     */
    tags?: any;
}
