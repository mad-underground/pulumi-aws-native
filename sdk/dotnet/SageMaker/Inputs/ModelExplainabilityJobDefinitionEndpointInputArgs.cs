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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html
    /// </summary>
    public sealed class ModelExplainabilityJobDefinitionEndpointInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-endpointname
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-featuresattribute
        /// </summary>
        [Input("featuresAttribute")]
        public Input<string>? FeaturesAttribute { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-inferenceattribute
        /// </summary>
        [Input("inferenceAttribute")]
        public Input<string>? InferenceAttribute { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-localpath
        /// </summary>
        [Input("localPath", required: true)]
        public Input<string> LocalPath { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-probabilityattribute
        /// </summary>
        [Input("probabilityAttribute")]
        public Input<string>? ProbabilityAttribute { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-s3datadistributiontype
        /// </summary>
        [Input("s3DataDistributionType")]
        public Input<string>? S3DataDistributionType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-endpointinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointinput-s3inputmode
        /// </summary>
        [Input("s3InputMode")]
        public Input<string>? S3InputMode { get; set; }

        public ModelExplainabilityJobDefinitionEndpointInputArgs()
        {
        }
    }
}
