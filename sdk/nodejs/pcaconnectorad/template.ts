// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a template that defines certificate configurations, both for issuance and client handling
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:pcaconnectorad:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    public readonly connectorArn!: pulumi.Output<string>;
    public readonly definition!: pulumi.Output<outputs.pcaconnectorad.TemplateDefinition0Properties | outputs.pcaconnectorad.TemplateDefinition1Properties | outputs.pcaconnectorad.TemplateDefinition2Properties>;
    public readonly name!: pulumi.Output<string>;
    public readonly reenrollAllCertificateHolders!: pulumi.Output<boolean | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly templateArn!: pulumi.Output<string>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectorArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorArn'");
            }
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            resourceInputs["connectorArn"] = args ? args.connectorArn : undefined;
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["reenrollAllCertificateHolders"] = args ? args.reenrollAllCertificateHolders : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateArn"] = undefined /*out*/;
        } else {
            resourceInputs["connectorArn"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["reenrollAllCertificateHolders"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["templateArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectorArn", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    connectorArn: pulumi.Input<string>;
    definition: pulumi.Input<inputs.pcaconnectorad.TemplateDefinition0PropertiesArgs | inputs.pcaconnectorad.TemplateDefinition1PropertiesArgs | inputs.pcaconnectorad.TemplateDefinition2PropertiesArgs>;
    name?: pulumi.Input<string>;
    reenrollAllCertificateHolders?: pulumi.Input<boolean>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
