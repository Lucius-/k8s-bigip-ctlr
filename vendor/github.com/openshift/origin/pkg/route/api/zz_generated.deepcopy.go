// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Route, InType: reflect.TypeOf(&Route{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteIngress, InType: reflect.TypeOf(&RouteIngress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteIngressCondition, InType: reflect.TypeOf(&RouteIngressCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteList, InType: reflect.TypeOf(&RouteList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RoutePort, InType: reflect.TypeOf(&RoutePort{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteSpec, InType: reflect.TypeOf(&RouteSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteStatus, InType: reflect.TypeOf(&RouteStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouteTargetReference, InType: reflect.TypeOf(&RouteTargetReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_RouterShard, InType: reflect.TypeOf(&RouterShard{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_TLSConfig, InType: reflect.TypeOf(&TLSConfig{})},
	)
}

func DeepCopy_api_Route(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Route)
		out := out.(*Route)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_api_RouteSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_api_RouteStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_api_RouteIngress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteIngress)
		out := out.(*RouteIngress)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]RouteIngressCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_api_RouteIngressCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_api_RouteIngressCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteIngressCondition)
		out := out.(*RouteIngressCondition)
		*out = *in
		if in.LastTransitionTime != nil {
			in, out := &in.LastTransitionTime, &out.LastTransitionTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

func DeepCopy_api_RouteList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteList)
		out := out.(*RouteList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Route, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Route(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_api_RoutePort(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoutePort)
		out := out.(*RoutePort)
		*out = *in
		return nil
	}
}

func DeepCopy_api_RouteSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteSpec)
		out := out.(*RouteSpec)
		*out = *in
		if err := DeepCopy_api_RouteTargetReference(&in.To, &out.To, c); err != nil {
			return err
		}
		if in.AlternateBackends != nil {
			in, out := &in.AlternateBackends, &out.AlternateBackends
			*out = make([]RouteTargetReference, len(*in))
			for i := range *in {
				if err := DeepCopy_api_RouteTargetReference(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Port != nil {
			in, out := &in.Port, &out.Port
			*out = new(RoutePort)
			**out = **in
		}
		if in.TLS != nil {
			in, out := &in.TLS, &out.TLS
			*out = new(TLSConfig)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_api_RouteStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteStatus)
		out := out.(*RouteStatus)
		*out = *in
		if in.Ingress != nil {
			in, out := &in.Ingress, &out.Ingress
			*out = make([]RouteIngress, len(*in))
			for i := range *in {
				if err := DeepCopy_api_RouteIngress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_api_RouteTargetReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteTargetReference)
		out := out.(*RouteTargetReference)
		*out = *in
		if in.Weight != nil {
			in, out := &in.Weight, &out.Weight
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_api_RouterShard(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouterShard)
		out := out.(*RouterShard)
		*out = *in
		return nil
	}
}

func DeepCopy_api_TLSConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TLSConfig)
		out := out.(*TLSConfig)
		*out = *in
		return nil
	}
}
