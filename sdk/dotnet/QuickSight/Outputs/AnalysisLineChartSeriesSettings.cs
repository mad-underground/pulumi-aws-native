// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisLineChartSeriesSettings
    {
        public readonly Outputs.AnalysisLineChartLineStyleSettings? LineStyleSettings;
        public readonly Outputs.AnalysisLineChartMarkerStyleSettings? MarkerStyleSettings;

        [OutputConstructor]
        private AnalysisLineChartSeriesSettings(
            Outputs.AnalysisLineChartLineStyleSettings? lineStyleSettings,

            Outputs.AnalysisLineChartMarkerStyleSettings? markerStyleSettings)
        {
            LineStyleSettings = lineStyleSettings;
            MarkerStyleSettings = markerStyleSettings;
        }
    }
}
