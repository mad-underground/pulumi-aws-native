// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableCellStyleArgs : global::Pulumi.ResourceArgs
    {
        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        [Input("border")]
        public Input<Inputs.AnalysisGlobalTableBorderOptionsArgs>? Border { get; set; }

        [Input("fontConfiguration")]
        public Input<Inputs.AnalysisFontConfigurationArgs>? FontConfiguration { get; set; }

        [Input("height")]
        public Input<double>? Height { get; set; }

        [Input("horizontalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisHorizontalTextAlignment>? HorizontalTextAlignment { get; set; }

        [Input("textWrap")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisTextWrap>? TextWrap { get; set; }

        [Input("verticalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVerticalTextAlignment>? VerticalTextAlignment { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisTableCellStyleArgs()
        {
        }
        public static new AnalysisTableCellStyleArgs Empty => new AnalysisTableCellStyleArgs();
    }
}