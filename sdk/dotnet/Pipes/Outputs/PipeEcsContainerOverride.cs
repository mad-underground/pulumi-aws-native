// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeEcsContainerOverride
    {
        public readonly ImmutableArray<string> Command;
        public readonly int? Cpu;
        public readonly ImmutableArray<Outputs.PipeEcsEnvironmentVariable> Environment;
        public readonly ImmutableArray<Outputs.PipeEcsEnvironmentFile> EnvironmentFiles;
        public readonly int? Memory;
        public readonly int? MemoryReservation;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.PipeEcsResourceRequirement> ResourceRequirements;

        [OutputConstructor]
        private PipeEcsContainerOverride(
            ImmutableArray<string> command,

            int? cpu,

            ImmutableArray<Outputs.PipeEcsEnvironmentVariable> environment,

            ImmutableArray<Outputs.PipeEcsEnvironmentFile> environmentFiles,

            int? memory,

            int? memoryReservation,

            string? name,

            ImmutableArray<Outputs.PipeEcsResourceRequirement> resourceRequirements)
        {
            Command = command;
            Cpu = cpu;
            Environment = environment;
            EnvironmentFiles = environmentFiles;
            Memory = memory;
            MemoryReservation = memoryReservation;
            Name = name;
            ResourceRequirements = resourceRequirements;
        }
    }
}
