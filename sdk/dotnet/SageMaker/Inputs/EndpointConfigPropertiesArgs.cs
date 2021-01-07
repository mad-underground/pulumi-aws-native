// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
    /// </summary>
    public sealed class EndpointConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-datacaptureconfig
        /// </summary>
        [Input("DataCaptureConfig")]
        public Input<Inputs.EndpointConfigDataCaptureConfigArgs>? DataCaptureConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-endpointconfigname
        /// </summary>
        [Input("EndpointConfigName")]
        public Input<string>? EndpointConfigName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-kmskeyid
        /// </summary>
        [Input("KmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("ProductionVariants", required: true)]
        private InputList<Inputs.EndpointConfigProductionVariantArgs>? _ProductionVariants;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-productionvariants
        /// </summary>
        public InputList<Inputs.EndpointConfigProductionVariantArgs> ProductionVariants
        {
            get => _ProductionVariants ?? (_ProductionVariants = new InputList<Inputs.EndpointConfigProductionVariantArgs>());
            set => _ProductionVariants = value;
        }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public EndpointConfigPropertiesArgs()
        {
        }
    }
}