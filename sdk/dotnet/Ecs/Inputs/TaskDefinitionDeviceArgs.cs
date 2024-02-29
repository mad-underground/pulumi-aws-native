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
    /// The ``Device`` property specifies an object representing a container instance host device.
    /// </summary>
    public sealed class TaskDefinitionDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path inside the container at which to expose the host device.
        /// </summary>
        [Input("containerPath")]
        public Input<string>? ContainerPath { get; set; }

        /// <summary>
        /// The path for the device on the host container instance.
        /// </summary>
        [Input("hostPath")]
        public Input<string>? HostPath { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// The explicit permissions to provide to the container for the device. By default, the container has permissions for ``read``, ``write``, and ``mknod`` for the device.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        public TaskDefinitionDeviceArgs()
        {
        }
        public static new TaskDefinitionDeviceArgs Empty => new TaskDefinitionDeviceArgs();
    }
}
