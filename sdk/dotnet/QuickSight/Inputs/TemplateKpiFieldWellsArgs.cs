// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateKpiFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetValues")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _targetValues;
        public InputList<Inputs.TemplateMeasureFieldArgs> TargetValues
        {
            get => _targetValues ?? (_targetValues = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _targetValues = value;
        }

        [Input("trendGroups")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _trendGroups;
        public InputList<Inputs.TemplateDimensionFieldArgs> TrendGroups
        {
            get => _trendGroups ?? (_trendGroups = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _trendGroups = value;
        }

        [Input("values")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _values;
        public InputList<Inputs.TemplateMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _values = value;
        }

        public TemplateKpiFieldWellsArgs()
        {
        }
        public static new TemplateKpiFieldWellsArgs Empty => new TemplateKpiFieldWellsArgs();
    }
}
