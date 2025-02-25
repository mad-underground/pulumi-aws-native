// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RolesAnywhere
{
    public static class GetProfile
    {
        /// <summary>
        /// Definition of AWS::RolesAnywhere::Profile Resource Type
        /// </summary>
        public static Task<GetProfileResult> InvokeAsync(GetProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProfileResult>("aws-native:rolesanywhere:getProfile", args ?? new GetProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::RolesAnywhere::Profile Resource Type
        /// </summary>
        public static Output<GetProfileResult> Invoke(GetProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProfileResult>("aws-native:rolesanywhere:getProfile", args ?? new GetProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("profileId", required: true)]
        public string ProfileId { get; set; } = null!;

        public GetProfileArgs()
        {
        }
        public static new GetProfileArgs Empty => new GetProfileArgs();
    }

    public sealed class GetProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("profileId", required: true)]
        public Input<string> ProfileId { get; set; } = null!;

        public GetProfileInvokeArgs()
        {
        }
        public static new GetProfileInvokeArgs Empty => new GetProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetProfileResult
    {
        public readonly double? DurationSeconds;
        public readonly bool? Enabled;
        public readonly ImmutableArray<string> ManagedPolicyArns;
        public readonly string? Name;
        public readonly string? ProfileArn;
        public readonly string? ProfileId;
        public readonly bool? RequireInstanceProperties;
        public readonly ImmutableArray<string> RoleArns;
        public readonly string? SessionPolicy;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetProfileResult(
            double? durationSeconds,

            bool? enabled,

            ImmutableArray<string> managedPolicyArns,

            string? name,

            string? profileArn,

            string? profileId,

            bool? requireInstanceProperties,

            ImmutableArray<string> roleArns,

            string? sessionPolicy,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DurationSeconds = durationSeconds;
            Enabled = enabled;
            ManagedPolicyArns = managedPolicyArns;
            Name = name;
            ProfileArn = profileArn;
            ProfileId = profileId;
            RequireInstanceProperties = requireInstanceProperties;
            RoleArns = roleArns;
            SessionPolicy = sessionPolicy;
            Tags = tags;
        }
    }
}
