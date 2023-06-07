// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Inputs
{

    /// <summary>
    /// specifies whether the scaling activities for a scalable target are in a suspended state
    /// </summary>
    public sealed class ScalableTargetSuspendedStateArgs : global::Pulumi.ResourceArgs
    {
        [Input("dynamicScalingInSuspended")]
        public Input<bool>? DynamicScalingInSuspended { get; set; }

        [Input("dynamicScalingOutSuspended")]
        public Input<bool>? DynamicScalingOutSuspended { get; set; }

        [Input("scheduledScalingSuspended")]
        public Input<bool>? ScheduledScalingSuspended { get; set; }

        public ScalableTargetSuspendedStateArgs()
        {
        }
        public static new ScalableTargetSuspendedStateArgs Empty => new ScalableTargetSuspendedStateArgs();
    }
}
