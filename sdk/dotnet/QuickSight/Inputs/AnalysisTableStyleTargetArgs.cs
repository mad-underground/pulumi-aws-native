// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableStyleTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cellType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisStyledCellType> CellType { get; set; } = null!;

        public AnalysisTableStyleTargetArgs()
        {
        }
        public static new AnalysisTableStyleTargetArgs Empty => new AnalysisTableStyleTargetArgs();
    }
}
