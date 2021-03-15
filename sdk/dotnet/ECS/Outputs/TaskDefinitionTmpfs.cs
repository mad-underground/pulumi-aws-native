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
    public sealed class TaskDefinitionTmpfs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-containerpath
        /// </summary>
        public readonly string? ContainerPath;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-mountoptions
        /// </summary>
        public readonly ImmutableArray<string> MountOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-size
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private TaskDefinitionTmpfs(
            string? containerPath,

            ImmutableArray<string> mountOptions,

            int size)
        {
            ContainerPath = containerPath;
            MountOptions = mountOptions;
            Size = size;
        }
    }
}
