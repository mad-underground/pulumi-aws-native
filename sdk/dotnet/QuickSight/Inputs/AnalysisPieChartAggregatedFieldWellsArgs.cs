// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPieChartAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("category")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _category;
        public InputList<Inputs.AnalysisDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("smallMultiples")]
        private InputList<Inputs.AnalysisDimensionFieldArgs>? _smallMultiples;
        public InputList<Inputs.AnalysisDimensionFieldArgs> SmallMultiples
        {
            get => _smallMultiples ?? (_smallMultiples = new InputList<Inputs.AnalysisDimensionFieldArgs>());
            set => _smallMultiples = value;
        }

        [Input("values")]
        private InputList<Inputs.AnalysisMeasureFieldArgs>? _values;
        public InputList<Inputs.AnalysisMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.AnalysisMeasureFieldArgs>());
            set => _values = value;
        }

        public AnalysisPieChartAggregatedFieldWellsArgs()
        {
        }
        public static new AnalysisPieChartAggregatedFieldWellsArgs Empty => new AnalysisPieChartAggregatedFieldWellsArgs();
    }
}
