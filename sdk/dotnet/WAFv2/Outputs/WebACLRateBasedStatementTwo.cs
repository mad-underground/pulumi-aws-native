// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementtwo.html
    /// </summary>
    [OutputType]
    public sealed class WebACLRateBasedStatementTwo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementtwo.html#cfn-wafv2-webacl-ratebasedstatementtwo-aggregatekeytype
        /// </summary>
        public readonly string AggregateKeyType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementtwo.html#cfn-wafv2-webacl-ratebasedstatementtwo-forwardedipconfig
        /// </summary>
        public readonly Outputs.WebACLForwardedIPConfiguration? ForwardedIPConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementtwo.html#cfn-wafv2-webacl-ratebasedstatementtwo-limit
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementtwo.html#cfn-wafv2-webacl-ratebasedstatementtwo-scopedownstatement
        /// </summary>
        public readonly Outputs.WebACLStatementThree? ScopeDownStatement;

        [OutputConstructor]
        private WebACLRateBasedStatementTwo(
            string aggregateKeyType,

            Outputs.WebACLForwardedIPConfiguration? forwardedIPConfig,

            int limit,

            Outputs.WebACLStatementThree? scopeDownStatement)
        {
            AggregateKeyType = aggregateKeyType;
            ForwardedIPConfig = forwardedIPConfig;
            Limit = limit;
            ScopeDownStatement = scopeDownStatement;
        }
    }
}
