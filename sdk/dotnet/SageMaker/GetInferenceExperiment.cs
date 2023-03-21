// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetInferenceExperiment
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::InferenceExperiment
        /// </summary>
        public static Task<GetInferenceExperimentResult> InvokeAsync(GetInferenceExperimentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInferenceExperimentResult>("aws-native:sagemaker:getInferenceExperiment", args ?? new GetInferenceExperimentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::InferenceExperiment
        /// </summary>
        public static Output<GetInferenceExperimentResult> Invoke(GetInferenceExperimentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInferenceExperimentResult>("aws-native:sagemaker:getInferenceExperiment", args ?? new GetInferenceExperimentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInferenceExperimentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for the inference experiment.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetInferenceExperimentArgs()
        {
        }
        public static new GetInferenceExperimentArgs Empty => new GetInferenceExperimentArgs();
    }

    public sealed class GetInferenceExperimentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for the inference experiment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetInferenceExperimentInvokeArgs()
        {
        }
        public static new GetInferenceExperimentInvokeArgs Empty => new GetInferenceExperimentInvokeArgs();
    }


    [OutputType]
    public sealed class GetInferenceExperimentResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the inference experiment.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The timestamp at which you created the inference experiment.
        /// </summary>
        public readonly string? CreationTime;
        public readonly Outputs.InferenceExperimentDataStorageConfig? DataStorageConfig;
        /// <summary>
        /// The description of the inference experiment.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The desired state of the experiment after starting or stopping operation.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.InferenceExperimentDesiredState? DesiredState;
        public readonly Outputs.InferenceExperimentEndpointMetadata? EndpointMetadata;
        /// <summary>
        /// The timestamp at which you last modified the inference experiment.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
        /// </summary>
        public readonly ImmutableArray<Outputs.InferenceExperimentModelVariantConfig> ModelVariants;
        public readonly Outputs.InferenceExperimentSchedule? Schedule;
        public readonly Outputs.InferenceExperimentShadowModeConfig? ShadowModeConfig;
        /// <summary>
        /// The status of the inference experiment.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.InferenceExperimentStatus? Status;
        /// <summary>
        /// The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
        /// </summary>
        public readonly string? StatusReason;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.InferenceExperimentTag> Tags;

        [OutputConstructor]
        private GetInferenceExperimentResult(
            string? arn,

            string? creationTime,

            Outputs.InferenceExperimentDataStorageConfig? dataStorageConfig,

            string? description,

            Pulumi.AwsNative.SageMaker.InferenceExperimentDesiredState? desiredState,

            Outputs.InferenceExperimentEndpointMetadata? endpointMetadata,

            string? lastModifiedTime,

            ImmutableArray<Outputs.InferenceExperimentModelVariantConfig> modelVariants,

            Outputs.InferenceExperimentSchedule? schedule,

            Outputs.InferenceExperimentShadowModeConfig? shadowModeConfig,

            Pulumi.AwsNative.SageMaker.InferenceExperimentStatus? status,

            string? statusReason,

            ImmutableArray<Outputs.InferenceExperimentTag> tags)
        {
            Arn = arn;
            CreationTime = creationTime;
            DataStorageConfig = dataStorageConfig;
            Description = description;
            DesiredState = desiredState;
            EndpointMetadata = endpointMetadata;
            LastModifiedTime = lastModifiedTime;
            ModelVariants = modelVariants;
            Schedule = schedule;
            ShadowModeConfig = shadowModeConfig;
            Status = status;
            StatusReason = statusReason;
            Tags = tags;
        }
    }
}
