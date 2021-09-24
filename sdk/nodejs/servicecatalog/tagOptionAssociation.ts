// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::TagOptionAssociation
 *
 * @deprecated TagOptionAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TagOptionAssociation extends pulumi.CustomResource {
    /**
     * Get an existing TagOptionAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagOptionAssociation {
        pulumi.log.warn("TagOptionAssociation is deprecated: TagOptionAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TagOptionAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicecatalog:TagOptionAssociation';

    /**
     * Returns true if the given object is an instance of TagOptionAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagOptionAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagOptionAssociation.__pulumiType;
    }

    public readonly resourceId!: pulumi.Output<string>;
    public readonly tagOptionId!: pulumi.Output<string>;

    /**
     * Create a TagOptionAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TagOptionAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: TagOptionAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TagOptionAssociation is deprecated: TagOptionAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.tagOptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagOptionId'");
            }
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["tagOptionId"] = args ? args.tagOptionId : undefined;
        } else {
            inputs["resourceId"] = undefined /*out*/;
            inputs["tagOptionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagOptionAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagOptionAssociation resource.
 */
export interface TagOptionAssociationArgs {
    resourceId: pulumi.Input<string>;
    tagOptionId: pulumi.Input<string>;
}