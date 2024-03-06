// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionEcsTaskProperties
    {
        public readonly ImmutableArray<Outputs.JobDefinitionTaskContainerProperties> Containers;
        public readonly Outputs.JobDefinitionEphemeralStorage? EphemeralStorage;
        public readonly string? ExecutionRoleArn;
        public readonly string? IpcMode;
        public readonly Outputs.JobDefinitionNetworkConfiguration? NetworkConfiguration;
        public readonly string? PidMode;
        public readonly string? PlatformVersion;
        public readonly Outputs.JobDefinitionRuntimePlatform? RuntimePlatform;
        public readonly string? TaskRoleArn;
        public readonly ImmutableArray<Outputs.JobDefinitionVolumes> Volumes;

        [OutputConstructor]
        private JobDefinitionEcsTaskProperties(
            ImmutableArray<Outputs.JobDefinitionTaskContainerProperties> containers,

            Outputs.JobDefinitionEphemeralStorage? ephemeralStorage,

            string? executionRoleArn,

            string? ipcMode,

            Outputs.JobDefinitionNetworkConfiguration? networkConfiguration,

            string? pidMode,

            string? platformVersion,

            Outputs.JobDefinitionRuntimePlatform? runtimePlatform,

            string? taskRoleArn,

            ImmutableArray<Outputs.JobDefinitionVolumes> volumes)
        {
            Containers = containers;
            EphemeralStorage = ephemeralStorage;
            ExecutionRoleArn = executionRoleArn;
            IpcMode = ipcMode;
            NetworkConfiguration = networkConfiguration;
            PidMode = pidMode;
            PlatformVersion = platformVersion;
            RuntimePlatform = runtimePlatform;
            TaskRoleArn = taskRoleArn;
            Volumes = volumes;
        }
    }
}
