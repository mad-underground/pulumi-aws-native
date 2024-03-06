// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::LoggerDefinition
 *
 * @deprecated LoggerDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class LoggerDefinition extends pulumi.CustomResource {
    /**
     * Get an existing LoggerDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoggerDefinition {
        pulumi.log.warn("LoggerDefinition is deprecated: LoggerDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new LoggerDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrass:LoggerDefinition';

    /**
     * Returns true if the given object is an instance of LoggerDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoggerDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoggerDefinition.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly initialVersion!: pulumi.Output<outputs.greengrass.LoggerDefinitionVersion | undefined>;
    public /*out*/ readonly latestVersionArn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Greengrass::LoggerDefinition` for more information about the expected schema for this property.
     */
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a LoggerDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated LoggerDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: LoggerDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoggerDefinition is deprecated: LoggerDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["initialVersion"] = args ? args.initialVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["latestVersionArn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["initialVersion"] = undefined /*out*/;
            resourceInputs["latestVersionArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["initialVersion"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(LoggerDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoggerDefinition resource.
 */
export interface LoggerDefinitionArgs {
    initialVersion?: pulumi.Input<inputs.greengrass.LoggerDefinitionVersionArgs>;
    name?: pulumi.Input<string>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Greengrass::LoggerDefinition` for more information about the expected schema for this property.
     */
    tags?: any;
}
