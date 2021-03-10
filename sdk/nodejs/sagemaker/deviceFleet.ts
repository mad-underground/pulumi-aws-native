// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html
 */
export class DeviceFleet extends pulumi.CustomResource {
    /**
     * Get an existing DeviceFleet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeviceFleet {
        return new DeviceFleet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:SageMaker:DeviceFleet';

    /**
     * Returns true if the given object is an instance of DeviceFleet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceFleet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceFleet.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-description
     */
    public readonly Description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly DeviceFleetName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-outputconfig
     */
    public readonly OutputConfig!: pulumi.Output<outputs.SageMaker.DeviceFleetEdgeOutputConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-rolearn
     */
    public readonly RoleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-tags
     */
    public readonly Tags!: pulumi.Output<outputs.Tag | undefined>;

    /**
     * Create a DeviceFleet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceFleetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.OutputConfig === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'OutputConfig'");
            }
            if ((!args || args.RoleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'RoleArn'");
            }
            inputs["Description"] = args ? args.Description : undefined;
            inputs["OutputConfig"] = args ? args.OutputConfig : undefined;
            inputs["RoleArn"] = args ? args.RoleArn : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["DeviceFleetName"] = undefined /*out*/;
        } else {
            inputs["Description"] = undefined /*out*/;
            inputs["DeviceFleetName"] = undefined /*out*/;
            inputs["OutputConfig"] = undefined /*out*/;
            inputs["RoleArn"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeviceFleet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeviceFleet resource.
 */
export interface DeviceFleetArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-description
     */
    readonly Description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-outputconfig
     */
    readonly OutputConfig: pulumi.Input<inputs.SageMaker.DeviceFleetEdgeOutputConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-rolearn
     */
    readonly RoleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-devicefleet.html#cfn-sagemaker-devicefleet-tags
     */
    readonly Tags?: pulumi.Input<inputs.Tag>;
}
