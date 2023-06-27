// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelPackage
 */
export class ModelPackage extends pulumi.CustomResource {
    /**
     * Get an existing ModelPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ModelPackage {
        return new ModelPackage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:ModelPackage';

    /**
     * Returns true if the given object is an instance of ModelPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModelPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModelPackage.__pulumiType;
    }

    public readonly additionalInferenceSpecifications!: pulumi.Output<outputs.sagemaker.ModelPackageAdditionalInferenceSpecificationDefinition[] | undefined>;
    public readonly additionalInferenceSpecificationsToAdd!: pulumi.Output<outputs.sagemaker.ModelPackageAdditionalInferenceSpecificationDefinition[] | undefined>;
    public readonly approvalDescription!: pulumi.Output<string | undefined>;
    public readonly certifyForMarketplace!: pulumi.Output<boolean | undefined>;
    public readonly clientToken!: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public readonly customerMetadataProperties!: pulumi.Output<outputs.sagemaker.ModelPackageCustomerMetadataProperties | undefined>;
    public readonly domain!: pulumi.Output<string | undefined>;
    public readonly driftCheckBaselines!: pulumi.Output<outputs.sagemaker.ModelPackageDriftCheckBaselines | undefined>;
    public readonly inferenceSpecification!: pulumi.Output<outputs.sagemaker.ModelPackageInferenceSpecification | undefined>;
    public readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    public readonly metadataProperties!: pulumi.Output<outputs.sagemaker.ModelPackageMetadataProperties | undefined>;
    public readonly modelApprovalStatus!: pulumi.Output<enums.sagemaker.ModelPackageModelApprovalStatus | undefined>;
    public readonly modelMetrics!: pulumi.Output<outputs.sagemaker.ModelPackageModelMetrics | undefined>;
    public /*out*/ readonly modelPackageArn!: pulumi.Output<string>;
    public readonly modelPackageDescription!: pulumi.Output<string | undefined>;
    public readonly modelPackageGroupName!: pulumi.Output<string | undefined>;
    public readonly modelPackageName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly modelPackageStatus!: pulumi.Output<enums.sagemaker.ModelPackageStatus>;
    public readonly modelPackageStatusDetails!: pulumi.Output<outputs.sagemaker.ModelPackageStatusDetails | undefined>;
    public readonly modelPackageVersion!: pulumi.Output<number | undefined>;
    public readonly samplePayloadUrl!: pulumi.Output<string | undefined>;
    public readonly sourceAlgorithmSpecification!: pulumi.Output<outputs.sagemaker.ModelPackageSourceAlgorithmSpecification | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.ModelPackageTag[] | undefined>;
    public readonly task!: pulumi.Output<string | undefined>;
    public readonly validationSpecification!: pulumi.Output<outputs.sagemaker.ModelPackageValidationSpecification | undefined>;

    /**
     * Create a ModelPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ModelPackageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["additionalInferenceSpecifications"] = args ? args.additionalInferenceSpecifications : undefined;
            resourceInputs["additionalInferenceSpecificationsToAdd"] = args ? args.additionalInferenceSpecificationsToAdd : undefined;
            resourceInputs["approvalDescription"] = args ? args.approvalDescription : undefined;
            resourceInputs["certifyForMarketplace"] = args ? args.certifyForMarketplace : undefined;
            resourceInputs["clientToken"] = args ? args.clientToken : undefined;
            resourceInputs["customerMetadataProperties"] = args ? args.customerMetadataProperties : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["driftCheckBaselines"] = args ? args.driftCheckBaselines : undefined;
            resourceInputs["inferenceSpecification"] = args ? args.inferenceSpecification : undefined;
            resourceInputs["lastModifiedTime"] = args ? args.lastModifiedTime : undefined;
            resourceInputs["metadataProperties"] = args ? args.metadataProperties : undefined;
            resourceInputs["modelApprovalStatus"] = args ? args.modelApprovalStatus : undefined;
            resourceInputs["modelMetrics"] = args ? args.modelMetrics : undefined;
            resourceInputs["modelPackageDescription"] = args ? args.modelPackageDescription : undefined;
            resourceInputs["modelPackageGroupName"] = args ? args.modelPackageGroupName : undefined;
            resourceInputs["modelPackageName"] = args ? args.modelPackageName : undefined;
            resourceInputs["modelPackageStatusDetails"] = args ? args.modelPackageStatusDetails : undefined;
            resourceInputs["modelPackageVersion"] = args ? args.modelPackageVersion : undefined;
            resourceInputs["samplePayloadUrl"] = args ? args.samplePayloadUrl : undefined;
            resourceInputs["sourceAlgorithmSpecification"] = args ? args.sourceAlgorithmSpecification : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["task"] = args ? args.task : undefined;
            resourceInputs["validationSpecification"] = args ? args.validationSpecification : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["modelPackageArn"] = undefined /*out*/;
            resourceInputs["modelPackageStatus"] = undefined /*out*/;
        } else {
            resourceInputs["additionalInferenceSpecifications"] = undefined /*out*/;
            resourceInputs["additionalInferenceSpecificationsToAdd"] = undefined /*out*/;
            resourceInputs["approvalDescription"] = undefined /*out*/;
            resourceInputs["certifyForMarketplace"] = undefined /*out*/;
            resourceInputs["clientToken"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["customerMetadataProperties"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["driftCheckBaselines"] = undefined /*out*/;
            resourceInputs["inferenceSpecification"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["metadataProperties"] = undefined /*out*/;
            resourceInputs["modelApprovalStatus"] = undefined /*out*/;
            resourceInputs["modelMetrics"] = undefined /*out*/;
            resourceInputs["modelPackageArn"] = undefined /*out*/;
            resourceInputs["modelPackageDescription"] = undefined /*out*/;
            resourceInputs["modelPackageGroupName"] = undefined /*out*/;
            resourceInputs["modelPackageName"] = undefined /*out*/;
            resourceInputs["modelPackageStatus"] = undefined /*out*/;
            resourceInputs["modelPackageStatusDetails"] = undefined /*out*/;
            resourceInputs["modelPackageVersion"] = undefined /*out*/;
            resourceInputs["samplePayloadUrl"] = undefined /*out*/;
            resourceInputs["sourceAlgorithmSpecification"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["task"] = undefined /*out*/;
            resourceInputs["validationSpecification"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModelPackage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModelPackage resource.
 */
export interface ModelPackageArgs {
    additionalInferenceSpecifications?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelPackageAdditionalInferenceSpecificationDefinitionArgs>[]>;
    additionalInferenceSpecificationsToAdd?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelPackageAdditionalInferenceSpecificationDefinitionArgs>[]>;
    approvalDescription?: pulumi.Input<string>;
    certifyForMarketplace?: pulumi.Input<boolean>;
    clientToken?: pulumi.Input<string>;
    customerMetadataProperties?: pulumi.Input<inputs.sagemaker.ModelPackageCustomerMetadataPropertiesArgs>;
    domain?: pulumi.Input<string>;
    driftCheckBaselines?: pulumi.Input<inputs.sagemaker.ModelPackageDriftCheckBaselinesArgs>;
    inferenceSpecification?: pulumi.Input<inputs.sagemaker.ModelPackageInferenceSpecificationArgs>;
    lastModifiedTime?: pulumi.Input<string>;
    metadataProperties?: pulumi.Input<inputs.sagemaker.ModelPackageMetadataPropertiesArgs>;
    modelApprovalStatus?: pulumi.Input<enums.sagemaker.ModelPackageModelApprovalStatus>;
    modelMetrics?: pulumi.Input<inputs.sagemaker.ModelPackageModelMetricsArgs>;
    modelPackageDescription?: pulumi.Input<string>;
    modelPackageGroupName?: pulumi.Input<string>;
    modelPackageName?: pulumi.Input<string>;
    modelPackageStatusDetails?: pulumi.Input<inputs.sagemaker.ModelPackageStatusDetailsArgs>;
    modelPackageVersion?: pulumi.Input<number>;
    samplePayloadUrl?: pulumi.Input<string>;
    sourceAlgorithmSpecification?: pulumi.Input<inputs.sagemaker.ModelPackageSourceAlgorithmSpecificationArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelPackageTagArgs>[]>;
    task?: pulumi.Input<string>;
    validationSpecification?: pulumi.Input<inputs.sagemaker.ModelPackageValidationSpecificationArgs>;
}
