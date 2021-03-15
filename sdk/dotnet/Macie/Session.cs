// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html
    /// </summary>
    [AwsNativeResourceType("aws-native:Macie:Session")]
    public partial class Session : Pulumi.CustomResource
    {
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
        /// </summary>
        [Output("findingPublishingFrequency")]
        public Output<string?> FindingPublishingFrequency { get; private set; } = null!;

        [Output("serviceRole")]
        public Output<string> ServiceRole { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Session resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Session(string name, SessionArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:Macie:Session", name, args ?? new SessionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Session(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:Macie:Session", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Session resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Session Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Session(name, id, options);
        }
    }

    public sealed class SessionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
        /// </summary>
        [Input("findingPublishingFrequency")]
        public Input<string>? FindingPublishingFrequency { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SessionArgs()
        {
        }
    }
}
