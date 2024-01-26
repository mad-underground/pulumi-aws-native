// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionContainerPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("command")]
        private InputList<string>? _command;
        public InputList<string> Command
        {
            get => _command ?? (_command = new InputList<string>());
            set => _command = value;
        }

        [Input("environment")]
        private InputList<Inputs.JobDefinitionEnvironmentArgs>? _environment;
        public InputList<Inputs.JobDefinitionEnvironmentArgs> Environment
        {
            get => _environment ?? (_environment = new InputList<Inputs.JobDefinitionEnvironmentArgs>());
            set => _environment = value;
        }

        [Input("ephemeralStorage")]
        public Input<Inputs.JobDefinitionEphemeralStorageArgs>? EphemeralStorage { get; set; }

        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        [Input("fargatePlatformConfiguration")]
        public Input<Inputs.JobDefinitionFargatePlatformConfigurationArgs>? FargatePlatformConfiguration { get; set; }

        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("jobRoleArn")]
        public Input<string>? JobRoleArn { get; set; }

        [Input("linuxParameters")]
        public Input<Inputs.JobDefinitionLinuxParametersArgs>? LinuxParameters { get; set; }

        [Input("logConfiguration")]
        public Input<Inputs.JobDefinitionLogConfigurationArgs>? LogConfiguration { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("mountPoints")]
        private InputList<Inputs.JobDefinitionMountPointsArgs>? _mountPoints;
        public InputList<Inputs.JobDefinitionMountPointsArgs> MountPoints
        {
            get => _mountPoints ?? (_mountPoints = new InputList<Inputs.JobDefinitionMountPointsArgs>());
            set => _mountPoints = value;
        }

        [Input("networkConfiguration")]
        public Input<Inputs.JobDefinitionNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        [Input("privileged")]
        public Input<bool>? Privileged { get; set; }

        [Input("readonlyRootFilesystem")]
        public Input<bool>? ReadonlyRootFilesystem { get; set; }

        [Input("resourceRequirements")]
        private InputList<Inputs.JobDefinitionResourceRequirementArgs>? _resourceRequirements;
        public InputList<Inputs.JobDefinitionResourceRequirementArgs> ResourceRequirements
        {
            get => _resourceRequirements ?? (_resourceRequirements = new InputList<Inputs.JobDefinitionResourceRequirementArgs>());
            set => _resourceRequirements = value;
        }

        [Input("runtimePlatform")]
        public Input<Inputs.JobDefinitionRuntimePlatformArgs>? RuntimePlatform { get; set; }

        [Input("secrets")]
        private InputList<Inputs.JobDefinitionSecretArgs>? _secrets;
        public InputList<Inputs.JobDefinitionSecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.JobDefinitionSecretArgs>());
            set => _secrets = value;
        }

        [Input("ulimits")]
        private InputList<Inputs.JobDefinitionUlimitArgs>? _ulimits;
        public InputList<Inputs.JobDefinitionUlimitArgs> Ulimits
        {
            get => _ulimits ?? (_ulimits = new InputList<Inputs.JobDefinitionUlimitArgs>());
            set => _ulimits = value;
        }

        [Input("user")]
        public Input<string>? User { get; set; }

        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        [Input("volumes")]
        private InputList<Inputs.JobDefinitionVolumesArgs>? _volumes;
        public InputList<Inputs.JobDefinitionVolumesArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.JobDefinitionVolumesArgs>());
            set => _volumes = value;
        }

        public JobDefinitionContainerPropertiesArgs()
        {
        }
        public static new JobDefinitionContainerPropertiesArgs Empty => new JobDefinitionContainerPropertiesArgs();
    }
}
