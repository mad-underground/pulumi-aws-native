// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceIpv6PrefixSpecification
    {
        public readonly string Ipv6Prefix;

        [OutputConstructor]
        private NetworkInterfaceIpv6PrefixSpecification(string ipv6Prefix)
        {
            Ipv6Prefix = ipv6Prefix;
        }
    }
}
