// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionContainerProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-command
        /// </summary>
        public readonly ImmutableArray<string> Command;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-environment
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionEnvironment> Environment;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-executionrolearn
        /// </summary>
        public readonly string? ExecutionRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-fargateplatformconfiguration
        /// </summary>
        public readonly Outputs.JobDefinitionFargatePlatformConfiguration? FargatePlatformConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-image
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-instancetype
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-jobrolearn
        /// </summary>
        public readonly string? JobRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-linuxparameters
        /// </summary>
        public readonly Outputs.JobDefinitionLinuxParameters? LinuxParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-logconfiguration
        /// </summary>
        public readonly Outputs.JobDefinitionLogConfiguration? LogConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-memory
        /// </summary>
        public readonly int? Memory;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-mountpoints
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionMountPoints> MountPoints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-networkconfiguration
        /// </summary>
        public readonly Outputs.JobDefinitionNetworkConfiguration? NetworkConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-privileged
        /// </summary>
        public readonly bool? Privileged;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-readonlyrootfilesystem
        /// </summary>
        public readonly bool? ReadonlyRootFilesystem;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-resourcerequirements
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionResourceRequirement> ResourceRequirements;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-secrets
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionSecret> Secrets;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-ulimits
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionUlimit> Ulimits;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-user
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-vcpus
        /// </summary>
        public readonly int? Vcpus;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-volumes
        /// </summary>
        public readonly ImmutableArray<Outputs.JobDefinitionVolumes> Volumes;

        [OutputConstructor]
        private JobDefinitionContainerProperties(
            ImmutableArray<string> Command,

            ImmutableArray<Outputs.JobDefinitionEnvironment> Environment,

            string? ExecutionRoleArn,

            Outputs.JobDefinitionFargatePlatformConfiguration? FargatePlatformConfiguration,

            string Image,

            string? InstanceType,

            string? JobRoleArn,

            Outputs.JobDefinitionLinuxParameters? LinuxParameters,

            Outputs.JobDefinitionLogConfiguration? LogConfiguration,

            int? Memory,

            ImmutableArray<Outputs.JobDefinitionMountPoints> MountPoints,

            Outputs.JobDefinitionNetworkConfiguration? NetworkConfiguration,

            bool? Privileged,

            bool? ReadonlyRootFilesystem,

            ImmutableArray<Outputs.JobDefinitionResourceRequirement> ResourceRequirements,

            ImmutableArray<Outputs.JobDefinitionSecret> Secrets,

            ImmutableArray<Outputs.JobDefinitionUlimit> Ulimits,

            string? User,

            int? Vcpus,

            ImmutableArray<Outputs.JobDefinitionVolumes> Volumes)
        {
            this.Command = Command;
            this.Environment = Environment;
            this.ExecutionRoleArn = ExecutionRoleArn;
            this.FargatePlatformConfiguration = FargatePlatformConfiguration;
            this.Image = Image;
            this.InstanceType = InstanceType;
            this.JobRoleArn = JobRoleArn;
            this.LinuxParameters = LinuxParameters;
            this.LogConfiguration = LogConfiguration;
            this.Memory = Memory;
            this.MountPoints = MountPoints;
            this.NetworkConfiguration = NetworkConfiguration;
            this.Privileged = Privileged;
            this.ReadonlyRootFilesystem = ReadonlyRootFilesystem;
            this.ResourceRequirements = ResourceRequirements;
            this.Secrets = Secrets;
            this.Ulimits = Ulimits;
            this.User = User;
            this.Vcpus = Vcpus;
            this.Volumes = Volumes;
        }
    }
}
