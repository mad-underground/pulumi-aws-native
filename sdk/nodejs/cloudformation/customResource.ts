// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFormation::CustomResource
 *
 * @deprecated CustomResource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class CustomResource extends pulumi.CustomResource {
    /**
     * Get an existing CustomResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomResource {
        pulumi.log.warn("CustomResource is deprecated: CustomResource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new CustomResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:CustomResource';

    /**
     * Returns true if the given object is an instance of CustomResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomResource.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly serviceToken!: pulumi.Output<string>;

    /**
     * Create a CustomResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated CustomResource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: CustomResourceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("CustomResource is deprecated: CustomResource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.serviceToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceToken'");
            }
            resourceInputs["serviceToken"] = args ? args.serviceToken : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["serviceToken"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["serviceToken"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(CustomResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomResource resource.
 */
export interface CustomResourceArgs {
    serviceToken: pulumi.Input<string>;
}
