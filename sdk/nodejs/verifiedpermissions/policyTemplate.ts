// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::VerifiedPermissions::PolicyTemplate Resource Type
 */
export class PolicyTemplate extends pulumi.CustomResource {
    /**
     * Get an existing PolicyTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyTemplate {
        return new PolicyTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:verifiedpermissions:PolicyTemplate';

    /**
     * Returns true if the given object is an instance of PolicyTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyTemplate.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly policyStoreId!: pulumi.Output<string>;
    public /*out*/ readonly policyTemplateId!: pulumi.Output<string>;
    public readonly statement!: pulumi.Output<string>;

    /**
     * Create a PolicyTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policyStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyStoreId'");
            }
            if ((!args || args.statement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statement'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["policyStoreId"] = args ? args.policyStoreId : undefined;
            resourceInputs["statement"] = args ? args.statement : undefined;
            resourceInputs["policyTemplateId"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["policyStoreId"] = undefined /*out*/;
            resourceInputs["policyTemplateId"] = undefined /*out*/;
            resourceInputs["statement"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["policyStoreId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PolicyTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyTemplate resource.
 */
export interface PolicyTemplateArgs {
    description?: pulumi.Input<string>;
    policyStoreId: pulumi.Input<string>;
    statement: pulumi.Input<string>;
}
