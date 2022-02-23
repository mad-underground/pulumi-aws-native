// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift
{
    public static class GetScheduledAction
    {
        /// <summary>
        /// The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.
        /// </summary>
        public static Task<GetScheduledActionResult> InvokeAsync(GetScheduledActionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScheduledActionResult>("aws-native:redshift:getScheduledAction", args ?? new GetScheduledActionArgs(), options.WithDefaults());

        /// <summary>
        /// The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.
        /// </summary>
        public static Output<GetScheduledActionResult> Invoke(GetScheduledActionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetScheduledActionResult>("aws-native:redshift:getScheduledAction", args ?? new GetScheduledActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScheduledActionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the scheduled action. The name must be unique within an account.
        /// </summary>
        [Input("scheduledActionName", required: true)]
        public string ScheduledActionName { get; set; } = null!;

        public GetScheduledActionArgs()
        {
        }
    }

    public sealed class GetScheduledActionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the scheduled action. The name must be unique within an account.
        /// </summary>
        [Input("scheduledActionName", required: true)]
        public Input<string> ScheduledActionName { get; set; } = null!;

        public GetScheduledActionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetScheduledActionResult
    {
        /// <summary>
        /// If true, the schedule is enabled. If false, the scheduled action does not trigger.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// The IAM role to assume to run the target action.
        /// </summary>
        public readonly string? IamRole;
        /// <summary>
        /// List of times when the scheduled action will run.
        /// </summary>
        public readonly ImmutableArray<string> NextInvocations;
        /// <summary>
        /// The schedule in `at( )` or `cron( )` format.
        /// </summary>
        public readonly string? Schedule;
        /// <summary>
        /// The description of the scheduled action.
        /// </summary>
        public readonly string? ScheduledActionDescription;
        /// <summary>
        /// The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// The state of the scheduled action.
        /// </summary>
        public readonly Pulumi.AwsNative.Redshift.ScheduledActionState? State;
        /// <summary>
        /// A JSON format string of the Amazon Redshift API operation with input parameters.
        /// </summary>
        public readonly Outputs.ScheduledActionType? TargetAction;

        [OutputConstructor]
        private GetScheduledActionResult(
            bool? enable,

            string? endTime,

            string? iamRole,

            ImmutableArray<string> nextInvocations,

            string? schedule,

            string? scheduledActionDescription,

            string? startTime,

            Pulumi.AwsNative.Redshift.ScheduledActionState? state,

            Outputs.ScheduledActionType? targetAction)
        {
            Enable = enable;
            EndTime = endTime;
            IamRole = iamRole;
            NextInvocations = nextInvocations;
            Schedule = schedule;
            ScheduledActionDescription = scheduledActionDescription;
            StartTime = startTime;
            State = state;
            TargetAction = targetAction;
        }
    }
}