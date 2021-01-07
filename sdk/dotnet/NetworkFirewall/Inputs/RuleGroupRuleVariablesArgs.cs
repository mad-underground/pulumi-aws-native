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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulevariables.html
    /// </summary>
    public sealed class RuleGroupRuleVariablesArgs : Pulumi.ResourceArgs
    {
        [Input("IPSets")]
        private InputMap<Inputs.RuleGroupIPSetArgs>? _IPSets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulevariables.html#cfn-networkfirewall-rulegroup-rulevariables-ipsets
        /// </summary>
        public InputMap<Inputs.RuleGroupIPSetArgs> IPSets
        {
            get => _IPSets ?? (_IPSets = new InputMap<Inputs.RuleGroupIPSetArgs>());
            set => _IPSets = value;
        }

        [Input("PortSets")]
        private InputMap<Inputs.RuleGroupPortSetArgs>? _PortSets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-rulevariables.html#cfn-networkfirewall-rulegroup-rulevariables-portsets
        /// </summary>
        public InputMap<Inputs.RuleGroupPortSetArgs> PortSets
        {
            get => _PortSets ?? (_PortSets = new InputMap<Inputs.RuleGroupPortSetArgs>());
            set => _PortSets = value;
        }

        public RuleGroupRuleVariablesArgs()
        {
        }
    }
}