// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Outputs
{

    [OutputType]
    public sealed class ResourceDefinitionVersionResourceDataContainer
    {
        public readonly Outputs.ResourceDefinitionVersionLocalDeviceResourceData? LocalDeviceResourceData;
        public readonly Outputs.ResourceDefinitionVersionLocalVolumeResourceData? LocalVolumeResourceData;
        public readonly Outputs.ResourceDefinitionVersionS3MachineLearningModelResourceData? S3MachineLearningModelResourceData;
        public readonly Outputs.ResourceDefinitionVersionSageMakerMachineLearningModelResourceData? SageMakerMachineLearningModelResourceData;
        public readonly Outputs.ResourceDefinitionVersionSecretsManagerSecretResourceData? SecretsManagerSecretResourceData;

        [OutputConstructor]
        private ResourceDefinitionVersionResourceDataContainer(
            Outputs.ResourceDefinitionVersionLocalDeviceResourceData? localDeviceResourceData,

            Outputs.ResourceDefinitionVersionLocalVolumeResourceData? localVolumeResourceData,

            Outputs.ResourceDefinitionVersionS3MachineLearningModelResourceData? s3MachineLearningModelResourceData,

            Outputs.ResourceDefinitionVersionSageMakerMachineLearningModelResourceData? sageMakerMachineLearningModelResourceData,

            Outputs.ResourceDefinitionVersionSecretsManagerSecretResourceData? secretsManagerSecretResourceData)
        {
            LocalDeviceResourceData = localDeviceResourceData;
            LocalVolumeResourceData = localVolumeResourceData;
            S3MachineLearningModelResourceData = s3MachineLearningModelResourceData;
            SageMakerMachineLearningModelResourceData = sageMakerMachineLearningModelResourceData;
            SecretsManagerSecretResourceData = secretsManagerSecretResourceData;
        }
    }
}