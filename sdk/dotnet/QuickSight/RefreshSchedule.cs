// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    /// <summary>
    /// Definition of the AWS::QuickSight::RefreshSchedule Resource Type.
    /// </summary>
    [AwsNativeResourceType("aws-native:quicksight:RefreshSchedule")]
    public partial class RefreshSchedule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the data source.&lt;/p&gt;
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsAccountId")]
        public Output<string?> AwsAccountId { get; private set; } = null!;

        [Output("dataSetId")]
        public Output<string?> DataSetId { get; private set; } = null!;

        [Output("schedule")]
        public Output<Outputs.RefreshScheduleMap?> Schedule { get; private set; } = null!;


        /// <summary>
        /// Create a RefreshSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RefreshSchedule(string name, RefreshScheduleArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:quicksight:RefreshSchedule", name, args ?? new RefreshScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RefreshSchedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:quicksight:RefreshSchedule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RefreshSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RefreshSchedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RefreshSchedule(name, id, options);
        }
    }

    public sealed class RefreshScheduleArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        [Input("dataSetId")]
        public Input<string>? DataSetId { get; set; }

        [Input("schedule")]
        public Input<Inputs.RefreshScheduleMapArgs>? Schedule { get; set; }

        public RefreshScheduleArgs()
        {
        }
        public static new RefreshScheduleArgs Empty => new RefreshScheduleArgs();
    }
}