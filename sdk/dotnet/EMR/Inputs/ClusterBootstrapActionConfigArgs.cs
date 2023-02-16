// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Inputs
{

    public sealed class ClusterBootstrapActionConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("scriptBootstrapAction", required: true)]
        public Input<Inputs.ClusterScriptBootstrapActionConfigArgs> ScriptBootstrapAction { get; set; } = null!;

        public ClusterBootstrapActionConfigArgs()
        {
        }
        public static new ClusterBootstrapActionConfigArgs Empty => new ClusterBootstrapActionConfigArgs();
    }
}
