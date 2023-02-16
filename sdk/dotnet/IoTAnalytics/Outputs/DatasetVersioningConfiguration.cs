// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatasetVersioningConfiguration
    {
        public readonly int? MaxVersions;
        public readonly bool? Unlimited;

        [OutputConstructor]
        private DatasetVersioningConfiguration(
            int? maxVersions,

            bool? unlimited)
        {
            MaxVersions = maxVersions;
            Unlimited = unlimited;
        }
    }
}
