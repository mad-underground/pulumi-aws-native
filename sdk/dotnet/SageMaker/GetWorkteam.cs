// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetWorkteam
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Workteam
        /// </summary>
        public static Task<GetWorkteamResult> InvokeAsync(GetWorkteamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkteamResult>("aws-native:sagemaker:getWorkteam", args ?? new GetWorkteamArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Workteam
        /// </summary>
        public static Output<GetWorkteamResult> Invoke(GetWorkteamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkteamResult>("aws-native:sagemaker:getWorkteam", args ?? new GetWorkteamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkteamArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWorkteamArgs()
        {
        }
        public static new GetWorkteamArgs Empty => new GetWorkteamArgs();
    }

    public sealed class GetWorkteamInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWorkteamInvokeArgs()
        {
        }
        public static new GetWorkteamInvokeArgs Empty => new GetWorkteamInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkteamResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.WorkteamMemberDefinition> MemberDefinitions;
        public readonly Outputs.WorkteamNotificationConfiguration? NotificationConfiguration;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetWorkteamResult(
            string? description,

            string? id,

            ImmutableArray<Outputs.WorkteamMemberDefinition> memberDefinitions,

            Outputs.WorkteamNotificationConfiguration? notificationConfiguration,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Description = description;
            Id = id;
            MemberDefinitions = memberDefinitions;
            NotificationConfiguration = notificationConfiguration;
            Tags = tags;
        }
    }
}
