// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDefaultFreeFormLayoutConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("canvasSizeOptions", required: true)]
        public Input<Inputs.AnalysisFreeFormLayoutCanvasSizeOptionsArgs> CanvasSizeOptions { get; set; } = null!;

        public AnalysisDefaultFreeFormLayoutConfigurationArgs()
        {
        }
        public static new AnalysisDefaultFreeFormLayoutConfigurationArgs Empty => new AnalysisDefaultFreeFormLayoutConfigurationArgs();
    }
}
