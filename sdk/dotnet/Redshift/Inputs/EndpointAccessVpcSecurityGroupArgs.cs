// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Inputs
{

    /// <summary>
    /// Describes the members of a VPC security group.
    /// </summary>
    public sealed class EndpointAccessVpcSecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the VPC security group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The identifier of the VPC security group.
        /// </summary>
        [Input("vpcSecurityGroupId")]
        public Input<string>? VpcSecurityGroupId { get; set; }

        public EndpointAccessVpcSecurityGroupArgs()
        {
        }
        public static new EndpointAccessVpcSecurityGroupArgs Empty => new EndpointAccessVpcSecurityGroupArgs();
    }
}
