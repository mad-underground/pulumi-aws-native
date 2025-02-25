// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainEndpointOptions
    {
        public readonly string? CustomEndpoint;
        public readonly string? CustomEndpointCertificateArn;
        public readonly bool? CustomEndpointEnabled;
        public readonly bool? EnforceHttps;
        public readonly string? TlsSecurityPolicy;

        [OutputConstructor]
        private DomainEndpointOptions(
            string? customEndpoint,

            string? customEndpointCertificateArn,

            bool? customEndpointEnabled,

            bool? enforceHttps,

            string? tlsSecurityPolicy)
        {
            CustomEndpoint = customEndpoint;
            CustomEndpointCertificateArn = customEndpointCertificateArn;
            CustomEndpointEnabled = customEndpointEnabled;
            EnforceHttps = enforceHttps;
            TlsSecurityPolicy = tlsSecurityPolicy;
        }
    }
}
