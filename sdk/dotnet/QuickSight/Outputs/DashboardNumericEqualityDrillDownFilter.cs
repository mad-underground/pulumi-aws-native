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
    public sealed class DashboardNumericEqualityDrillDownFilter
    {
        public readonly Outputs.DashboardColumnIdentifier Column;
        public readonly double Value;

        [OutputConstructor]
        private DashboardNumericEqualityDrillDownFilter(
            Outputs.DashboardColumnIdentifier column,

            double value)
        {
            Column = column;
            Value = value;
        }
    }
}
