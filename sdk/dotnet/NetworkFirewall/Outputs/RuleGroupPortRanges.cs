// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupPortRanges
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portranges.html#cfn-networkfirewall-rulegroup-portranges-portranges
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupPortRange> PortRanges;

        [OutputConstructor]
        private RuleGroupPortRanges(ImmutableArray<Outputs.RuleGroupPortRange> portRanges)
        {
            PortRanges = portRanges;
        }
    }
}
