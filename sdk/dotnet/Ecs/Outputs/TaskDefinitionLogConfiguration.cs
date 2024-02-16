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
    public sealed class TaskDefinitionLogConfiguration
    {
        public readonly string LogDriver;
        public readonly ImmutableDictionary<string, string>? Options;
        public readonly ImmutableArray<Outputs.TaskDefinitionSecret> SecretOptions;

        [OutputConstructor]
        private TaskDefinitionLogConfiguration(
            string logDriver,

            ImmutableDictionary<string, string>? options,

            ImmutableArray<Outputs.TaskDefinitionSecret> secretOptions)
        {
            LogDriver = logDriver;
            Options = options;
            SecretOptions = secretOptions;
        }
    }
}
