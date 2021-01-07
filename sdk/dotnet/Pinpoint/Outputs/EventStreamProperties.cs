// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Pinpoint.Outputs
{

    [OutputType]
    public sealed class EventStreamProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-applicationid
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-destinationstreamarn
        /// </summary>
        public readonly string DestinationStreamArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-rolearn
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private EventStreamProperties(
            string ApplicationId,

            string DestinationStreamArn,

            string RoleArn)
        {
            this.ApplicationId = ApplicationId;
            this.DestinationStreamArn = DestinationStreamArn;
            this.RoleArn = RoleArn;
        }
    }
}