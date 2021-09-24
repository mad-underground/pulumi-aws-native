// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class ConnectionPhysicalConnectionRequirementsArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("securityGroupIdList")]
        private InputList<string>? _securityGroupIdList;
        public InputList<string> SecurityGroupIdList
        {
            get => _securityGroupIdList ?? (_securityGroupIdList = new InputList<string>());
            set => _securityGroupIdList = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public ConnectionPhysicalConnectionRequirementsArgs()
        {
        }
    }
}