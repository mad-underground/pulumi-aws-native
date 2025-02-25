// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cellStyle")]
        public Input<Inputs.TemplateTableCellStyleArgs>? CellStyle { get; set; }

        [Input("headerStyle")]
        public Input<Inputs.TemplateTableCellStyleArgs>? HeaderStyle { get; set; }

        [Input("orientation")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateTableOrientation>? Orientation { get; set; }

        [Input("rowAlternateColorOptions")]
        public Input<Inputs.TemplateRowAlternateColorOptionsArgs>? RowAlternateColorOptions { get; set; }

        public TemplateTableOptionsArgs()
        {
        }
        public static new TemplateTableOptionsArgs Empty => new TemplateTableOptionsArgs();
    }
}
