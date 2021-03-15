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
    public sealed class GrantFilter
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-filter.html#cfn-licensemanager-grant-filter-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-grant-filter.html#cfn-licensemanager-grant-filter-values
        /// </summary>
        public readonly Outputs.GrantStringList Values;

        [OutputConstructor]
        private GrantFilter(
            string name,

            Outputs.GrantStringList values)
        {
            Name = name;
            Values = values;
        }
    }
}
