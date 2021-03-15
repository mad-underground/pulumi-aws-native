// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html
 */
export class CompositeAlarm extends pulumi.CustomResource {
    /**
     * Get an existing CompositeAlarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CompositeAlarm {
        return new CompositeAlarm(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:CloudWatch:CompositeAlarm';

    /**
     * Returns true if the given object is an instance of CompositeAlarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CompositeAlarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CompositeAlarm.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionsenabled
     */
    public readonly actionsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmactions
     */
    public readonly alarmActions!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmdescription
     */
    public readonly alarmDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmname
     */
    public readonly alarmName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmrule
     */
    public readonly alarmRule!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-insufficientdataactions
     */
    public readonly insufficientDataActions!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-okactions
     */
    public readonly oKActions!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CompositeAlarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CompositeAlarmArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.alarmName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'alarmName'");
            }
            if ((!args || args.alarmRule === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'alarmRule'");
            }
            inputs["actionsEnabled"] = args ? args.actionsEnabled : undefined;
            inputs["alarmActions"] = args ? args.alarmActions : undefined;
            inputs["alarmDescription"] = args ? args.alarmDescription : undefined;
            inputs["alarmName"] = args ? args.alarmName : undefined;
            inputs["alarmRule"] = args ? args.alarmRule : undefined;
            inputs["insufficientDataActions"] = args ? args.insufficientDataActions : undefined;
            inputs["oKActions"] = args ? args.oKActions : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["actionsEnabled"] = undefined /*out*/;
            inputs["alarmActions"] = undefined /*out*/;
            inputs["alarmDescription"] = undefined /*out*/;
            inputs["alarmName"] = undefined /*out*/;
            inputs["alarmRule"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["insufficientDataActions"] = undefined /*out*/;
            inputs["oKActions"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CompositeAlarm.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CompositeAlarm resource.
 */
export interface CompositeAlarmArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-actionsenabled
     */
    readonly actionsEnabled?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmactions
     */
    readonly alarmActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmdescription
     */
    readonly alarmDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmname
     */
    readonly alarmName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-alarmrule
     */
    readonly alarmRule: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-insufficientdataactions
     */
    readonly insufficientDataActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html#cfn-cloudwatch-compositealarm-okactions
     */
    readonly oKActions?: pulumi.Input<pulumi.Input<string>[]>;
}
