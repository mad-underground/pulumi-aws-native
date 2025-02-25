// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    public sealed class ConfigurationRecorderRecordingModeArgs : global::Pulumi.ResourceArgs
    {
        [Input("recordingFrequency", required: true)]
        public Input<string> RecordingFrequency { get; set; } = null!;

        [Input("recordingModeOverrides")]
        private InputList<Inputs.ConfigurationRecorderRecordingModeOverrideArgs>? _recordingModeOverrides;
        public InputList<Inputs.ConfigurationRecorderRecordingModeOverrideArgs> RecordingModeOverrides
        {
            get => _recordingModeOverrides ?? (_recordingModeOverrides = new InputList<Inputs.ConfigurationRecorderRecordingModeOverrideArgs>());
            set => _recordingModeOverrides = value;
        }

        public ConfigurationRecorderRecordingModeArgs()
        {
        }
        public static new ConfigurationRecorderRecordingModeArgs Empty => new ConfigurationRecorderRecordingModeArgs();
    }
}
