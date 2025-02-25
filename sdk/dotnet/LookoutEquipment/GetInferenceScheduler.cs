// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutEquipment
{
    public static class GetInferenceScheduler
    {
        /// <summary>
        /// Resource schema for LookoutEquipment InferenceScheduler.
        /// </summary>
        public static Task<GetInferenceSchedulerResult> InvokeAsync(GetInferenceSchedulerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInferenceSchedulerResult>("aws-native:lookoutequipment:getInferenceScheduler", args ?? new GetInferenceSchedulerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for LookoutEquipment InferenceScheduler.
        /// </summary>
        public static Output<GetInferenceSchedulerResult> Invoke(GetInferenceSchedulerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInferenceSchedulerResult>("aws-native:lookoutequipment:getInferenceScheduler", args ?? new GetInferenceSchedulerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInferenceSchedulerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the inference scheduler being created.
        /// </summary>
        [Input("inferenceSchedulerName", required: true)]
        public string InferenceSchedulerName { get; set; } = null!;

        public GetInferenceSchedulerArgs()
        {
        }
        public static new GetInferenceSchedulerArgs Empty => new GetInferenceSchedulerArgs();
    }

    public sealed class GetInferenceSchedulerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the inference scheduler being created.
        /// </summary>
        [Input("inferenceSchedulerName", required: true)]
        public Input<string> InferenceSchedulerName { get; set; } = null!;

        public GetInferenceSchedulerInvokeArgs()
        {
        }
        public static new GetInferenceSchedulerInvokeArgs Empty => new GetInferenceSchedulerInvokeArgs();
    }


    [OutputType]
    public sealed class GetInferenceSchedulerResult
    {
        /// <summary>
        /// A period of time (in minutes) by which inference on the data is delayed after the data starts.
        /// </summary>
        public readonly int? DataDelayOffsetInMinutes;
        /// <summary>
        /// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
        /// </summary>
        public readonly Outputs.DataInputConfigurationProperties? DataInputConfiguration;
        /// <summary>
        /// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
        /// </summary>
        public readonly Outputs.DataOutputConfigurationProperties? DataOutputConfiguration;
        /// <summary>
        /// How often data is uploaded to the source S3 bucket for the input data.
        /// </summary>
        public readonly Pulumi.AwsNative.LookoutEquipment.InferenceSchedulerDataUploadFrequency? DataUploadFrequency;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the inference scheduler being created.
        /// </summary>
        public readonly string? InferenceSchedulerArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// Any tags associated with the inference scheduler.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetInferenceSchedulerResult(
            int? dataDelayOffsetInMinutes,

            Outputs.DataInputConfigurationProperties? dataInputConfiguration,

            Outputs.DataOutputConfigurationProperties? dataOutputConfiguration,

            Pulumi.AwsNative.LookoutEquipment.InferenceSchedulerDataUploadFrequency? dataUploadFrequency,

            string? inferenceSchedulerArn,

            string? roleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DataDelayOffsetInMinutes = dataDelayOffsetInMinutes;
            DataInputConfiguration = dataInputConfiguration;
            DataOutputConfiguration = dataOutputConfiguration;
            DataUploadFrequency = dataUploadFrequency;
            InferenceSchedulerArn = inferenceSchedulerArn;
            RoleArn = roleArn;
            Tags = tags;
        }
    }
}
