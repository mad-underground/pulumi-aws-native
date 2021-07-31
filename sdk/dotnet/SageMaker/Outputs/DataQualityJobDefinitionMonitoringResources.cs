// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-monitoringresources.html
    /// </summary>
    [OutputType]
    public sealed class DataQualityJobDefinitionMonitoringResources
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-monitoringresources.html#cfn-sagemaker-dataqualityjobdefinition-monitoringresources-clusterconfig
        /// </summary>
        public readonly Outputs.DataQualityJobDefinitionClusterConfig ClusterConfig;

        [OutputConstructor]
        private DataQualityJobDefinitionMonitoringResources(Outputs.DataQualityJobDefinitionClusterConfig clusterConfig)
        {
            ClusterConfig = clusterConfig;
        }
    }
}
