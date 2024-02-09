// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppConfig::Environment
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appconfig:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * The application ID.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * A description of the environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The environment ID.
     */
    public /*out*/ readonly environmentId!: pulumi.Output<string>;
    /**
     * Amazon CloudWatch alarms to monitor during the deployment process.
     */
    public readonly monitors!: pulumi.Output<outputs.appconfig.EnvironmentMonitor[] | undefined>;
    /**
     * A name for the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
     */
    public readonly tags!: pulumi.Output<outputs.appconfig.EnvironmentTag[] | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["monitors"] = args ? args.monitors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["environmentId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["monitors"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * The application ID.
     */
    applicationId: pulumi.Input<string>;
    /**
     * A description of the environment.
     */
    description?: pulumi.Input<string>;
    /**
     * Amazon CloudWatch alarms to monitor during the deployment process.
     */
    monitors?: pulumi.Input<pulumi.Input<inputs.appconfig.EnvironmentMonitorArgs>[]>;
    /**
     * A name for the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.appconfig.EnvironmentTagArgs>[]>;
}
