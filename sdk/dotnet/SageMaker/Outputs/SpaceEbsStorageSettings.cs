// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Properties related to the space's Amazon Elastic Block Store volume.
    /// </summary>
    [OutputType]
    public sealed class SpaceEbsStorageSettings
    {
        /// <summary>
        /// Size of the Amazon EBS volume in Gb
        /// </summary>
        public readonly int EbsVolumeSizeInGb;

        [OutputConstructor]
        private SpaceEbsStorageSettings(int ebsVolumeSizeInGb)
        {
            EbsVolumeSizeInGb = ebsVolumeSizeInGb;
        }
    }
}