// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ipset.html
    /// </summary>
    [OutputType]
    public sealed class RuleGroupIPSet
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ipset.html#cfn-networkfirewall-rulegroup-ipset-definition
        /// </summary>
        public readonly Outputs.RuleGroupVariableDefinitionList? Definition;

        [OutputConstructor]
        private RuleGroupIPSet(Outputs.RuleGroupVariableDefinitionList? definition)
        {
            Definition = definition;
        }
    }
}
