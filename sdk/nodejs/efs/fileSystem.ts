// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
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
    public static readonly __pulumiType = 'aws-native:EFS:FileSystem';

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
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-backuppolicy
     */
    public readonly backupPolicy!: pulumi.Output<outputs.EFS.FileSystemBackupPolicy | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-encrypted
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly fileSystemId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystempolicy
     */
    public readonly fileSystemPolicy!: pulumi.Output<any | string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
     */
    public readonly fileSystemTags!: pulumi.Output<outputs.EFS.FileSystemElasticFileSystemTag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-kmskeyid
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-lifecyclepolicies
     */
    public readonly lifecyclePolicies!: pulumi.Output<outputs.EFS.FileSystemLifecyclePolicy[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
     */
    public readonly performanceMode!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-provisionedthroughputinmibps
     */
    public readonly provisionedThroughputInMibps!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-throughputmode
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["backupPolicy"] = args ? args.backupPolicy : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["fileSystemPolicy"] = args ? args.fileSystemPolicy : undefined;
            inputs["fileSystemTags"] = args ? args.fileSystemTags : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["lifecyclePolicies"] = args ? args.lifecyclePolicies : undefined;
            inputs["performanceMode"] = args ? args.performanceMode : undefined;
            inputs["provisionedThroughputInMibps"] = args ? args.provisionedThroughputInMibps : undefined;
            inputs["throughputMode"] = args ? args.throughputMode : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["fileSystemId"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["backupPolicy"] = undefined /*out*/;
            inputs["encrypted"] = undefined /*out*/;
            inputs["fileSystemId"] = undefined /*out*/;
            inputs["fileSystemPolicy"] = undefined /*out*/;
            inputs["fileSystemTags"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["lifecyclePolicies"] = undefined /*out*/;
            inputs["performanceMode"] = undefined /*out*/;
            inputs["provisionedThroughputInMibps"] = undefined /*out*/;
            inputs["throughputMode"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-backuppolicy
     */
    backupPolicy?: pulumi.Input<inputs.EFS.FileSystemBackupPolicyArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-encrypted
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystempolicy
     */
    fileSystemPolicy?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
     */
    fileSystemTags?: pulumi.Input<pulumi.Input<inputs.EFS.FileSystemElasticFileSystemTagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-kmskeyid
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-lifecyclepolicies
     */
    lifecyclePolicies?: pulumi.Input<pulumi.Input<inputs.EFS.FileSystemLifecyclePolicyArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
     */
    performanceMode?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-provisionedthroughputinmibps
     */
    provisionedThroughputInMibps?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-throughputmode
     */
    throughputMode?: pulumi.Input<string>;
}
