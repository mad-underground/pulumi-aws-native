// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.LicenseManager.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-rulelist.html
    /// </summary>
    public sealed class LicenseRuleListArgs : Pulumi.ResourceArgs
    {
        [Input("RuleList")]
        private InputList<Inputs.LicenseRuleArgs>? _RuleList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-rulelist.html#cfn-licensemanager-license-rulelist-rulelist
        /// </summary>
        public InputList<Inputs.LicenseRuleArgs> RuleList
        {
            get => _RuleList ?? (_RuleList = new InputList<Inputs.LicenseRuleArgs>());
            set => _RuleList = value;
        }

        public LicenseRuleListArgs()
        {
        }
    }
}
