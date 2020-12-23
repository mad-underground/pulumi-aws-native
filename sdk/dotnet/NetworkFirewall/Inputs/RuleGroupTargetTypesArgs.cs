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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-targettypes.html
    /// </summary>
    public sealed class RuleGroupTargetTypesArgs : Pulumi.ResourceArgs
    {
        [Input("TargetTypes")]
        private InputList<string>? _TargetTypes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-targettypes.html#cfn-networkfirewall-rulegroup-targettypes-targettypes
        /// </summary>
        public InputList<string> TargetTypes
        {
            get => _TargetTypes ?? (_TargetTypes = new InputList<string>());
            set => _TargetTypes = value;
        }

        public RuleGroupTargetTypesArgs()
        {
        }
    }
}
