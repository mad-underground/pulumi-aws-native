// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaPackage.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html
    /// </summary>
    public sealed class PackagingGroupPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-authorization
        /// </summary>
        [Input("Authorization")]
        public Input<Inputs.PackagingGroupAuthorizationArgs>? Authorization { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-id
        /// </summary>
        [Input("Id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public PackagingGroupPropertiesArgs()
        {
        }
    }
}