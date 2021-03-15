// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class DataQualityJobDefinitionDataQualityBaselineConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-baseliningjobname
        /// </summary>
        public readonly string? BaseliningJobName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-constraintsresource
        /// </summary>
        public readonly Outputs.DataQualityJobDefinitionConstraintsResource? ConstraintsResource;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig-statisticsresource
        /// </summary>
        public readonly Outputs.DataQualityJobDefinitionStatisticsResource? StatisticsResource;

        [OutputConstructor]
        private DataQualityJobDefinitionDataQualityBaselineConfig(
            string? baseliningJobName,

            Outputs.DataQualityJobDefinitionConstraintsResource? constraintsResource,

            Outputs.DataQualityJobDefinitionStatisticsResource? statisticsResource)
        {
            BaseliningJobName = baseliningJobName;
            ConstraintsResource = constraintsResource;
            StatisticsResource = statisticsResource;
        }
    }
}
