// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Publishes new or first hook version to AWS CloudFormation Registry.
 */
export class HookVersion extends pulumi.CustomResource {
    /**
     * Get an existing HookVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HookVersion {
        return new HookVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:HookVersion';

    /**
     * Returns true if the given object is an instance of HookVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HookVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HookVersion.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
     */
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Indicates if this type version is the current default version
     */
    public /*out*/ readonly isDefaultVersion!: pulumi.Output<boolean>;
    /**
     * Specifies logging configuration information for a type.
     */
    public readonly loggingConfig!: pulumi.Output<outputs.cloudformation.HookVersionLoggingConfig | undefined>;
    /**
     * A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
     *
     * For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
     */
    public readonly schemaHandlerPackage!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the type without the versionID.
     */
    public /*out*/ readonly typeArn!: pulumi.Output<string>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    public readonly typeName!: pulumi.Output<string>;
    /**
     * The ID of the version of the type represented by this hook instance.
     */
    public /*out*/ readonly versionId!: pulumi.Output<string>;
    /**
     * The scope at which the type is visible and usable in CloudFormation operations.
     *
     * Valid values include:
     *
     * PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
     *
     * PUBLIC: The type is publically visible and usable within any Amazon account.
     */
    public /*out*/ readonly visibility!: pulumi.Output<enums.cloudformation.HookVersionVisibility>;

    /**
     * Create a HookVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HookVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.schemaHandlerPackage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemaHandlerPackage'");
            }
            if ((!args || args.typeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'typeName'");
            }
            resourceInputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            resourceInputs["loggingConfig"] = args ? args.loggingConfig : undefined;
            resourceInputs["schemaHandlerPackage"] = args ? args.schemaHandlerPackage : undefined;
            resourceInputs["typeName"] = args ? args.typeName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["isDefaultVersion"] = undefined /*out*/;
            resourceInputs["typeArn"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["executionRoleArn"] = undefined /*out*/;
            resourceInputs["isDefaultVersion"] = undefined /*out*/;
            resourceInputs["loggingConfig"] = undefined /*out*/;
            resourceInputs["schemaHandlerPackage"] = undefined /*out*/;
            resourceInputs["typeArn"] = undefined /*out*/;
            resourceInputs["typeName"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HookVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a HookVersion resource.
 */
export interface HookVersionArgs {
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
     */
    executionRoleArn?: pulumi.Input<string>;
    /**
     * Specifies logging configuration information for a type.
     */
    loggingConfig?: pulumi.Input<inputs.cloudformation.HookVersionLoggingConfigArgs>;
    /**
     * A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
     *
     * For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
     */
    schemaHandlerPackage: pulumi.Input<string>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    typeName: pulumi.Input<string>;
}
