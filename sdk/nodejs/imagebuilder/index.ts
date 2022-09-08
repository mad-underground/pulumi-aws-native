// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ComponentArgs } from "./component";
export type Component = import("./component").Component;
export const Component: typeof import("./component").Component = null as any;

export { ContainerRecipeArgs } from "./containerRecipe";
export type ContainerRecipe = import("./containerRecipe").ContainerRecipe;
export const ContainerRecipe: typeof import("./containerRecipe").ContainerRecipe = null as any;

export { DistributionConfigurationArgs } from "./distributionConfiguration";
export type DistributionConfiguration = import("./distributionConfiguration").DistributionConfiguration;
export const DistributionConfiguration: typeof import("./distributionConfiguration").DistributionConfiguration = null as any;

export { GetComponentArgs, GetComponentResult, GetComponentOutputArgs } from "./getComponent";
export const getComponent: typeof import("./getComponent").getComponent = null as any;
export const getComponentOutput: typeof import("./getComponent").getComponentOutput = null as any;

export { GetContainerRecipeArgs, GetContainerRecipeResult, GetContainerRecipeOutputArgs } from "./getContainerRecipe";
export const getContainerRecipe: typeof import("./getContainerRecipe").getContainerRecipe = null as any;
export const getContainerRecipeOutput: typeof import("./getContainerRecipe").getContainerRecipeOutput = null as any;

export { GetDistributionConfigurationArgs, GetDistributionConfigurationResult, GetDistributionConfigurationOutputArgs } from "./getDistributionConfiguration";
export const getDistributionConfiguration: typeof import("./getDistributionConfiguration").getDistributionConfiguration = null as any;
export const getDistributionConfigurationOutput: typeof import("./getDistributionConfiguration").getDistributionConfigurationOutput = null as any;

export { GetImageArgs, GetImageResult, GetImageOutputArgs } from "./getImage";
export const getImage: typeof import("./getImage").getImage = null as any;
export const getImageOutput: typeof import("./getImage").getImageOutput = null as any;

export { GetImagePipelineArgs, GetImagePipelineResult, GetImagePipelineOutputArgs } from "./getImagePipeline";
export const getImagePipeline: typeof import("./getImagePipeline").getImagePipeline = null as any;
export const getImagePipelineOutput: typeof import("./getImagePipeline").getImagePipelineOutput = null as any;

export { GetImageRecipeArgs, GetImageRecipeResult, GetImageRecipeOutputArgs } from "./getImageRecipe";
export const getImageRecipe: typeof import("./getImageRecipe").getImageRecipe = null as any;
export const getImageRecipeOutput: typeof import("./getImageRecipe").getImageRecipeOutput = null as any;

export { GetInfrastructureConfigurationArgs, GetInfrastructureConfigurationResult, GetInfrastructureConfigurationOutputArgs } from "./getInfrastructureConfiguration";
export const getInfrastructureConfiguration: typeof import("./getInfrastructureConfiguration").getInfrastructureConfiguration = null as any;
export const getInfrastructureConfigurationOutput: typeof import("./getInfrastructureConfiguration").getInfrastructureConfigurationOutput = null as any;

export { ImageArgs } from "./image";
export type Image = import("./image").Image;
export const Image: typeof import("./image").Image = null as any;

export { ImagePipelineArgs } from "./imagePipeline";
export type ImagePipeline = import("./imagePipeline").ImagePipeline;
export const ImagePipeline: typeof import("./imagePipeline").ImagePipeline = null as any;

export { ImageRecipeArgs } from "./imageRecipe";
export type ImageRecipe = import("./imageRecipe").ImageRecipe;
export const ImageRecipe: typeof import("./imageRecipe").ImageRecipe = null as any;

export { InfrastructureConfigurationArgs } from "./infrastructureConfiguration";
export type InfrastructureConfiguration = import("./infrastructureConfiguration").InfrastructureConfiguration;
export const InfrastructureConfiguration: typeof import("./infrastructureConfiguration").InfrastructureConfiguration = null as any;

utilities.lazyLoad(exports, ["Component"], () => require("./component"));
utilities.lazyLoad(exports, ["ContainerRecipe"], () => require("./containerRecipe"));
utilities.lazyLoad(exports, ["DistributionConfiguration"], () => require("./distributionConfiguration"));
utilities.lazyLoad(exports, ["getComponent","getComponentOutput"], () => require("./getComponent"));
utilities.lazyLoad(exports, ["getContainerRecipe","getContainerRecipeOutput"], () => require("./getContainerRecipe"));
utilities.lazyLoad(exports, ["getDistributionConfiguration","getDistributionConfigurationOutput"], () => require("./getDistributionConfiguration"));
utilities.lazyLoad(exports, ["getImage","getImageOutput"], () => require("./getImage"));
utilities.lazyLoad(exports, ["getImagePipeline","getImagePipelineOutput"], () => require("./getImagePipeline"));
utilities.lazyLoad(exports, ["getImageRecipe","getImageRecipeOutput"], () => require("./getImageRecipe"));
utilities.lazyLoad(exports, ["getInfrastructureConfiguration","getInfrastructureConfigurationOutput"], () => require("./getInfrastructureConfiguration"));
utilities.lazyLoad(exports, ["Image"], () => require("./image"));
utilities.lazyLoad(exports, ["ImagePipeline"], () => require("./imagePipeline"));
utilities.lazyLoad(exports, ["ImageRecipe"], () => require("./imageRecipe"));
utilities.lazyLoad(exports, ["InfrastructureConfiguration"], () => require("./infrastructureConfiguration"));

// Export enums:
export * from "../types/enums/imagebuilder";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:imagebuilder:Component":
                return new Component(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:ContainerRecipe":
                return new ContainerRecipe(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:DistributionConfiguration":
                return new DistributionConfiguration(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:Image":
                return new Image(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:ImagePipeline":
                return new ImagePipeline(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:ImageRecipe":
                return new ImageRecipe(name, <any>undefined, { urn })
            case "aws-native:imagebuilder:InfrastructureConfiguration":
                return new InfrastructureConfiguration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "imagebuilder", _module)
