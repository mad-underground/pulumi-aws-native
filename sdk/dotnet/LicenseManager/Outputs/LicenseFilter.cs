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
    public sealed class LicenseFilter
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-filter.html#cfn-licensemanager-license-filter-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-filter.html#cfn-licensemanager-license-filter-values
        /// </summary>
        public readonly Outputs.LicenseStringList Values;

        [OutputConstructor]
        private LicenseFilter(
            string Name,

            Outputs.LicenseStringList Values)
        {
            this.Name = Name;
            this.Values = Values;
        }
    }
}
