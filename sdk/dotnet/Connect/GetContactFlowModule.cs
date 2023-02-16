// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetContactFlowModule
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::ContactFlowModule.
        /// </summary>
        public static Task<GetContactFlowModuleResult> InvokeAsync(GetContactFlowModuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContactFlowModuleResult>("aws-native:connect:getContactFlowModule", args ?? new GetContactFlowModuleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::ContactFlowModule.
        /// </summary>
        public static Output<GetContactFlowModuleResult> Invoke(GetContactFlowModuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContactFlowModuleResult>("aws-native:connect:getContactFlowModule", args ?? new GetContactFlowModuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContactFlowModuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the contact flow module (ARN).
        /// </summary>
        [Input("contactFlowModuleArn", required: true)]
        public string ContactFlowModuleArn { get; set; } = null!;

        public GetContactFlowModuleArgs()
        {
        }
        public static new GetContactFlowModuleArgs Empty => new GetContactFlowModuleArgs();
    }

    public sealed class GetContactFlowModuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the contact flow module (ARN).
        /// </summary>
        [Input("contactFlowModuleArn", required: true)]
        public Input<string> ContactFlowModuleArn { get; set; } = null!;

        public GetContactFlowModuleInvokeArgs()
        {
        }
        public static new GetContactFlowModuleInvokeArgs Empty => new GetContactFlowModuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetContactFlowModuleResult
    {
        /// <summary>
        /// The identifier of the contact flow module (ARN).
        /// </summary>
        public readonly string? ContactFlowModuleArn;
        /// <summary>
        /// The content of the contact flow module in JSON format.
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// The description of the contact flow module.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The identifier of the Amazon Connect instance (ARN).
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// The name of the contact flow module.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The state of the contact flow module.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.ContactFlowModuleState? State;
        /// <summary>
        /// The status of the contact flow module.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.ContactFlowModuleStatus? Status;
        /// <summary>
        /// One or more tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContactFlowModuleTag> Tags;

        [OutputConstructor]
        private GetContactFlowModuleResult(
            string? contactFlowModuleArn,

            string? content,

            string? description,

            string? instanceArn,

            string? name,

            Pulumi.AwsNative.Connect.ContactFlowModuleState? state,

            Pulumi.AwsNative.Connect.ContactFlowModuleStatus? status,

            ImmutableArray<Outputs.ContactFlowModuleTag> tags)
        {
            ContactFlowModuleArn = contactFlowModuleArn;
            Content = content;
            Description = description;
            InstanceArn = instanceArn;
            Name = name;
            State = state;
            Status = status;
            Tags = tags;
        }
    }
}
