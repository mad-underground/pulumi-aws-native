// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisNumericAxisOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("range")]
        public Input<Inputs.AnalysisAxisDisplayRangeArgs>? Range { get; set; }

        [Input("scale")]
        public Input<Inputs.AnalysisAxisScaleArgs>? Scale { get; set; }

        public AnalysisNumericAxisOptionsArgs()
        {
        }
        public static new AnalysisNumericAxisOptionsArgs Empty => new AnalysisNumericAxisOptionsArgs();
    }
}
