// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodeDeploy::DeploymentGroup
 *
 * @deprecated DeploymentGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DeploymentGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeploymentGroup {
        pulumi.log.warn("DeploymentGroup is deprecated: DeploymentGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DeploymentGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:codedeploy:DeploymentGroup';

    /**
     * Returns true if the given object is an instance of DeploymentGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentGroup.__pulumiType;
    }

    public readonly alarmConfiguration!: pulumi.Output<outputs.codedeploy.DeploymentGroupAlarmConfiguration | undefined>;
    public readonly applicationName!: pulumi.Output<string>;
    public readonly autoRollbackConfiguration!: pulumi.Output<outputs.codedeploy.DeploymentGroupAutoRollbackConfiguration | undefined>;
    public readonly autoScalingGroups!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly blueGreenDeploymentConfiguration!: pulumi.Output<outputs.codedeploy.DeploymentGroupBlueGreenDeploymentConfiguration | undefined>;
    public readonly deployment!: pulumi.Output<outputs.codedeploy.DeploymentGroupDeployment | undefined>;
    public readonly deploymentConfigName!: pulumi.Output<string | undefined>;
    public readonly deploymentGroupName!: pulumi.Output<string | undefined>;
    public readonly deploymentStyle!: pulumi.Output<outputs.codedeploy.DeploymentGroupDeploymentStyle | undefined>;
    public readonly ec2TagFilters!: pulumi.Output<outputs.codedeploy.DeploymentGroupEc2TagFilter[] | undefined>;
    public readonly ec2TagSet!: pulumi.Output<outputs.codedeploy.DeploymentGroupEc2TagSet | undefined>;
    public readonly ecsServices!: pulumi.Output<outputs.codedeploy.DeploymentGroupEcsService[] | undefined>;
    public readonly loadBalancerInfo!: pulumi.Output<outputs.codedeploy.DeploymentGroupLoadBalancerInfo | undefined>;
    public readonly onPremisesInstanceTagFilters!: pulumi.Output<outputs.codedeploy.DeploymentGroupTagFilter[] | undefined>;
    public readonly onPremisesTagSet!: pulumi.Output<outputs.codedeploy.DeploymentGroupOnPremisesTagSet | undefined>;
    public readonly outdatedInstancesStrategy!: pulumi.Output<string | undefined>;
    public readonly serviceRoleArn!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly terminationHookEnabled!: pulumi.Output<boolean | undefined>;
    public readonly triggerConfigurations!: pulumi.Output<outputs.codedeploy.DeploymentGroupTriggerConfig[] | undefined>;

    /**
     * Create a DeploymentGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DeploymentGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DeploymentGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DeploymentGroup is deprecated: DeploymentGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationName'");
            }
            if ((!args || args.serviceRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            resourceInputs["alarmConfiguration"] = args ? args.alarmConfiguration : undefined;
            resourceInputs["applicationName"] = args ? args.applicationName : undefined;
            resourceInputs["autoRollbackConfiguration"] = args ? args.autoRollbackConfiguration : undefined;
            resourceInputs["autoScalingGroups"] = args ? args.autoScalingGroups : undefined;
            resourceInputs["blueGreenDeploymentConfiguration"] = args ? args.blueGreenDeploymentConfiguration : undefined;
            resourceInputs["deployment"] = args ? args.deployment : undefined;
            resourceInputs["deploymentConfigName"] = args ? args.deploymentConfigName : undefined;
            resourceInputs["deploymentGroupName"] = args ? args.deploymentGroupName : undefined;
            resourceInputs["deploymentStyle"] = args ? args.deploymentStyle : undefined;
            resourceInputs["ec2TagFilters"] = args ? args.ec2TagFilters : undefined;
            resourceInputs["ec2TagSet"] = args ? args.ec2TagSet : undefined;
            resourceInputs["ecsServices"] = args ? args.ecsServices : undefined;
            resourceInputs["loadBalancerInfo"] = args ? args.loadBalancerInfo : undefined;
            resourceInputs["onPremisesInstanceTagFilters"] = args ? args.onPremisesInstanceTagFilters : undefined;
            resourceInputs["onPremisesTagSet"] = args ? args.onPremisesTagSet : undefined;
            resourceInputs["outdatedInstancesStrategy"] = args ? args.outdatedInstancesStrategy : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["terminationHookEnabled"] = args ? args.terminationHookEnabled : undefined;
            resourceInputs["triggerConfigurations"] = args ? args.triggerConfigurations : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["alarmConfiguration"] = undefined /*out*/;
            resourceInputs["applicationName"] = undefined /*out*/;
            resourceInputs["autoRollbackConfiguration"] = undefined /*out*/;
            resourceInputs["autoScalingGroups"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["blueGreenDeploymentConfiguration"] = undefined /*out*/;
            resourceInputs["deployment"] = undefined /*out*/;
            resourceInputs["deploymentConfigName"] = undefined /*out*/;
            resourceInputs["deploymentGroupName"] = undefined /*out*/;
            resourceInputs["deploymentStyle"] = undefined /*out*/;
            resourceInputs["ec2TagFilters"] = undefined /*out*/;
            resourceInputs["ec2TagSet"] = undefined /*out*/;
            resourceInputs["ecsServices"] = undefined /*out*/;
            resourceInputs["loadBalancerInfo"] = undefined /*out*/;
            resourceInputs["onPremisesInstanceTagFilters"] = undefined /*out*/;
            resourceInputs["onPremisesTagSet"] = undefined /*out*/;
            resourceInputs["outdatedInstancesStrategy"] = undefined /*out*/;
            resourceInputs["serviceRoleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["terminationHookEnabled"] = undefined /*out*/;
            resourceInputs["triggerConfigurations"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationName", "deploymentGroupName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DeploymentGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeploymentGroup resource.
 */
export interface DeploymentGroupArgs {
    alarmConfiguration?: pulumi.Input<inputs.codedeploy.DeploymentGroupAlarmConfigurationArgs>;
    applicationName: pulumi.Input<string>;
    autoRollbackConfiguration?: pulumi.Input<inputs.codedeploy.DeploymentGroupAutoRollbackConfigurationArgs>;
    autoScalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
    blueGreenDeploymentConfiguration?: pulumi.Input<inputs.codedeploy.DeploymentGroupBlueGreenDeploymentConfigurationArgs>;
    deployment?: pulumi.Input<inputs.codedeploy.DeploymentGroupDeploymentArgs>;
    deploymentConfigName?: pulumi.Input<string>;
    deploymentGroupName?: pulumi.Input<string>;
    deploymentStyle?: pulumi.Input<inputs.codedeploy.DeploymentGroupDeploymentStyleArgs>;
    ec2TagFilters?: pulumi.Input<pulumi.Input<inputs.codedeploy.DeploymentGroupEc2TagFilterArgs>[]>;
    ec2TagSet?: pulumi.Input<inputs.codedeploy.DeploymentGroupEc2TagSetArgs>;
    ecsServices?: pulumi.Input<pulumi.Input<inputs.codedeploy.DeploymentGroupEcsServiceArgs>[]>;
    loadBalancerInfo?: pulumi.Input<inputs.codedeploy.DeploymentGroupLoadBalancerInfoArgs>;
    onPremisesInstanceTagFilters?: pulumi.Input<pulumi.Input<inputs.codedeploy.DeploymentGroupTagFilterArgs>[]>;
    onPremisesTagSet?: pulumi.Input<inputs.codedeploy.DeploymentGroupOnPremisesTagSetArgs>;
    outdatedInstancesStrategy?: pulumi.Input<string>;
    serviceRoleArn: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    terminationHookEnabled?: pulumi.Input<boolean>;
    triggerConfigurations?: pulumi.Input<pulumi.Input<inputs.codedeploy.DeploymentGroupTriggerConfigArgs>[]>;
}
