// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion
 *
 * @deprecated ApplicationVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ApplicationVersion extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationVersion {
        pulumi.log.warn("ApplicationVersion is deprecated: ApplicationVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ApplicationVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticbeanstalk:ApplicationVersion';

    /**
     * Returns true if the given object is an instance of ApplicationVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationVersion.__pulumiType;
    }

    public readonly applicationName!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly sourceBundle!: pulumi.Output<outputs.elasticbeanstalk.ApplicationVersionSourceBundle>;

    /**
     * Create a ApplicationVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ApplicationVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ApplicationVersionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ApplicationVersion is deprecated: ApplicationVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationName'");
            }
            if ((!args || args.sourceBundle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceBundle'");
            }
            resourceInputs["applicationName"] = args ? args.applicationName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["sourceBundle"] = args ? args.sourceBundle : undefined;
        } else {
            resourceInputs["applicationName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["sourceBundle"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationVersion resource.
 */
export interface ApplicationVersionArgs {
    applicationName: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    sourceBundle: pulumi.Input<inputs.elasticbeanstalk.ApplicationVersionSourceBundleArgs>;
}
