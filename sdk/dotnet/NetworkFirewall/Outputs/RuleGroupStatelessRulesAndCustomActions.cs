// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupStatelessRulesAndCustomActions
    {
        public readonly ImmutableArray<Outputs.RuleGroupCustomAction> CustomActions;
        public readonly ImmutableArray<Outputs.RuleGroupStatelessRule> StatelessRules;

        [OutputConstructor]
        private RuleGroupStatelessRulesAndCustomActions(
            ImmutableArray<Outputs.RuleGroupCustomAction> customActions,

            ImmutableArray<Outputs.RuleGroupStatelessRule> statelessRules)
        {
            CustomActions = customActions;
            StatelessRules = statelessRules;
        }
    }
}
