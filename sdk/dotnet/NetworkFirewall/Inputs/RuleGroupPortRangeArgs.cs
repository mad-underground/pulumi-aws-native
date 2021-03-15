// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html
    /// </summary>
    public sealed class RuleGroupPortRangeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html#cfn-networkfirewall-rulegroup-portrange-fromport
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html#cfn-networkfirewall-rulegroup-portrange-toport
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public RuleGroupPortRangeArgs()
        {
        }
    }
}
