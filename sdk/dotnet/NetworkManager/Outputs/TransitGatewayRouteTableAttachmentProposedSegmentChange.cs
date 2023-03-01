// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    /// <summary>
    /// The attachment to move from one segment to another.
    /// </summary>
    [OutputType]
    public sealed class TransitGatewayRouteTableAttachmentProposedSegmentChange
    {
        /// <summary>
        /// The rule number in the policy document that applies to this change.
        /// </summary>
        public readonly int? AttachmentPolicyRuleNumber;
        /// <summary>
        /// The name of the segment to change.
        /// </summary>
        public readonly string? SegmentName;
        /// <summary>
        /// The key-value tags that changed for the segment.
        /// </summary>
        public readonly ImmutableArray<Outputs.TransitGatewayRouteTableAttachmentTag> Tags;

        [OutputConstructor]
        private TransitGatewayRouteTableAttachmentProposedSegmentChange(
            int? attachmentPolicyRuleNumber,

            string? segmentName,

            ImmutableArray<Outputs.TransitGatewayRouteTableAttachmentTag> tags)
        {
            AttachmentPolicyRuleNumber = attachmentPolicyRuleNumber;
            SegmentName = segmentName;
            Tags = tags;
        }
    }
}
