// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppStream::StackFleetAssociation
 *
 * @deprecated StackFleetAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class StackFleetAssociation extends pulumi.CustomResource {
    /**
     * Get an existing StackFleetAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StackFleetAssociation {
        pulumi.log.warn("StackFleetAssociation is deprecated: StackFleetAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new StackFleetAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appstream:StackFleetAssociation';

    /**
     * Returns true if the given object is an instance of StackFleetAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackFleetAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackFleetAssociation.__pulumiType;
    }

    public readonly fleetName!: pulumi.Output<string>;
    public readonly stackName!: pulumi.Output<string>;

    /**
     * Create a StackFleetAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated StackFleetAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: StackFleetAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("StackFleetAssociation is deprecated: StackFleetAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.fleetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fleetName'");
            }
            if ((!args || args.stackName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackName'");
            }
            inputs["fleetName"] = args ? args.fleetName : undefined;
            inputs["stackName"] = args ? args.stackName : undefined;
        } else {
            inputs["fleetName"] = undefined /*out*/;
            inputs["stackName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StackFleetAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StackFleetAssociation resource.
 */
export interface StackFleetAssociationArgs {
    fleetName: pulumi.Input<string>;
    stackName: pulumi.Input<string>;
}