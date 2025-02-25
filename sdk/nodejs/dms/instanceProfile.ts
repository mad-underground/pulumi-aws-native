// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DMS::InstanceProfile.
 */
export class InstanceProfile extends pulumi.CustomResource {
    /**
     * Get an existing InstanceProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceProfile {
        return new InstanceProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dms:InstanceProfile';

    /**
     * Returns true if the given object is an instance of InstanceProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceProfile.__pulumiType;
    }

    /**
     * The property describes an availability zone of the instance profile.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * The optional description of the instance profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The property describes an ARN of the instance profile.
     */
    public /*out*/ readonly instanceProfileArn!: pulumi.Output<string>;
    /**
     * The property describes a creating time of the instance profile.
     */
    public /*out*/ readonly instanceProfileCreationTime!: pulumi.Output<string>;
    /**
     * The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn
     */
    public readonly instanceProfileIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The property describes a name for the instance profile.
     */
    public readonly instanceProfileName!: pulumi.Output<string | undefined>;
    /**
     * The property describes kms key arn for the instance profile.
     */
    public readonly kmsKeyArn!: pulumi.Output<string | undefined>;
    /**
     * The property describes a network type for the instance profile.
     */
    public readonly networkType!: pulumi.Output<enums.dms.InstanceProfileNetworkType | undefined>;
    /**
     * The property describes the publicly accessible of the instance profile
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * The property describes a subnet group identifier for the instance profile.
     */
    public readonly subnetGroupIdentifier!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The property describes vps security groups for the instance profile.
     */
    public readonly vpcSecurityGroups!: pulumi.Output<string[] | undefined>;

    /**
     * Create a InstanceProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceProfileIdentifier"] = args ? args.instanceProfileIdentifier : undefined;
            resourceInputs["instanceProfileName"] = args ? args.instanceProfileName : undefined;
            resourceInputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            resourceInputs["subnetGroupIdentifier"] = args ? args.subnetGroupIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcSecurityGroups"] = args ? args.vpcSecurityGroups : undefined;
            resourceInputs["instanceProfileArn"] = undefined /*out*/;
            resourceInputs["instanceProfileCreationTime"] = undefined /*out*/;
        } else {
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["instanceProfileArn"] = undefined /*out*/;
            resourceInputs["instanceProfileCreationTime"] = undefined /*out*/;
            resourceInputs["instanceProfileIdentifier"] = undefined /*out*/;
            resourceInputs["instanceProfileName"] = undefined /*out*/;
            resourceInputs["kmsKeyArn"] = undefined /*out*/;
            resourceInputs["networkType"] = undefined /*out*/;
            resourceInputs["publiclyAccessible"] = undefined /*out*/;
            resourceInputs["subnetGroupIdentifier"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroups"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceProfile resource.
 */
export interface InstanceProfileArgs {
    /**
     * The property describes an availability zone of the instance profile.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The optional description of the instance profile.
     */
    description?: pulumi.Input<string>;
    /**
     * The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn
     */
    instanceProfileIdentifier?: pulumi.Input<string>;
    /**
     * The property describes a name for the instance profile.
     */
    instanceProfileName?: pulumi.Input<string>;
    /**
     * The property describes kms key arn for the instance profile.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * The property describes a network type for the instance profile.
     */
    networkType?: pulumi.Input<enums.dms.InstanceProfileNetworkType>;
    /**
     * The property describes the publicly accessible of the instance profile
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * The property describes a subnet group identifier for the instance profile.
     */
    subnetGroupIdentifier?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The property describes vps security groups for the instance profile.
     */
    vpcSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
}
