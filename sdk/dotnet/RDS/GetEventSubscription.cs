// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetEventSubscription
    {
        /// <summary>
        /// The AWS::RDS::EventSubscription resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see Using Amazon RDS Event Notification in the Amazon RDS User Guide.
        /// </summary>
        public static Task<GetEventSubscriptionResult> InvokeAsync(GetEventSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventSubscriptionResult>("aws-native:rds:getEventSubscription", args ?? new GetEventSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::RDS::EventSubscription resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see Using Amazon RDS Event Notification in the Amazon RDS User Guide.
        /// </summary>
        public static Output<GetEventSubscriptionResult> Invoke(GetEventSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEventSubscriptionResult>("aws-native:rds:getEventSubscription", args ?? new GetEventSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public string SubscriptionName { get; set; } = null!;

        public GetEventSubscriptionArgs()
        {
        }
    }

    public sealed class GetEventSubscriptionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the subscription.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        public GetEventSubscriptionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventSubscriptionResult
    {
        /// <summary>
        /// A Boolean value; set to true to activate the subscription, set to false to create the subscription but not active it.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// A list of event categories for a SourceType that you want to subscribe to. You can see a list of the categories for a given SourceType in the Events topic in the Amazon RDS User Guide or by using the DescribeEventCategories action.
        /// </summary>
        public readonly ImmutableArray<string> EventCategories;
        /// <summary>
        /// The list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it cannot end with a hyphen or contain two consecutive hyphens.
        /// </summary>
        public readonly ImmutableArray<string> SourceIds;
        /// <summary>
        /// The type of source that will be generating the events. For example, if you want to be notified of events generated by a DB instance, you would set this parameter to db-instance. if this value is not specified, all events are returned.
        /// </summary>
        public readonly Pulumi.AwsNative.RDS.EventSubscriptionSourceType? SourceType;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventSubscriptionTag> Tags;

        [OutputConstructor]
        private GetEventSubscriptionResult(
            bool? enabled,

            ImmutableArray<string> eventCategories,

            ImmutableArray<string> sourceIds,

            Pulumi.AwsNative.RDS.EventSubscriptionSourceType? sourceType,

            ImmutableArray<Outputs.EventSubscriptionTag> tags)
        {
            Enabled = enabled;
            EventCategories = eventCategories;
            SourceIds = sourceIds;
            SourceType = sourceType;
            Tags = tags;
        }
    }
}
