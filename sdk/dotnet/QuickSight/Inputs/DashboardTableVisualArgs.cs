// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTableVisualArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.DashboardVisualCustomActionArgs>? _actions;
        public InputList<Inputs.DashboardVisualCustomActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.DashboardVisualCustomActionArgs>());
            set => _actions = value;
        }

        [Input("chartConfiguration")]
        public Input<Inputs.DashboardTableConfigurationArgs>? ChartConfiguration { get; set; }

        [Input("conditionalFormatting")]
        public Input<Inputs.DashboardTableConditionalFormattingArgs>? ConditionalFormatting { get; set; }

        [Input("subtitle")]
        public Input<Inputs.DashboardVisualSubtitleLabelOptionsArgs>? Subtitle { get; set; }

        [Input("title")]
        public Input<Inputs.DashboardVisualTitleLabelOptionsArgs>? Title { get; set; }

        [Input("visualId", required: true)]
        public Input<string> VisualId { get; set; } = null!;

        public DashboardTableVisualArgs()
        {
        }
        public static new DashboardTableVisualArgs Empty => new DashboardTableVisualArgs();
    }
}
