// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisArcAxisConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("range")]
        public Input<Inputs.AnalysisArcAxisDisplayRangeArgs>? Range { get; set; }

        [Input("reserveRange")]
        public Input<double>? ReserveRange { get; set; }

        public AnalysisArcAxisConfigurationArgs()
        {
        }
        public static new AnalysisArcAxisConfigurationArgs Empty => new AnalysisArcAxisConfigurationArgs();
    }
}
