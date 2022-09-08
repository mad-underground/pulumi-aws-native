// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::S3::Bucket
 */
export function getBucket(args: GetBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:s3:getBucket", {
        "bucketName": args.bucketName,
    }, opts);
}

export interface GetBucketArgs {
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
     */
    bucketName: string;
}

export interface GetBucketResult {
    /**
     * Configuration for the transfer acceleration state.
     */
    readonly accelerateConfiguration?: outputs.s3.BucketAccelerateConfiguration;
    /**
     * A canned access control list (ACL) that grants predefined permissions to the bucket.
     */
    readonly accessControl?: enums.s3.BucketAccessControl;
    /**
     * The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
     */
    readonly analyticsConfigurations?: outputs.s3.BucketAnalyticsConfiguration[];
    /**
     * The Amazon Resource Name (ARN) of the specified bucket.
     */
    readonly arn?: string;
    readonly bucketEncryption?: outputs.s3.BucketEncryption;
    /**
     * Rules that define cross-origin resource sharing of objects in this bucket.
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
     * Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
     */
    readonly intelligentTieringConfigurations?: outputs.s3.BucketIntelligentTieringConfiguration[];
    /**
     * The inventory configuration for an Amazon S3 bucket.
     */
    readonly inventoryConfigurations?: outputs.s3.BucketInventoryConfiguration[];
    /**
     * Rules that define how Amazon S3 manages objects during their lifetime.
     */
    readonly lifecycleConfiguration?: outputs.s3.BucketLifecycleConfiguration;
    /**
     * Settings that define where logs are stored.
     */
    readonly loggingConfiguration?: outputs.s3.BucketLoggingConfiguration;
    /**
     * Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
     */
    readonly metricsConfigurations?: outputs.s3.BucketMetricsConfiguration[];
    /**
     * Configuration that defines how Amazon S3 handles bucket notifications.
     */
    readonly notificationConfiguration?: outputs.s3.BucketNotificationConfiguration;
    /**
     * Places an Object Lock configuration on the specified bucket.
     */
    readonly objectLockConfiguration?: outputs.s3.BucketObjectLockConfiguration;
    /**
     * Specifies the container element for object ownership rules.
     */
    readonly ownershipControls?: outputs.s3.BucketOwnershipControls;
    readonly publicAccessBlockConfiguration?: outputs.s3.BucketPublicAccessBlockConfiguration;
    /**
     * Returns the regional domain name of the specified bucket.
     */
    readonly regionalDomainName?: string;
    /**
     * Configuration for replicating objects in an S3 bucket.
     */
    readonly replicationConfiguration?: outputs.s3.BucketReplicationConfiguration;
    /**
     * An arbitrary set of tags (key-value pairs) for this S3 bucket.
     */
    readonly tags?: outputs.s3.BucketTag[];
    readonly versioningConfiguration?: outputs.s3.BucketVersioningConfiguration;
    readonly websiteConfiguration?: outputs.s3.BucketWebsiteConfiguration;
    /**
     * The Amazon S3 website endpoint for the specified bucket.
     */
    readonly websiteURL?: string;
}

export function getBucketOutput(args: GetBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketResult> {
    return pulumi.output(args).apply(a => getBucket(a, opts))
}

export interface GetBucketOutputArgs {
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
     */
    bucketName: pulumi.Input<string>;
}
