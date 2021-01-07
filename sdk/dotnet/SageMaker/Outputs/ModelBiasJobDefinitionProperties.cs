// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class ModelBiasJobDefinitionProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-jobdefinitionname
        /// </summary>
        public readonly string? JobDefinitionName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-jobresources
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionMonitoringResources JobResources;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionModelBiasAppSpecification ModelBiasAppSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasbaselineconfig
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionModelBiasBaselineConfig? ModelBiasBaselineConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionModelBiasJobInput ModelBiasJobInput;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjoboutputconfig
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionMonitoringOutputConfig ModelBiasJobOutputConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-networkconfig
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionNetworkConfig? NetworkConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-rolearn
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-stoppingcondition
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionStoppingCondition? StoppingCondition;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private ModelBiasJobDefinitionProperties(
            string? JobDefinitionName,

            Outputs.ModelBiasJobDefinitionMonitoringResources JobResources,

            Outputs.ModelBiasJobDefinitionModelBiasAppSpecification ModelBiasAppSpecification,

            Outputs.ModelBiasJobDefinitionModelBiasBaselineConfig? ModelBiasBaselineConfig,

            Outputs.ModelBiasJobDefinitionModelBiasJobInput ModelBiasJobInput,

            Outputs.ModelBiasJobDefinitionMonitoringOutputConfig ModelBiasJobOutputConfig,

            Outputs.ModelBiasJobDefinitionNetworkConfig? NetworkConfig,

            string RoleArn,

            Outputs.ModelBiasJobDefinitionStoppingCondition? StoppingCondition,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.JobDefinitionName = JobDefinitionName;
            this.JobResources = JobResources;
            this.ModelBiasAppSpecification = ModelBiasAppSpecification;
            this.ModelBiasBaselineConfig = ModelBiasBaselineConfig;
            this.ModelBiasJobInput = ModelBiasJobInput;
            this.ModelBiasJobOutputConfig = ModelBiasJobOutputConfig;
            this.NetworkConfig = NetworkConfig;
            this.RoleArn = RoleArn;
            this.StoppingCondition = StoppingCondition;
            this.Tags = Tags;
        }
    }
}
