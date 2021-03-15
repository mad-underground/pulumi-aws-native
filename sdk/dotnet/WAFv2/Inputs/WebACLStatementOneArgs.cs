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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html
    /// </summary>
    public sealed class WebACLStatementOneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-andstatement
        /// </summary>
        [Input("andStatement")]
        public Input<Inputs.WebACLAndStatementOneArgs>? AndStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-bytematchstatement
        /// </summary>
        [Input("byteMatchStatement")]
        public Input<Inputs.WebACLByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-geomatchstatement
        /// </summary>
        [Input("geoMatchStatement")]
        public Input<Inputs.WebACLGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-ipsetreferencestatement
        /// </summary>
        [Input("iPSetReferenceStatement")]
        public Input<Inputs.WebACLIPSetReferenceStatementArgs>? IPSetReferenceStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-managedrulegroupstatement
        /// </summary>
        [Input("managedRuleGroupStatement")]
        public Input<Inputs.WebACLManagedRuleGroupStatementArgs>? ManagedRuleGroupStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-notstatement
        /// </summary>
        [Input("notStatement")]
        public Input<Inputs.WebACLNotStatementOneArgs>? NotStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-orstatement
        /// </summary>
        [Input("orStatement")]
        public Input<Inputs.WebACLOrStatementOneArgs>? OrStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-ratebasedstatement
        /// </summary>
        [Input("rateBasedStatement")]
        public Input<Inputs.WebACLRateBasedStatementOneArgs>? RateBasedStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-regexpatternsetreferencestatement
        /// </summary>
        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebACLRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-rulegroupreferencestatement
        /// </summary>
        [Input("ruleGroupReferenceStatement")]
        public Input<Inputs.WebACLRuleGroupReferenceStatementArgs>? RuleGroupReferenceStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-sizeconstraintstatement
        /// </summary>
        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebACLSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-sqlimatchstatement
        /// </summary>
        [Input("sqliMatchStatement")]
        public Input<Inputs.WebACLSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-statementone.html#cfn-wafv2-webacl-statementone-xssmatchstatement
        /// </summary>
        [Input("xssMatchStatement")]
        public Input<Inputs.WebACLXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebACLStatementOneArgs()
        {
        }
    }
}
