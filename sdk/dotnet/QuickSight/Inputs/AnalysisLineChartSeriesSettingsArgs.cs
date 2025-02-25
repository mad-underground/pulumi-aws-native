// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisLineChartSeriesSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("lineStyleSettings")]
        public Input<Inputs.AnalysisLineChartLineStyleSettingsArgs>? LineStyleSettings { get; set; }

        [Input("markerStyleSettings")]
        public Input<Inputs.AnalysisLineChartMarkerStyleSettingsArgs>? MarkerStyleSettings { get; set; }

        public AnalysisLineChartSeriesSettingsArgs()
        {
        }
        public static new AnalysisLineChartSeriesSettingsArgs Empty => new AnalysisLineChartSeriesSettingsArgs();
    }
}
