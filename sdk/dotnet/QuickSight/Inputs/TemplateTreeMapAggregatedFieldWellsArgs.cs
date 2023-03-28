// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTreeMapAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("colors")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _colors;
        public InputList<Inputs.TemplateMeasureFieldArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _colors = value;
        }

        [Input("groups")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _groups;
        public InputList<Inputs.TemplateDimensionFieldArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _groups = value;
        }

        [Input("sizes")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _sizes;
        public InputList<Inputs.TemplateMeasureFieldArgs> Sizes
        {
            get => _sizes ?? (_sizes = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _sizes = value;
        }

        public TemplateTreeMapAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateTreeMapAggregatedFieldWellsArgs Empty => new TemplateTreeMapAggregatedFieldWellsArgs();
    }
}