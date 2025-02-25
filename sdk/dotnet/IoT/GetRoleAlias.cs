// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetRoleAlias
    {
        /// <summary>
        /// Use the AWS::IoT::RoleAlias resource to declare an AWS IoT RoleAlias.
        /// </summary>
        public static Task<GetRoleAliasResult> InvokeAsync(GetRoleAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoleAliasResult>("aws-native:iot:getRoleAlias", args ?? new GetRoleAliasArgs(), options.WithDefaults());

        /// <summary>
        /// Use the AWS::IoT::RoleAlias resource to declare an AWS IoT RoleAlias.
        /// </summary>
        public static Output<GetRoleAliasResult> Invoke(GetRoleAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoleAliasResult>("aws-native:iot:getRoleAlias", args ?? new GetRoleAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoleAliasArgs : global::Pulumi.InvokeArgs
    {
        [Input("roleAlias", required: true)]
        public string RoleAliasValue { get; set; } = null!;

        public GetRoleAliasArgs()
        {
        }
        public static new GetRoleAliasArgs Empty => new GetRoleAliasArgs();
    }

    public sealed class GetRoleAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("roleAlias", required: true)]
        public Input<string> RoleAlias { get; set; } = null!;

        public GetRoleAliasInvokeArgs()
        {
        }
        public static new GetRoleAliasInvokeArgs Empty => new GetRoleAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoleAliasResult
    {
        public readonly int? CredentialDurationSeconds;
        public readonly string? RoleAliasArn;
        public readonly string? RoleArn;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetRoleAliasResult(
            int? credentialDurationSeconds,

            string? roleAliasArn,

            string? roleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            CredentialDurationSeconds = credentialDurationSeconds;
            RoleAliasArn = roleAliasArn;
            RoleArn = roleArn;
            Tags = tags;
        }
    }
}
