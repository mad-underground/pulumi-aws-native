// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies an instance's Capacity Reservation targeting option.
    /// </summary>
    public sealed class LaunchTemplateCapacityReservationSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the instance's Capacity Reservation preferences.
        /// </summary>
        [Input("capacityReservationPreference")]
        public Input<string>? CapacityReservationPreference { get; set; }

        [Input("capacityReservationTarget")]
        public Input<Inputs.LaunchTemplateCapacityReservationTargetArgs>? CapacityReservationTarget { get; set; }

        public LaunchTemplateCapacityReservationSpecificationArgs()
        {
        }
        public static new LaunchTemplateCapacityReservationSpecificationArgs Empty => new LaunchTemplateCapacityReservationSpecificationArgs();
    }
}