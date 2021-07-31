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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-monitoringoutput.html
    /// </summary>
    [OutputType]
    public sealed class ModelBiasJobDefinitionMonitoringOutput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-monitoringoutput.html#cfn-sagemaker-modelbiasjobdefinition-monitoringoutput-s3output
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionS3Output S3Output;

        [OutputConstructor]
        private ModelBiasJobDefinitionMonitoringOutput(Outputs.ModelBiasJobDefinitionS3Output s3Output)
        {
            S3Output = s3Output;
        }
    }
}
