// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3
{
    /// <summary>
    /// The ``AWS::S3::Bucket`` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
    ///  To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
    ///   You can only delete empty buckets. Deletion fails for buckets that have contents.
    /// </summary>
    [AwsNativeResourceType("aws-native:s3:Bucket")]
    public partial class Bucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Output("accelerateConfiguration")]
        public Output<Outputs.BucketAccelerateConfiguration?> AccelerateConfiguration { get; private set; } = null!;

        /// <summary>
        /// This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
        ///   A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.
        ///   S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.
        ///   The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Output("accessControl")]
        public Output<Pulumi.AwsNative.S3.BucketAccessControl?> AccessControl { get; private set; } = null!;

        /// <summary>
        /// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
        /// </summary>
        [Output("analyticsConfigurations")]
        public Output<ImmutableArray<Outputs.BucketAnalyticsConfiguration>> AnalyticsConfigurations { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Output("bucketEncryption")]
        public Output<Outputs.BucketEncryption?> BucketEncryption { get; private set; } = null!;

        /// <summary>
        /// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*. 
        ///   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
        /// </summary>
        [Output("bucketName")]
        public Output<string?> BucketName { get; private set; } = null!;

        /// <summary>
        /// Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Output("corsConfiguration")]
        public Output<Outputs.BucketCorsConfiguration?> CorsConfiguration { get; private set; } = null!;

        /// <summary>
        /// The IPv4 DNS name of the specified bucket.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
        /// </summary>
        [Output("dualStackDomainName")]
        public Output<string> DualStackDomainName { get; private set; } = null!;

        /// <summary>
        /// Defines how Amazon S3 handles Intelligent-Tiering storage.
        /// </summary>
        [Output("intelligentTieringConfigurations")]
        public Output<ImmutableArray<Outputs.BucketIntelligentTieringConfiguration>> IntelligentTieringConfigurations { get; private set; } = null!;

        /// <summary>
        /// Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
        /// </summary>
        [Output("inventoryConfigurations")]
        public Output<ImmutableArray<Outputs.BucketInventoryConfiguration>> InventoryConfigurations { get; private set; } = null!;

        /// <summary>
        /// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Output("lifecycleConfiguration")]
        public Output<Outputs.BucketLifecycleConfiguration?> LifecycleConfiguration { get; private set; } = null!;

        /// <summary>
        /// Settings that define where logs are stored.
        /// </summary>
        [Output("loggingConfiguration")]
        public Output<Outputs.BucketLoggingConfiguration?> LoggingConfiguration { get; private set; } = null!;

        /// <summary>
        /// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
        /// </summary>
        [Output("metricsConfigurations")]
        public Output<ImmutableArray<Outputs.BucketMetricsConfiguration>> MetricsConfigurations { get; private set; } = null!;

        /// <summary>
        /// Configuration that defines how Amazon S3 handles bucket notifications.
        /// </summary>
        [Output("notificationConfiguration")]
        public Output<Outputs.BucketNotificationConfiguration?> NotificationConfiguration { get; private set; } = null!;

        /// <summary>
        /// This operation is not supported by directory buckets.
        ///   Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). 
        ///    +  The ``DefaultRetention`` settings require both a mode and a period.
        ///   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
        ///   +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
        /// </summary>
        [Output("objectLockConfiguration")]
        public Output<Outputs.BucketObjectLockConfiguration?> ObjectLockConfiguration { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket.
        /// </summary>
        [Output("objectLockEnabled")]
        public Output<bool?> ObjectLockEnabled { get; private set; } = null!;

        /// <summary>
        /// Configuration that defines how Amazon S3 handles Object Ownership rules.
        /// </summary>
        [Output("ownershipControls")]
        public Output<Outputs.BucketOwnershipControls?> OwnershipControls { get; private set; } = null!;

        /// <summary>
        /// Configuration that defines how Amazon S3 handles public access.
        /// </summary>
        [Output("publicAccessBlockConfiguration")]
        public Output<Outputs.BucketPublicAccessBlockConfiguration?> PublicAccessBlockConfiguration { get; private set; } = null!;

        /// <summary>
        /// Returns the regional domain name of the specified bucket.
        /// </summary>
        [Output("regionalDomainName")]
        public Output<string> RegionalDomainName { get; private set; } = null!;

        /// <summary>
        /// Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
        ///  Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
        /// </summary>
        [Output("replicationConfiguration")]
        public Output<Outputs.BucketReplicationConfiguration?> ReplicationConfiguration { get; private set; } = null!;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this S3 bucket.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
        /// </summary>
        [Output("versioningConfiguration")]
        public Output<Outputs.BucketVersioningConfiguration?> VersioningConfiguration { get; private set; } = null!;

        /// <summary>
        /// Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
        /// </summary>
        [Output("websiteConfiguration")]
        public Output<Outputs.BucketWebsiteConfiguration?> WebsiteConfiguration { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 website endpoint for the specified bucket.
        /// </summary>
        [Output("websiteUrl")]
        public Output<string> WebsiteUrl { get; private set; } = null!;


        /// <summary>
        /// Create a Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bucket(string name, BucketArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:s3:Bucket", name, args ?? new BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bucket(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3:Bucket", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "bucketName",
                    "objectLockEnabled",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bucket Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Bucket(name, id, options);
        }
    }

    public sealed class BucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("accelerateConfiguration")]
        public Input<Inputs.BucketAccelerateConfigurationArgs>? AccelerateConfiguration { get; set; }

        /// <summary>
        /// This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
        ///   A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.
        ///   S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.
        ///   The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("accessControl")]
        public Input<Pulumi.AwsNative.S3.BucketAccessControl>? AccessControl { get; set; }

        [Input("analyticsConfigurations")]
        private InputList<Inputs.BucketAnalyticsConfigurationArgs>? _analyticsConfigurations;

        /// <summary>
        /// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
        /// </summary>
        public InputList<Inputs.BucketAnalyticsConfigurationArgs> AnalyticsConfigurations
        {
            get => _analyticsConfigurations ?? (_analyticsConfigurations = new InputList<Inputs.BucketAnalyticsConfigurationArgs>());
            set => _analyticsConfigurations = value;
        }

        /// <summary>
        /// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("bucketEncryption")]
        public Input<Inputs.BucketEncryptionArgs>? BucketEncryption { get; set; }

        /// <summary>
        /// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*. 
        ///   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("corsConfiguration")]
        public Input<Inputs.BucketCorsConfigurationArgs>? CorsConfiguration { get; set; }

        [Input("intelligentTieringConfigurations")]
        private InputList<Inputs.BucketIntelligentTieringConfigurationArgs>? _intelligentTieringConfigurations;

        /// <summary>
        /// Defines how Amazon S3 handles Intelligent-Tiering storage.
        /// </summary>
        public InputList<Inputs.BucketIntelligentTieringConfigurationArgs> IntelligentTieringConfigurations
        {
            get => _intelligentTieringConfigurations ?? (_intelligentTieringConfigurations = new InputList<Inputs.BucketIntelligentTieringConfigurationArgs>());
            set => _intelligentTieringConfigurations = value;
        }

        [Input("inventoryConfigurations")]
        private InputList<Inputs.BucketInventoryConfigurationArgs>? _inventoryConfigurations;

        /// <summary>
        /// Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
        /// </summary>
        public InputList<Inputs.BucketInventoryConfigurationArgs> InventoryConfigurations
        {
            get => _inventoryConfigurations ?? (_inventoryConfigurations = new InputList<Inputs.BucketInventoryConfigurationArgs>());
            set => _inventoryConfigurations = value;
        }

        /// <summary>
        /// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
        /// </summary>
        [Input("lifecycleConfiguration")]
        public Input<Inputs.BucketLifecycleConfigurationArgs>? LifecycleConfiguration { get; set; }

        /// <summary>
        /// Settings that define where logs are stored.
        /// </summary>
        [Input("loggingConfiguration")]
        public Input<Inputs.BucketLoggingConfigurationArgs>? LoggingConfiguration { get; set; }

        [Input("metricsConfigurations")]
        private InputList<Inputs.BucketMetricsConfigurationArgs>? _metricsConfigurations;

        /// <summary>
        /// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
        /// </summary>
        public InputList<Inputs.BucketMetricsConfigurationArgs> MetricsConfigurations
        {
            get => _metricsConfigurations ?? (_metricsConfigurations = new InputList<Inputs.BucketMetricsConfigurationArgs>());
            set => _metricsConfigurations = value;
        }

        /// <summary>
        /// Configuration that defines how Amazon S3 handles bucket notifications.
        /// </summary>
        [Input("notificationConfiguration")]
        public Input<Inputs.BucketNotificationConfigurationArgs>? NotificationConfiguration { get; set; }

        /// <summary>
        /// This operation is not supported by directory buckets.
        ///   Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). 
        ///    +  The ``DefaultRetention`` settings require both a mode and a period.
        ///   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
        ///   +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
        /// </summary>
        [Input("objectLockConfiguration")]
        public Input<Inputs.BucketObjectLockConfigurationArgs>? ObjectLockConfiguration { get; set; }

        /// <summary>
        /// Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket.
        /// </summary>
        [Input("objectLockEnabled")]
        public Input<bool>? ObjectLockEnabled { get; set; }

        /// <summary>
        /// Configuration that defines how Amazon S3 handles Object Ownership rules.
        /// </summary>
        [Input("ownershipControls")]
        public Input<Inputs.BucketOwnershipControlsArgs>? OwnershipControls { get; set; }

        /// <summary>
        /// Configuration that defines how Amazon S3 handles public access.
        /// </summary>
        [Input("publicAccessBlockConfiguration")]
        public Input<Inputs.BucketPublicAccessBlockConfigurationArgs>? PublicAccessBlockConfiguration { get; set; }

        /// <summary>
        /// Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
        ///  Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
        /// </summary>
        [Input("replicationConfiguration")]
        public Input<Inputs.BucketReplicationConfigurationArgs>? ReplicationConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this S3 bucket.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
        /// </summary>
        [Input("versioningConfiguration")]
        public Input<Inputs.BucketVersioningConfigurationArgs>? VersioningConfiguration { get; set; }

        /// <summary>
        /// Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
        /// </summary>
        [Input("websiteConfiguration")]
        public Input<Inputs.BucketWebsiteConfigurationArgs>? WebsiteConfiguration { get; set; }

        public BucketArgs()
        {
        }
        public static new BucketArgs Empty => new BucketArgs();
    }
}
