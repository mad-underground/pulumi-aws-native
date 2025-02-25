// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.
 */
export class LinkAssociation extends pulumi.CustomResource {
    /**
     * Get an existing LinkAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LinkAssociation {
        return new LinkAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:networkmanager:LinkAssociation';

    /**
     * Returns true if the given object is an instance of LinkAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkAssociation.__pulumiType;
    }

    /**
     * The ID of the device
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * The ID of the global network.
     */
    public readonly globalNetworkId!: pulumi.Output<string>;
    /**
     * The ID of the link
     */
    public readonly linkId!: pulumi.Output<string>;

    /**
     * Create a LinkAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            if ((!args || args.globalNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalNetworkId'");
            }
            if ((!args || args.linkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'linkId'");
            }
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["globalNetworkId"] = args ? args.globalNetworkId : undefined;
            resourceInputs["linkId"] = args ? args.linkId : undefined;
        } else {
            resourceInputs["deviceId"] = undefined /*out*/;
            resourceInputs["globalNetworkId"] = undefined /*out*/;
            resourceInputs["linkId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["deviceId", "globalNetworkId", "linkId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(LinkAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LinkAssociation resource.
 */
export interface LinkAssociationArgs {
    /**
     * The ID of the device
     */
    deviceId: pulumi.Input<string>;
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
    /**
     * The ID of the link
     */
    linkId: pulumi.Input<string>;
}
