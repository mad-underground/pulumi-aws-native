// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTreeMapFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("treeMapAggregatedFieldWells")]
        public Input<Inputs.DashboardTreeMapAggregatedFieldWellsArgs>? TreeMapAggregatedFieldWells { get; set; }

        public DashboardTreeMapFieldWellsArgs()
        {
        }
        public static new DashboardTreeMapFieldWellsArgs Empty => new DashboardTreeMapFieldWellsArgs();
    }
}
