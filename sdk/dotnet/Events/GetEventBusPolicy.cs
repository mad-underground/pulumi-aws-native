// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events
{
    public static class GetEventBusPolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::Events::EventBusPolicy
        /// </summary>
        public static Task<GetEventBusPolicyResult> InvokeAsync(GetEventBusPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventBusPolicyResult>("aws-native:events:getEventBusPolicy", args ?? new GetEventBusPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Events::EventBusPolicy
        /// </summary>
        public static Output<GetEventBusPolicyResult> Invoke(GetEventBusPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventBusPolicyResult>("aws-native:events:getEventBusPolicy", args ?? new GetEventBusPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventBusPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEventBusPolicyArgs()
        {
        }
        public static new GetEventBusPolicyArgs Empty => new GetEventBusPolicyArgs();
    }

    public sealed class GetEventBusPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEventBusPolicyInvokeArgs()
        {
        }
        public static new GetEventBusPolicyInvokeArgs Empty => new GetEventBusPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventBusPolicyResult
    {
        public readonly string? Action;
        public readonly Outputs.EventBusPolicyCondition? Condition;
        public readonly string? Id;
        public readonly string? Principal;
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? Statement;

        [OutputConstructor]
        private GetEventBusPolicyResult(
            string? action,

            Outputs.EventBusPolicyCondition? condition,

            string? id,

            string? principal,

            object? statement)
        {
            Action = action;
            Condition = condition;
            Id = id;
            Principal = principal;
            Statement = statement;
        }
    }
}
