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
    public sealed class LicenseConsumptionConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-borrowconfiguration
        /// </summary>
        public readonly Outputs.LicenseBorrowConfiguration? BorrowConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-provisionalconfiguration
        /// </summary>
        public readonly Outputs.LicenseProvisionalConfiguration? ProvisionalConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-consumptionconfiguration.html#cfn-licensemanager-license-consumptionconfiguration-renewtype
        /// </summary>
        public readonly string? RenewType;

        [OutputConstructor]
        private LicenseConsumptionConfiguration(
            Outputs.LicenseBorrowConfiguration? borrowConfiguration,

            Outputs.LicenseProvisionalConfiguration? provisionalConfiguration,

            string? renewType)
        {
            BorrowConfiguration = borrowConfiguration;
            ProvisionalConfiguration = provisionalConfiguration;
            RenewType = renewType;
        }
    }
}
