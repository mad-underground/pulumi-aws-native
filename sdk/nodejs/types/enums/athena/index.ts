// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DataCatalogType = {
    Lambda: "LAMBDA",
    Glue: "GLUE",
    Hive: "HIVE",
} as const;

/**
 * The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. 
 */
export type DataCatalogType = (typeof DataCatalogType)[keyof typeof DataCatalogType];

export const WorkGroupEncryptionOption = {
    SseS3: "SSE_S3",
    SseKms: "SSE_KMS",
    CseKms: "CSE_KMS",
} as const;

/**
 * Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.
 */
export type WorkGroupEncryptionOption = (typeof WorkGroupEncryptionOption)[keyof typeof WorkGroupEncryptionOption];

export const WorkGroupS3AclOption = {
    BucketOwnerFullControl: "BUCKET_OWNER_FULL_CONTROL",
} as const;

/**
 * The Amazon S3 canned ACL that Athena should specify when storing query results. Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL
 */
export type WorkGroupS3AclOption = (typeof WorkGroupS3AclOption)[keyof typeof WorkGroupS3AclOption];

export const WorkGroupState = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * The state of the workgroup: ENABLED or DISABLED.
 */
export type WorkGroupState = (typeof WorkGroupState)[keyof typeof WorkGroupState];
