// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisBoxPlotStyleOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("fillStyle")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisBoxPlotFillStyle>? FillStyle { get; set; }

        public AnalysisBoxPlotStyleOptionsArgs()
        {
        }
        public static new AnalysisBoxPlotStyleOptionsArgs Empty => new AnalysisBoxPlotStyleOptionsArgs();
    }
}
