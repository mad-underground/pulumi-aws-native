// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Batch.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html
    /// </summary>
    public sealed class JobDefinitionTmpfsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-containerpath
        /// </summary>
        [Input("ContainerPath", required: true)]
        public Input<string> ContainerPath { get; set; } = null!;

        [Input("MountOptions")]
        private InputList<string>? _MountOptions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-mountoptions
        /// </summary>
        public InputList<string> MountOptions
        {
            get => _MountOptions ?? (_MountOptions = new InputList<string>());
            set => _MountOptions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-size
        /// </summary>
        [Input("Size", required: true)]
        public Input<int> Size { get; set; } = null!;

        public JobDefinitionTmpfsArgs()
        {
        }
    }
}