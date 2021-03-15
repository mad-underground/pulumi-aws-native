// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html
    /// </summary>
    public sealed class ModelBiasJobDefinitionModelBiasJobInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput-endpointinput
        /// </summary>
        [Input("endpointInput", required: true)]
        public Input<Inputs.ModelBiasJobDefinitionEndpointInputArgs> EndpointInput { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-modelbiasjobinput.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput-groundtruths3input
        /// </summary>
        [Input("groundTruthS3Input", required: true)]
        public Input<Inputs.ModelBiasJobDefinitionMonitoringGroundTruthS3InputArgs> GroundTruthS3Input { get; set; } = null!;

        public ModelBiasJobDefinitionModelBiasJobInputArgs()
        {
        }
    }
}
