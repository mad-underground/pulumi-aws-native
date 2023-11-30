// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class InferenceComponentContainerSpecification
    {
        public readonly string? ArtifactUrl;
        public readonly Outputs.InferenceComponentDeployedImage? DeployedImage;
        public readonly Outputs.InferenceComponentEnvironmentMap? Environment;
        public readonly string? Image;

        [OutputConstructor]
        private InferenceComponentContainerSpecification(
            string? artifactUrl,

            Outputs.InferenceComponentDeployedImage? deployedImage,

            Outputs.InferenceComponentEnvironmentMap? environment,

            string? image)
        {
            ArtifactUrl = artifactUrl;
            DeployedImage = deployedImage;
            Environment = environment;
            Image = image;
        }
    }
}
