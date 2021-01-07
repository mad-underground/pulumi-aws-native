// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.WAFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupGeoMatchStatement
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-geomatchstatement.html#cfn-wafv2-rulegroup-geomatchstatement-countrycodes
        /// </summary>
        public readonly ImmutableArray<string> CountryCodes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-geomatchstatement.html#cfn-wafv2-rulegroup-geomatchstatement-forwardedipconfig
        /// </summary>
        public readonly Outputs.RuleGroupForwardedIPConfiguration? ForwardedIPConfig;

        [OutputConstructor]
        private RuleGroupGeoMatchStatement(
            ImmutableArray<string> CountryCodes,

            Outputs.RuleGroupForwardedIPConfiguration? ForwardedIPConfig)
        {
            this.CountryCodes = CountryCodes;
            this.ForwardedIPConfig = ForwardedIPConfig;
        }
    }
}