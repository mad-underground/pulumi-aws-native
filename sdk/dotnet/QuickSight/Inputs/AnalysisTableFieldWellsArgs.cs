// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("tableAggregatedFieldWells")]
        public Input<Inputs.AnalysisTableAggregatedFieldWellsArgs>? TableAggregatedFieldWells { get; set; }

        [Input("tableUnaggregatedFieldWells")]
        public Input<Inputs.AnalysisTableUnaggregatedFieldWellsArgs>? TableUnaggregatedFieldWells { get; set; }

        public AnalysisTableFieldWellsArgs()
        {
        }
        public static new AnalysisTableFieldWellsArgs Empty => new AnalysisTableFieldWellsArgs();
    }
}
