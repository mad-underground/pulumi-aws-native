// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
    /// </summary>
    public sealed class BucketLifecycleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.BucketRuleArgs>? _rules;

        /// <summary>
        /// A lifecycle rule for individual objects in an Amazon S3 bucket.
        /// </summary>
        public InputList<Inputs.BucketRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketRuleArgs>());
            set => _rules = value;
        }

        public BucketLifecycleConfigurationArgs()
        {
        }
        public static new BucketLifecycleConfigurationArgs Empty => new BucketLifecycleConfigurationArgs();
    }
}
