// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisClusterMarkerConfiguration
    {
        public readonly Outputs.AnalysisClusterMarker? ClusterMarker;

        [OutputConstructor]
        private AnalysisClusterMarkerConfiguration(Outputs.AnalysisClusterMarker? clusterMarker)
        {
            ClusterMarker = clusterMarker;
        }
    }
}