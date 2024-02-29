// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The details for a volume mount point that's used in a container definition.
    /// </summary>
    public sealed class TaskDefinitionMountPointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path on the container to mount the host volume at.
        /// </summary>
        [Input("containerPath")]
        public Input<string>? ContainerPath { get; set; }

        /// <summary>
        /// If this value is ``true``, the container has read-only access to the volume. If this value is ``false``, then the container can write to the volume. The default value is ``false``.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The name of the volume to mount. Must be a volume name referenced in the ``name`` parameter of task definition ``volume``.
        /// </summary>
        [Input("sourceVolume")]
        public Input<string>? SourceVolume { get; set; }

        public TaskDefinitionMountPointArgs()
        {
        }
        public static new TaskDefinitionMountPointArgs Empty => new TaskDefinitionMountPointArgs();
    }
}
