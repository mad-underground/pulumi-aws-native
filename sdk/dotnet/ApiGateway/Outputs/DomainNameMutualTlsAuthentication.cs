// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApiGateway.Outputs
{

    [OutputType]
    public sealed class DomainNameMutualTlsAuthentication
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreuri
        /// </summary>
        public readonly string? TruststoreUri;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreversion
        /// </summary>
        public readonly string? TruststoreVersion;

        [OutputConstructor]
        private DomainNameMutualTlsAuthentication(
            string? TruststoreUri,

            string? TruststoreVersion)
        {
            this.TruststoreUri = TruststoreUri;
            this.TruststoreVersion = TruststoreVersion;
        }
    }
}