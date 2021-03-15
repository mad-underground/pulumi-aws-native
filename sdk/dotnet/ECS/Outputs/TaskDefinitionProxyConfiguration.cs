// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionProxyConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-proxyconfiguration.html#cfn-ecs-taskdefinition-proxyconfiguration-containername
        /// </summary>
        public readonly string ContainerName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-proxyconfiguration.html#cfn-ecs-taskdefinition-proxyconfiguration-proxyconfigurationproperties
        /// </summary>
        public readonly ImmutableArray<Outputs.TaskDefinitionKeyValuePair> ProxyConfigurationProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-proxyconfiguration.html#cfn-ecs-taskdefinition-proxyconfiguration-type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private TaskDefinitionProxyConfiguration(
            string containerName,

            ImmutableArray<Outputs.TaskDefinitionKeyValuePair> proxyConfigurationProperties,

            string? type)
        {
            ContainerName = containerName;
            ProxyConfigurationProperties = proxyConfigurationProperties;
            Type = type;
        }
    }
}
