// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelMaintenanceCreateSettings
    {
        public readonly string? MaintenanceDay;
        public readonly string? MaintenanceStartTime;

        [OutputConstructor]
        private ChannelMaintenanceCreateSettings(
            string? maintenanceDay,

            string? maintenanceStartTime)
        {
            MaintenanceDay = maintenanceDay;
            MaintenanceStartTime = maintenanceStartTime;
        }
    }
}
