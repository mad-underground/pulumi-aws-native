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
    public sealed class RuleGroupIPSetForwardedIPConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ipsetforwardedipconfiguration.html#cfn-wafv2-rulegroup-ipsetforwardedipconfiguration-fallbackbehavior
        /// </summary>
        public readonly string FallbackBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ipsetforwardedipconfiguration.html#cfn-wafv2-rulegroup-ipsetforwardedipconfiguration-headername
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ipsetforwardedipconfiguration.html#cfn-wafv2-rulegroup-ipsetforwardedipconfiguration-position
        /// </summary>
        public readonly string Position;

        [OutputConstructor]
        private RuleGroupIPSetForwardedIPConfiguration(
            string fallbackBehavior,

            string headerName,

            string position)
        {
            FallbackBehavior = fallbackBehavior;
            HeaderName = headerName;
            Position = position;
        }
    }
}
