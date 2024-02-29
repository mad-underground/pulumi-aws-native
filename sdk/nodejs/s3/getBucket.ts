// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::S3::Bucket`` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
 *  To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
 *   You can only delete empty buckets. Deletion fails for buckets that have contents.
 */
export function getBucket(args: GetBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:s3:getBucket", {
        "bucketName": args.bucketName,
    }, opts);
}

export interface GetBucketArgs {
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*. 
     *   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
     */
    bucketName: string;
}

export interface GetBucketResult {
    /**
     * Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
     */
    readonly accelerateConfiguration?: outputs.s3.BucketAccelerateConfiguration;
    /**
     * Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
     */
    readonly analyticsConfigurations?: outputs.s3.BucketAnalyticsConfiguration[];
    /**
     * The Amazon Resource Name (ARN) of the specified bucket.
     */
    readonly arn?: string;
    /**
     * Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
     */
    readonly bucketEncryption?: outputs.s3.BucketEncryption;
    /**
     * Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
     */
    readonly corsConfiguration?: outputs.s3.BucketCorsConfiguration;
    /**
     * The IPv4 DNS name of the specified bucket.
     */
    readonly domainName?: string;
    /**
     * The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
     */
    readonly dualStackDomainName?: string;
    /**
     * Defines how Amazon S3 handles Intelligent-Tiering storage.
     */
    readonly intelligentTieringConfigurations?: outputs.s3.BucketIntelligentTieringConfiguration[];
    /**
     * Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
     */
    readonly inventoryConfigurations?: outputs.s3.BucketInventoryConfiguration[];
    /**
     * Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
     */
    readonly lifecycleConfiguration?: outputs.s3.BucketLifecycleConfiguration;
    /**
     * Settings that define where logs are stored.
     */
    readonly loggingConfiguration?: outputs.s3.BucketLoggingConfiguration;
    /**
     * Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
     */
    readonly metricsConfigurations?: outputs.s3.BucketMetricsConfiguration[];
    /**
     * Configuration that defines how Amazon S3 handles bucket notifications.
     */
    readonly notificationConfiguration?: outputs.s3.BucketNotificationConfiguration;
    /**
     * This operation is not supported by directory buckets.
     *   Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). 
     *    +  The ``DefaultRetention`` settings require both a mode and a period.
     *   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
     *   +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
     */
    readonly objectLockConfiguration?: outputs.s3.BucketObjectLockConfiguration;
    /**
     * Configuration that defines how Amazon S3 handles Object Ownership rules.
     */
    readonly ownershipControls?: outputs.s3.BucketOwnershipControls;
    /**
     * Configuration that defines how Amazon S3 handles public access.
     */
    readonly publicAccessBlockConfiguration?: outputs.s3.BucketPublicAccessBlockConfiguration;
    /**
     * Returns the regional domain name of the specified bucket.
     */
    readonly regionalDomainName?: string;
    /**
     * Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
     *  Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
     */
    readonly replicationConfiguration?: outputs.s3.BucketReplicationConfiguration;
    /**
     * An arbitrary set of tags (key-value pairs) for this S3 bucket.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
     */
    readonly versioningConfiguration?: outputs.s3.BucketVersioningConfiguration;
    /**
     * Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
     */
    readonly websiteConfiguration?: outputs.s3.BucketWebsiteConfiguration;
    /**
     * The Amazon S3 website endpoint for the specified bucket.
     */
    readonly websiteUrl?: string;
}
/**
 * The ``AWS::S3::Bucket`` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
 *  To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
 *   You can only delete empty buckets. Deletion fails for buckets that have contents.
 */
export function getBucketOutput(args: GetBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketResult> {
    return pulumi.output(args).apply((a: any) => getBucket(a, opts))
}

export interface GetBucketOutputArgs {
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*. 
     *   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
     */
    bucketName: pulumi.Input<string>;
}
