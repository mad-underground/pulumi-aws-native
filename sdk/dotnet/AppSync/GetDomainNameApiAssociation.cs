// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    public static class GetDomainNameApiAssociation
    {
        /// <summary>
        /// Resource Type definition for AWS::AppSync::DomainNameApiAssociation
        /// </summary>
        public static Task<GetDomainNameApiAssociationResult> InvokeAsync(GetDomainNameApiAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainNameApiAssociationResult>("aws-native:appsync:getDomainNameApiAssociation", args ?? new GetDomainNameApiAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppSync::DomainNameApiAssociation
        /// </summary>
        public static Output<GetDomainNameApiAssociationResult> Invoke(GetDomainNameApiAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainNameApiAssociationResult>("aws-native:appsync:getDomainNameApiAssociation", args ?? new GetDomainNameApiAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainNameApiAssociationArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiAssociationIdentifier", required: true)]
        public string ApiAssociationIdentifier { get; set; } = null!;

        public GetDomainNameApiAssociationArgs()
        {
        }
        public static new GetDomainNameApiAssociationArgs Empty => new GetDomainNameApiAssociationArgs();
    }

    public sealed class GetDomainNameApiAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiAssociationIdentifier", required: true)]
        public Input<string> ApiAssociationIdentifier { get; set; } = null!;

        public GetDomainNameApiAssociationInvokeArgs()
        {
        }
        public static new GetDomainNameApiAssociationInvokeArgs Empty => new GetDomainNameApiAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainNameApiAssociationResult
    {
        public readonly string? ApiAssociationIdentifier;
        public readonly string? ApiId;

        [OutputConstructor]
        private GetDomainNameApiAssociationResult(
            string? apiAssociationIdentifier,

            string? apiId)
        {
            ApiAssociationIdentifier = apiAssociationIdentifier;
            ApiId = apiId;
        }
    }
}
