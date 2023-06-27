// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardScatterPlotUnaggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _category;
        public InputList<Inputs.DashboardDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("label")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _label;
        public InputList<Inputs.DashboardDimensionFieldArgs> Label
        {
            get => _label ?? (_label = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _label = value;
        }

        [Input("size")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _size;
        public InputList<Inputs.DashboardMeasureFieldArgs> Size
        {
            get => _size ?? (_size = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _size = value;
        }

        [Input("xAxis")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _xAxis;
        public InputList<Inputs.DashboardDimensionFieldArgs> XAxis
        {
            get => _xAxis ?? (_xAxis = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _xAxis = value;
        }

        [Input("yAxis")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _yAxis;
        public InputList<Inputs.DashboardDimensionFieldArgs> YAxis
        {
            get => _yAxis ?? (_yAxis = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _yAxis = value;
        }

        public DashboardScatterPlotUnaggregatedFieldWellsArgs()
        {
        }
        public static new DashboardScatterPlotUnaggregatedFieldWellsArgs Empty => new DashboardScatterPlotUnaggregatedFieldWellsArgs();
    }
}
