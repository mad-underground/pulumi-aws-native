// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentMutationActionSetStateParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("componentName", required: true)]
        public Input<string> ComponentName { get; set; } = null!;

        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("set", required: true)]
        public Input<Inputs.ComponentPropertyArgs> Set { get; set; } = null!;

        public ComponentMutationActionSetStateParameterArgs()
        {
        }
        public static new ComponentMutationActionSetStateParameterArgs Empty => new ComponentMutationActionSetStateParameterArgs();
    }
}
