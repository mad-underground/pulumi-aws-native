// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
    /// </summary>
    [OutputType]
    public sealed class TaskDefinitionContainerDefinition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-command
        /// </summary>
        public readonly ImmutableArray<string> Command;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-cpu
        /// </summary>
        public readonly int? Cpu;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dependson
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionContainerDependency> DependsOn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-disablenetworking
        /// </summary>
        public readonly bool? DisableNetworking;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dnssearchdomains
        /// </summary>
        public readonly ImmutableArray<string> DnsSearchDomains;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dnsservers
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dockerlabels
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DockerLabels;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dockersecurityoptions
        /// </summary>
        public readonly ImmutableArray<string> DockerSecurityOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-entrypoint
        /// </summary>
        public readonly ImmutableArray<string> EntryPoint;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-environment
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionKeyValuePair> Environment;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-environmentfiles
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionEnvironmentFile> EnvironmentFiles;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-essential
        /// </summary>
        public readonly bool? Essential;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-extrahosts
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionHostEntry> ExtraHosts;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-firelensconfiguration
        /// </summary>
        public readonly Outputs.TaskDefinitionFirelensConfiguration? FirelensConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-healthcheck
        /// </summary>
        public readonly Outputs.TaskDefinitionHealthCheck? HealthCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-hostname
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-image
        /// </summary>
        public readonly string? Image;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-interactive
        /// </summary>
        public readonly bool? Interactive;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-links
        /// </summary>
        public readonly ImmutableArray<string> Links;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-linuxparameters
        /// </summary>
        public readonly Outputs.TaskDefinitionLinuxParameters? LinuxParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration
        /// </summary>
        public readonly Outputs.TaskDefinitionLogConfiguration? LogConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-memory
        /// </summary>
        public readonly int? Memory;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-memoryreservation
        /// </summary>
        public readonly int? MemoryReservation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionMountPoint> MountPoints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-portmappings
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionPortMapping> PortMappings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-privileged
        /// </summary>
        public readonly bool? Privileged;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-pseudoterminal
        /// </summary>
        public readonly bool? PseudoTerminal;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-readonlyrootfilesystem
        /// </summary>
        public readonly bool? ReadonlyRootFilesystem;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-repositorycredentials
        /// </summary>
        public readonly Outputs.TaskDefinitionRepositoryCredentials? RepositoryCredentials;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-resourcerequirements
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionResourceRequirement> ResourceRequirements;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-secrets
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionSecret> Secrets;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-starttimeout
        /// </summary>
        public readonly int? StartTimeout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-stoptimeout
        /// </summary>
        public readonly int? StopTimeout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-systemcontrols
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionSystemControl> SystemControls;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-ulimits
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionUlimit> Ulimits;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-user
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionVolumeFrom> VolumesFrom;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-workingdirectory
        /// </summary>
        public readonly string? WorkingDirectory;

        [OutputConstructor]
        private TaskDefinitionContainerDefinition(
            ImmutableArray<string> command,

            int? cpu,

            ImmutableArray<Outputs.TaskDefinitionContainerDependency> dependsOn,

            bool? disableNetworking,

            ImmutableArray<string> dnsSearchDomains,

            ImmutableArray<string> dnsServers,

            ImmutableDictionary<string, string>? dockerLabels,

            ImmutableArray<string> dockerSecurityOptions,

            ImmutableArray<string> entryPoint,

            ImmutableArray<Outputs.TaskDefinitionKeyValuePair> environment,

            ImmutableArray<Outputs.TaskDefinitionEnvironmentFile> environmentFiles,

            bool? essential,

            ImmutableArray<Outputs.TaskDefinitionHostEntry> extraHosts,

            Outputs.TaskDefinitionFirelensConfiguration? firelensConfiguration,

            Outputs.TaskDefinitionHealthCheck? healthCheck,

            string? hostname,

            string? image,

            bool? interactive,

            ImmutableArray<string> links,

            Outputs.TaskDefinitionLinuxParameters? linuxParameters,

            Outputs.TaskDefinitionLogConfiguration? logConfiguration,

            int? memory,

            int? memoryReservation,

            ImmutableArray<Outputs.TaskDefinitionMountPoint> mountPoints,

            string? name,

            ImmutableArray<Outputs.TaskDefinitionPortMapping> portMappings,

            bool? privileged,

            bool? pseudoTerminal,

            bool? readonlyRootFilesystem,

            Outputs.TaskDefinitionRepositoryCredentials? repositoryCredentials,

            ImmutableArray<Outputs.TaskDefinitionResourceRequirement> resourceRequirements,

            ImmutableArray<Outputs.TaskDefinitionSecret> secrets,

            int? startTimeout,

            int? stopTimeout,

            ImmutableArray<Outputs.TaskDefinitionSystemControl> systemControls,

            ImmutableArray<Outputs.TaskDefinitionUlimit> ulimits,

            string? user,

            ImmutableArray<Outputs.TaskDefinitionVolumeFrom> volumesFrom,

            string? workingDirectory)
        {
            Command = command;
            Cpu = cpu;
            DependsOn = dependsOn;
            DisableNetworking = disableNetworking;
            DnsSearchDomains = dnsSearchDomains;
            DnsServers = dnsServers;
            DockerLabels = dockerLabels;
            DockerSecurityOptions = dockerSecurityOptions;
            EntryPoint = entryPoint;
            Environment = environment;
            EnvironmentFiles = environmentFiles;
            Essential = essential;
            ExtraHosts = extraHosts;
            FirelensConfiguration = firelensConfiguration;
            HealthCheck = healthCheck;
            Hostname = hostname;
            Image = image;
            Interactive = interactive;
            Links = links;
            LinuxParameters = linuxParameters;
            LogConfiguration = logConfiguration;
            Memory = memory;
            MemoryReservation = memoryReservation;
            MountPoints = mountPoints;
            Name = name;
            PortMappings = portMappings;
            Privileged = privileged;
            PseudoTerminal = pseudoTerminal;
            ReadonlyRootFilesystem = readonlyRootFilesystem;
            RepositoryCredentials = repositoryCredentials;
            ResourceRequirements = resourceRequirements;
            Secrets = secrets;
            StartTimeout = startTimeout;
            StopTimeout = stopTimeout;
            SystemControls = systemControls;
            Ulimits = ulimits;
            User = user;
            VolumesFrom = volumesFrom;
            WorkingDirectory = workingDirectory;
        }
    }
}
