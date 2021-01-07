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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html
    /// </summary>
    public sealed class LicenseValidityDateFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-begin
        /// </summary>
        [Input("Begin", required: true)]
        public Input<string> Begin { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-validitydateformat.html#cfn-licensemanager-license-validitydateformat-end
        /// </summary>
        [Input("End", required: true)]
        public Input<string> End { get; set; } = null!;

        public LicenseValidityDateFormatArgs()
        {
        }
    }
}
