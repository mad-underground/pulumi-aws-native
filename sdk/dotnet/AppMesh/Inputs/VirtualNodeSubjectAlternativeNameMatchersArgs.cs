// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualNodeSubjectAlternativeNameMatchersArgs : global::Pulumi.ResourceArgs
    {
        [Input("exact")]
        private InputList<string>? _exact;
        public InputList<string> Exact
        {
            get => _exact ?? (_exact = new InputList<string>());
            set => _exact = value;
        }

        public VirtualNodeSubjectAlternativeNameMatchersArgs()
        {
        }
        public static new VirtualNodeSubjectAlternativeNameMatchersArgs Empty => new VirtualNodeSubjectAlternativeNameMatchersArgs();
    }
}
