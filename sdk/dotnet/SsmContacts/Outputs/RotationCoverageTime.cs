// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmContacts.Outputs
{

    /// <summary>
    /// StartTime and EndTime for the Shift
    /// </summary>
    [OutputType]
    public sealed class RotationCoverageTime
    {
        public readonly string EndTime;
        public readonly string StartTime;

        [OutputConstructor]
        private RotationCoverageTime(
            string endTime,

            string startTime)
        {
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
