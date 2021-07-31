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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrules.html
    /// </summary>
    [OutputType]
    public sealed class RuleGroupStatefulRules
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-statefulrules.html#cfn-networkfirewall-rulegroup-statefulrules-statefulrules
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupStatefulRule> StatefulRules;

        [OutputConstructor]
        private RuleGroupStatefulRules(ImmutableArray<Outputs.RuleGroupStatefulRule> statefulRules)
        {
            StatefulRules = statefulRules;
        }
    }
}
