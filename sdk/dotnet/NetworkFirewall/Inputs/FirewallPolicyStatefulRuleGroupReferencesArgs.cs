// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreferences.html
    /// </summary>
    public sealed class FirewallPolicyStatefulRuleGroupReferencesArgs : Pulumi.ResourceArgs
    {
        [Input("StatefulRuleGroupReferences")]
        private InputList<Inputs.FirewallPolicyStatefulRuleGroupReferenceArgs>? _StatefulRuleGroupReferences;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreferences.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupreferences-statefulrulegroupreferences
        /// </summary>
        public InputList<Inputs.FirewallPolicyStatefulRuleGroupReferenceArgs> StatefulRuleGroupReferences
        {
            get => _StatefulRuleGroupReferences ?? (_StatefulRuleGroupReferences = new InputList<Inputs.FirewallPolicyStatefulRuleGroupReferenceArgs>());
            set => _StatefulRuleGroupReferences = value;
        }

        public FirewallPolicyStatefulRuleGroupReferencesArgs()
        {
        }
    }
}
