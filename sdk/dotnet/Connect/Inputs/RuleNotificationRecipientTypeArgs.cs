// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The type of notification recipient.
    /// </summary>
    public sealed class RuleNotificationRecipientTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("userArns")]
        private InputList<string>? _userArns;

        /// <summary>
        /// The list of recipients by user arns.
        /// </summary>
        public InputList<string> UserArns
        {
            get => _userArns ?? (_userArns = new InputList<string>());
            set => _userArns = value;
        }

        /// <summary>
        /// The collection of recipients who are identified by user tags
        /// </summary>
        [Input("userTags")]
        public Input<object>? UserTags { get; set; }

        public RuleNotificationRecipientTypeArgs()
        {
        }
        public static new RuleNotificationRecipientTypeArgs Empty => new RuleNotificationRecipientTypeArgs();
    }
}
