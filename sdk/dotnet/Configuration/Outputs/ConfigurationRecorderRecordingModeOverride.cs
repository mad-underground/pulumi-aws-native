// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    [OutputType]
    public sealed class ConfigurationRecorderRecordingModeOverride
    {
        public readonly string? Description;
        public readonly string RecordingFrequency;
        public readonly ImmutableArray<string> ResourceTypes;

        [OutputConstructor]
        private ConfigurationRecorderRecordingModeOverride(
            string? description,

            string recordingFrequency,

            ImmutableArray<string> resourceTypes)
        {
            Description = description;
            RecordingFrequency = recordingFrequency;
            ResourceTypes = resourceTypes;
        }
    }
}