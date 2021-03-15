// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html
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
    public static readonly __pulumiType = 'aws-native:SageMaker:ModelQualityJobDefinition';

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

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly jobDefinitionArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobdefinitionname
     */
    public readonly jobDefinitionName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobresources
     */
    public readonly jobResources!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionMonitoringResources>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification
     */
    public readonly modelQualityAppSpecification!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionModelQualityAppSpecification>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig
     */
    public readonly modelQualityBaselineConfig!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionModelQualityBaselineConfig | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput
     */
    public readonly modelQualityJobInput!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionModelQualityJobInput>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjoboutputconfig
     */
    public readonly modelQualityJobOutputConfig!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionMonitoringOutputConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-networkconfig
     */
    public readonly networkConfig!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionNetworkConfig | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-rolearn
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-stoppingcondition
     */
    public readonly stoppingCondition!: pulumi.Output<outputs.SageMaker.ModelQualityJobDefinitionStoppingCondition | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a ModelQualityJobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelQualityJobDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.jobResources === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'jobResources'");
            }
            if ((!args || args.modelQualityAppSpecification === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'modelQualityAppSpecification'");
            }
            if ((!args || args.modelQualityJobInput === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'modelQualityJobInput'");
            }
            if ((!args || args.modelQualityJobOutputConfig === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'modelQualityJobOutputConfig'");
            }
            if ((!args || args.roleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["jobDefinitionName"] = args ? args.jobDefinitionName : undefined;
            inputs["jobResources"] = args ? args.jobResources : undefined;
            inputs["modelQualityAppSpecification"] = args ? args.modelQualityAppSpecification : undefined;
            inputs["modelQualityBaselineConfig"] = args ? args.modelQualityBaselineConfig : undefined;
            inputs["modelQualityJobInput"] = args ? args.modelQualityJobInput : undefined;
            inputs["modelQualityJobOutputConfig"] = args ? args.modelQualityJobOutputConfig : undefined;
            inputs["networkConfig"] = args ? args.networkConfig : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["stoppingCondition"] = args ? args.stoppingCondition : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["jobDefinitionArn"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["jobDefinitionArn"] = undefined /*out*/;
            inputs["jobDefinitionName"] = undefined /*out*/;
            inputs["jobResources"] = undefined /*out*/;
            inputs["modelQualityAppSpecification"] = undefined /*out*/;
            inputs["modelQualityBaselineConfig"] = undefined /*out*/;
            inputs["modelQualityJobInput"] = undefined /*out*/;
            inputs["modelQualityJobOutputConfig"] = undefined /*out*/;
            inputs["networkConfig"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
            inputs["stoppingCondition"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ModelQualityJobDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModelQualityJobDefinition resource.
 */
export interface ModelQualityJobDefinitionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobdefinitionname
     */
    readonly jobDefinitionName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobresources
     */
    readonly jobResources: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionMonitoringResources>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification
     */
    readonly modelQualityAppSpecification: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionModelQualityAppSpecification>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig
     */
    readonly modelQualityBaselineConfig?: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionModelQualityBaselineConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput
     */
    readonly modelQualityJobInput: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionModelQualityJobInput>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjoboutputconfig
     */
    readonly modelQualityJobOutputConfig: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionMonitoringOutputConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-networkconfig
     */
    readonly networkConfig?: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionNetworkConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-rolearn
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-stoppingcondition
     */
    readonly stoppingCondition?: pulumi.Input<inputs.SageMaker.ModelQualityJobDefinitionStoppingCondition>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
