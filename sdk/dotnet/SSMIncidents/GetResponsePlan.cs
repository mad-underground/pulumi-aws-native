// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSMIncidents
{
    public static class GetResponsePlan
    {
        /// <summary>
        /// Resource type definition for AWS::SSMIncidents::ResponsePlan
        /// </summary>
        public static Task<GetResponsePlanResult> InvokeAsync(GetResponsePlanArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResponsePlanResult>("aws-native:ssmincidents:getResponsePlan", args ?? new GetResponsePlanArgs(), options.WithDefaults());

        /// <summary>
        /// Resource type definition for AWS::SSMIncidents::ResponsePlan
        /// </summary>
        public static Output<GetResponsePlanResult> Invoke(GetResponsePlanInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResponsePlanResult>("aws-native:ssmincidents:getResponsePlan", args ?? new GetResponsePlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResponsePlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the response plan.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetResponsePlanArgs()
        {
        }
        public static new GetResponsePlanArgs Empty => new GetResponsePlanArgs();
    }

    public sealed class GetResponsePlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the response plan.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetResponsePlanInvokeArgs()
        {
        }
        public static new GetResponsePlanInvokeArgs Empty => new GetResponsePlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetResponsePlanResult
    {
        /// <summary>
        /// The list of actions.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResponsePlanAction> Actions;
        /// <summary>
        /// The ARN of the response plan.
        /// </summary>
        public readonly string? Arn;
        public readonly Outputs.ResponsePlanChatChannel? ChatChannel;
        /// <summary>
        /// The display name of the response plan.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The list of engagements to use.
        /// </summary>
        public readonly ImmutableArray<string> Engagements;
        public readonly Outputs.ResponsePlanIncidentTemplate? IncidentTemplate;
        /// <summary>
        /// The tags to apply to the response plan.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResponsePlanTag> Tags;

        [OutputConstructor]
        private GetResponsePlanResult(
            ImmutableArray<Outputs.ResponsePlanAction> actions,

            string? arn,

            Outputs.ResponsePlanChatChannel? chatChannel,

            string? displayName,

            ImmutableArray<string> engagements,

            Outputs.ResponsePlanIncidentTemplate? incidentTemplate,

            ImmutableArray<Outputs.ResponsePlanTag> tags)
        {
            Actions = actions;
            Arn = arn;
            ChatChannel = chatChannel;
            DisplayName = displayName;
            Engagements = engagements;
            IncidentTemplate = incidentTemplate;
            Tags = tags;
        }
    }
}
