// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldSortOptions")]
        private InputList<Inputs.DashboardPivotFieldSortOptionsArgs>? _fieldSortOptions;
        public InputList<Inputs.DashboardPivotFieldSortOptionsArgs> FieldSortOptions
        {
            get => _fieldSortOptions ?? (_fieldSortOptions = new InputList<Inputs.DashboardPivotFieldSortOptionsArgs>());
            set => _fieldSortOptions = value;
        }

        public DashboardPivotTableSortConfigurationArgs()
        {
        }
        public static new DashboardPivotTableSortConfigurationArgs Empty => new DashboardPivotTableSortConfigurationArgs();
    }
}
