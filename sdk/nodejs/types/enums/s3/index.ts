// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessPointNetworkOrigin = {
    Internet: "Internet",
    Vpc: "VPC",
} as const;

/**
 * Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.
 */
export type AccessPointNetworkOrigin = (typeof AccessPointNetworkOrigin)[keyof typeof AccessPointNetworkOrigin];

export const BucketAccelerateConfigurationAccelerationStatus = {
    Enabled: "Enabled",
    Suspended: "Suspended",
} as const;

/**
 * Configures the transfer acceleration state for an Amazon S3 bucket.
 */
export type BucketAccelerateConfigurationAccelerationStatus = (typeof BucketAccelerateConfigurationAccelerationStatus)[keyof typeof BucketAccelerateConfigurationAccelerationStatus];

export const BucketAccessControl = {
    AuthenticatedRead: "AuthenticatedRead",
    AwsExecRead: "AwsExecRead",
    BucketOwnerFullControl: "BucketOwnerFullControl",
    BucketOwnerRead: "BucketOwnerRead",
    LogDeliveryWrite: "LogDeliveryWrite",
    Private: "Private",
    PublicRead: "PublicRead",
    PublicReadWrite: "PublicReadWrite",
} as const;

/**
 * A canned access control list (ACL) that grants predefined permissions to the bucket.
 */
export type BucketAccessControl = (typeof BucketAccessControl)[keyof typeof BucketAccessControl];

export const BucketCorsRuleAllowedMethodsItem = {
    Get: "GET",
    Put: "PUT",
    Head: "HEAD",
    Post: "POST",
    Delete: "DELETE",
} as const;

export type BucketCorsRuleAllowedMethodsItem = (typeof BucketCorsRuleAllowedMethodsItem)[keyof typeof BucketCorsRuleAllowedMethodsItem];

export const BucketDefaultRetentionMode = {
    Compliance: "COMPLIANCE",
    Governance: "GOVERNANCE",
} as const;

export type BucketDefaultRetentionMode = (typeof BucketDefaultRetentionMode)[keyof typeof BucketDefaultRetentionMode];

export const BucketDeleteMarkerReplicationStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

export type BucketDeleteMarkerReplicationStatus = (typeof BucketDeleteMarkerReplicationStatus)[keyof typeof BucketDeleteMarkerReplicationStatus];

export const BucketDestinationFormat = {
    Csv: "CSV",
    Orc: "ORC",
    Parquet: "Parquet",
} as const;

/**
 * Specifies the file format used when exporting data to Amazon S3.
 */
export type BucketDestinationFormat = (typeof BucketDestinationFormat)[keyof typeof BucketDestinationFormat];

export const BucketIntelligentTieringConfigurationStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies the status of the configuration.
 */
export type BucketIntelligentTieringConfigurationStatus = (typeof BucketIntelligentTieringConfigurationStatus)[keyof typeof BucketIntelligentTieringConfigurationStatus];

export const BucketInventoryConfigurationIncludedObjectVersions = {
    All: "All",
    Current: "Current",
} as const;

/**
 * Object versions to include in the inventory list.
 */
export type BucketInventoryConfigurationIncludedObjectVersions = (typeof BucketInventoryConfigurationIncludedObjectVersions)[keyof typeof BucketInventoryConfigurationIncludedObjectVersions];

export const BucketInventoryConfigurationOptionalFieldsItem = {
    Size: "Size",
    LastModifiedDate: "LastModifiedDate",
    StorageClass: "StorageClass",
    ETag: "ETag",
    IsMultipartUploaded: "IsMultipartUploaded",
    ReplicationStatus: "ReplicationStatus",
    EncryptionStatus: "EncryptionStatus",
    ObjectLockRetainUntilDate: "ObjectLockRetainUntilDate",
    ObjectLockMode: "ObjectLockMode",
    ObjectLockLegalHoldStatus: "ObjectLockLegalHoldStatus",
    IntelligentTieringAccessTier: "IntelligentTieringAccessTier",
    BucketKeyStatus: "BucketKeyStatus",
} as const;

export type BucketInventoryConfigurationOptionalFieldsItem = (typeof BucketInventoryConfigurationOptionalFieldsItem)[keyof typeof BucketInventoryConfigurationOptionalFieldsItem];

export const BucketInventoryConfigurationScheduleFrequency = {
    Daily: "Daily",
    Weekly: "Weekly",
} as const;

/**
 * Specifies the schedule for generating inventory results.
 */
export type BucketInventoryConfigurationScheduleFrequency = (typeof BucketInventoryConfigurationScheduleFrequency)[keyof typeof BucketInventoryConfigurationScheduleFrequency];

export const BucketMetricsStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

export type BucketMetricsStatus = (typeof BucketMetricsStatus)[keyof typeof BucketMetricsStatus];

export const BucketNoncurrentVersionTransitionStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    StandardIa: "STANDARD_IA",
} as const;

/**
 * The class of storage used to store the object.
 */
export type BucketNoncurrentVersionTransitionStorageClass = (typeof BucketNoncurrentVersionTransitionStorageClass)[keyof typeof BucketNoncurrentVersionTransitionStorageClass];

export const BucketOwnershipControlsRuleObjectOwnership = {
    ObjectWriter: "ObjectWriter",
    BucketOwnerPreferred: "BucketOwnerPreferred",
    BucketOwnerEnforced: "BucketOwnerEnforced",
} as const;

/**
 * Specifies an object ownership rule.
 */
export type BucketOwnershipControlsRuleObjectOwnership = (typeof BucketOwnershipControlsRuleObjectOwnership)[keyof typeof BucketOwnershipControlsRuleObjectOwnership];

export const BucketRedirectAllRequestsToProtocol = {
    Http: "http",
    Https: "https",
} as const;

/**
 * Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
 */
export type BucketRedirectAllRequestsToProtocol = (typeof BucketRedirectAllRequestsToProtocol)[keyof typeof BucketRedirectAllRequestsToProtocol];

export const BucketRedirectRuleProtocol = {
    Http: "http",
    Https: "https",
} as const;

/**
 * Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
 */
export type BucketRedirectRuleProtocol = (typeof BucketRedirectRuleProtocol)[keyof typeof BucketRedirectRuleProtocol];

export const BucketReplicaModificationsStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Specifies whether Amazon S3 replicates modifications on replicas.
 */
export type BucketReplicaModificationsStatus = (typeof BucketReplicaModificationsStatus)[keyof typeof BucketReplicaModificationsStatus];

export const BucketReplicationDestinationStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    ReducedRedundancy: "REDUCED_REDUNDANCY",
    Standard: "STANDARD",
    StandardIa: "STANDARD_IA",
} as const;

/**
 * The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
 */
export type BucketReplicationDestinationStorageClass = (typeof BucketReplicationDestinationStorageClass)[keyof typeof BucketReplicationDestinationStorageClass];

export const BucketReplicationRuleStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether the rule is enabled.
 */
export type BucketReplicationRuleStatus = (typeof BucketReplicationRuleStatus)[keyof typeof BucketReplicationRuleStatus];

export const BucketReplicationTimeStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

export type BucketReplicationTimeStatus = (typeof BucketReplicationTimeStatus)[keyof typeof BucketReplicationTimeStatus];

export const BucketRuleStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

export type BucketRuleStatus = (typeof BucketRuleStatus)[keyof typeof BucketRuleStatus];

export const BucketServerSideEncryptionByDefaultSSEAlgorithm = {
    Awskms: "aws:kms",
    Aes256: "AES256",
    Awskmsdsse: "aws:kms:dsse",
} as const;

export type BucketServerSideEncryptionByDefaultSSEAlgorithm = (typeof BucketServerSideEncryptionByDefaultSSEAlgorithm)[keyof typeof BucketServerSideEncryptionByDefaultSSEAlgorithm];

export const BucketSseKmsEncryptedObjectsStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * Specifies whether Amazon S3 replicates objects created with server-side encryption using a customer master key (CMK) stored in AWS Key Management Service.
 */
export type BucketSseKmsEncryptedObjectsStatus = (typeof BucketSseKmsEncryptedObjectsStatus)[keyof typeof BucketSseKmsEncryptedObjectsStatus];

export const BucketTieringAccessTier = {
    ArchiveAccess: "ARCHIVE_ACCESS",
    DeepArchiveAccess: "DEEP_ARCHIVE_ACCESS",
} as const;

/**
 * S3 Intelligent-Tiering access tier. See Storage class for automatically optimizing frequently and infrequently accessed objects for a list of access tiers in the S3 Intelligent-Tiering storage class.
 */
export type BucketTieringAccessTier = (typeof BucketTieringAccessTier)[keyof typeof BucketTieringAccessTier];

export const BucketTransitionStorageClass = {
    DeepArchive: "DEEP_ARCHIVE",
    Glacier: "GLACIER",
    GlacierIr: "GLACIER_IR",
    IntelligentTiering: "INTELLIGENT_TIERING",
    OnezoneIa: "ONEZONE_IA",
    StandardIa: "STANDARD_IA",
} as const;

export type BucketTransitionStorageClass = (typeof BucketTransitionStorageClass)[keyof typeof BucketTransitionStorageClass];

export const BucketVersioningConfigurationStatus = {
    Enabled: "Enabled",
    Suspended: "Suspended",
} as const;

/**
 * The versioning state of the bucket.
 */
export type BucketVersioningConfigurationStatus = (typeof BucketVersioningConfigurationStatus)[keyof typeof BucketVersioningConfigurationStatus];

export const MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic = {
    True: "true",
    False: "false",
} as const;

/**
 * Specifies whether the policy is public or not.
 */
export type MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic = (typeof MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic)[keyof typeof MultiRegionAccessPointPolicyPolicyStatusPropertiesIsPublic];

export const StorageLensS3BucketDestinationFormat = {
    Csv: "CSV",
    Parquet: "Parquet",
} as const;

/**
 * Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.
 */
export type StorageLensS3BucketDestinationFormat = (typeof StorageLensS3BucketDestinationFormat)[keyof typeof StorageLensS3BucketDestinationFormat];

export const StorageLensS3BucketDestinationOutputSchemaVersion = {
    V1: "V_1",
} as const;

/**
 * The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.
 */
export type StorageLensS3BucketDestinationOutputSchemaVersion = (typeof StorageLensS3BucketDestinationOutputSchemaVersion)[keyof typeof StorageLensS3BucketDestinationOutputSchemaVersion];
