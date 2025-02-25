// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class FirewallPolicyStatefulRuleGroupReferenceArgs : global::Pulumi.ResourceArgs
    {
        [Input("override")]
        public Input<Inputs.FirewallPolicyStatefulRuleGroupOverrideArgs>? Override { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        public FirewallPolicyStatefulRuleGroupReferenceArgs()
        {
        }
        public static new FirewallPolicyStatefulRuleGroupReferenceArgs Empty => new FirewallPolicyStatefulRuleGroupReferenceArgs();
    }
}
