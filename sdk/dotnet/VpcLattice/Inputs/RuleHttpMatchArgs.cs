// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Inputs
{

    public sealed class RuleHttpMatchArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerMatches")]
        private InputList<Inputs.RuleHeaderMatchArgs>? _headerMatches;
        public InputList<Inputs.RuleHeaderMatchArgs> HeaderMatches
        {
            get => _headerMatches ?? (_headerMatches = new InputList<Inputs.RuleHeaderMatchArgs>());
            set => _headerMatches = value;
        }

        [Input("method")]
        public Input<Pulumi.AwsNative.VpcLattice.RuleHttpMatchMethod>? Method { get; set; }

        [Input("pathMatch")]
        public Input<Inputs.RulePathMatchArgs>? PathMatch { get; set; }

        public RuleHttpMatchArgs()
        {
        }
        public static new RuleHttpMatchArgs Empty => new RuleHttpMatchArgs();
    }
}