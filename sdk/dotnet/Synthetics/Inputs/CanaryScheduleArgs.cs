// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html
    /// </summary>
    public sealed class CanaryScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-durationinseconds
        /// </summary>
        [Input("durationInSeconds")]
        public Input<string>? DurationInSeconds { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-expression
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        public CanaryScheduleArgs()
        {
        }
    }
}
