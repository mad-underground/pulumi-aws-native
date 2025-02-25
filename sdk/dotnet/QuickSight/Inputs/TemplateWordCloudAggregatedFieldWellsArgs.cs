// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateWordCloudAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBy")]
        private InputList<Inputs.TemplateDimensionFieldArgs>? _groupBy;
        public InputList<Inputs.TemplateDimensionFieldArgs> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<Inputs.TemplateDimensionFieldArgs>());
            set => _groupBy = value;
        }

        [Input("size")]
        private InputList<Inputs.TemplateMeasureFieldArgs>? _size;
        public InputList<Inputs.TemplateMeasureFieldArgs> Size
        {
            get => _size ?? (_size = new InputList<Inputs.TemplateMeasureFieldArgs>());
            set => _size = value;
        }

        public TemplateWordCloudAggregatedFieldWellsArgs()
        {
        }
        public static new TemplateWordCloudAggregatedFieldWellsArgs Empty => new TemplateWordCloudAggregatedFieldWellsArgs();
    }
}
