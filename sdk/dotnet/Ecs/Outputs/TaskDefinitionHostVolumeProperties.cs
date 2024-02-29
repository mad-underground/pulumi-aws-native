// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    /// <summary>
    /// The ``HostVolumeProperties`` property specifies details on a container instance bind mount host volume.
    /// </summary>
    [OutputType]
    public sealed class TaskDefinitionHostVolumeProperties
    {
        /// <summary>
        /// When the ``host`` parameter is used, specify a ``sourcePath`` to declare the path on the host container instance that's presented to the container. If this parameter is empty, then the Docker daemon has assigned a host path for you. If the ``host`` parameter contains a ``sourcePath`` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the ``sourcePath`` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
        ///  If you're using the Fargate launch type, the ``sourcePath`` parameter is not supported.
        /// </summary>
        public readonly string? SourcePath;

        [OutputConstructor]
        private TaskDefinitionHostVolumeProperties(string? sourcePath)
        {
            SourcePath = sourcePath;
        }
    }
}
