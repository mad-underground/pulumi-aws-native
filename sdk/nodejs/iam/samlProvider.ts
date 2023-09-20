// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::SAMLProvider
 */
export class SamlProvider extends pulumi.CustomResource {
    /**
     * Get an existing SamlProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SamlProvider {
        return new SamlProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:SamlProvider';

    /**
     * Returns true if the given object is an instance of SamlProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlProvider.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the SAML provider
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly samlMetadataDocument!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.iam.SamlProviderTag[] | undefined>;

    /**
     * Create a SamlProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.samlMetadataDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlMetadataDocument'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["samlMetadataDocument"] = args ? args.samlMetadataDocument : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["samlMetadataDocument"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SamlProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SamlProvider resource.
 */
export interface SamlProviderArgs {
    name?: pulumi.Input<string>;
    samlMetadataDocument: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.iam.SamlProviderTagArgs>[]>;
}