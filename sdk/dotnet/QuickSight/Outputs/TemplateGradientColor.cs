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
    public sealed class TemplateGradientColor
    {
        public readonly ImmutableArray<Outputs.TemplateGradientStop> Stops;

        [OutputConstructor]
        private TemplateGradientColor(ImmutableArray<Outputs.TemplateGradientStop> stops)
        {
            Stops = stops;
        }
    }
}
