// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
 *
 * @deprecated AnomalyMonitor is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class AnomalyMonitor extends pulumi.CustomResource {
    /**
     * Get an existing AnomalyMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AnomalyMonitor {
        pulumi.log.warn("AnomalyMonitor is deprecated: AnomalyMonitor is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new AnomalyMonitor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ce:AnomalyMonitor';

    /**
     * Returns true if the given object is an instance of AnomalyMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnomalyMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnomalyMonitor.__pulumiType;
    }

    /**
     * The date when the monitor was created. 
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The value for evaluated dimensions.
     */
    public /*out*/ readonly dimensionalValueCount!: pulumi.Output<number>;
    /**
     * The date when the monitor last evaluated for anomalies.
     */
    public /*out*/ readonly lastEvaluatedDate!: pulumi.Output<string>;
    /**
     * The date when the monitor was last updated.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    public /*out*/ readonly monitorArn!: pulumi.Output<string>;
    /**
     * The dimensions to evaluate
     */
    public readonly monitorDimension!: pulumi.Output<enums.ce.AnomalyMonitorMonitorDimension | undefined>;
    /**
     * The name of the monitor.
     */
    public readonly monitorName!: pulumi.Output<string>;
    public readonly monitorSpecification!: pulumi.Output<string | undefined>;
    public readonly monitorType!: pulumi.Output<enums.ce.AnomalyMonitorMonitorType>;
    /**
     * Tags to assign to monitor.
     */
    public readonly resourceTags!: pulumi.Output<outputs.ce.AnomalyMonitorResourceTag[] | undefined>;

    /**
     * Create a AnomalyMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated AnomalyMonitor is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: AnomalyMonitorArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("AnomalyMonitor is deprecated: AnomalyMonitor is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.monitorName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorName'");
            }
            if ((!args || args.monitorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorType'");
            }
            resourceInputs["monitorDimension"] = args ? args.monitorDimension : undefined;
            resourceInputs["monitorName"] = args ? args.monitorName : undefined;
            resourceInputs["monitorSpecification"] = args ? args.monitorSpecification : undefined;
            resourceInputs["monitorType"] = args ? args.monitorType : undefined;
            resourceInputs["resourceTags"] = args ? args.resourceTags : undefined;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["dimensionalValueCount"] = undefined /*out*/;
            resourceInputs["lastEvaluatedDate"] = undefined /*out*/;
            resourceInputs["lastUpdatedDate"] = undefined /*out*/;
            resourceInputs["monitorArn"] = undefined /*out*/;
        } else {
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["dimensionalValueCount"] = undefined /*out*/;
            resourceInputs["lastEvaluatedDate"] = undefined /*out*/;
            resourceInputs["lastUpdatedDate"] = undefined /*out*/;
            resourceInputs["monitorArn"] = undefined /*out*/;
            resourceInputs["monitorDimension"] = undefined /*out*/;
            resourceInputs["monitorName"] = undefined /*out*/;
            resourceInputs["monitorSpecification"] = undefined /*out*/;
            resourceInputs["monitorType"] = undefined /*out*/;
            resourceInputs["resourceTags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnomalyMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AnomalyMonitor resource.
 */
export interface AnomalyMonitorArgs {
    /**
     * The dimensions to evaluate
     */
    monitorDimension?: pulumi.Input<enums.ce.AnomalyMonitorMonitorDimension>;
    /**
     * The name of the monitor.
     */
    monitorName: pulumi.Input<string>;
    monitorSpecification?: pulumi.Input<string>;
    monitorType: pulumi.Input<enums.ce.AnomalyMonitorMonitorType>;
    /**
     * Tags to assign to monitor.
     */
    resourceTags?: pulumi.Input<pulumi.Input<inputs.ce.AnomalyMonitorResourceTagArgs>[]>;
}
