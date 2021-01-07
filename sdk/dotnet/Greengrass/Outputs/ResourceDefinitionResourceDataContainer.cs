// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Outputs
{

    [OutputType]
    public sealed class ResourceDefinitionResourceDataContainer
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html#cfn-greengrass-resourcedefinition-resourcedatacontainer-localdeviceresourcedata
        /// </summary>
        public readonly Outputs.ResourceDefinitionLocalDeviceResourceData? LocalDeviceResourceData;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html#cfn-greengrass-resourcedefinition-resourcedatacontainer-localvolumeresourcedata
        /// </summary>
        public readonly Outputs.ResourceDefinitionLocalVolumeResourceData? LocalVolumeResourceData;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html#cfn-greengrass-resourcedefinition-resourcedatacontainer-s3machinelearningmodelresourcedata
        /// </summary>
        public readonly Outputs.ResourceDefinitionS3MachineLearningModelResourceData? S3MachineLearningModelResourceData;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html#cfn-greengrass-resourcedefinition-resourcedatacontainer-sagemakermachinelearningmodelresourcedata
        /// </summary>
        public readonly Outputs.ResourceDefinitionSageMakerMachineLearningModelResourceData? SageMakerMachineLearningModelResourceData;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html#cfn-greengrass-resourcedefinition-resourcedatacontainer-secretsmanagersecretresourcedata
        /// </summary>
        public readonly Outputs.ResourceDefinitionSecretsManagerSecretResourceData? SecretsManagerSecretResourceData;

        [OutputConstructor]
        private ResourceDefinitionResourceDataContainer(
            Outputs.ResourceDefinitionLocalDeviceResourceData? LocalDeviceResourceData,

            Outputs.ResourceDefinitionLocalVolumeResourceData? LocalVolumeResourceData,

            Outputs.ResourceDefinitionS3MachineLearningModelResourceData? S3MachineLearningModelResourceData,

            Outputs.ResourceDefinitionSageMakerMachineLearningModelResourceData? SageMakerMachineLearningModelResourceData,

            Outputs.ResourceDefinitionSecretsManagerSecretResourceData? SecretsManagerSecretResourceData)
        {
            this.LocalDeviceResourceData = LocalDeviceResourceData;
            this.LocalVolumeResourceData = LocalVolumeResourceData;
            this.S3MachineLearningModelResourceData = S3MachineLearningModelResourceData;
            this.SageMakerMachineLearningModelResourceData = SageMakerMachineLearningModelResourceData;
            this.SecretsManagerSecretResourceData = SecretsManagerSecretResourceData;
        }
    }
}