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
    public sealed class TemplateColumnSort
    {
        public readonly Outputs.TemplateAggregationFunction? AggregationFunction;
        public readonly Pulumi.AwsNative.QuickSight.TemplateSortDirection Direction;
        public readonly Outputs.TemplateColumnIdentifier SortBy;

        [OutputConstructor]
        private TemplateColumnSort(
            Outputs.TemplateAggregationFunction? aggregationFunction,

            Pulumi.AwsNative.QuickSight.TemplateSortDirection direction,

            Outputs.TemplateColumnIdentifier sortBy)
        {
            AggregationFunction = aggregationFunction;
            Direction = direction;
            SortBy = sortBy;
        }
    }
}
