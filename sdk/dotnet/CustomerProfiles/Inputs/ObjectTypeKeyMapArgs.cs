// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    public sealed class ObjectTypeKeyMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("objectTypeKeyList")]
        private InputList<Inputs.ObjectTypeKeyArgs>? _objectTypeKeyList;
        public InputList<Inputs.ObjectTypeKeyArgs> ObjectTypeKeyList
        {
            get => _objectTypeKeyList ?? (_objectTypeKeyList = new InputList<Inputs.ObjectTypeKeyArgs>());
            set => _objectTypeKeyList = value;
        }

        public ObjectTypeKeyMapArgs()
        {
        }
        public static new ObjectTypeKeyMapArgs Empty => new ObjectTypeKeyMapArgs();
    }
}
