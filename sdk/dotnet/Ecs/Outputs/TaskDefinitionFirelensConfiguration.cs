// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionFirelensConfiguration
    {
        public readonly object? Options;
        public readonly string? Type;

        [OutputConstructor]
        private TaskDefinitionFirelensConfiguration(
            object? options,

            string? type)
        {
            Options = options;
            Type = type;
        }
    }
}