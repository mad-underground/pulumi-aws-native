// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./codeRepository";
export * from "./dataQualityJobDefinition";
export * from "./device";
export * from "./deviceFleet";
export * from "./endpoint";
export * from "./endpointConfig";
export * from "./model";
export * from "./modelBiasJobDefinition";
export * from "./modelExplainabilityJobDefinition";
export * from "./modelPackageGroup";
export * from "./modelQualityJobDefinition";
export * from "./monitoringSchedule";
export * from "./notebookInstance";
export * from "./notebookInstanceLifecycleConfig";
export * from "./pipeline";
export * from "./project";
export * from "./workteam";

// Import resources to register:
import { CodeRepository } from "./codeRepository";
import { DataQualityJobDefinition } from "./dataQualityJobDefinition";
import { Device } from "./device";
import { DeviceFleet } from "./deviceFleet";
import { Endpoint } from "./endpoint";
import { EndpointConfig } from "./endpointConfig";
import { Model } from "./model";
import { ModelBiasJobDefinition } from "./modelBiasJobDefinition";
import { ModelExplainabilityJobDefinition } from "./modelExplainabilityJobDefinition";
import { ModelPackageGroup } from "./modelPackageGroup";
import { ModelQualityJobDefinition } from "./modelQualityJobDefinition";
import { MonitoringSchedule } from "./monitoringSchedule";
import { NotebookInstance } from "./notebookInstance";
import { NotebookInstanceLifecycleConfig } from "./notebookInstanceLifecycleConfig";
import { Pipeline } from "./pipeline";
import { Project } from "./project";
import { Workteam } from "./workteam";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:SageMaker:CodeRepository":
                return new CodeRepository(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:DataQualityJobDefinition":
                return new DataQualityJobDefinition(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Device":
                return new Device(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:DeviceFleet":
                return new DeviceFleet(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:EndpointConfig":
                return new EndpointConfig(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Model":
                return new Model(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:ModelBiasJobDefinition":
                return new ModelBiasJobDefinition(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:ModelExplainabilityJobDefinition":
                return new ModelExplainabilityJobDefinition(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:ModelPackageGroup":
                return new ModelPackageGroup(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:ModelQualityJobDefinition":
                return new ModelQualityJobDefinition(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:MonitoringSchedule":
                return new MonitoringSchedule(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:NotebookInstance":
                return new NotebookInstance(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:NotebookInstanceLifecycleConfig":
                return new NotebookInstanceLifecycleConfig(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Project":
                return new Project(name, <any>undefined, { urn })
            case "cloudformation:SageMaker:Workteam":
                return new Workteam(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "SageMaker", _module)