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
    public sealed class DashboardSectionStyle
    {
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? Height;
        public readonly Outputs.DashboardSpacing? Padding;

        [OutputConstructor]
        private DashboardSectionStyle(
            string? height,

            Outputs.DashboardSpacing? padding)
        {
            Height = height;
            Padding = padding;
        }
    }
}
