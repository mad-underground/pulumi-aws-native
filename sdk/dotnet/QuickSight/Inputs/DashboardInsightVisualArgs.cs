// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardInsightVisualArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.DashboardVisualCustomActionArgs>? _actions;
        public InputList<Inputs.DashboardVisualCustomActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.DashboardVisualCustomActionArgs>());
            set => _actions = value;
        }

        [Input("dataSetIdentifier", required: true)]
        public Input<string> DataSetIdentifier { get; set; } = null!;

        [Input("insightConfiguration")]
        public Input<Inputs.DashboardInsightConfigurationArgs>? InsightConfiguration { get; set; }

        [Input("subtitle")]
        public Input<Inputs.DashboardVisualSubtitleLabelOptionsArgs>? Subtitle { get; set; }

        [Input("title")]
        public Input<Inputs.DashboardVisualTitleLabelOptionsArgs>? Title { get; set; }

        [Input("visualId", required: true)]
        public Input<string> VisualId { get; set; } = null!;

        public DashboardInsightVisualArgs()
        {
        }
        public static new DashboardInsightVisualArgs Empty => new DashboardInsightVisualArgs();
    }
}