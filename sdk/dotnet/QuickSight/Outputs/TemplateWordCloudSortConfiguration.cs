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
    public sealed class TemplateWordCloudSortConfiguration
    {
        public readonly Outputs.TemplateItemsLimitConfiguration? CategoryItemsLimit;
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> CategorySort;

        [OutputConstructor]
        private TemplateWordCloudSortConfiguration(
            Outputs.TemplateItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.TemplateFieldSortOptions> categorySort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
        }
    }
}
