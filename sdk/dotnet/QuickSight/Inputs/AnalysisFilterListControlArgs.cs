// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFilterListControlArgs : global::Pulumi.ResourceArgs
    {
        [Input("cascadingControlConfiguration")]
        public Input<Inputs.AnalysisCascadingControlConfigurationArgs>? CascadingControlConfiguration { get; set; }

        [Input("displayOptions")]
        public Input<Inputs.AnalysisListControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        [Input("filterControlId", required: true)]
        public Input<string> FilterControlId { get; set; } = null!;

        [Input("selectableValues")]
        public Input<Inputs.AnalysisFilterSelectableValuesArgs>? SelectableValues { get; set; }

        [Input("sourceFilterId", required: true)]
        public Input<string> SourceFilterId { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisSheetControlListType>? Type { get; set; }

        public AnalysisFilterListControlArgs()
        {
        }
        public static new AnalysisFilterListControlArgs Empty => new AnalysisFilterListControlArgs();
    }
}
