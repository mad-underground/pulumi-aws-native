// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
 */
export class ModelQualityJobDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ModelQualityJobDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ModelQualityJobDefinition {
        return new ModelQualityJobDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:ModelQualityJobDefinition';

    /**
     * Returns true if the given object is an instance of ModelQualityJobDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModelQualityJobDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModelQualityJobDefinition.__pulumiType;
    }

    /**
     * The time at which the job definition was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public readonly endpointName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    public /*out*/ readonly jobDefinitionArn!: pulumi.Output<string>;
    public readonly jobDefinitionName!: pulumi.Output<string | undefined>;
    public readonly jobResources!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionMonitoringResources>;
    public readonly modelQualityAppSpecification!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionModelQualityAppSpecification>;
    public readonly modelQualityBaselineConfig!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionModelQualityBaselineConfig | undefined>;
    public readonly modelQualityJobInput!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionModelQualityJobInput>;
    public readonly modelQualityJobOutputConfig!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionMonitoringOutputConfig>;
    public readonly networkConfig!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionNetworkConfig | undefined>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    public readonly stoppingCondition!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionStoppingCondition | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.ModelQualityJobDefinitionTag[] | undefined>;

    /**
     * Create a ModelQualityJobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelQualityJobDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobResources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobResources'");
            }
            if ((!args || args.modelQualityAppSpecification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelQualityAppSpecification'");
            }
            if ((!args || args.modelQualityJobInput === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelQualityJobInput'");
            }
            if ((!args || args.modelQualityJobOutputConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelQualityJobOutputConfig'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["jobDefinitionName"] = args ? args.jobDefinitionName : undefined;
            resourceInputs["jobResources"] = args ? args.jobResources : undefined;
            resourceInputs["modelQualityAppSpecification"] = args ? args.modelQualityAppSpecification : undefined;
            resourceInputs["modelQualityBaselineConfig"] = args ? args.modelQualityBaselineConfig : undefined;
            resourceInputs["modelQualityJobInput"] = args ? args.modelQualityJobInput : undefined;
            resourceInputs["modelQualityJobOutputConfig"] = args ? args.modelQualityJobOutputConfig : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stoppingCondition"] = args ? args.stoppingCondition : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["jobDefinitionArn"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["endpointName"] = undefined /*out*/;
            resourceInputs["jobDefinitionArn"] = undefined /*out*/;
            resourceInputs["jobDefinitionName"] = undefined /*out*/;
            resourceInputs["jobResources"] = undefined /*out*/;
            resourceInputs["modelQualityAppSpecification"] = undefined /*out*/;
            resourceInputs["modelQualityBaselineConfig"] = undefined /*out*/;
            resourceInputs["modelQualityJobInput"] = undefined /*out*/;
            resourceInputs["modelQualityJobOutputConfig"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["stoppingCondition"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModelQualityJobDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModelQualityJobDefinition resource.
 */
export interface ModelQualityJobDefinitionArgs {
    endpointName?: pulumi.Input<string>;
    jobDefinitionName?: pulumi.Input<string>;
    jobResources: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionMonitoringResourcesArgs>;
    modelQualityAppSpecification: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionModelQualityAppSpecificationArgs>;
    modelQualityBaselineConfig?: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionModelQualityBaselineConfigArgs>;
    modelQualityJobInput: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionModelQualityJobInputArgs>;
    modelQualityJobOutputConfig: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionMonitoringOutputConfigArgs>;
    networkConfig?: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionNetworkConfigArgs>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    roleArn: pulumi.Input<string>;
    stoppingCondition?: pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionStoppingConditionArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelQualityJobDefinitionTagArgs>[]>;
}
