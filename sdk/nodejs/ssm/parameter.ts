// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::SSM::Parameter`` resource creates an SSM parameter in SYSlong Parameter Store.
 *   To create an SSM parameter, you must have the IAMlong (IAM) permissions ``ssm:PutParameter`` and ``ssm:AddTagsToResource``. On stack creation, CFNlong adds the following three tags to the parameter: ``aws:cloudformation:stack-name``, ``aws:cloudformation:logical-id``, and ``aws:cloudformation:stack-id``, in addition to any custom tags you specify.
 *  To add, update, or remove tags during stack update, you must have IAM permissions for both ``ssm:AddTagsToResource`` and ``ssm:RemoveTagsFromResource``. For more information, see [Managing Access Using Policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/security-iam.html#security_iam_access-manage) in the *User Guide*.
 *   For information about valid values for parameters, see [Requirements and Constraints for Parameter Names](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-su-create.html#sysman-paramete
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Parameter {
        return new Parameter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ssm:Parameter';

    /**
     * Returns true if the given object is an instance of Parameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Parameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Parameter.__pulumiType;
    }

    /**
     * A regular expression used to validate the parameter value. For example, for ``String`` types with values restricted to numbers, you can specify the following: ``AllowedPattern=^\d+$``
     */
    public readonly allowedPattern!: pulumi.Output<string | undefined>;
    /**
     * The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
     */
    public readonly dataType!: pulumi.Output<enums.ssm.ParameterDataType | undefined>;
    /**
     * Information about the parameter.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the parameter.
     *  The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: ``arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName``
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Information about the policies assigned to a parameter.
     *   [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *User Guide*.
     */
    public readonly policies!: pulumi.Output<string | undefined>;
    /**
     * Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The parameter tier.
     */
    public readonly tier!: pulumi.Output<enums.ssm.ParameterTier | undefined>;
    /**
     * The type of parameter.
     *   Although ``SecureString`` is included in the list of valid values, CFNlong does *not* currently support creating a ``SecureString`` parameter type.
     */
    public readonly type!: pulumi.Output<enums.ssm.ParameterType>;
    /**
     * The parameter value.
     *   If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        } else {
            resourceInputs["allowedPattern"] = undefined /*out*/;
            resourceInputs["dataType"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["policies"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Parameter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    /**
     * A regular expression used to validate the parameter value. For example, for ``String`` types with values restricted to numbers, you can specify the following: ``AllowedPattern=^\d+$``
     */
    allowedPattern?: pulumi.Input<string>;
    /**
     * The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
     */
    dataType?: pulumi.Input<enums.ssm.ParameterDataType>;
    /**
     * Information about the parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the parameter.
     *  The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: ``arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName``
     */
    name?: pulumi.Input<string>;
    /**
     * Information about the policies assigned to a parameter.
     *   [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *User Guide*.
     */
    policies?: pulumi.Input<string>;
    /**
     * Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The parameter tier.
     */
    tier?: pulumi.Input<enums.ssm.ParameterTier>;
    /**
     * The type of parameter.
     *   Although ``SecureString`` is included in the list of valid values, CFNlong does *not* currently support creating a ``SecureString`` parameter type.
     */
    type: pulumi.Input<enums.ssm.ParameterType>;
    /**
     * The parameter value.
     *   If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
     */
    value: pulumi.Input<string>;
}
