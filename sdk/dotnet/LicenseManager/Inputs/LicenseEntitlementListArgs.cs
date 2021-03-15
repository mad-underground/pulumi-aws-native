// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlementlist.html
    /// </summary>
    public sealed class LicenseEntitlementListArgs : Pulumi.ResourceArgs
    {
        [Input("entitlementList")]
        private InputList<Inputs.LicenseEntitlementArgs>? _entitlementList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlementlist.html#cfn-licensemanager-license-entitlementlist-entitlementlist
        /// </summary>
        public InputList<Inputs.LicenseEntitlementArgs> EntitlementList
        {
            get => _entitlementList ?? (_entitlementList = new InputList<Inputs.LicenseEntitlementArgs>());
            set => _entitlementList = value;
        }

        public LicenseEntitlementListArgs()
        {
        }
    }
}
