// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class FileSystemOpenZFSConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("automaticBackupRetentionDays")]
        public Input<int>? AutomaticBackupRetentionDays { get; set; }

        [Input("copyTagsToBackups")]
        public Input<bool>? CopyTagsToBackups { get; set; }

        [Input("copyTagsToVolumes")]
        public Input<bool>? CopyTagsToVolumes { get; set; }

        [Input("dailyAutomaticBackupStartTime")]
        public Input<string>? DailyAutomaticBackupStartTime { get; set; }

        [Input("deploymentType", required: true)]
        public Input<string> DeploymentType { get; set; } = null!;

        [Input("diskIopsConfiguration")]
        public Input<Inputs.FileSystemDiskIopsConfigurationArgs>? DiskIopsConfiguration { get; set; }

        [Input("options")]
        private InputList<string>? _options;
        public InputList<string> Options
        {
            get => _options ?? (_options = new InputList<string>());
            set => _options = value;
        }

        [Input("rootVolumeConfiguration")]
        public Input<Inputs.FileSystemRootVolumeConfigurationArgs>? RootVolumeConfiguration { get; set; }

        [Input("throughputCapacity")]
        public Input<int>? ThroughputCapacity { get; set; }

        [Input("weeklyMaintenanceStartTime")]
        public Input<string>? WeeklyMaintenanceStartTime { get; set; }

        public FileSystemOpenZFSConfigurationArgs()
        {
        }
    }
}
