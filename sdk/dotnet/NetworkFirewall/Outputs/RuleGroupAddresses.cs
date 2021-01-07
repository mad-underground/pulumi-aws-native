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
    public sealed class RuleGroupAddresses
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-addresses.html#cfn-networkfirewall-rulegroup-addresses-addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupAddress> Addresses;

        [OutputConstructor]
        private RuleGroupAddresses(ImmutableArray<Outputs.RuleGroupAddress> Addresses)
        {
            this.Addresses = Addresses;
        }
    }
}