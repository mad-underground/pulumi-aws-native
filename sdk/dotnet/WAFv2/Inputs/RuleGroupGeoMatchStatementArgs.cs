// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.WAFv2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-geomatchstatement.html
    /// </summary>
    public sealed class RuleGroupGeoMatchStatementArgs : Pulumi.ResourceArgs
    {
        [Input("CountryCodes")]
        private InputList<string>? _CountryCodes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-geomatchstatement.html#cfn-wafv2-rulegroup-geomatchstatement-countrycodes
        /// </summary>
        public InputList<string> CountryCodes
        {
            get => _CountryCodes ?? (_CountryCodes = new InputList<string>());
            set => _CountryCodes = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-geomatchstatement.html#cfn-wafv2-rulegroup-geomatchstatement-forwardedipconfig
        /// </summary>
        [Input("ForwardedIPConfig")]
        public Input<Inputs.RuleGroupForwardedIPConfigurationArgs>? ForwardedIPConfig { get; set; }

        public RuleGroupGeoMatchStatementArgs()
        {
        }
    }
}
