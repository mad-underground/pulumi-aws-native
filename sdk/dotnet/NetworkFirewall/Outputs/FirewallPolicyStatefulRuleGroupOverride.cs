// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class FirewallPolicyStatefulRuleGroupOverride
    {
        public readonly Pulumi.AwsNative.NetworkFirewall.FirewallPolicyOverrideAction? Action;

        [OutputConstructor]
        private FirewallPolicyStatefulRuleGroupOverride(Pulumi.AwsNative.NetworkFirewall.FirewallPolicyOverrideAction? action)
        {
            Action = action;
        }
    }
}
