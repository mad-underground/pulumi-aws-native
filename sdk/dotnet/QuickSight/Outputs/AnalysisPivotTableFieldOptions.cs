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
    public sealed class AnalysisPivotTableFieldOptions
    {
        public readonly ImmutableArray<Outputs.AnalysisPivotTableDataPathOption> DataPathOptions;
        public readonly ImmutableArray<Outputs.AnalysisPivotTableFieldOption> SelectedFieldOptions;

        [OutputConstructor]
        private AnalysisPivotTableFieldOptions(
            ImmutableArray<Outputs.AnalysisPivotTableDataPathOption> dataPathOptions,

            ImmutableArray<Outputs.AnalysisPivotTableFieldOption> selectedFieldOptions)
        {
            DataPathOptions = dataPathOptions;
            SelectedFieldOptions = selectedFieldOptions;
        }
    }
}