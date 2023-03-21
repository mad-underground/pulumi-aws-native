// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisRadarChartAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _category;
        public InputList<Inputs.AnalysisDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("color")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _color;
        public InputList<Inputs.AnalysisDimensionFieldArgs> Color
        {
            get => _color ?? (_color = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _color = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisRadarChartAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisRadarChartAggregatedFieldWellsArgs Empty => new AnalysisRadarChartAggregatedFieldWellsArgs();
    }
}
