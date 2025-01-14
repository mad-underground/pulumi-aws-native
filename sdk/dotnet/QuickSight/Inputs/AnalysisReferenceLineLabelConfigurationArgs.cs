// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisReferenceLineLabelConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("customLabelConfiguration")]
        public Input<Inputs.AnalysisReferenceLineCustomLabelConfigurationArgs>? CustomLabelConfiguration { get; set; }

        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        [Input("fontConfiguration")]
        public Input<Inputs.AnalysisFontConfigurationArgs>? FontConfiguration { get; set; }

        [Input("horizontalPosition")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisReferenceLineLabelHorizontalPosition>? HorizontalPosition { get; set; }

        [Input("valueLabelConfiguration")]
        public Input<Inputs.AnalysisReferenceLineValueLabelConfigurationArgs>? ValueLabelConfiguration { get; set; }

        [Input("verticalPosition")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisReferenceLineLabelVerticalPosition>? VerticalPosition { get; set; }

        public AnalysisReferenceLineLabelConfigurationArgs()
        {
        }
        public static new AnalysisReferenceLineLabelConfigurationArgs Empty => new AnalysisReferenceLineLabelConfigurationArgs();
    }
}
