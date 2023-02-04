// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package different

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Component struct {
	pulumi.ResourceState

	A   pulumi.BoolOutput      `pulumi:"a"`
	B   pulumi.BoolPtrOutput   `pulumi:"b"`
	Bar FooPtrOutput           `pulumi:"bar"`
	Baz FooArrayOutput         `pulumi:"baz"`
	C   pulumi.IntOutput       `pulumi:"c"`
	D   pulumi.IntPtrOutput    `pulumi:"d"`
	E   pulumi.StringOutput    `pulumi:"e"`
	F   pulumi.StringPtrOutput `pulumi:"f"`
	Foo FooPtrOutput           `pulumi:"foo"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	var resource Component
	err := ctx.RegisterRemoteComponentResource("example::Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type componentArgs struct {
	A   bool    `pulumi:"a"`
	B   *bool   `pulumi:"b"`
	Bar *Foo    `pulumi:"bar"`
	Baz []Foo   `pulumi:"baz"`
	C   int     `pulumi:"c"`
	D   *int    `pulumi:"d"`
	E   string  `pulumi:"e"`
	F   *string `pulumi:"f"`
	Foo *Foo    `pulumi:"foo"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	A   bool
	B   *bool
	Bar *FooArgs
	Baz []FooInput
	C   int
	D   *int
	E   string
	F   *string
	Foo FooPtrInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func (o ComponentOutput) A() pulumi.BoolOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolOutput { return v.A }).(pulumi.BoolOutput)
}

func (o ComponentOutput) B() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.BoolPtrOutput { return v.B }).(pulumi.BoolPtrOutput)
}

func (o ComponentOutput) Bar() FooPtrOutput {
	return o.ApplyT(func(v *Component) FooPtrOutput { return v.Bar }).(FooPtrOutput)
}

func (o ComponentOutput) Baz() FooArrayOutput {
	return o.ApplyT(func(v *Component) FooArrayOutput { return v.Baz }).(FooArrayOutput)
}

func (o ComponentOutput) C() pulumi.IntOutput {
	return o.ApplyT(func(v *Component) pulumi.IntOutput { return v.C }).(pulumi.IntOutput)
}

func (o ComponentOutput) D() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.IntPtrOutput { return v.D }).(pulumi.IntPtrOutput)
}

func (o ComponentOutput) E() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.E }).(pulumi.StringOutput)
}

func (o ComponentOutput) F() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.F }).(pulumi.StringPtrOutput)
}

func (o ComponentOutput) Foo() FooPtrOutput {
	return o.ApplyT(func(v *Component) FooPtrOutput { return v.Foo }).(FooPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	pulumi.RegisterOutputType(ComponentOutput{})
}
