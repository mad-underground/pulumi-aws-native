// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppMesh::Route
//
// Deprecated: Route is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Route struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput    `pulumi:"arn"`
	MeshName          pulumi.StringOutput    `pulumi:"meshName"`
	MeshOwner         pulumi.StringPtrOutput `pulumi:"meshOwner"`
	ResourceOwner     pulumi.StringOutput    `pulumi:"resourceOwner"`
	RouteName         pulumi.StringPtrOutput `pulumi:"routeName"`
	Spec              RouteSpecOutput        `pulumi:"spec"`
	Tags              RouteTagArrayOutput    `pulumi:"tags"`
	Uid               pulumi.StringOutput    `pulumi:"uid"`
	VirtualRouterName pulumi.StringOutput    `pulumi:"virtualRouterName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualRouterName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualRouterName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("aws-native:appmesh:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws-native:appmesh:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	MeshName          string     `pulumi:"meshName"`
	MeshOwner         *string    `pulumi:"meshOwner"`
	RouteName         *string    `pulumi:"routeName"`
	Spec              RouteSpec  `pulumi:"spec"`
	Tags              []RouteTag `pulumi:"tags"`
	VirtualRouterName string     `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	MeshName          pulumi.StringInput
	MeshOwner         pulumi.StringPtrInput
	RouteName         pulumi.StringPtrInput
	Spec              RouteSpecInput
	Tags              RouteTagArrayInput
	VirtualRouterName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o RouteOutput) MeshName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.MeshName }).(pulumi.StringOutput)
}

func (o RouteOutput) MeshOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.MeshOwner }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ResourceOwner }).(pulumi.StringOutput)
}

func (o RouteOutput) RouteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.RouteName }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) Spec() RouteSpecOutput {
	return o.ApplyT(func(v *Route) RouteSpecOutput { return v.Spec }).(RouteSpecOutput)
}

func (o RouteOutput) Tags() RouteTagArrayOutput {
	return o.ApplyT(func(v *Route) RouteTagArrayOutput { return v.Tags }).(RouteTagArrayOutput)
}

func (o RouteOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func (o RouteOutput) VirtualRouterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.VirtualRouterName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
