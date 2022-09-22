// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeStarNotifications
{
    public static class GetNotificationRule
    {
        /// <summary>
        /// Resource Type definition for AWS::CodeStarNotifications::NotificationRule
        /// </summary>
        public static Task<GetNotificationRuleResult> InvokeAsync(GetNotificationRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotificationRuleResult>("aws-native:codestarnotifications:getNotificationRule", args ?? new GetNotificationRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CodeStarNotifications::NotificationRule
        /// </summary>
        public static Output<GetNotificationRuleResult> Invoke(GetNotificationRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotificationRuleResult>("aws-native:codestarnotifications:getNotificationRule", args ?? new GetNotificationRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetNotificationRuleArgs()
        {
        }
        public static new GetNotificationRuleArgs Empty => new GetNotificationRuleArgs();
    }

    public sealed class GetNotificationRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetNotificationRuleInvokeArgs()
        {
        }
        public static new GetNotificationRuleInvokeArgs Empty => new GetNotificationRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotificationRuleResult
    {
        public readonly string? Arn;
        public readonly string? CreatedBy;
        public readonly Pulumi.AwsNative.CodeStarNotifications.NotificationRuleDetailType? DetailType;
        public readonly ImmutableArray<string> EventTypeIds;
        public readonly string? Name;
        public readonly Pulumi.AwsNative.CodeStarNotifications.NotificationRuleStatus? Status;
        public readonly ImmutableArray<Outputs.NotificationRuleTarget> Targets;

        [OutputConstructor]
        private GetNotificationRuleResult(
            string? arn,

            string? createdBy,

            Pulumi.AwsNative.CodeStarNotifications.NotificationRuleDetailType? detailType,

            ImmutableArray<string> eventTypeIds,

            string? name,

            Pulumi.AwsNative.CodeStarNotifications.NotificationRuleStatus? status,

            ImmutableArray<Outputs.NotificationRuleTarget> targets)
        {
            Arn = arn;
            CreatedBy = createdBy;
            DetailType = detailType;
            EventTypeIds = eventTypeIds;
            Name = name;
            Status = status;
            Targets = targets;
        }
    }
}
