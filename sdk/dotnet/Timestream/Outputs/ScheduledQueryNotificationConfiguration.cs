// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
    /// </summary>
    [OutputType]
    public sealed class ScheduledQueryNotificationConfiguration
    {
        public readonly Outputs.ScheduledQuerySnsConfiguration SnsConfiguration;

        [OutputConstructor]
        private ScheduledQueryNotificationConfiguration(Outputs.ScheduledQuerySnsConfiguration snsConfiguration)
        {
            SnsConfiguration = snsConfiguration;
        }
    }
}
