// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type
 */
export class IpamResourceDiscoveryAssociation extends pulumi.CustomResource {
    /**
     * Get an existing IpamResourceDiscoveryAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IpamResourceDiscoveryAssociation {
        return new IpamResourceDiscoveryAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:IpamResourceDiscoveryAssociation';

    /**
     * Returns true if the given object is an instance of IpamResourceDiscoveryAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpamResourceDiscoveryAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpamResourceDiscoveryAssociation.__pulumiType;
    }

    /**
     * Arn of the IPAM.
     */
    public /*out*/ readonly ipamArn!: pulumi.Output<string>;
    /**
     * The Id of the IPAM this Resource Discovery is associated to.
     */
    public readonly ipamId!: pulumi.Output<string>;
    /**
     * The home region of the IPAM.
     */
    public /*out*/ readonly ipamRegion!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the resource discovery association is a part of.
     */
    public /*out*/ readonly ipamResourceDiscoveryAssociationArn!: pulumi.Output<string>;
    /**
     * Id of the IPAM Resource Discovery Association.
     */
    public /*out*/ readonly ipamResourceDiscoveryAssociationId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
     */
    public readonly ipamResourceDiscoveryId!: pulumi.Output<string>;
    /**
     * If the Resource Discovery Association exists due as part of CreateIpam.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * The AWS Account ID for the account where the shared IPAM exists.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The status of the resource discovery.
     */
    public /*out*/ readonly resourceDiscoveryStatus!: pulumi.Output<string>;
    /**
     * The operational state of the Resource Discovery Association. Related to Create/Delete activities.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a IpamResourceDiscoveryAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpamResourceDiscoveryAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ipamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipamId'");
            }
            if ((!args || args.ipamResourceDiscoveryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipamResourceDiscoveryId'");
            }
            resourceInputs["ipamId"] = args ? args.ipamId : undefined;
            resourceInputs["ipamResourceDiscoveryId"] = args ? args.ipamResourceDiscoveryId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ipamArn"] = undefined /*out*/;
            resourceInputs["ipamRegion"] = undefined /*out*/;
            resourceInputs["ipamResourceDiscoveryAssociationArn"] = undefined /*out*/;
            resourceInputs["ipamResourceDiscoveryAssociationId"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["resourceDiscoveryStatus"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["ipamArn"] = undefined /*out*/;
            resourceInputs["ipamId"] = undefined /*out*/;
            resourceInputs["ipamRegion"] = undefined /*out*/;
            resourceInputs["ipamResourceDiscoveryAssociationArn"] = undefined /*out*/;
            resourceInputs["ipamResourceDiscoveryAssociationId"] = undefined /*out*/;
            resourceInputs["ipamResourceDiscoveryId"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["resourceDiscoveryStatus"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["ipamId", "ipamResourceDiscoveryId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(IpamResourceDiscoveryAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IpamResourceDiscoveryAssociation resource.
 */
export interface IpamResourceDiscoveryAssociationArgs {
    /**
     * The Id of the IPAM this Resource Discovery is associated to.
     */
    ipamId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
     */
    ipamResourceDiscoveryId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
