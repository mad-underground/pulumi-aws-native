// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateColumnConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorsConfiguration")]
        public Input<Inputs.TemplateColorsConfigurationArgs>? ColorsConfiguration { get; set; }

        [Input("column", required: true)]
        public Input<Inputs.TemplateColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("formatConfiguration")]
        public Input<Inputs.TemplateFormatConfigurationArgs>? FormatConfiguration { get; set; }

        [Input("role")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateColumnRole>? Role { get; set; }

        public TemplateColumnConfigurationArgs()
        {
        }
        public static new TemplateColumnConfigurationArgs Empty => new TemplateColumnConfigurationArgs();
    }
}
