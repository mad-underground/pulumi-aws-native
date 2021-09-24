// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoTThingsGraph::FlowTemplate
 *
 * @deprecated FlowTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class FlowTemplate extends pulumi.CustomResource {
    /**
     * Get an existing FlowTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowTemplate {
        pulumi.log.warn("FlowTemplate is deprecated: FlowTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new FlowTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotthingsgraph:FlowTemplate';

    /**
     * Returns true if the given object is an instance of FlowTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowTemplate.__pulumiType;
    }

    public readonly compatibleNamespaceVersion!: pulumi.Output<number | undefined>;
    public readonly definition!: pulumi.Output<outputs.iotthingsgraph.FlowTemplateDefinitionDocument>;

    /**
     * Create a FlowTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated FlowTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: FlowTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("FlowTemplate is deprecated: FlowTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            inputs["compatibleNamespaceVersion"] = args ? args.compatibleNamespaceVersion : undefined;
            inputs["definition"] = args ? args.definition : undefined;
        } else {
            inputs["compatibleNamespaceVersion"] = undefined /*out*/;
            inputs["definition"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FlowTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FlowTemplate resource.
 */
export interface FlowTemplateArgs {
    compatibleNamespaceVersion?: pulumi.Input<number>;
    definition: pulumi.Input<inputs.iotthingsgraph.FlowTemplateDefinitionDocumentArgs>;
}