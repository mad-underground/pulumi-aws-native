// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::FSx::FileSystem
 *
 * @deprecated FileSystem is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
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
        pulumi.log.warn("FileSystem is deprecated: FileSystem is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new FileSystem(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:fsx:FileSystem';

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

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly backupId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    public readonly fileSystemType!: pulumi.Output<string>;
    public readonly fileSystemTypeVersion!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly lustreConfiguration!: pulumi.Output<outputs.fsx.FileSystemLustreConfiguration | undefined>;
    public /*out*/ readonly lustreMountName!: pulumi.Output<string>;
    public readonly ontapConfiguration!: pulumi.Output<outputs.fsx.FileSystemOntapConfiguration | undefined>;
    public readonly openZfsConfiguration!: pulumi.Output<outputs.fsx.FileSystemOpenZfsConfiguration | undefined>;
    public /*out*/ readonly resourceArn!: pulumi.Output<string>;
    public /*out*/ readonly rootVolumeId!: pulumi.Output<string>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly storageCapacity!: pulumi.Output<number | undefined>;
    public readonly storageType!: pulumi.Output<string | undefined>;
    public readonly subnetIds!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly windowsConfiguration!: pulumi.Output<outputs.fsx.FileSystemWindowsConfiguration | undefined>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated FileSystem is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: FileSystemArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("FileSystem is deprecated: FileSystem is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.fileSystemType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemType'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["fileSystemType"] = args ? args.fileSystemType : undefined;
            resourceInputs["fileSystemTypeVersion"] = args ? args.fileSystemTypeVersion : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["lustreConfiguration"] = args ? args.lustreConfiguration : undefined;
            resourceInputs["ontapConfiguration"] = args ? args.ontapConfiguration : undefined;
            resourceInputs["openZfsConfiguration"] = args ? args.openZfsConfiguration : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["storageCapacity"] = args ? args.storageCapacity : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["windowsConfiguration"] = args ? args.windowsConfiguration : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["lustreMountName"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["rootVolumeId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["backupId"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["fileSystemType"] = undefined /*out*/;
            resourceInputs["fileSystemTypeVersion"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["lustreConfiguration"] = undefined /*out*/;
            resourceInputs["lustreMountName"] = undefined /*out*/;
            resourceInputs["ontapConfiguration"] = undefined /*out*/;
            resourceInputs["openZfsConfiguration"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["rootVolumeId"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["storageCapacity"] = undefined /*out*/;
            resourceInputs["storageType"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["windowsConfiguration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["backupId", "fileSystemType", "fileSystemTypeVersion", "kmsKeyId", "securityGroupIds[*]", "subnetIds[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    backupId?: pulumi.Input<string>;
    fileSystemType: pulumi.Input<string>;
    fileSystemTypeVersion?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    lustreConfiguration?: pulumi.Input<inputs.fsx.FileSystemLustreConfigurationArgs>;
    ontapConfiguration?: pulumi.Input<inputs.fsx.FileSystemOntapConfigurationArgs>;
    openZfsConfiguration?: pulumi.Input<inputs.fsx.FileSystemOpenZfsConfigurationArgs>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    storageCapacity?: pulumi.Input<number>;
    storageType?: pulumi.Input<string>;
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    windowsConfiguration?: pulumi.Input<inputs.fsx.FileSystemWindowsConfigurationArgs>;
}
