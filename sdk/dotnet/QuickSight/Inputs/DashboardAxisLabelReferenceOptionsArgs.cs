// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardAxisLabelReferenceOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        public DashboardAxisLabelReferenceOptionsArgs()
        {
        }
        public static new DashboardAxisLabelReferenceOptionsArgs Empty => new DashboardAxisLabelReferenceOptionsArgs();
    }
}
