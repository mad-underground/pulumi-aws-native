// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::MonitoringSchedule
 */
export class MonitoringSchedule extends pulumi.CustomResource {
    /**
     * Get an existing MonitoringSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MonitoringSchedule {
        return new MonitoringSchedule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:MonitoringSchedule';

    /**
     * Returns true if the given object is an instance of MonitoringSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoringSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoringSchedule.__pulumiType;
    }

    /**
     * The time at which the schedule was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public readonly endpointName!: pulumi.Output<string | undefined>;
    /**
     * Contains the reason a monitoring job failed, if it failed.
     */
    public readonly failureReason!: pulumi.Output<string | undefined>;
    /**
     * A timestamp that indicates the last time the monitoring job was modified.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Describes metadata on the last execution to run, if there was one.
     */
    public readonly lastMonitoringExecutionSummary!: pulumi.Output<outputs.sagemaker.MonitoringScheduleMonitoringExecutionSummary | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the monitoring schedule.
     */
    public /*out*/ readonly monitoringScheduleArn!: pulumi.Output<string>;
    public readonly monitoringScheduleConfig!: pulumi.Output<outputs.sagemaker.MonitoringScheduleConfig>;
    public readonly monitoringScheduleName!: pulumi.Output<string>;
    /**
     * The status of a schedule job.
     */
    public readonly monitoringScheduleStatus!: pulumi.Output<enums.sagemaker.MonitoringScheduleStatus | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.MonitoringScheduleTag[] | undefined>;

    /**
     * Create a MonitoringSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitoringScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.monitoringScheduleConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitoringScheduleConfig'");
            }
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["failureReason"] = args ? args.failureReason : undefined;
            resourceInputs["lastMonitoringExecutionSummary"] = args ? args.lastMonitoringExecutionSummary : undefined;
            resourceInputs["monitoringScheduleConfig"] = args ? args.monitoringScheduleConfig : undefined;
            resourceInputs["monitoringScheduleName"] = args ? args.monitoringScheduleName : undefined;
            resourceInputs["monitoringScheduleStatus"] = args ? args.monitoringScheduleStatus : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["monitoringScheduleArn"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["endpointName"] = undefined /*out*/;
            resourceInputs["failureReason"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["lastMonitoringExecutionSummary"] = undefined /*out*/;
            resourceInputs["monitoringScheduleArn"] = undefined /*out*/;
            resourceInputs["monitoringScheduleConfig"] = undefined /*out*/;
            resourceInputs["monitoringScheduleName"] = undefined /*out*/;
            resourceInputs["monitoringScheduleStatus"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MonitoringSchedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MonitoringSchedule resource.
 */
export interface MonitoringScheduleArgs {
    endpointName?: pulumi.Input<string>;
    /**
     * Contains the reason a monitoring job failed, if it failed.
     */
    failureReason?: pulumi.Input<string>;
    /**
     * Describes metadata on the last execution to run, if there was one.
     */
    lastMonitoringExecutionSummary?: pulumi.Input<inputs.sagemaker.MonitoringScheduleMonitoringExecutionSummaryArgs>;
    monitoringScheduleConfig: pulumi.Input<inputs.sagemaker.MonitoringScheduleConfigArgs>;
    monitoringScheduleName?: pulumi.Input<string>;
    /**
     * The status of a schedule job.
     */
    monitoringScheduleStatus?: pulumi.Input<enums.sagemaker.MonitoringScheduleStatus>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.MonitoringScheduleTagArgs>[]>;
}
