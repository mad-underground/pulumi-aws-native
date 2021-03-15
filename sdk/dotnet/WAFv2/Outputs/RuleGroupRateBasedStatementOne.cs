// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRateBasedStatementOne
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratebasedstatementone.html#cfn-wafv2-rulegroup-ratebasedstatementone-aggregatekeytype
        /// </summary>
        public readonly string AggregateKeyType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratebasedstatementone.html#cfn-wafv2-rulegroup-ratebasedstatementone-forwardedipconfig
        /// </summary>
        public readonly Outputs.RuleGroupForwardedIPConfiguration? ForwardedIPConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratebasedstatementone.html#cfn-wafv2-rulegroup-ratebasedstatementone-limit
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratebasedstatementone.html#cfn-wafv2-rulegroup-ratebasedstatementone-scopedownstatement
        /// </summary>
        public readonly Outputs.RuleGroupStatementTwo? ScopeDownStatement;

        [OutputConstructor]
        private RuleGroupRateBasedStatementOne(
            string aggregateKeyType,

            Outputs.RuleGroupForwardedIPConfiguration? forwardedIPConfig,

            int limit,

            Outputs.RuleGroupStatementTwo? scopeDownStatement)
        {
            AggregateKeyType = aggregateKeyType;
            ForwardedIPConfig = forwardedIPConfig;
            Limit = limit;
            ScopeDownStatement = scopeDownStatement;
        }
    }
}
