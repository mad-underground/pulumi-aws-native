// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    public static class GetDomainName
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::DomainName
        /// </summary>
        public static Task<GetDomainNameResult> InvokeAsync(GetDomainNameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainNameResult>("aws-native:apigatewayv2:getDomainName", args ?? new GetDomainNameArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::DomainName
        /// </summary>
        public static Output<GetDomainNameResult> Invoke(GetDomainNameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainNameResult>("aws-native:apigatewayv2:getDomainName", args ?? new GetDomainNameInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainNameArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainNameArgs()
        {
        }
        public static new GetDomainNameArgs Empty => new GetDomainNameArgs();
    }

    public sealed class GetDomainNameInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainNameInvokeArgs()
        {
        }
        public static new GetDomainNameInvokeArgs Empty => new GetDomainNameInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainNameResult
    {
        public readonly ImmutableArray<Outputs.DomainNameConfiguration> DomainNameConfigurations;
        public readonly string? Id;
        public readonly Outputs.DomainNameMutualTlsAuthentication? MutualTlsAuthentication;
        public readonly string? RegionalDomainName;
        public readonly string? RegionalHostedZoneId;
        public readonly object? Tags;

        [OutputConstructor]
        private GetDomainNameResult(
            ImmutableArray<Outputs.DomainNameConfiguration> domainNameConfigurations,

            string? id,

            Outputs.DomainNameMutualTlsAuthentication? mutualTlsAuthentication,

            string? regionalDomainName,

            string? regionalHostedZoneId,

            object? tags)
        {
            DomainNameConfigurations = domainNameConfigurations;
            Id = id;
            MutualTlsAuthentication = mutualTlsAuthentication;
            RegionalDomainName = regionalDomainName;
            RegionalHostedZoneId = regionalHostedZoneId;
            Tags = tags;
        }
    }
}
