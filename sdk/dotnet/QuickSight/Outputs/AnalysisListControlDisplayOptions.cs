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
    public sealed class AnalysisListControlDisplayOptions
    {
        public readonly Outputs.AnalysisListControlSearchOptions? SearchOptions;
        public readonly Outputs.AnalysisListControlSelectAllOptions? SelectAllOptions;
        public readonly Outputs.AnalysisLabelOptions? TitleOptions;

        [OutputConstructor]
        private AnalysisListControlDisplayOptions(
            Outputs.AnalysisListControlSearchOptions? searchOptions,

            Outputs.AnalysisListControlSelectAllOptions? selectAllOptions,

            Outputs.AnalysisLabelOptions? titleOptions)
        {
            SearchOptions = searchOptions;
            SelectAllOptions = selectAllOptions;
            TitleOptions = titleOptions;
        }
    }
}