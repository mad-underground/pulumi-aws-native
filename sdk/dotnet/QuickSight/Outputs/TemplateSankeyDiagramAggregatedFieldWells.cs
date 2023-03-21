// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateSankeyDiagramAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.TemplateDimensionField> Destination;
        public readonly ImmutableArray<Outputs.TemplateDimensionField> Source;
        public readonly ImmutableArray<Outputs.TemplateMeasureField> Weight;

        [OutputConstructor]
        private TemplateSankeyDiagramAggregatedFieldWells(
            ImmutableArray<Outputs.TemplateDimensionField> destination,

            ImmutableArray<Outputs.TemplateDimensionField> source,

            ImmutableArray<Outputs.TemplateMeasureField> weight)
        {
            Destination = destination;
            Source = source;
            Weight = weight;
        }
    }
}
