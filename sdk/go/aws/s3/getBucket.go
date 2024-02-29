// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::S3::Bucket“ resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
//
//	To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
//	 You can only delete empty buckets. Deletion fails for buckets that have contents.
func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBucketResult
	err := ctx.Invoke("aws-native:s3:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketArgs struct {
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*.
	//   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName string `pulumi:"bucketName"`
}

type LookupBucketResult struct {
	// Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
	AccelerateConfiguration *BucketAccelerateConfiguration `pulumi:"accelerateConfiguration"`
	// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations []BucketAnalyticsConfiguration `pulumi:"analyticsConfigurations"`
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn *string `pulumi:"arn"`
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
	BucketEncryption *BucketEncryption `pulumi:"bucketEncryption"`
	// Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
	CorsConfiguration *BucketCorsConfiguration `pulumi:"corsConfiguration"`
	// The IPv4 DNS name of the specified bucket.
	DomainName *string `pulumi:"domainName"`
	// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
	DualStackDomainName *string `pulumi:"dualStackDomainName"`
	// Defines how Amazon S3 handles Intelligent-Tiering storage.
	IntelligentTieringConfigurations []BucketIntelligentTieringConfiguration `pulumi:"intelligentTieringConfigurations"`
	// Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
	InventoryConfigurations []BucketInventoryConfiguration `pulumi:"inventoryConfigurations"`
	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
	LifecycleConfiguration *BucketLifecycleConfiguration `pulumi:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	LoggingConfiguration *BucketLoggingConfiguration `pulumi:"loggingConfiguration"`
	// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
	MetricsConfigurations []BucketMetricsConfiguration `pulumi:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration *BucketNotificationConfiguration `pulumi:"notificationConfiguration"`
	// This operation is not supported by directory buckets.
	//   Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	//    +  The ``DefaultRetention`` settings require both a mode and a period.
	//   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
	//   +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// Configuration that defines how Amazon S3 handles Object Ownership rules.
	OwnershipControls *BucketOwnershipControls `pulumi:"ownershipControls"`
	// Configuration that defines how Amazon S3 handles public access.
	PublicAccessBlockConfiguration *BucketPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Returns the regional domain name of the specified bucket.
	RegionalDomainName *string `pulumi:"regionalDomainName"`
	// Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
	//  Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags []aws.Tag `pulumi:"tags"`
	// Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
	VersioningConfiguration *BucketVersioningConfiguration `pulumi:"versioningConfiguration"`
	// Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
	WebsiteConfiguration *BucketWebsiteConfiguration `pulumi:"websiteConfiguration"`
	// The Amazon S3 website endpoint for the specified bucket.
	WebsiteUrl *string `pulumi:"websiteUrl"`
}

func LookupBucketOutput(ctx *pulumi.Context, args LookupBucketOutputArgs, opts ...pulumi.InvokeOption) LookupBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketResult, error) {
			args := v.(LookupBucketArgs)
			r, err := LookupBucket(ctx, &args, opts...)
			var s LookupBucketResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketResultOutput)
}

type LookupBucketOutputArgs struct {
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide*.
	//   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName pulumi.StringInput `pulumi:"bucketName"`
}

func (LookupBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketArgs)(nil)).Elem()
}

type LookupBucketResultOutput struct{ *pulumi.OutputState }

func (LookupBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketResult)(nil)).Elem()
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutput() LookupBucketResultOutput {
	return o
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutputWithContext(ctx context.Context) LookupBucketResultOutput {
	return o
}

// Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
func (o LookupBucketResultOutput) AccelerateConfiguration() BucketAccelerateConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketAccelerateConfiguration { return v.AccelerateConfiguration }).(BucketAccelerateConfigurationPtrOutput)
}

// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
func (o LookupBucketResultOutput) AnalyticsConfigurations() BucketAnalyticsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketAnalyticsConfiguration { return v.AnalyticsConfigurations }).(BucketAnalyticsConfigurationArrayOutput)
}

// The Amazon Resource Name (ARN) of the specified bucket.
func (o LookupBucketResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
func (o LookupBucketResultOutput) BucketEncryption() BucketEncryptionPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketEncryption { return v.BucketEncryption }).(BucketEncryptionPtrOutput)
}

// Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
func (o LookupBucketResultOutput) CorsConfiguration() BucketCorsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketCorsConfiguration { return v.CorsConfiguration }).(BucketCorsConfigurationPtrOutput)
}

// The IPv4 DNS name of the specified bucket.
func (o LookupBucketResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
func (o LookupBucketResultOutput) DualStackDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.DualStackDomainName }).(pulumi.StringPtrOutput)
}

// Defines how Amazon S3 handles Intelligent-Tiering storage.
func (o LookupBucketResultOutput) IntelligentTieringConfigurations() BucketIntelligentTieringConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketIntelligentTieringConfiguration {
		return v.IntelligentTieringConfigurations
	}).(BucketIntelligentTieringConfigurationArrayOutput)
}

// Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
func (o LookupBucketResultOutput) InventoryConfigurations() BucketInventoryConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketInventoryConfiguration { return v.InventoryConfigurations }).(BucketInventoryConfigurationArrayOutput)
}

// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
func (o LookupBucketResultOutput) LifecycleConfiguration() BucketLifecycleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketLifecycleConfiguration { return v.LifecycleConfiguration }).(BucketLifecycleConfigurationPtrOutput)
}

// Settings that define where logs are stored.
func (o LookupBucketResultOutput) LoggingConfiguration() BucketLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketLoggingConfiguration { return v.LoggingConfiguration }).(BucketLoggingConfigurationPtrOutput)
}

// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
func (o LookupBucketResultOutput) MetricsConfigurations() BucketMetricsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketMetricsConfiguration { return v.MetricsConfigurations }).(BucketMetricsConfigurationArrayOutput)
}

// Configuration that defines how Amazon S3 handles bucket notifications.
func (o LookupBucketResultOutput) NotificationConfiguration() BucketNotificationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketNotificationConfiguration { return v.NotificationConfiguration }).(BucketNotificationConfigurationPtrOutput)
}

// This operation is not supported by directory buckets.
//
//	Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
//	 +  The ``DefaultRetention`` settings require both a mode and a period.
//	+  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
//	+  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
func (o LookupBucketResultOutput) ObjectLockConfiguration() BucketObjectLockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketObjectLockConfiguration { return v.ObjectLockConfiguration }).(BucketObjectLockConfigurationPtrOutput)
}

// Configuration that defines how Amazon S3 handles Object Ownership rules.
func (o LookupBucketResultOutput) OwnershipControls() BucketOwnershipControlsPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketOwnershipControls { return v.OwnershipControls }).(BucketOwnershipControlsPtrOutput)
}

// Configuration that defines how Amazon S3 handles public access.
func (o LookupBucketResultOutput) PublicAccessBlockConfiguration() BucketPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketPublicAccessBlockConfiguration {
		return v.PublicAccessBlockConfiguration
	}).(BucketPublicAccessBlockConfigurationPtrOutput)
}

// Returns the regional domain name of the specified bucket.
func (o LookupBucketResultOutput) RegionalDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.RegionalDomainName }).(pulumi.StringPtrOutput)
}

// Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the “VersioningConfiguration“ property.
//
//	Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
func (o LookupBucketResultOutput) ReplicationConfiguration() BucketReplicationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketReplicationConfiguration { return v.ReplicationConfiguration }).(BucketReplicationConfigurationPtrOutput)
}

// An arbitrary set of tags (key-value pairs) for this S3 bucket.
func (o LookupBucketResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
func (o LookupBucketResultOutput) VersioningConfiguration() BucketVersioningConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketVersioningConfiguration { return v.VersioningConfiguration }).(BucketVersioningConfigurationPtrOutput)
}

// Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
func (o LookupBucketResultOutput) WebsiteConfiguration() BucketWebsiteConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketWebsiteConfiguration { return v.WebsiteConfiguration }).(BucketWebsiteConfigurationPtrOutput)
}

// The Amazon S3 website endpoint for the specified bucket.
func (o LookupBucketResultOutput) WebsiteUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.WebsiteUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketResultOutput{})
}
