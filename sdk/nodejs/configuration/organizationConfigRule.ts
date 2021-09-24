// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Config::OrganizationConfigRule
 *
 * @deprecated OrganizationConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class OrganizationConfigRule extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationConfigRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationConfigRule {
        pulumi.log.warn("OrganizationConfigRule is deprecated: OrganizationConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new OrganizationConfigRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:configuration:OrganizationConfigRule';

    /**
     * Returns true if the given object is an instance of OrganizationConfigRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationConfigRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationConfigRule.__pulumiType;
    }

    public readonly excludedAccounts!: pulumi.Output<string[] | undefined>;
    public readonly organizationConfigRuleName!: pulumi.Output<string>;
    public readonly organizationCustomRuleMetadata!: pulumi.Output<outputs.configuration.OrganizationConfigRuleOrganizationCustomRuleMetadata | undefined>;
    public readonly organizationManagedRuleMetadata!: pulumi.Output<outputs.configuration.OrganizationConfigRuleOrganizationManagedRuleMetadata | undefined>;

    /**
     * Create a OrganizationConfigRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated OrganizationConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: OrganizationConfigRuleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("OrganizationConfigRule is deprecated: OrganizationConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationConfigRuleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationConfigRuleName'");
            }
            inputs["excludedAccounts"] = args ? args.excludedAccounts : undefined;
            inputs["organizationConfigRuleName"] = args ? args.organizationConfigRuleName : undefined;
            inputs["organizationCustomRuleMetadata"] = args ? args.organizationCustomRuleMetadata : undefined;
            inputs["organizationManagedRuleMetadata"] = args ? args.organizationManagedRuleMetadata : undefined;
        } else {
            inputs["excludedAccounts"] = undefined /*out*/;
            inputs["organizationConfigRuleName"] = undefined /*out*/;
            inputs["organizationCustomRuleMetadata"] = undefined /*out*/;
            inputs["organizationManagedRuleMetadata"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationConfigRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationConfigRule resource.
 */
export interface OrganizationConfigRuleArgs {
    excludedAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    organizationConfigRuleName: pulumi.Input<string>;
    organizationCustomRuleMetadata?: pulumi.Input<inputs.configuration.OrganizationConfigRuleOrganizationCustomRuleMetadataArgs>;
    organizationManagedRuleMetadata?: pulumi.Input<inputs.configuration.OrganizationConfigRuleOrganizationManagedRuleMetadataArgs>;
}