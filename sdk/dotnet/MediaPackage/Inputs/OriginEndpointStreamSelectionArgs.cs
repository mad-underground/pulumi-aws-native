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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-streamselection.html
    /// </summary>
    public sealed class OriginEndpointStreamSelectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-streamselection.html#cfn-mediapackage-originendpoint-streamselection-maxvideobitspersecond
        /// </summary>
        [Input("MaxVideoBitsPerSecond")]
        public Input<int>? MaxVideoBitsPerSecond { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-streamselection.html#cfn-mediapackage-originendpoint-streamselection-minvideobitspersecond
        /// </summary>
        [Input("MinVideoBitsPerSecond")]
        public Input<int>? MinVideoBitsPerSecond { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-streamselection.html#cfn-mediapackage-originendpoint-streamselection-streamorder
        /// </summary>
        [Input("StreamOrder")]
        public Input<string>? StreamOrder { get; set; }

        public OriginEndpointStreamSelectionArgs()
        {
        }
    }
}