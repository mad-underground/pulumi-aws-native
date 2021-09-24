// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Model
 *
 * @deprecated Model is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Model extends pulumi.CustomResource {
    /**
     * Get an existing Model resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Model {
        pulumi.log.warn("Model is deprecated: Model is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Model(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:Model';

    /**
     * Returns true if the given object is an instance of Model.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Model {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Model.__pulumiType;
    }

    public readonly containers!: pulumi.Output<outputs.sagemaker.ModelContainerDefinition[] | undefined>;
    public readonly enableNetworkIsolation!: pulumi.Output<boolean | undefined>;
    public readonly executionRoleArn!: pulumi.Output<string>;
    public readonly inferenceExecutionConfig!: pulumi.Output<outputs.sagemaker.ModelInferenceExecutionConfig | undefined>;
    public readonly modelName!: pulumi.Output<string | undefined>;
    public readonly primaryContainer!: pulumi.Output<outputs.sagemaker.ModelContainerDefinition | undefined>;
    public readonly tags!: pulumi.Output<outputs.sagemaker.ModelTag[] | undefined>;
    public readonly vpcConfig!: pulumi.Output<outputs.sagemaker.ModelVpcConfig | undefined>;

    /**
     * Create a Model resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Model is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ModelArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Model is deprecated: Model is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.executionRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["enableNetworkIsolation"] = args ? args.enableNetworkIsolation : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["inferenceExecutionConfig"] = args ? args.inferenceExecutionConfig : undefined;
            inputs["modelName"] = args ? args.modelName : undefined;
            inputs["primaryContainer"] = args ? args.primaryContainer : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
        } else {
            inputs["containers"] = undefined /*out*/;
            inputs["enableNetworkIsolation"] = undefined /*out*/;
            inputs["executionRoleArn"] = undefined /*out*/;
            inputs["inferenceExecutionConfig"] = undefined /*out*/;
            inputs["modelName"] = undefined /*out*/;
            inputs["primaryContainer"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["vpcConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Model.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Model resource.
 */
export interface ModelArgs {
    containers?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelContainerDefinitionArgs>[]>;
    enableNetworkIsolation?: pulumi.Input<boolean>;
    executionRoleArn: pulumi.Input<string>;
    inferenceExecutionConfig?: pulumi.Input<inputs.sagemaker.ModelInferenceExecutionConfigArgs>;
    modelName?: pulumi.Input<string>;
    primaryContainer?: pulumi.Input<inputs.sagemaker.ModelContainerDefinitionArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelTagArgs>[]>;
    vpcConfig?: pulumi.Input<inputs.sagemaker.ModelVpcConfigArgs>;
}