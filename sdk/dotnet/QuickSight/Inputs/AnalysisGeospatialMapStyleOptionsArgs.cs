// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisGeospatialMapStyleOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("baseMapStyle")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisBaseMapStyleType>? BaseMapStyle { get; set; }

        public AnalysisGeospatialMapStyleOptionsArgs()
        {
        }
        public static new AnalysisGeospatialMapStyleOptionsArgs Empty => new AnalysisGeospatialMapStyleOptionsArgs();
    }
}
