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
    public sealed class DashboardTableStyleTarget
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardStyledCellType CellType;

        [OutputConstructor]
        private DashboardTableStyleTarget(Pulumi.AwsNative.QuickSight.DashboardStyledCellType cellType)
        {
            CellType = cellType;
        }
    }
}
