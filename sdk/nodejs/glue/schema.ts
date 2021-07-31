// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html
 */
export class Schema extends pulumi.CustomResource {
    /**
     * Get an existing Schema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schema {
        return new Schema(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Glue:Schema';

    /**
     * Returns true if the given object is an instance of Schema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schema.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-checkpointversion
     */
    public readonly checkpointVersion!: pulumi.Output<outputs.Glue.SchemaSchemaVersion | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-compatibility
     */
    public readonly compatibility!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-dataformat
     */
    public readonly dataFormat!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly initialSchemaVersionId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-registry
     */
    public readonly registry!: pulumi.Output<outputs.Glue.SchemaRegistry | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-schemadefinition
     */
    public readonly schemaDefinition!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Schema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SchemaArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.compatibility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compatibility'");
            }
            if ((!args || args.dataFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFormat'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.schemaDefinition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemaDefinition'");
            }
            inputs["checkpointVersion"] = args ? args.checkpointVersion : undefined;
            inputs["compatibility"] = args ? args.compatibility : undefined;
            inputs["dataFormat"] = args ? args.dataFormat : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["registry"] = args ? args.registry : undefined;
            inputs["schemaDefinition"] = args ? args.schemaDefinition : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["initialSchemaVersionId"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["checkpointVersion"] = undefined /*out*/;
            inputs["compatibility"] = undefined /*out*/;
            inputs["dataFormat"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["initialSchemaVersionId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["registry"] = undefined /*out*/;
            inputs["schemaDefinition"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Schema.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schema resource.
 */
export interface SchemaArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-checkpointversion
     */
    checkpointVersion?: pulumi.Input<inputs.Glue.SchemaSchemaVersionArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-compatibility
     */
    compatibility: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-dataformat
     */
    dataFormat: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-description
     */
    description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-name
     */
    name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-registry
     */
    registry?: pulumi.Input<inputs.Glue.SchemaRegistryArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-schemadefinition
     */
    schemaDefinition: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
