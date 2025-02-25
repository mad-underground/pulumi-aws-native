// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A domain defined for 3rd party data source in Profile Service
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:customerprofiles:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * The time of this integration got created
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The URL of the SQS dead letter queue
     */
    public readonly deadLetterQueueUrl!: pulumi.Output<string | undefined>;
    /**
     * The default encryption key
     */
    public readonly defaultEncryptionKey!: pulumi.Output<string | undefined>;
    /**
     * The default number of days until the data within the domain expires.
     */
    public readonly defaultExpirationDays!: pulumi.Output<number>;
    /**
     * The unique name of the domain.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The time of this integration got last updated at
     */
    public /*out*/ readonly lastUpdatedAt!: pulumi.Output<string>;
    public readonly matching!: pulumi.Output<outputs.customerprofiles.DomainMatching | undefined>;
    public readonly ruleBasedMatching!: pulumi.Output<outputs.customerprofiles.DomainRuleBasedMatching | undefined>;
    public /*out*/ readonly stats!: pulumi.Output<outputs.customerprofiles.DomainStats>;
    /**
     * The tags (keys and values) associated with the domain
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.defaultExpirationDays === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultExpirationDays'");
            }
            resourceInputs["deadLetterQueueUrl"] = args ? args.deadLetterQueueUrl : undefined;
            resourceInputs["defaultEncryptionKey"] = args ? args.defaultEncryptionKey : undefined;
            resourceInputs["defaultExpirationDays"] = args ? args.defaultExpirationDays : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["matching"] = args ? args.matching : undefined;
            resourceInputs["ruleBasedMatching"] = args ? args.ruleBasedMatching : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastUpdatedAt"] = undefined /*out*/;
            resourceInputs["stats"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deadLetterQueueUrl"] = undefined /*out*/;
            resourceInputs["defaultEncryptionKey"] = undefined /*out*/;
            resourceInputs["defaultExpirationDays"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["lastUpdatedAt"] = undefined /*out*/;
            resourceInputs["matching"] = undefined /*out*/;
            resourceInputs["ruleBasedMatching"] = undefined /*out*/;
            resourceInputs["stats"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["domainName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The URL of the SQS dead letter queue
     */
    deadLetterQueueUrl?: pulumi.Input<string>;
    /**
     * The default encryption key
     */
    defaultEncryptionKey?: pulumi.Input<string>;
    /**
     * The default number of days until the data within the domain expires.
     */
    defaultExpirationDays: pulumi.Input<number>;
    /**
     * The unique name of the domain.
     */
    domainName?: pulumi.Input<string>;
    matching?: pulumi.Input<inputs.customerprofiles.DomainMatchingArgs>;
    ruleBasedMatching?: pulumi.Input<inputs.customerprofiles.DomainRuleBasedMatchingArgs>;
    /**
     * The tags (keys and values) associated with the domain
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
