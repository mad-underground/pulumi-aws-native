// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisGradientColorArgs : global::Pulumi.ResourceArgs
    {
        [Input("stops")]
        private InputList<Inputs.AnalysisGradientStopArgs>? _stops;
        public InputList<Inputs.AnalysisGradientStopArgs> Stops
        {
            get => _stops ?? (_stops = new InputList<Inputs.AnalysisGradientStopArgs>());
            set => _stops = value;
        }

        public AnalysisGradientColorArgs()
        {
        }
        public static new AnalysisGradientColorArgs Empty => new AnalysisGradientColorArgs();
    }
}
