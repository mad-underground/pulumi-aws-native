// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions
{
    public static class GetIdentitySource
    {
        /// <summary>
        /// Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
        /// </summary>
        public static Task<GetIdentitySourceResult> InvokeAsync(GetIdentitySourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentitySourceResult>("aws-native:verifiedpermissions:getIdentitySource", args ?? new GetIdentitySourceArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
        /// </summary>
        public static Output<GetIdentitySourceResult> Invoke(GetIdentitySourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentitySourceResult>("aws-native:verifiedpermissions:getIdentitySource", args ?? new GetIdentitySourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentitySourceArgs : global::Pulumi.InvokeArgs
    {
        [Input("identitySourceId", required: true)]
        public string IdentitySourceId { get; set; } = null!;

        [Input("policyStoreId", required: true)]
        public string PolicyStoreId { get; set; } = null!;

        public GetIdentitySourceArgs()
        {
        }
        public static new GetIdentitySourceArgs Empty => new GetIdentitySourceArgs();
    }

    public sealed class GetIdentitySourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("identitySourceId", required: true)]
        public Input<string> IdentitySourceId { get; set; } = null!;

        [Input("policyStoreId", required: true)]
        public Input<string> PolicyStoreId { get; set; } = null!;

        public GetIdentitySourceInvokeArgs()
        {
        }
        public static new GetIdentitySourceInvokeArgs Empty => new GetIdentitySourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentitySourceResult
    {
        public readonly Outputs.IdentitySourceConfiguration? Configuration;
        public readonly Outputs.IdentitySourceDetails? Details;
        public readonly string? IdentitySourceId;
        public readonly string? PrincipalEntityType;

        [OutputConstructor]
        private GetIdentitySourceResult(
            Outputs.IdentitySourceConfiguration? configuration,

            Outputs.IdentitySourceDetails? details,

            string? identitySourceId,

            string? principalEntityType)
        {
            Configuration = configuration;
            Details = details;
            IdentitySourceId = identitySourceId;
            PrincipalEntityType = principalEntityType;
        }
    }
}
