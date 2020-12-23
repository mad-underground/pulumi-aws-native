// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.LicenseManager.Outputs
{

    [OutputType]
    public sealed class GrantFilterList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-filterlist.html#cfn-licensemanager-grant-filterlist-filterlist
        /// </summary>
        public readonly ImmutableArray<Outputs.GrantFilter> FilterList;

        [OutputConstructor]
        private GrantFilterList(ImmutableArray<Outputs.GrantFilter> FilterList)
        {
            this.FilterList = FilterList;
        }
    }
}
