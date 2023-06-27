// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Describes the Docker container for the model package.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageContainerDefinition
    {
        /// <summary>
        /// The DNS host name for the Docker container.
        /// </summary>
        public readonly string? ContainerHostname;
        public readonly Outputs.ModelPackageEnvironment? Environment;
        /// <summary>
        /// The machine learning framework of the model package container image.
        /// </summary>
        public readonly string? Framework;
        /// <summary>
        /// The framework version of the Model Package Container Image.
        /// </summary>
        public readonly string? FrameworkVersion;
        /// <summary>
        /// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// An MD5 hash of the training algorithm that identifies the Docker image used for training.
        /// </summary>
        public readonly string? ImageDigest;
        /// <summary>
        /// A structure with Model Input details.
        /// </summary>
        public readonly string? ModelDataUrl;
        public readonly Outputs.ModelPackageContainerDefinitionModelInputProperties? ModelInput;
        /// <summary>
        /// The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.
        /// </summary>
        public readonly string? NearestModelName;

        [OutputConstructor]
        private ModelPackageContainerDefinition(
            string? containerHostname,

            Outputs.ModelPackageEnvironment? environment,

            string? framework,

            string? frameworkVersion,

            string image,

            string? imageDigest,

            string? modelDataUrl,

            Outputs.ModelPackageContainerDefinitionModelInputProperties? modelInput,

            string? nearestModelName)
        {
            ContainerHostname = containerHostname;
            Environment = environment;
            Framework = framework;
            FrameworkVersion = frameworkVersion;
            Image = image;
            ImageDigest = imageDigest;
            ModelDataUrl = modelDataUrl;
            ModelInput = modelInput;
            NearestModelName = nearestModelName;
        }
    }
}
