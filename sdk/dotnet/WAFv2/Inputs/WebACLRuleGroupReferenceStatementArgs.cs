// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html
    /// </summary>
    public sealed class WebACLRuleGroupReferenceStatementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-arn
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("excludedRules")]
        private InputList<Inputs.WebACLExcludedRuleArgs>? _excludedRules;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-rulegroupreferencestatement.html#cfn-wafv2-webacl-rulegroupreferencestatement-excludedrules
        /// </summary>
        public InputList<Inputs.WebACLExcludedRuleArgs> ExcludedRules
        {
            get => _excludedRules ?? (_excludedRules = new InputList<Inputs.WebACLExcludedRuleArgs>());
            set => _excludedRules = value;
        }

        public WebACLRuleGroupReferenceStatementArgs()
        {
        }
    }
}
