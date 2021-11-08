// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Inputs
{

    public sealed class ScheduledActionTypeArgs : Pulumi.ResourceArgs
    {
        [Input("pauseCluster")]
        public Input<Inputs.ScheduledActionPauseClusterMessageArgs>? PauseCluster { get; set; }

        [Input("resizeCluster")]
        public Input<Inputs.ScheduledActionResizeClusterMessageArgs>? ResizeCluster { get; set; }

        [Input("resumeCluster")]
        public Input<Inputs.ScheduledActionResumeClusterMessageArgs>? ResumeCluster { get; set; }

        public ScheduledActionTypeArgs()
        {
        }
    }
}