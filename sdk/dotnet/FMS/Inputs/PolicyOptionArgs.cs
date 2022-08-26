// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FMS.Inputs
{

    /// <summary>
    /// Firewall policy option.
    /// </summary>
    public sealed class PolicyOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("networkFirewallPolicy")]
        public Input<Inputs.PolicyNetworkFirewallPolicyArgs>? NetworkFirewallPolicy { get; set; }

        [Input("thirdPartyFirewallPolicy")]
        public Input<Inputs.PolicyThirdPartyFirewallPolicyArgs>? ThirdPartyFirewallPolicy { get; set; }

        public PolicyOptionArgs()
        {
        }
        public static new PolicyOptionArgs Empty => new PolicyOptionArgs();
    }
}