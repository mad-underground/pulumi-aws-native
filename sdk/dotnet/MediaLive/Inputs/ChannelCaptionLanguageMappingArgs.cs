// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html
    /// </summary>
    public sealed class ChannelCaptionLanguageMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-captionchannel
        /// </summary>
        [Input("CaptionChannel")]
        public Input<int>? CaptionChannel { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-languagecode
        /// </summary>
        [Input("LanguageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionlanguagemapping.html#cfn-medialive-channel-captionlanguagemapping-languagedescription
        /// </summary>
        [Input("LanguageDescription")]
        public Input<string>? LanguageDescription { get; set; }

        public ChannelCaptionLanguageMappingArgs()
        {
        }
    }
}
