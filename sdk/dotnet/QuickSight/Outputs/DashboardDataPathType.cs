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
    public sealed class DashboardDataPathType
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardPivotTableDataPathType? PivotTableDataPathType;

        [OutputConstructor]
        private DashboardDataPathType(Pulumi.AwsNative.QuickSight.DashboardPivotTableDataPathType? pivotTableDataPathType)
        {
            PivotTableDataPathType = pivotTableDataPathType;
        }
    }
}
