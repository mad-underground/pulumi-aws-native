// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// An address space to be inserted into this pool. All allocations must be made from this address space.
    /// </summary>
    [OutputType]
    public sealed class IpamPoolProvisionedCidr
    {
        public readonly string Cidr;

        [OutputConstructor]
        private IpamPoolProvisionedCidr(string cidr)
        {
            Cidr = cidr;
        }
    }
}
