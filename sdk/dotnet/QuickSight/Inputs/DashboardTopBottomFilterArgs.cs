// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTopBottomFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregationSortConfigurations", required: true)]
        private InputList<Inputs.DashboardAggregationSortConfigurationArgs>? _aggregationSortConfigurations;
        public InputList<Inputs.DashboardAggregationSortConfigurationArgs> AggregationSortConfigurations
        {
            get => _aggregationSortConfigurations ?? (_aggregationSortConfigurations = new InputList<Inputs.DashboardAggregationSortConfigurationArgs>());
            set => _aggregationSortConfigurations = value;
        }

        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("filterId", required: true)]
        public Input<string> FilterId { get; set; } = null!;

        [Input("limit")]
        public Input<double>? Limit { get; set; }

        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        [Input("timeGranularity")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardTimeGranularity>? TimeGranularity { get; set; }

        public DashboardTopBottomFilterArgs()
        {
        }
        public static new DashboardTopBottomFilterArgs Empty => new DashboardTopBottomFilterArgs();
    }
}