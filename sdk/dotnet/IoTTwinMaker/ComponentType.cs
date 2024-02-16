// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker
{
    /// <summary>
    /// Resource schema for AWS::IoTTwinMaker::ComponentType
    /// </summary>
    [AwsNativeResourceType("aws-native:iottwinmaker:ComponentType")]
    public partial class ComponentType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the component type.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the component type.
        /// </summary>
        [Output("componentTypeId")]
        public Output<string> ComponentTypeId { get; private set; } = null!;

        /// <summary>
        /// An map of the composite component types in the component type. Each composite component type's key must be unique to this map.
        /// </summary>
        [Output("compositeComponentTypes")]
        public Output<ImmutableDictionary<string, Outputs.ComponentTypeCompositeComponentType>?> CompositeComponentTypes { get; private set; } = null!;

        /// <summary>
        /// The date and time when the component type was created.
        /// </summary>
        [Output("creationDateTime")]
        public Output<string> CreationDateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the component type.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the parent component type to extend.
        /// </summary>
        [Output("extendsFrom")]
        public Output<ImmutableArray<string>> ExtendsFrom { get; private set; } = null!;

        /// <summary>
        /// a Map of functions in the component type. Each function's key must be unique to this map.
        /// </summary>
        [Output("functions")]
        public Output<ImmutableDictionary<string, Outputs.ComponentTypeFunction>?> Functions { get; private set; } = null!;

        /// <summary>
        /// A Boolean value that specifies whether the component type is abstract.
        /// </summary>
        [Output("isAbstract")]
        public Output<bool> IsAbstract { get; private set; } = null!;

        /// <summary>
        /// A Boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
        /// </summary>
        [Output("isSchemaInitialized")]
        public Output<bool> IsSchemaInitialized { get; private set; } = null!;

        /// <summary>
        /// A Boolean value that specifies whether an entity can have more than one component of this type.
        /// </summary>
        [Output("isSingleton")]
        public Output<bool?> IsSingleton { get; private set; } = null!;

        /// <summary>
        /// An map of the property definitions in the component type. Each property definition's key must be unique to this map.
        /// </summary>
        [Output("propertyDefinitions")]
        public Output<ImmutableDictionary<string, Outputs.ComponentTypePropertyDefinition>?> PropertyDefinitions { get; private set; } = null!;

        /// <summary>
        /// An map of the property groups in the component type. Each property group's key must be unique to this map.
        /// </summary>
        [Output("propertyGroups")]
        public Output<ImmutableDictionary<string, Outputs.ComponentTypePropertyGroup>?> PropertyGroups { get; private set; } = null!;

        /// <summary>
        /// The current status of the component type.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ComponentTypeStatus> Status { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs to associate with a resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The last date and time when the component type was updated.
        /// </summary>
        [Output("updateDateTime")]
        public Output<string> UpdateDateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the workspace that contains the component type.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a ComponentType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComponentType(string name, ComponentTypeArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iottwinmaker:ComponentType", name, args ?? new ComponentTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComponentType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iottwinmaker:ComponentType", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "componentTypeId",
                    "workspaceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ComponentType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComponentType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ComponentType(name, id, options);
        }
    }

    public sealed class ComponentTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the component type.
        /// </summary>
        [Input("componentTypeId", required: true)]
        public Input<string> ComponentTypeId { get; set; } = null!;

        [Input("compositeComponentTypes")]
        private InputMap<Inputs.ComponentTypeCompositeComponentTypeArgs>? _compositeComponentTypes;

        /// <summary>
        /// An map of the composite component types in the component type. Each composite component type's key must be unique to this map.
        /// </summary>
        public InputMap<Inputs.ComponentTypeCompositeComponentTypeArgs> CompositeComponentTypes
        {
            get => _compositeComponentTypes ?? (_compositeComponentTypes = new InputMap<Inputs.ComponentTypeCompositeComponentTypeArgs>());
            set => _compositeComponentTypes = value;
        }

        /// <summary>
        /// The description of the component type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extendsFrom")]
        private InputList<string>? _extendsFrom;

        /// <summary>
        /// Specifies the parent component type to extend.
        /// </summary>
        public InputList<string> ExtendsFrom
        {
            get => _extendsFrom ?? (_extendsFrom = new InputList<string>());
            set => _extendsFrom = value;
        }

        [Input("functions")]
        private InputMap<Inputs.ComponentTypeFunctionArgs>? _functions;

        /// <summary>
        /// a Map of functions in the component type. Each function's key must be unique to this map.
        /// </summary>
        public InputMap<Inputs.ComponentTypeFunctionArgs> Functions
        {
            get => _functions ?? (_functions = new InputMap<Inputs.ComponentTypeFunctionArgs>());
            set => _functions = value;
        }

        /// <summary>
        /// A Boolean value that specifies whether an entity can have more than one component of this type.
        /// </summary>
        [Input("isSingleton")]
        public Input<bool>? IsSingleton { get; set; }

        [Input("propertyDefinitions")]
        private InputMap<Inputs.ComponentTypePropertyDefinitionArgs>? _propertyDefinitions;

        /// <summary>
        /// An map of the property definitions in the component type. Each property definition's key must be unique to this map.
        /// </summary>
        public InputMap<Inputs.ComponentTypePropertyDefinitionArgs> PropertyDefinitions
        {
            get => _propertyDefinitions ?? (_propertyDefinitions = new InputMap<Inputs.ComponentTypePropertyDefinitionArgs>());
            set => _propertyDefinitions = value;
        }

        [Input("propertyGroups")]
        private InputMap<Inputs.ComponentTypePropertyGroupArgs>? _propertyGroups;

        /// <summary>
        /// An map of the property groups in the component type. Each property group's key must be unique to this map.
        /// </summary>
        public InputMap<Inputs.ComponentTypePropertyGroupArgs> PropertyGroups
        {
            get => _propertyGroups ?? (_propertyGroups = new InputMap<Inputs.ComponentTypePropertyGroupArgs>());
            set => _propertyGroups = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs to associate with a resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the workspace that contains the component type.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public ComponentTypeArgs()
        {
        }
        public static new ComponentTypeArgs Empty => new ComponentTypeArgs();
    }
}
