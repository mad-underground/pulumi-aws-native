// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::FirewallDomainList.
 */
export class FirewallDomainList extends pulumi.CustomResource {
    /**
     * Get an existing FirewallDomainList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FirewallDomainList {
        return new FirewallDomainList(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53resolver:FirewallDomainList';

    /**
     * Returns true if the given object is an instance of FirewallDomainList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallDomainList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallDomainList.__pulumiType;
    }

    /**
     * Arn
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ResourceId
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The id of the creator request.
     */
    public /*out*/ readonly creatorRequestId!: pulumi.Output<string>;
    /**
     * Count
     */
    public /*out*/ readonly domainCount!: pulumi.Output<number>;
    /**
     * S3 URL to import domains from.
     */
    public readonly domainFileUrl!: pulumi.Output<string | undefined>;
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * ServicePrincipal
     */
    public /*out*/ readonly managedOwnerName!: pulumi.Output<string>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly modificationTime!: pulumi.Output<string>;
    /**
     * FirewallDomainListName
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.route53resolver.FirewallDomainListStatus>;
    /**
     * FirewallDomainListAssociationStatus
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a FirewallDomainList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallDomainListArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["domainFileUrl"] = args ? args.domainFileUrl : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["domainCount"] = undefined /*out*/;
            resourceInputs["managedOwnerName"] = undefined /*out*/;
            resourceInputs["modificationTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["domainCount"] = undefined /*out*/;
            resourceInputs["domainFileUrl"] = undefined /*out*/;
            resourceInputs["domains"] = undefined /*out*/;
            resourceInputs["managedOwnerName"] = undefined /*out*/;
            resourceInputs["modificationTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FirewallDomainList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FirewallDomainList resource.
 */
export interface FirewallDomainListArgs {
    /**
     * S3 URL to import domains from.
     */
    domainFileUrl?: pulumi.Input<string>;
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * FirewallDomainListName
     */
    name?: pulumi.Input<string>;
    /**
     * Tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
