// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class FirewallPolicyDimension
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-dimension.html#cfn-networkfirewall-firewallpolicy-dimension-value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private FirewallPolicyDimension(string Value)
        {
            this.Value = Value;
        }
    }
}