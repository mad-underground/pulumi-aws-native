// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CapacityProviderArgs } from "./capacityProvider";
export type CapacityProvider = import("./capacityProvider").CapacityProvider;
export const CapacityProvider: typeof import("./capacityProvider").CapacityProvider = null as any;
utilities.lazyLoad(exports, ["CapacityProvider"], () => require("./capacityProvider"));

export { ClusterArgs } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { ClusterCapacityProviderAssociationsArgs } from "./clusterCapacityProviderAssociations";
export type ClusterCapacityProviderAssociations = import("./clusterCapacityProviderAssociations").ClusterCapacityProviderAssociations;
export const ClusterCapacityProviderAssociations: typeof import("./clusterCapacityProviderAssociations").ClusterCapacityProviderAssociations = null as any;
utilities.lazyLoad(exports, ["ClusterCapacityProviderAssociations"], () => require("./clusterCapacityProviderAssociations"));

export { GetCapacityProviderArgs, GetCapacityProviderResult, GetCapacityProviderOutputArgs } from "./getCapacityProvider";
export const getCapacityProvider: typeof import("./getCapacityProvider").getCapacityProvider = null as any;
export const getCapacityProviderOutput: typeof import("./getCapacityProvider").getCapacityProviderOutput = null as any;
utilities.lazyLoad(exports, ["getCapacityProvider","getCapacityProviderOutput"], () => require("./getCapacityProvider"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetClusterCapacityProviderAssociationsArgs, GetClusterCapacityProviderAssociationsResult, GetClusterCapacityProviderAssociationsOutputArgs } from "./getClusterCapacityProviderAssociations";
export const getClusterCapacityProviderAssociations: typeof import("./getClusterCapacityProviderAssociations").getClusterCapacityProviderAssociations = null as any;
export const getClusterCapacityProviderAssociationsOutput: typeof import("./getClusterCapacityProviderAssociations").getClusterCapacityProviderAssociationsOutput = null as any;
utilities.lazyLoad(exports, ["getClusterCapacityProviderAssociations","getClusterCapacityProviderAssociationsOutput"], () => require("./getClusterCapacityProviderAssociations"));

export { GetPrimaryTaskSetArgs, GetPrimaryTaskSetResult, GetPrimaryTaskSetOutputArgs } from "./getPrimaryTaskSet";
export const getPrimaryTaskSet: typeof import("./getPrimaryTaskSet").getPrimaryTaskSet = null as any;
export const getPrimaryTaskSetOutput: typeof import("./getPrimaryTaskSet").getPrimaryTaskSetOutput = null as any;
utilities.lazyLoad(exports, ["getPrimaryTaskSet","getPrimaryTaskSetOutput"], () => require("./getPrimaryTaskSet"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { GetTaskDefinitionArgs, GetTaskDefinitionResult, GetTaskDefinitionOutputArgs } from "./getTaskDefinition";
export const getTaskDefinition: typeof import("./getTaskDefinition").getTaskDefinition = null as any;
export const getTaskDefinitionOutput: typeof import("./getTaskDefinition").getTaskDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getTaskDefinition","getTaskDefinitionOutput"], () => require("./getTaskDefinition"));

export { GetTaskSetArgs, GetTaskSetResult, GetTaskSetOutputArgs } from "./getTaskSet";
export const getTaskSet: typeof import("./getTaskSet").getTaskSet = null as any;
export const getTaskSetOutput: typeof import("./getTaskSet").getTaskSetOutput = null as any;
utilities.lazyLoad(exports, ["getTaskSet","getTaskSetOutput"], () => require("./getTaskSet"));

export { PrimaryTaskSetArgs } from "./primaryTaskSet";
export type PrimaryTaskSet = import("./primaryTaskSet").PrimaryTaskSet;
export const PrimaryTaskSet: typeof import("./primaryTaskSet").PrimaryTaskSet = null as any;
utilities.lazyLoad(exports, ["PrimaryTaskSet"], () => require("./primaryTaskSet"));

export { ServiceArgs } from "./service";
export type Service = import("./service").Service;
export const Service: typeof import("./service").Service = null as any;
utilities.lazyLoad(exports, ["Service"], () => require("./service"));

export { TaskDefinitionArgs } from "./taskDefinition";
export type TaskDefinition = import("./taskDefinition").TaskDefinition;
export const TaskDefinition: typeof import("./taskDefinition").TaskDefinition = null as any;
utilities.lazyLoad(exports, ["TaskDefinition"], () => require("./taskDefinition"));

export { TaskSetArgs } from "./taskSet";
export type TaskSet = import("./taskSet").TaskSet;
export const TaskSet: typeof import("./taskSet").TaskSet = null as any;
utilities.lazyLoad(exports, ["TaskSet"], () => require("./taskSet"));


// Export enums:
export * from "../types/enums/ecs";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ecs:CapacityProvider":
                return new CapacityProvider(name, <any>undefined, { urn })
            case "aws-native:ecs:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:ecs:ClusterCapacityProviderAssociations":
                return new ClusterCapacityProviderAssociations(name, <any>undefined, { urn })
            case "aws-native:ecs:PrimaryTaskSet":
                return new PrimaryTaskSet(name, <any>undefined, { urn })
            case "aws-native:ecs:Service":
                return new Service(name, <any>undefined, { urn })
            case "aws-native:ecs:TaskDefinition":
                return new TaskDefinition(name, <any>undefined, { urn })
            case "aws-native:ecs:TaskSet":
                return new TaskSet(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ecs", _module)
