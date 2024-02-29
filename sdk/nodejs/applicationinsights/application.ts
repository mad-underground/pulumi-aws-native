// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ApplicationInsights::Application
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:applicationinsights:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * The ARN of the ApplicationInsights application.
     */
    public /*out*/ readonly applicationArn!: pulumi.Output<string>;
    /**
     * If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing
     */
    public readonly attachMissingPermission!: pulumi.Output<boolean | undefined>;
    /**
     * If set to true, application will be configured with recommended monitoring configuration.
     */
    public readonly autoConfigurationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The monitoring settings of the components.
     */
    public readonly componentMonitoringSettings!: pulumi.Output<outputs.applicationinsights.ApplicationComponentMonitoringSetting[] | undefined>;
    /**
     * The custom grouped components.
     */
    public readonly customComponents!: pulumi.Output<outputs.applicationinsights.ApplicationCustomComponent[] | undefined>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources.
     */
    public readonly cweMonitorEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The grouping type of the application
     */
    public readonly groupingType!: pulumi.Output<enums.applicationinsights.ApplicationGroupingType | undefined>;
    /**
     * The log pattern sets.
     */
    public readonly logPatternSets!: pulumi.Output<outputs.applicationinsights.ApplicationLogPatternSet[] | undefined>;
    /**
     * When set to true, creates opsItems for any problems detected on an application.
     */
    public readonly opsCenterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The SNS topic provided to Application Insights that is associated to the created opsItem.
     */
    public readonly opsItemSnsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The tags of Application Insights application.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["attachMissingPermission"] = args ? args.attachMissingPermission : undefined;
            resourceInputs["autoConfigurationEnabled"] = args ? args.autoConfigurationEnabled : undefined;
            resourceInputs["componentMonitoringSettings"] = args ? args.componentMonitoringSettings : undefined;
            resourceInputs["customComponents"] = args ? args.customComponents : undefined;
            resourceInputs["cweMonitorEnabled"] = args ? args.cweMonitorEnabled : undefined;
            resourceInputs["groupingType"] = args ? args.groupingType : undefined;
            resourceInputs["logPatternSets"] = args ? args.logPatternSets : undefined;
            resourceInputs["opsCenterEnabled"] = args ? args.opsCenterEnabled : undefined;
            resourceInputs["opsItemSnsTopicArn"] = args ? args.opsItemSnsTopicArn : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["applicationArn"] = undefined /*out*/;
        } else {
            resourceInputs["applicationArn"] = undefined /*out*/;
            resourceInputs["attachMissingPermission"] = undefined /*out*/;
            resourceInputs["autoConfigurationEnabled"] = undefined /*out*/;
            resourceInputs["componentMonitoringSettings"] = undefined /*out*/;
            resourceInputs["customComponents"] = undefined /*out*/;
            resourceInputs["cweMonitorEnabled"] = undefined /*out*/;
            resourceInputs["groupingType"] = undefined /*out*/;
            resourceInputs["logPatternSets"] = undefined /*out*/;
            resourceInputs["opsCenterEnabled"] = undefined /*out*/;
            resourceInputs["opsItemSnsTopicArn"] = undefined /*out*/;
            resourceInputs["resourceGroupName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["groupingType", "resourceGroupName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing
     */
    attachMissingPermission?: pulumi.Input<boolean>;
    /**
     * If set to true, application will be configured with recommended monitoring configuration.
     */
    autoConfigurationEnabled?: pulumi.Input<boolean>;
    /**
     * The monitoring settings of the components.
     */
    componentMonitoringSettings?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationComponentMonitoringSettingArgs>[]>;
    /**
     * The custom grouped components.
     */
    customComponents?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationCustomComponentArgs>[]>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources.
     */
    cweMonitorEnabled?: pulumi.Input<boolean>;
    /**
     * The grouping type of the application
     */
    groupingType?: pulumi.Input<enums.applicationinsights.ApplicationGroupingType>;
    /**
     * The log pattern sets.
     */
    logPatternSets?: pulumi.Input<pulumi.Input<inputs.applicationinsights.ApplicationLogPatternSetArgs>[]>;
    /**
     * When set to true, creates opsItems for any problems detected on an application.
     */
    opsCenterEnabled?: pulumi.Input<boolean>;
    /**
     * The SNS topic provided to Application Insights that is associated to the created opsItem.
     */
    opsItemSnsTopicArn?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of Application Insights application.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
