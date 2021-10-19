// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// The options for the transit gateway vpc attachment.
    /// </summary>
    [OutputType]
    public sealed class OptionsProperties
    {
        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        public readonly string? ApplianceModeSupport;
        /// <summary>
        /// Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        public readonly string? DnsSupport;
        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        public readonly string? Ipv6Support;

        [OutputConstructor]
        private OptionsProperties(
            string? applianceModeSupport,

            string? dnsSupport,

            string? ipv6Support)
        {
            ApplianceModeSupport = applianceModeSupport;
            DnsSupport = dnsSupport;
            Ipv6Support = ipv6Support;
        }
    }
}