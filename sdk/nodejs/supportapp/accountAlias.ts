// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.
 */
export class AccountAlias extends pulumi.CustomResource {
    /**
     * Get an existing AccountAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountAlias {
        return new AccountAlias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:supportapp:AccountAlias';

    /**
     * Returns true if the given object is an instance of AccountAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountAlias.__pulumiType;
    }

    /**
     * An account alias associated with a customer's account.
     */
    public readonly accountAlias!: pulumi.Output<string>;
    /**
     * Unique identifier representing an alias tied to an account
     */
    public /*out*/ readonly accountAliasResourceId!: pulumi.Output<string>;

    /**
     * Create a AccountAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountAliasArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountAlias'");
            }
            resourceInputs["accountAlias"] = args ? args.accountAlias : undefined;
            resourceInputs["accountAliasResourceId"] = undefined /*out*/;
        } else {
            resourceInputs["accountAlias"] = undefined /*out*/;
            resourceInputs["accountAliasResourceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountAlias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountAlias resource.
 */
export interface AccountAliasArgs {
    /**
     * An account alias associated with a customer's account.
     */
    accountAlias: pulumi.Input<string>;
}
