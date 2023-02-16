// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Disk associated with the Instance.
    /// </summary>
    public sealed class InstanceDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance attached to the disk.
        /// </summary>
        [Input("attachedTo")]
        public Input<string>? AttachedTo { get; set; }

        /// <summary>
        /// Attachment state of the disk.
        /// </summary>
        [Input("attachmentState")]
        public Input<string>? AttachmentState { get; set; }

        /// <summary>
        /// The names to use for your new Lightsail disk.
        /// </summary>
        [Input("diskName", required: true)]
        public Input<string> DiskName { get; set; } = null!;

        /// <summary>
        /// IOPS of disk.
        /// </summary>
        [Input("iOPS")]
        public Input<int>? IOPS { get; set; }

        /// <summary>
        /// Is the Attached disk is the system disk of the Instance.
        /// </summary>
        [Input("isSystemDisk")]
        public Input<bool>? IsSystemDisk { get; set; }

        /// <summary>
        /// Path of the disk attached to the instance.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Size of the disk attached to the Instance.
        /// </summary>
        [Input("sizeInGb")]
        public Input<string>? SizeInGb { get; set; }

        public InstanceDiskArgs()
        {
        }
        public static new InstanceDiskArgs Empty => new InstanceDiskArgs();
    }
}
