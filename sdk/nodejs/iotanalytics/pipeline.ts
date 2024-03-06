// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoTAnalytics::Pipeline
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotanalytics:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly pipelineActivities!: pulumi.Output<outputs.iotanalytics.PipelineActivity[]>;
    public readonly pipelineName!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.pipelineActivities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineActivities'");
            }
            resourceInputs["pipelineActivities"] = args ? args.pipelineActivities : undefined;
            resourceInputs["pipelineName"] = args ? args.pipelineName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["pipelineActivities"] = undefined /*out*/;
            resourceInputs["pipelineName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["pipelineName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    pipelineActivities: pulumi.Input<pulumi.Input<inputs.iotanalytics.PipelineActivityArgs>[]>;
    pipelineName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
