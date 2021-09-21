// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * AWS::RoboMaker::Robot resource creates an AWS RoboMaker fleet.
 */
export class Robot extends pulumi.CustomResource {
    /**
     * Get an existing Robot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Robot {
        return new Robot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:robomaker:Robot';

    /**
     * Returns true if the given object is an instance of Robot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Robot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Robot.__pulumiType;
    }

    /**
     * The target architecture of the robot.
     */
    public readonly architecture!: pulumi.Output<enums.robomaker.RobotArchitecture>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the fleet.
     */
    public readonly fleet!: pulumi.Output<string | undefined>;
    /**
     * The Greengrass group id.
     */
    public readonly greengrassGroupId!: pulumi.Output<string>;
    /**
     * The name for the robot.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.robomaker.RobotTags | undefined>;

    /**
     * Create a Robot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RobotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.architecture === undefined) && !opts.urn) {
                throw new Error("Missing required property 'architecture'");
            }
            if ((!args || args.greengrassGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'greengrassGroupId'");
            }
            inputs["architecture"] = args ? args.architecture : undefined;
            inputs["fleet"] = args ? args.fleet : undefined;
            inputs["greengrassGroupId"] = args ? args.greengrassGroupId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["architecture"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["fleet"] = undefined /*out*/;
            inputs["greengrassGroupId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Robot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Robot resource.
 */
export interface RobotArgs {
    /**
     * The target architecture of the robot.
     */
    architecture: pulumi.Input<enums.robomaker.RobotArchitecture>;
    /**
     * The Amazon Resource Name (ARN) of the fleet.
     */
    fleet?: pulumi.Input<string>;
    /**
     * The Greengrass group id.
     */
    greengrassGroupId: pulumi.Input<string>;
    /**
     * The name for the robot.
     */
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<inputs.robomaker.RobotTagsArgs>;
}