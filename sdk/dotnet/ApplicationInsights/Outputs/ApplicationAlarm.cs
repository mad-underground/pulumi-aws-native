// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-alarm.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationAlarm
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-alarm.html#cfn-applicationinsights-application-alarm-alarmname
        /// </summary>
        public readonly string AlarmName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-alarm.html#cfn-applicationinsights-application-alarm-severity
        /// </summary>
        public readonly string? Severity;

        [OutputConstructor]
        private ApplicationAlarm(
            string alarmName,

            string? severity)
        {
            AlarmName = alarmName;
            Severity = severity;
        }
    }
}
