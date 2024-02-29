// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::EFS::FileSystem`` resource creates a new, empty file system in EFSlong (EFS). You must create a mount target ([AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an EC2 or other AWS cloud compute resource.
 */
export class FileSystem extends pulumi.CustomResource {
    /**
     * Get an existing FileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FileSystem {
        return new FileSystem(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:efs:FileSystem';

    /**
     * Returns true if the given object is an instance of FileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileSystem.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*.
     *   One Zone file systems are not available in all Availability Zones in AWS-Regions where Amazon EFS is available.
     */
    public readonly availabilityZoneName!: pulumi.Output<string | undefined>;
    /**
     * Use the ``BackupPolicy`` to turn automatic backups on or off for the file system.
     */
    public readonly backupPolicy!: pulumi.Output<outputs.efs.FileSystemBackupPolicy | undefined>;
    /**
     * (Optional) A boolean that specifies whether or not to bypass the ``FileSystemPolicy`` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ``PutFileSystemPolicy`` requests on this file system. Set ``BypassPolicyLockoutSafetyCheck`` to ``True`` only when you intend to prevent the IAM principal that is making the request from making subsequent ``PutFileSystemPolicy`` requests on this file system. The default value is ``False``.
     */
    public readonly bypassPolicyLockoutSafetyCheck!: pulumi.Output<boolean | undefined>;
    /**
     * A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ``/aws/elasticfilesystem``, is used to protect the encrypted file system.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly fileSystemId!: pulumi.Output<string>;
    /**
     * The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EFS::FileSystem` for more information about the expected schema for this property.
     */
    public readonly fileSystemPolicy!: pulumi.Output<any | undefined>;
    /**
     * Describes the protection on the file system.
     */
    public readonly fileSystemProtection!: pulumi.Output<outputs.efs.FileSystemProtection | undefined>;
    /**
     * Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``"Key":"Name","Value":"{value}"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.
     */
    public readonly fileSystemTags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats:
     *   +  Key ID - A unique identifier of the key, for example ``1234abcd-12ab-34cd-56ef-1234567890ab``.
     *   +  ARN - An Amazon Resource Name (ARN) for the key, for example ``arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``.
     *   +  Key alias - A previously created display name for a key, for example ``alias/projectKey1``.
     *   +  Key alias ARN - An ARN for a key alias, for example ``arn:aws:kms:us-west-2:444455556666:alias/projectKey1``.
     *   
     *  If ``KmsKeyId`` is specified, the ``Encrypted`` parameter must be set to true.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following:
     *   +  When to move files in the file system from primary storage to IA storage.
     *   + When to move files in the file system from primary storage or IA storage to Archive storage.
     *  +  When to move files that are in IA or Archive storage to primary storage.
     *   
     *   EFS requires that each ``LifecyclePolicy`` object have only a single transition. This means that in a request body, ``LifecyclePolicies`` needs to be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive`` ``TransitionToPrimaryStorageClass``. See the example requests in the following section for more information.
     */
    public readonly lifecyclePolicies!: pulumi.Output<outputs.efs.FileSystemLifecyclePolicy[] | undefined>;
    /**
     * The Performance mode of the file system. We recommend ``generalPurpose`` performance mode for all file systems. File systems using the ``maxIO`` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ``maxIO`` mode is not supported on One Zone file systems.
     *   Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.
     *   Default is ``generalPurpose``.
     */
    public readonly performanceMode!: pulumi.Output<string | undefined>;
    /**
     * The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.
     */
    public readonly provisionedThroughputInMibps!: pulumi.Output<number | undefined>;
    /**
     * Describes the replication configuration for a specific file system.
     */
    public readonly replicationConfiguration!: pulumi.Output<outputs.efs.FileSystemReplicationConfiguration | undefined>;
    /**
     * Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. 
     *  Default is ``bursting``.
     */
    public readonly throughputMode!: pulumi.Output<string | undefined>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FileSystemArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["availabilityZoneName"] = args ? args.availabilityZoneName : undefined;
            resourceInputs["backupPolicy"] = args ? args.backupPolicy : undefined;
            resourceInputs["bypassPolicyLockoutSafetyCheck"] = args ? args.bypassPolicyLockoutSafetyCheck : undefined;
            resourceInputs["encrypted"] = args ? args.encrypted : undefined;
            resourceInputs["fileSystemPolicy"] = args ? args.fileSystemPolicy : undefined;
            resourceInputs["fileSystemProtection"] = args ? args.fileSystemProtection : undefined;
            resourceInputs["fileSystemTags"] = args ? args.fileSystemTags : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["lifecyclePolicies"] = args ? args.lifecyclePolicies : undefined;
            resourceInputs["performanceMode"] = args ? args.performanceMode : undefined;
            resourceInputs["provisionedThroughputInMibps"] = args ? args.provisionedThroughputInMibps : undefined;
            resourceInputs["replicationConfiguration"] = args ? args.replicationConfiguration : undefined;
            resourceInputs["throughputMode"] = args ? args.throughputMode : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["fileSystemId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["availabilityZoneName"] = undefined /*out*/;
            resourceInputs["backupPolicy"] = undefined /*out*/;
            resourceInputs["bypassPolicyLockoutSafetyCheck"] = undefined /*out*/;
            resourceInputs["encrypted"] = undefined /*out*/;
            resourceInputs["fileSystemId"] = undefined /*out*/;
            resourceInputs["fileSystemPolicy"] = undefined /*out*/;
            resourceInputs["fileSystemProtection"] = undefined /*out*/;
            resourceInputs["fileSystemTags"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["lifecyclePolicies"] = undefined /*out*/;
            resourceInputs["performanceMode"] = undefined /*out*/;
            resourceInputs["provisionedThroughputInMibps"] = undefined /*out*/;
            resourceInputs["replicationConfiguration"] = undefined /*out*/;
            resourceInputs["throughputMode"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["availabilityZoneName", "encrypted", "kmsKeyId", "performanceMode"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*.
     *   One Zone file systems are not available in all Availability Zones in AWS-Regions where Amazon EFS is available.
     */
    availabilityZoneName?: pulumi.Input<string>;
    /**
     * Use the ``BackupPolicy`` to turn automatic backups on or off for the file system.
     */
    backupPolicy?: pulumi.Input<inputs.efs.FileSystemBackupPolicyArgs>;
    /**
     * (Optional) A boolean that specifies whether or not to bypass the ``FileSystemPolicy`` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ``PutFileSystemPolicy`` requests on this file system. Set ``BypassPolicyLockoutSafetyCheck`` to ``True`` only when you intend to prevent the IAM principal that is making the request from making subsequent ``PutFileSystemPolicy`` requests on this file system. The default value is ``False``.
     */
    bypassPolicyLockoutSafetyCheck?: pulumi.Input<boolean>;
    /**
     * A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ``/aws/elasticfilesystem``, is used to protect the encrypted file system.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EFS::FileSystem` for more information about the expected schema for this property.
     */
    fileSystemPolicy?: any;
    /**
     * Describes the protection on the file system.
     */
    fileSystemProtection?: pulumi.Input<inputs.efs.FileSystemProtectionArgs>;
    /**
     * Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``"Key":"Name","Value":"{value}"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.
     */
    fileSystemTags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats:
     *   +  Key ID - A unique identifier of the key, for example ``1234abcd-12ab-34cd-56ef-1234567890ab``.
     *   +  ARN - An Amazon Resource Name (ARN) for the key, for example ``arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``.
     *   +  Key alias - A previously created display name for a key, for example ``alias/projectKey1``.
     *   +  Key alias ARN - An ARN for a key alias, for example ``arn:aws:kms:us-west-2:444455556666:alias/projectKey1``.
     *   
     *  If ``KmsKeyId`` is specified, the ``Encrypted`` parameter must be set to true.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following:
     *   +  When to move files in the file system from primary storage to IA storage.
     *   + When to move files in the file system from primary storage or IA storage to Archive storage.
     *  +  When to move files that are in IA or Archive storage to primary storage.
     *   
     *   EFS requires that each ``LifecyclePolicy`` object have only a single transition. This means that in a request body, ``LifecyclePolicies`` needs to be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive`` ``TransitionToPrimaryStorageClass``. See the example requests in the following section for more information.
     */
    lifecyclePolicies?: pulumi.Input<pulumi.Input<inputs.efs.FileSystemLifecyclePolicyArgs>[]>;
    /**
     * The Performance mode of the file system. We recommend ``generalPurpose`` performance mode for all file systems. File systems using the ``maxIO`` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ``maxIO`` mode is not supported on One Zone file systems.
     *   Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.
     *   Default is ``generalPurpose``.
     */
    performanceMode?: pulumi.Input<string>;
    /**
     * The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.
     */
    provisionedThroughputInMibps?: pulumi.Input<number>;
    /**
     * Describes the replication configuration for a specific file system.
     */
    replicationConfiguration?: pulumi.Input<inputs.efs.FileSystemReplicationConfigurationArgs>;
    /**
     * Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. 
     *  Default is ``bursting``.
     */
    throughputMode?: pulumi.Input<string>;
}
