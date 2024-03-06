// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::Crawler
 *
 * @deprecated Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Crawler extends pulumi.CustomResource {
    /**
     * Get an existing Crawler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Crawler {
        pulumi.log.warn("Crawler is deprecated: Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Crawler(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:glue:Crawler';

    /**
     * Returns true if the given object is an instance of Crawler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Crawler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Crawler.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly classifiers!: pulumi.Output<string[] | undefined>;
    public readonly configuration!: pulumi.Output<string | undefined>;
    public readonly crawlerSecurityConfiguration!: pulumi.Output<string | undefined>;
    public readonly databaseName!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly recrawlPolicy!: pulumi.Output<outputs.glue.CrawlerRecrawlPolicy | undefined>;
    public readonly role!: pulumi.Output<string>;
    public readonly schedule!: pulumi.Output<outputs.glue.CrawlerSchedule | undefined>;
    public readonly schemaChangePolicy!: pulumi.Output<outputs.glue.CrawlerSchemaChangePolicy | undefined>;
    public readonly tablePrefix!: pulumi.Output<string | undefined>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Crawler` for more information about the expected schema for this property.
     */
    public readonly tags!: pulumi.Output<any | undefined>;
    public readonly targets!: pulumi.Output<outputs.glue.CrawlerTargets>;

    /**
     * Create a Crawler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: CrawlerArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Crawler is deprecated: Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            resourceInputs["classifiers"] = args ? args.classifiers : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["crawlerSecurityConfiguration"] = args ? args.crawlerSecurityConfiguration : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recrawlPolicy"] = args ? args.recrawlPolicy : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["schemaChangePolicy"] = args ? args.schemaChangePolicy : undefined;
            resourceInputs["tablePrefix"] = args ? args.tablePrefix : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["classifiers"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["crawlerSecurityConfiguration"] = undefined /*out*/;
            resourceInputs["databaseName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["recrawlPolicy"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["schemaChangePolicy"] = undefined /*out*/;
            resourceInputs["tablePrefix"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Crawler.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Crawler resource.
 */
export interface CrawlerArgs {
    classifiers?: pulumi.Input<pulumi.Input<string>[]>;
    configuration?: pulumi.Input<string>;
    crawlerSecurityConfiguration?: pulumi.Input<string>;
    databaseName?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    recrawlPolicy?: pulumi.Input<inputs.glue.CrawlerRecrawlPolicyArgs>;
    role: pulumi.Input<string>;
    schedule?: pulumi.Input<inputs.glue.CrawlerScheduleArgs>;
    schemaChangePolicy?: pulumi.Input<inputs.glue.CrawlerSchemaChangePolicyArgs>;
    tablePrefix?: pulumi.Input<string>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Crawler` for more information about the expected schema for this property.
     */
    tags?: any;
    targets: pulumi.Input<inputs.glue.CrawlerTargetsArgs>;
}
