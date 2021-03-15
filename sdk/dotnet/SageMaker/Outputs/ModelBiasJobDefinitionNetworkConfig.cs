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
    public sealed class ModelBiasJobDefinitionNetworkConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-networkconfig.html#cfn-sagemaker-modelbiasjobdefinition-networkconfig-enableintercontainertrafficencryption
        /// </summary>
        public readonly bool? EnableInterContainerTrafficEncryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-networkconfig.html#cfn-sagemaker-modelbiasjobdefinition-networkconfig-enablenetworkisolation
        /// </summary>
        public readonly bool? EnableNetworkIsolation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-networkconfig.html#cfn-sagemaker-modelbiasjobdefinition-networkconfig-vpcconfig
        /// </summary>
        public readonly Outputs.ModelBiasJobDefinitionVpcConfig? VpcConfig;

        [OutputConstructor]
        private ModelBiasJobDefinitionNetworkConfig(
            bool? enableInterContainerTrafficEncryption,

            bool? enableNetworkIsolation,

            Outputs.ModelBiasJobDefinitionVpcConfig? vpcConfig)
        {
            EnableInterContainerTrafficEncryption = enableInterContainerTrafficEncryption;
            EnableNetworkIsolation = enableNetworkIsolation;
            VpcConfig = vpcConfig;
        }
    }
}
