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
    public sealed class FirewallPolicyStatelessRuleGroupReferences
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessrulegroupreferences.html#cfn-networkfirewall-firewallpolicy-statelessrulegroupreferences-statelessrulegroupreferences
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyStatelessRuleGroupReference> StatelessRuleGroupReferences;

        [OutputConstructor]
        private FirewallPolicyStatelessRuleGroupReferences(ImmutableArray<Outputs.FirewallPolicyStatelessRuleGroupReference> statelessRuleGroupReferences)
        {
            StatelessRuleGroupReferences = statelessRuleGroupReferences;
        }
    }
}
