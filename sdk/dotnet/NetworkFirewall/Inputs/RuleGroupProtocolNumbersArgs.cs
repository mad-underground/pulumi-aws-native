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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-protocolnumbers.html
    /// </summary>
    public sealed class RuleGroupProtocolNumbersArgs : Pulumi.ResourceArgs
    {
        [Input("protocolNumbers")]
        private InputList<int>? _protocolNumbers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-protocolnumbers.html#cfn-networkfirewall-rulegroup-protocolnumbers-protocolnumbers
        /// </summary>
        public InputList<int> ProtocolNumbers
        {
            get => _protocolNumbers ?? (_protocolNumbers = new InputList<int>());
            set => _protocolNumbers = value;
        }

        public RuleGroupProtocolNumbersArgs()
        {
        }
    }
}
