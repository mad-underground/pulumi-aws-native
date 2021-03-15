// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager.Outputs
{

    [OutputType]
    public sealed class LicenseTagList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-taglist.html#cfn-licensemanager-license-taglist-taglist
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> TagList;

        [OutputConstructor]
        private LicenseTagList(ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tagList)
        {
            TagList = tagList;
        }
    }
}
