// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
 */
export class TaskSet extends pulumi.CustomResource {
    /**
     * Get an existing TaskSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TaskSet {
        return new TaskSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ecs:TaskSet';

    /**
     * Returns true if the given object is an instance of TaskSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TaskSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TaskSet.__pulumiType;
    }

    /**
     * The ID of the task set.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. 
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. 
     */
    public readonly launchType!: pulumi.Output<enums.ecs.TaskSetLaunchType | undefined>;
    public readonly loadBalancers!: pulumi.Output<outputs.ecs.TaskSetLoadBalancer[] | undefined>;
    public readonly networkConfiguration!: pulumi.Output<outputs.ecs.TaskSetNetworkConfiguration | undefined>;
    /**
     * The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.
     */
    public readonly platformVersion!: pulumi.Output<string | undefined>;
    /**
     * A floating-point percentage of the desired number of tasks to place and keep running in the task set.
     */
    public readonly scale!: pulumi.Output<outputs.ecs.TaskSetScale | undefined>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.
     */
    public readonly serviceRegistries!: pulumi.Output<outputs.ecs.TaskSetServiceRegistry[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.
     */
    public readonly taskDefinition!: pulumi.Output<string>;

    /**
     * Create a TaskSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cluster'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.taskDefinition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskDefinition'");
            }
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["launchType"] = args ? args.launchType : undefined;
            resourceInputs["loadBalancers"] = args ? args.loadBalancers : undefined;
            resourceInputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            resourceInputs["platformVersion"] = args ? args.platformVersion : undefined;
            resourceInputs["scale"] = args ? args.scale : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["serviceRegistries"] = args ? args.serviceRegistries : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["taskDefinition"] = args ? args.taskDefinition : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["cluster"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["launchType"] = undefined /*out*/;
            resourceInputs["loadBalancers"] = undefined /*out*/;
            resourceInputs["networkConfiguration"] = undefined /*out*/;
            resourceInputs["platformVersion"] = undefined /*out*/;
            resourceInputs["scale"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
            resourceInputs["serviceRegistries"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["taskDefinition"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["cluster", "externalId", "launchType", "loadBalancers[*]", "networkConfiguration", "platformVersion", "service", "serviceRegistries[*]", "taskDefinition"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TaskSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TaskSet resource.
 */
export interface TaskSetArgs {
    /**
     * The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
     */
    cluster: pulumi.Input<string>;
    /**
     * An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. 
     */
    externalId?: pulumi.Input<string>;
    /**
     * The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. 
     */
    launchType?: pulumi.Input<enums.ecs.TaskSetLaunchType>;
    loadBalancers?: pulumi.Input<pulumi.Input<inputs.ecs.TaskSetLoadBalancerArgs>[]>;
    networkConfiguration?: pulumi.Input<inputs.ecs.TaskSetNetworkConfigurationArgs>;
    /**
     * The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.
     */
    platformVersion?: pulumi.Input<string>;
    /**
     * A floating-point percentage of the desired number of tasks to place and keep running in the task set.
     */
    scale?: pulumi.Input<inputs.ecs.TaskSetScaleArgs>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
     */
    service: pulumi.Input<string>;
    /**
     * The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.
     */
    serviceRegistries?: pulumi.Input<pulumi.Input<inputs.ecs.TaskSetServiceRegistryArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.
     */
    taskDefinition: pulumi.Input<string>;
}
