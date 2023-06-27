// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateScatterPlotCategoricallyAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _category;
        public InputList<Inputs.TemplateDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("label")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _label;
        public InputList<Inputs.TemplateDimensionFieldArgs> Label
        {
            get => _label ?? (_label = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _label = value;
        }

        [Input("size")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _size;
        public InputList<Inputs.TemplateMeasureFieldArgs> Size
        {
            get => _size ?? (_size = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _size = value;
        }

        [Input("xAxis")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _xAxis;
        public InputList<Inputs.TemplateMeasureFieldArgs> XAxis
        {
            get => _xAxis ?? (_xAxis = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _xAxis = value;
        }

        [Input("yAxis")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _yAxis;
        public InputList<Inputs.TemplateMeasureFieldArgs> YAxis
        {
            get => _yAxis ?? (_yAxis = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _yAxis = value;
        }

        public TemplateScatterPlotCategoricallyAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateScatterPlotCategoricallyAggregatedFieldWellsArgs Empty => new TemplateScatterPlotCategoricallyAggregatedFieldWellsArgs();
    }
}
