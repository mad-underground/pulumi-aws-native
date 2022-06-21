// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class CidrCollectionLocationArgs : Pulumi.ResourceArgs
    {
        [Input("cidrList", required: true)]
        private InputList<string>? _cidrList;

        /// <summary>
        /// A list of CIDR blocks.
        /// </summary>
        public InputList<string> CidrList
        {
            get => _cidrList ?? (_cidrList = new InputList<string>());
            set => _cidrList = value;
        }

        /// <summary>
        /// The name of the location that is associated with the CIDR collection.
        /// </summary>
        [Input("locationName", required: true)]
        public Input<string> LocationName { get; set; } = null!;

        public CidrCollectionLocationArgs()
        {
        }
    }
}
