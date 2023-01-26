// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RUM.Outputs
{

    /// <summary>
    /// AppMonitor custom events configuration
    /// </summary>
    [OutputType]
    public sealed class AppMonitorCustomEvents
    {
        /// <summary>
        /// Indicates whether AppMonitor accepts custom events.
        /// </summary>
        public readonly Pulumi.AwsNative.RUM.AppMonitorCustomEventsStatus? Status;

        [OutputConstructor]
        private AppMonitorCustomEvents(Pulumi.AwsNative.RUM.AppMonitorCustomEventsStatus? status)
        {
            Status = status;
        }
    }
}
