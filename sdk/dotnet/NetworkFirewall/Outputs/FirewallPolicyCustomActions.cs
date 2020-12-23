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
    public sealed class FirewallPolicyCustomActions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-customactions.html#cfn-networkfirewall-firewallpolicy-customactions-customactions
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyCustomAction> CustomActions;

        [OutputConstructor]
        private FirewallPolicyCustomActions(ImmutableArray<Outputs.FirewallPolicyCustomAction> CustomActions)
        {
            this.CustomActions = CustomActions;
        }
    }
}
