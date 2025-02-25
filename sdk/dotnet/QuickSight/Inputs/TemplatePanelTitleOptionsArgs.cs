// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePanelTitleOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("fontConfiguration")]
        public Input<Inputs.TemplateFontConfigurationArgs>? FontConfiguration { get; set; }

        [Input("horizontalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateHorizontalTextAlignment>? HorizontalTextAlignment { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? Visibility { get; set; }

        public TemplatePanelTitleOptionsArgs()
        {
        }
        public static new TemplatePanelTitleOptionsArgs Empty => new TemplatePanelTitleOptionsArgs();
    }
}
