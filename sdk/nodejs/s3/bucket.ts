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
export class Bucket extends pulumi.CustomResource {
    /**
     * Get an existing Bucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Bucket {
        return new Bucket(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:s3:Bucket';

    /**
     * Returns true if the given object is an instance of Bucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bucket.__pulumiType;
    }

    /**
     * Configuration for the transfer acceleration state.
     */
    public readonly accelerateConfiguration!: pulumi.Output<outputs.s3.BucketAccelerateConfiguration | undefined>;
    /**
     * A canned access control list (ACL) that grants predefined permissions to the bucket.
     */
    public readonly accessControl!: pulumi.Output<enums.s3.BucketAccessControl | undefined>;
    /**
     * The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
     */
    public readonly analyticsConfigurations!: pulumi.Output<outputs.s3.BucketAnalyticsConfiguration[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the specified bucket.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly bucketEncryption!: pulumi.Output<outputs.s3.BucketEncryption | undefined>;
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
     */
    public readonly bucketName!: pulumi.Output<string | undefined>;
    /**
     * Rules that define cross-origin resource sharing of objects in this bucket.
     */
    public readonly corsConfiguration!: pulumi.Output<outputs.s3.BucketCorsConfiguration | undefined>;
    /**
     * The IPv4 DNS name of the specified bucket.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
     */
    public /*out*/ readonly dualStackDomainName!: pulumi.Output<string>;
    /**
     * Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
     */
    public readonly intelligentTieringConfigurations!: pulumi.Output<outputs.s3.BucketIntelligentTieringConfiguration[] | undefined>;
    /**
     * The inventory configuration for an Amazon S3 bucket.
     */
    public readonly inventoryConfigurations!: pulumi.Output<outputs.s3.BucketInventoryConfiguration[] | undefined>;
    /**
     * Rules that define how Amazon S3 manages objects during their lifetime.
     */
    public readonly lifecycleConfiguration!: pulumi.Output<outputs.s3.BucketLifecycleConfiguration | undefined>;
    /**
     * Settings that define where logs are stored.
     */
    public readonly loggingConfiguration!: pulumi.Output<outputs.s3.BucketLoggingConfiguration | undefined>;
    /**
     * Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
     */
    public readonly metricsConfigurations!: pulumi.Output<outputs.s3.BucketMetricsConfiguration[] | undefined>;
    /**
     * Configuration that defines how Amazon S3 handles bucket notifications.
     */
    public readonly notificationConfiguration!: pulumi.Output<outputs.s3.BucketNotificationConfiguration | undefined>;
    /**
     * Places an Object Lock configuration on the specified bucket.
     */
    public readonly objectLockConfiguration!: pulumi.Output<outputs.s3.BucketObjectLockConfiguration | undefined>;
    /**
     * Indicates whether this bucket has an Object Lock configuration enabled.
     */
    public readonly objectLockEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the container element for object ownership rules.
     */
    public readonly ownershipControls!: pulumi.Output<outputs.s3.BucketOwnershipControls | undefined>;
    public readonly publicAccessBlockConfiguration!: pulumi.Output<outputs.s3.BucketPublicAccessBlockConfiguration | undefined>;
    /**
     * Returns the regional domain name of the specified bucket.
     */
    public /*out*/ readonly regionalDomainName!: pulumi.Output<string>;
    /**
     * Configuration for replicating objects in an S3 bucket.
     */
    public readonly replicationConfiguration!: pulumi.Output<outputs.s3.BucketReplicationConfiguration | undefined>;
    /**
     * An arbitrary set of tags (key-value pairs) for this S3 bucket.
     */
    public readonly tags!: pulumi.Output<outputs.s3.BucketTag[] | undefined>;
    public readonly versioningConfiguration!: pulumi.Output<outputs.s3.BucketVersioningConfiguration | undefined>;
    public readonly websiteConfiguration!: pulumi.Output<outputs.s3.BucketWebsiteConfiguration | undefined>;
    /**
     * The Amazon S3 website endpoint for the specified bucket.
     */
    public /*out*/ readonly websiteURL!: pulumi.Output<string>;

    /**
     * Create a Bucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BucketArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["accelerateConfiguration"] = args ? args.accelerateConfiguration : undefined;
            resourceInputs["accessControl"] = args ? args.accessControl : undefined;
            resourceInputs["analyticsConfigurations"] = args ? args.analyticsConfigurations : undefined;
            resourceInputs["bucketEncryption"] = args ? args.bucketEncryption : undefined;
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["corsConfiguration"] = args ? args.corsConfiguration : undefined;
            resourceInputs["intelligentTieringConfigurations"] = args ? args.intelligentTieringConfigurations : undefined;
            resourceInputs["inventoryConfigurations"] = args ? args.inventoryConfigurations : undefined;
            resourceInputs["lifecycleConfiguration"] = args ? args.lifecycleConfiguration : undefined;
            resourceInputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            resourceInputs["metricsConfigurations"] = args ? args.metricsConfigurations : undefined;
            resourceInputs["notificationConfiguration"] = args ? args.notificationConfiguration : undefined;
            resourceInputs["objectLockConfiguration"] = args ? args.objectLockConfiguration : undefined;
            resourceInputs["objectLockEnabled"] = args ? args.objectLockEnabled : undefined;
            resourceInputs["ownershipControls"] = args ? args.ownershipControls : undefined;
            resourceInputs["publicAccessBlockConfiguration"] = args ? args.publicAccessBlockConfiguration : undefined;
            resourceInputs["replicationConfiguration"] = args ? args.replicationConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versioningConfiguration"] = args ? args.versioningConfiguration : undefined;
            resourceInputs["websiteConfiguration"] = args ? args.websiteConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["dualStackDomainName"] = undefined /*out*/;
            resourceInputs["regionalDomainName"] = undefined /*out*/;
            resourceInputs["websiteURL"] = undefined /*out*/;
        } else {
            resourceInputs["accelerateConfiguration"] = undefined /*out*/;
            resourceInputs["accessControl"] = undefined /*out*/;
            resourceInputs["analyticsConfigurations"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["bucketEncryption"] = undefined /*out*/;
            resourceInputs["bucketName"] = undefined /*out*/;
            resourceInputs["corsConfiguration"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["dualStackDomainName"] = undefined /*out*/;
            resourceInputs["intelligentTieringConfigurations"] = undefined /*out*/;
            resourceInputs["inventoryConfigurations"] = undefined /*out*/;
            resourceInputs["lifecycleConfiguration"] = undefined /*out*/;
            resourceInputs["loggingConfiguration"] = undefined /*out*/;
            resourceInputs["metricsConfigurations"] = undefined /*out*/;
            resourceInputs["notificationConfiguration"] = undefined /*out*/;
            resourceInputs["objectLockConfiguration"] = undefined /*out*/;
            resourceInputs["objectLockEnabled"] = undefined /*out*/;
            resourceInputs["ownershipControls"] = undefined /*out*/;
            resourceInputs["publicAccessBlockConfiguration"] = undefined /*out*/;
            resourceInputs["regionalDomainName"] = undefined /*out*/;
            resourceInputs["replicationConfiguration"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["versioningConfiguration"] = undefined /*out*/;
            resourceInputs["websiteConfiguration"] = undefined /*out*/;
            resourceInputs["websiteURL"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Bucket resource.
 */
export interface BucketArgs {
    /**
     * Configuration for the transfer acceleration state.
     */
    accelerateConfiguration?: pulumi.Input<inputs.s3.BucketAccelerateConfigurationArgs>;
    /**
     * A canned access control list (ACL) that grants predefined permissions to the bucket.
     */
    accessControl?: pulumi.Input<enums.s3.BucketAccessControl>;
    /**
     * The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
     */
    analyticsConfigurations?: pulumi.Input<pulumi.Input<inputs.s3.BucketAnalyticsConfigurationArgs>[]>;
    bucketEncryption?: pulumi.Input<inputs.s3.BucketEncryptionArgs>;
    /**
     * A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * Rules that define cross-origin resource sharing of objects in this bucket.
     */
    corsConfiguration?: pulumi.Input<inputs.s3.BucketCorsConfigurationArgs>;
    /**
     * Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
     */
    intelligentTieringConfigurations?: pulumi.Input<pulumi.Input<inputs.s3.BucketIntelligentTieringConfigurationArgs>[]>;
    /**
     * The inventory configuration for an Amazon S3 bucket.
     */
    inventoryConfigurations?: pulumi.Input<pulumi.Input<inputs.s3.BucketInventoryConfigurationArgs>[]>;
    /**
     * Rules that define how Amazon S3 manages objects during their lifetime.
     */
    lifecycleConfiguration?: pulumi.Input<inputs.s3.BucketLifecycleConfigurationArgs>;
    /**
     * Settings that define where logs are stored.
     */
    loggingConfiguration?: pulumi.Input<inputs.s3.BucketLoggingConfigurationArgs>;
    /**
     * Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
     */
    metricsConfigurations?: pulumi.Input<pulumi.Input<inputs.s3.BucketMetricsConfigurationArgs>[]>;
    /**
     * Configuration that defines how Amazon S3 handles bucket notifications.
     */
    notificationConfiguration?: pulumi.Input<inputs.s3.BucketNotificationConfigurationArgs>;
    /**
     * Places an Object Lock configuration on the specified bucket.
     */
    objectLockConfiguration?: pulumi.Input<inputs.s3.BucketObjectLockConfigurationArgs>;
    /**
     * Indicates whether this bucket has an Object Lock configuration enabled.
     */
    objectLockEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the container element for object ownership rules.
     */
    ownershipControls?: pulumi.Input<inputs.s3.BucketOwnershipControlsArgs>;
    publicAccessBlockConfiguration?: pulumi.Input<inputs.s3.BucketPublicAccessBlockConfigurationArgs>;
    /**
     * Configuration for replicating objects in an S3 bucket.
     */
    replicationConfiguration?: pulumi.Input<inputs.s3.BucketReplicationConfigurationArgs>;
    /**
     * An arbitrary set of tags (key-value pairs) for this S3 bucket.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.s3.BucketTagArgs>[]>;
    versioningConfiguration?: pulumi.Input<inputs.s3.BucketVersioningConfigurationArgs>;
    websiteConfiguration?: pulumi.Input<inputs.s3.BucketWebsiteConfigurationArgs>;
}
