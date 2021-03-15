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
    public sealed class GrantAllowedOperationList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-allowedoperationlist.html#cfn-licensemanager-grant-allowedoperationlist-allowedoperationlist
        /// </summary>
        public readonly ImmutableArray<string> AllowedOperationList;

        [OutputConstructor]
        private GrantAllowedOperationList(ImmutableArray<string> allowedOperationList)
        {
            AllowedOperationList = allowedOperationList;
        }
    }
}
