// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoT.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html
    /// </summary>
    public sealed class ThingAttributePayloadArgs : Pulumi.ResourceArgs
    {
        [Input("Attributes")]
        private InputMap<string>? _Attributes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html#cfn-iot-thing-attributepayload-attributes
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _Attributes ?? (_Attributes = new InputMap<string>());
            set => _Attributes = value;
        }

        public ThingAttributePayloadArgs()
        {
        }
    }
}