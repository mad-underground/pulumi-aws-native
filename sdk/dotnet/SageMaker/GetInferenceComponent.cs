// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetInferenceComponent
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::InferenceComponent
        /// </summary>
        public static Task<GetInferenceComponentResult> InvokeAsync(GetInferenceComponentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInferenceComponentResult>("aws-native:sagemaker:getInferenceComponent", args ?? new GetInferenceComponentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::InferenceComponent
        /// </summary>
        public static Output<GetInferenceComponentResult> Invoke(GetInferenceComponentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInferenceComponentResult>("aws-native:sagemaker:getInferenceComponent", args ?? new GetInferenceComponentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInferenceComponentArgs : global::Pulumi.InvokeArgs
    {
        [Input("inferenceComponentArn", required: true)]
        public string InferenceComponentArn { get; set; } = null!;

        public GetInferenceComponentArgs()
        {
        }
        public static new GetInferenceComponentArgs Empty => new GetInferenceComponentArgs();
    }

    public sealed class GetInferenceComponentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("inferenceComponentArn", required: true)]
        public Input<string> InferenceComponentArn { get; set; } = null!;

        public GetInferenceComponentInvokeArgs()
        {
        }
        public static new GetInferenceComponentInvokeArgs Empty => new GetInferenceComponentInvokeArgs();
    }


    [OutputType]
    public sealed class GetInferenceComponentResult
    {
        public readonly string? CreationTime;
        public readonly string? EndpointArn;
        public readonly string? EndpointName;
        public readonly string? FailureReason;
        public readonly string? InferenceComponentArn;
        public readonly string? InferenceComponentName;
        public readonly Pulumi.AwsNative.SageMaker.InferenceComponentStatus? InferenceComponentStatus;
        public readonly string? LastModifiedTime;
        public readonly Outputs.InferenceComponentRuntimeConfig? RuntimeConfig;
        public readonly Outputs.InferenceComponentSpecification? Specification;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly string? VariantName;

        [OutputConstructor]
        private GetInferenceComponentResult(
            string? creationTime,

            string? endpointArn,

            string? endpointName,

            string? failureReason,

            string? inferenceComponentArn,

            string? inferenceComponentName,

            Pulumi.AwsNative.SageMaker.InferenceComponentStatus? inferenceComponentStatus,

            string? lastModifiedTime,

            Outputs.InferenceComponentRuntimeConfig? runtimeConfig,

            Outputs.InferenceComponentSpecification? specification,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? variantName)
        {
            CreationTime = creationTime;
            EndpointArn = endpointArn;
            EndpointName = endpointName;
            FailureReason = failureReason;
            InferenceComponentArn = inferenceComponentArn;
            InferenceComponentName = inferenceComponentName;
            InferenceComponentStatus = inferenceComponentStatus;
            LastModifiedTime = lastModifiedTime;
            RuntimeConfig = runtimeConfig;
            Specification = specification;
            Tags = tags;
            VariantName = variantName;
        }
    }
}
