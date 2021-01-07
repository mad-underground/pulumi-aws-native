// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.RDS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
    /// </summary>
    public sealed class EventSubscriptionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-enabled
        /// </summary>
        [Input("Enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("EventCategories")]
        private InputList<string>? _EventCategories;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-eventcategories
        /// </summary>
        public InputList<string> EventCategories
        {
            get => _EventCategories ?? (_EventCategories = new InputList<string>());
            set => _EventCategories = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-snstopicarn
        /// </summary>
        [Input("SnsTopicArn", required: true)]
        public Input<string> SnsTopicArn { get; set; } = null!;

        [Input("SourceIds")]
        private InputList<string>? _SourceIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourceids
        /// </summary>
        public InputList<string> SourceIds
        {
            get => _SourceIds ?? (_SourceIds = new InputList<string>());
            set => _SourceIds = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourcetype
        /// </summary>
        [Input("SourceType")]
        public Input<string>? SourceType { get; set; }

        public EventSubscriptionPropertiesArgs()
        {
        }
    }
}