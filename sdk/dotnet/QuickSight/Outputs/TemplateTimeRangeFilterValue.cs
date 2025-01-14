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
    public sealed class TemplateTimeRangeFilterValue
    {
        public readonly string? Parameter;
        public readonly Outputs.TemplateRollingDateConfiguration? RollingDate;
        public readonly string? StaticValue;

        [OutputConstructor]
        private TemplateTimeRangeFilterValue(
            string? parameter,

            Outputs.TemplateRollingDateConfiguration? rollingDate,

            string? staticValue)
        {
            Parameter = parameter;
            RollingDate = rollingDate;
            StaticValue = staticValue;
        }
    }
}
