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
    public sealed class LicenseArnList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-arnlist.html#cfn-licensemanager-license-arnlist-arnlist
        /// </summary>
        public readonly ImmutableArray<string> ArnList;

        [OutputConstructor]
        private LicenseArnList(ImmutableArray<string> arnList)
        {
            ArnList = arnList;
        }
    }
}
