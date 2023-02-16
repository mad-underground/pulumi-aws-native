// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace
 *
 * @deprecated PublicDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class PublicDnsNamespace extends pulumi.CustomResource {
    /**
     * Get an existing PublicDnsNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublicDnsNamespace {
        pulumi.log.warn("PublicDnsNamespace is deprecated: PublicDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new PublicDnsNamespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicediscovery:PublicDnsNamespace';

    /**
     * Returns true if the given object is an instance of PublicDnsNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicDnsNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicDnsNamespace.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly hostedZoneId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.servicediscovery.PublicDnsNamespaceProperties | undefined>;
    public readonly tags!: pulumi.Output<outputs.servicediscovery.PublicDnsNamespaceTag[] | undefined>;

    /**
     * Create a PublicDnsNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated PublicDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: PublicDnsNamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PublicDnsNamespace is deprecated: PublicDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicDnsNamespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublicDnsNamespace resource.
 */
export interface PublicDnsNamespaceArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    properties?: pulumi.Input<inputs.servicediscovery.PublicDnsNamespacePropertiesArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.servicediscovery.PublicDnsNamespaceTagArgs>[]>;
}
