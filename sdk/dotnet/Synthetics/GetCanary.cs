// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics
{
    public static class GetCanary
    {
        /// <summary>
        /// Resource Type definition for AWS::Synthetics::Canary
        /// </summary>
        public static Task<GetCanaryResult> InvokeAsync(GetCanaryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCanaryResult>("aws-native:synthetics:getCanary", args ?? new GetCanaryArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Synthetics::Canary
        /// </summary>
        public static Output<GetCanaryResult> Invoke(GetCanaryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCanaryResult>("aws-native:synthetics:getCanary", args ?? new GetCanaryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCanaryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the canary.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCanaryArgs()
        {
        }
        public static new GetCanaryArgs Empty => new GetCanaryArgs();
    }

    public sealed class GetCanaryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the canary.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCanaryInvokeArgs()
        {
        }
        public static new GetCanaryInvokeArgs Empty => new GetCanaryInvokeArgs();
    }


    [OutputType]
    public sealed class GetCanaryResult
    {
        /// <summary>
        /// Provide artifact configuration
        /// </summary>
        public readonly Outputs.CanaryArtifactConfig? ArtifactConfig;
        /// <summary>
        /// Provide the s3 bucket output location for test results
        /// </summary>
        public readonly string? ArtifactS3Location;
        /// <summary>
        /// Provide the canary script source
        /// </summary>
        public readonly Outputs.CanaryCode? Code;
        /// <summary>
        /// Deletes associated lambda resources created by Synthetics if set to True. Default is False
        /// </summary>
        public readonly bool? DeleteLambdaResourcesOnCanaryDeletion;
        /// <summary>
        /// Lambda Execution role used to run your canaries
        /// </summary>
        public readonly string? ExecutionRoleArn;
        /// <summary>
        /// Retention period of failed canary runs represented in number of days
        /// </summary>
        public readonly int? FailureRetentionPeriod;
        /// <summary>
        /// Id of the canary
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Provide canary run configuration
        /// </summary>
        public readonly Outputs.CanaryRunConfig? RunConfig;
        /// <summary>
        /// Runtime version of Synthetics Library
        /// </summary>
        public readonly string? RuntimeVersion;
        /// <summary>
        /// Frequency to run your canaries
        /// </summary>
        public readonly Outputs.CanarySchedule? Schedule;
        /// <summary>
        /// Runs canary if set to True. Default is False
        /// </summary>
        public readonly bool? StartCanaryAfterCreation;
        /// <summary>
        /// State of the canary
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Retention period of successful canary runs represented in number of days
        /// </summary>
        public readonly int? SuccessRetentionPeriod;
        public readonly ImmutableArray<Outputs.CanaryTag> Tags;
        /// <summary>
        /// Provide VPC Configuration if enabled.
        /// </summary>
        public readonly Outputs.CanaryVPCConfig? VPCConfig;
        /// <summary>
        /// Visual reference configuration for visual testing
        /// </summary>
        public readonly Outputs.CanaryVisualReference? VisualReference;

        [OutputConstructor]
        private GetCanaryResult(
            Outputs.CanaryArtifactConfig? artifactConfig,

            string? artifactS3Location,

            Outputs.CanaryCode? code,

            bool? deleteLambdaResourcesOnCanaryDeletion,

            string? executionRoleArn,

            int? failureRetentionPeriod,

            string? id,

            Outputs.CanaryRunConfig? runConfig,

            string? runtimeVersion,

            Outputs.CanarySchedule? schedule,

            bool? startCanaryAfterCreation,

            string? state,

            int? successRetentionPeriod,

            ImmutableArray<Outputs.CanaryTag> tags,

            Outputs.CanaryVPCConfig? vPCConfig,

            Outputs.CanaryVisualReference? visualReference)
        {
            ArtifactConfig = artifactConfig;
            ArtifactS3Location = artifactS3Location;
            Code = code;
            DeleteLambdaResourcesOnCanaryDeletion = deleteLambdaResourcesOnCanaryDeletion;
            ExecutionRoleArn = executionRoleArn;
            FailureRetentionPeriod = failureRetentionPeriod;
            Id = id;
            RunConfig = runConfig;
            RuntimeVersion = runtimeVersion;
            Schedule = schedule;
            StartCanaryAfterCreation = startCanaryAfterCreation;
            State = state;
            SuccessRetentionPeriod = successRetentionPeriod;
            Tags = tags;
            VPCConfig = vPCConfig;
            VisualReference = visualReference;
        }
    }
}
