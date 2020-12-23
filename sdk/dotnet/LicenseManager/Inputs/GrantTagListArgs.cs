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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-taglist.html
    /// </summary>
    public sealed class GrantTagListArgs : Pulumi.ResourceArgs
    {
        [Input("TagList")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _TagList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-taglist.html#cfn-licensemanager-grant-taglist-taglist
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> TagList
        {
            get => _TagList ?? (_TagList = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _TagList = value;
        }

        public GrantTagListArgs()
        {
        }
    }
}
