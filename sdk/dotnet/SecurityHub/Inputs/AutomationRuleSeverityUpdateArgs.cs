// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    public sealed class AutomationRuleSeverityUpdateArgs : global::Pulumi.ResourceArgs
    {
        [Input("label")]
        public Input<Pulumi.AwsNative.SecurityHub.AutomationRuleSeverityUpdateLabel>? Label { get; set; }

        [Input("normalized")]
        public Input<int>? Normalized { get; set; }

        [Input("product")]
        public Input<double>? Product { get; set; }

        public AutomationRuleSeverityUpdateArgs()
        {
        }
        public static new AutomationRuleSeverityUpdateArgs Empty => new AutomationRuleSeverityUpdateArgs();
    }
}
