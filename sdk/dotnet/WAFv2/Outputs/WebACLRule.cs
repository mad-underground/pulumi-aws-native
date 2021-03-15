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
    public sealed class WebACLRule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-action
        /// </summary>
        public readonly Outputs.WebACLRuleAction? Action;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-overrideaction
        /// </summary>
        public readonly Outputs.WebACLOverrideAction? OverrideAction;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-priority
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-statement
        /// </summary>
        public readonly Outputs.WebACLStatementOne Statement;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rule.html#cfn-wafv2-webacl-rule-visibilityconfig
        /// </summary>
        public readonly Outputs.WebACLVisibilityConfig VisibilityConfig;

        [OutputConstructor]
        private WebACLRule(
            Outputs.WebACLRuleAction? action,

            string name,

            Outputs.WebACLOverrideAction? overrideAction,

            int priority,

            Outputs.WebACLStatementOne statement,

            Outputs.WebACLVisibilityConfig visibilityConfig)
        {
            Action = action;
            Name = name;
            OverrideAction = overrideAction;
            Priority = priority;
            Statement = statement;
            VisibilityConfig = visibilityConfig;
        }
    }
}
