// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::ResourceDefinitionVersion
 *
 * @deprecated ResourceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ResourceDefinitionVersion extends pulumi.CustomResource {
    /**
     * Get an existing ResourceDefinitionVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourceDefinitionVersion {
        pulumi.log.warn("ResourceDefinitionVersion is deprecated: ResourceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ResourceDefinitionVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrass:ResourceDefinitionVersion';

    /**
     * Returns true if the given object is an instance of ResourceDefinitionVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceDefinitionVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceDefinitionVersion.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly resourceDefinitionId!: pulumi.Output<string>;
    public readonly resources!: pulumi.Output<outputs.greengrass.ResourceDefinitionVersionResourceInstance[]>;

    /**
     * Create a ResourceDefinitionVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ResourceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ResourceDefinitionVersionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ResourceDefinitionVersion is deprecated: ResourceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceDefinitionId'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            resourceInputs["resourceDefinitionId"] = args ? args.resourceDefinitionId : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["resourceDefinitionId"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["resourceDefinitionId", "resources[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ResourceDefinitionVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourceDefinitionVersion resource.
 */
export interface ResourceDefinitionVersionArgs {
    resourceDefinitionId: pulumi.Input<string>;
    resources: pulumi.Input<pulumi.Input<inputs.greengrass.ResourceDefinitionVersionResourceInstanceArgs>[]>;
}
