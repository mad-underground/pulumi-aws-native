// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTSiteWise.Outputs
{

    [OutputType]
    public sealed class DashboardAttributes
    {
        public readonly string DashboardArn;
        public readonly string DashboardId;

        [OutputConstructor]
        private DashboardAttributes(
            string DashboardArn,

            string DashboardId)
        {
            this.DashboardArn = DashboardArn;
            this.DashboardId = DashboardId;
        }
    }
}
