// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RolesAnywhere
{
    public static class GetCRL
    {
        /// <summary>
        /// Definition of AWS::RolesAnywhere::CRL Resource Type
        /// </summary>
        public static Task<GetCRLResult> InvokeAsync(GetCRLArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCRLResult>("aws-native:rolesanywhere:getCRL", args ?? new GetCRLArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::RolesAnywhere::CRL Resource Type
        /// </summary>
        public static Output<GetCRLResult> Invoke(GetCRLInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCRLResult>("aws-native:rolesanywhere:getCRL", args ?? new GetCRLInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCRLArgs : global::Pulumi.InvokeArgs
    {
        [Input("crlId", required: true)]
        public string CrlId { get; set; } = null!;

        public GetCRLArgs()
        {
        }
        public static new GetCRLArgs Empty => new GetCRLArgs();
    }

    public sealed class GetCRLInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("crlId", required: true)]
        public Input<string> CrlId { get; set; } = null!;

        public GetCRLInvokeArgs()
        {
        }
        public static new GetCRLInvokeArgs Empty => new GetCRLInvokeArgs();
    }


    [OutputType]
    public sealed class GetCRLResult
    {
        public readonly string? CrlData;
        public readonly string? CrlId;
        public readonly bool? Enabled;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.CRLTag> Tags;
        public readonly string? TrustAnchorArn;

        [OutputConstructor]
        private GetCRLResult(
            string? crlData,

            string? crlId,

            bool? enabled,

            string? name,

            ImmutableArray<Outputs.CRLTag> tags,

            string? trustAnchorArn)
        {
            CrlData = crlData;
            CrlId = crlId;
            Enabled = enabled;
            Name = name;
            Tags = tags;
            TrustAnchorArn = trustAnchorArn;
        }
    }
}