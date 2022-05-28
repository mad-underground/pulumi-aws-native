// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
    /// </summary>
    public sealed class DistributionConfigurationFastLaunchSnapshotConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
        /// </summary>
        [Input("targetResourceCount")]
        public Input<int>? TargetResourceCount { get; set; }

        public DistributionConfigurationFastLaunchSnapshotConfigurationArgs()
        {
        }
    }
}
