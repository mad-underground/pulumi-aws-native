// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRateBasedStatement
    {
        public readonly Pulumi.AwsNative.WaFv2.RuleGroupRateBasedStatementAggregateKeyType AggregateKeyType;
        /// <summary>
        /// Specifies the aggregate keys to use in a rate-base rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupRateBasedStatementCustomKey> CustomKeys;
        public readonly Outputs.RuleGroupForwardedIpConfiguration? ForwardedIpConfig;
        public readonly int Limit;
        public readonly Outputs.RuleGroupStatement? ScopeDownStatement;

        [OutputConstructor]
        private RuleGroupRateBasedStatement(
            Pulumi.AwsNative.WaFv2.RuleGroupRateBasedStatementAggregateKeyType aggregateKeyType,

            ImmutableArray<Outputs.RuleGroupRateBasedStatementCustomKey> customKeys,

            Outputs.RuleGroupForwardedIpConfiguration? forwardedIpConfig,

            int limit,

            Outputs.RuleGroupStatement? scopeDownStatement)
        {
            AggregateKeyType = aggregateKeyType;
            CustomKeys = customKeys;
            ForwardedIpConfig = forwardedIpConfig;
            Limit = limit;
            ScopeDownStatement = scopeDownStatement;
        }
    }
}
