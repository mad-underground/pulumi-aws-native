// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateValidationStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.TemplateValidationStrategyMode> Mode { get; set; } = null!;

        public TemplateValidationStrategyArgs()
        {
        }
        public static new TemplateValidationStrategyArgs Empty => new TemplateValidationStrategyArgs();
    }
}
