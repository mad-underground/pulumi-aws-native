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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html
    /// </summary>
    [OutputType]
    public sealed class WebACLRuleGroupReferenceStatement
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-arn
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-excludedrules
        /// </summary>
        public readonly ImmutableArray<Outputs.WebACLExcludedRule> ExcludedRules;

        [OutputConstructor]
        private WebACLRuleGroupReferenceStatement(
            string arn,

            ImmutableArray<Outputs.WebACLExcludedRule> excludedRules)
        {
            Arn = arn;
            ExcludedRules = excludedRules;
        }
    }
}
