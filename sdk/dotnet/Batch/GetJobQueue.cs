// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch
{
    public static class GetJobQueue
    {
        /// <summary>
        /// Resource Type definition for AWS::Batch::JobQueue
        /// </summary>
        public static Task<GetJobQueueResult> InvokeAsync(GetJobQueueArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobQueueResult>("aws-native:batch:getJobQueue", args ?? new GetJobQueueArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Batch::JobQueue
        /// </summary>
        public static Output<GetJobQueueResult> Invoke(GetJobQueueInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetJobQueueResult>("aws-native:batch:getJobQueue", args ?? new GetJobQueueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobQueueArgs : Pulumi.InvokeArgs
    {
        [Input("jobQueueArn", required: true)]
        public string JobQueueArn { get; set; } = null!;

        public GetJobQueueArgs()
        {
        }
    }

    public sealed class GetJobQueueInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("jobQueueArn", required: true)]
        public Input<string> JobQueueArn { get; set; } = null!;

        public GetJobQueueInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobQueueResult
    {
        public readonly ImmutableArray<Outputs.JobQueueComputeEnvironmentOrder> ComputeEnvironmentOrder;
        public readonly string? JobQueueArn;
        public readonly int? Priority;
        public readonly string? SchedulingPolicyArn;
        public readonly Pulumi.AwsNative.Batch.JobQueueState? State;

        [OutputConstructor]
        private GetJobQueueResult(
            ImmutableArray<Outputs.JobQueueComputeEnvironmentOrder> computeEnvironmentOrder,

            string? jobQueueArn,

            int? priority,

            string? schedulingPolicyArn,

            Pulumi.AwsNative.Batch.JobQueueState? state)
        {
            ComputeEnvironmentOrder = computeEnvironmentOrder;
            JobQueueArn = jobQueueArn;
            Priority = priority;
            SchedulingPolicyArn = schedulingPolicyArn;
            State = state;
        }
    }
}