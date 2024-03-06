// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::EFS::MountTarget`` resource is an Amazon EFS resource that creates a mount target for an EFS file system. You can then mount the file system on Amazon EC2 instances or other resources by using the mount target.
 */
export class MountTarget extends pulumi.CustomResource {
    /**
     * Get an existing MountTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MountTarget {
        return new MountTarget(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:efs:MountTarget';

    /**
     * Returns true if the given object is an instance of MountTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MountTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MountTarget.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The ID of the file system for which to create the mount target.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * Valid IPv4 address within the address range of the specified subnet.
     */
    public readonly ipAddress!: pulumi.Output<string | undefined>;
    /**
     * Up to five VPC security group IDs, of the form ``sg-xxxxxxxx``. These must be for the same VPC as subnet specified.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a MountTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MountTargetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            if ((!args || args.securityGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroups'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["fileSystemId"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["fileSystemId", "ipAddress", "subnetId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(MountTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MountTarget resource.
 */
export interface MountTargetArgs {
    /**
     * The ID of the file system for which to create the mount target.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * Valid IPv4 address within the address range of the specified subnet.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Up to five VPC security group IDs, of the form ``sg-xxxxxxxx``. These must be for the same VPC as subnet specified.
     */
    securityGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone.
     */
    subnetId: pulumi.Input<string>;
}
