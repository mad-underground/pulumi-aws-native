// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPieChartFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("pieChartAggregatedFieldWells")]
        public Input<Inputs.AnalysisPieChartAggregatedFieldWellsArgs>? PieChartAggregatedFieldWells { get; set; }

        public AnalysisPieChartFieldWellsArgs()
        {
        }
        public static new AnalysisPieChartFieldWellsArgs Empty => new AnalysisPieChartFieldWellsArgs();
    }
}