// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisArcConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("arcAngle")]
        public Input<double>? ArcAngle { get; set; }

        [Input("arcThickness")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisArcThicknessOptions>? ArcThickness { get; set; }

        public AnalysisArcConfigurationArgs()
        {
        }
        public static new AnalysisArcConfigurationArgs Empty => new AnalysisArcConfigurationArgs();
    }
}
