// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html
    /// </summary>
    public sealed class ConnectionPhysicalConnectionRequirementsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-availabilityzone
        /// </summary>
        [Input("AvailabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("SecurityGroupIdList")]
        private InputList<string>? _SecurityGroupIdList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-securitygroupidlist
        /// </summary>
        public InputList<string> SecurityGroupIdList
        {
            get => _SecurityGroupIdList ?? (_SecurityGroupIdList = new InputList<string>());
            set => _SecurityGroupIdList = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-subnetid
        /// </summary>
        [Input("SubnetId")]
        public Input<string>? SubnetId { get; set; }

        public ConnectionPhysicalConnectionRequirementsArgs()
        {
        }
    }
}