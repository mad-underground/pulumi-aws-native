// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFontConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        [Input("fontDecoration")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateFontDecoration>? FontDecoration { get; set; }

        [Input("fontSize")]
        public Input<Inputs.TemplateFontSizeArgs>? FontSize { get; set; }

        [Input("fontStyle")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateFontStyle>? FontStyle { get; set; }

        [Input("fontWeight")]
        public Input<Inputs.TemplateFontWeightArgs>? FontWeight { get; set; }

        public TemplateFontConfigurationArgs()
        {
        }
        public static new TemplateFontConfigurationArgs Empty => new TemplateFontConfigurationArgs();
    }
}