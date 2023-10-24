// Package corehr code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkcorehr

import (
	"context"
)

// 消息处理器定义
type P2ContractCreatedV1Handler struct {
	handler func(context.Context, *P2ContractCreatedV1) error
}

func NewP2ContractCreatedV1Handler(handler func(context.Context, *P2ContractCreatedV1) error) *P2ContractCreatedV1Handler {
	h := &P2ContractCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ContractCreatedV1Handler) Event() interface{} {
	return &P2ContractCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2ContractCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ContractCreatedV1))
}

// 消息处理器定义
type P2DepartmentCreatedV1Handler struct {
	handler func(context.Context, *P2DepartmentCreatedV1) error
}

func NewP2DepartmentCreatedV1Handler(handler func(context.Context, *P2DepartmentCreatedV1) error) *P2DepartmentCreatedV1Handler {
	h := &P2DepartmentCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2DepartmentCreatedV1Handler) Event() interface{} {
	return &P2DepartmentCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2DepartmentCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2DepartmentCreatedV1))
}

// 消息处理器定义
type P2DepartmentDeletedV1Handler struct {
	handler func(context.Context, *P2DepartmentDeletedV1) error
}

func NewP2DepartmentDeletedV1Handler(handler func(context.Context, *P2DepartmentDeletedV1) error) *P2DepartmentDeletedV1Handler {
	h := &P2DepartmentDeletedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2DepartmentDeletedV1Handler) Event() interface{} {
	return &P2DepartmentDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2DepartmentDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2DepartmentDeletedV1))
}

// 消息处理器定义
type P2DepartmentUpdatedV1Handler struct {
	handler func(context.Context, *P2DepartmentUpdatedV1) error
}

func NewP2DepartmentUpdatedV1Handler(handler func(context.Context, *P2DepartmentUpdatedV1) error) *P2DepartmentUpdatedV1Handler {
	h := &P2DepartmentUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2DepartmentUpdatedV1Handler) Event() interface{} {
	return &P2DepartmentUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2DepartmentUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2DepartmentUpdatedV1))
}

// 消息处理器定义
type P2EmploymentConvertedV1Handler struct {
	handler func(context.Context, *P2EmploymentConvertedV1) error
}

func NewP2EmploymentConvertedV1Handler(handler func(context.Context, *P2EmploymentConvertedV1) error) *P2EmploymentConvertedV1Handler {
	h := &P2EmploymentConvertedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2EmploymentConvertedV1Handler) Event() interface{} {
	return &P2EmploymentConvertedV1{}
}

// 回调开发者注册的handle
func (h *P2EmploymentConvertedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2EmploymentConvertedV1))
}

// 消息处理器定义
type P2EmploymentCreatedV1Handler struct {
	handler func(context.Context, *P2EmploymentCreatedV1) error
}

func NewP2EmploymentCreatedV1Handler(handler func(context.Context, *P2EmploymentCreatedV1) error) *P2EmploymentCreatedV1Handler {
	h := &P2EmploymentCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2EmploymentCreatedV1Handler) Event() interface{} {
	return &P2EmploymentCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2EmploymentCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2EmploymentCreatedV1))
}

// 消息处理器定义
type P2EmploymentDeletedV1Handler struct {
	handler func(context.Context, *P2EmploymentDeletedV1) error
}

func NewP2EmploymentDeletedV1Handler(handler func(context.Context, *P2EmploymentDeletedV1) error) *P2EmploymentDeletedV1Handler {
	h := &P2EmploymentDeletedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2EmploymentDeletedV1Handler) Event() interface{} {
	return &P2EmploymentDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2EmploymentDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2EmploymentDeletedV1))
}

// 消息处理器定义
type P2EmploymentResignedV1Handler struct {
	handler func(context.Context, *P2EmploymentResignedV1) error
}

func NewP2EmploymentResignedV1Handler(handler func(context.Context, *P2EmploymentResignedV1) error) *P2EmploymentResignedV1Handler {
	h := &P2EmploymentResignedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2EmploymentResignedV1Handler) Event() interface{} {
	return &P2EmploymentResignedV1{}
}

// 回调开发者注册的handle
func (h *P2EmploymentResignedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2EmploymentResignedV1))
}

// 消息处理器定义
type P2EmploymentUpdatedV1Handler struct {
	handler func(context.Context, *P2EmploymentUpdatedV1) error
}

func NewP2EmploymentUpdatedV1Handler(handler func(context.Context, *P2EmploymentUpdatedV1) error) *P2EmploymentUpdatedV1Handler {
	h := &P2EmploymentUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2EmploymentUpdatedV1Handler) Event() interface{} {
	return &P2EmploymentUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2EmploymentUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2EmploymentUpdatedV1))
}

// 消息处理器定义
type P2JobChangeUpdatedV1Handler struct {
	handler func(context.Context, *P2JobChangeUpdatedV1) error
}

func NewP2JobChangeUpdatedV1Handler(handler func(context.Context, *P2JobChangeUpdatedV1) error) *P2JobChangeUpdatedV1Handler {
	h := &P2JobChangeUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2JobChangeUpdatedV1Handler) Event() interface{} {
	return &P2JobChangeUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2JobChangeUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2JobChangeUpdatedV1))
}

// 消息处理器定义
type P2JobDataChangedV1Handler struct {
	handler func(context.Context, *P2JobDataChangedV1) error
}

func NewP2JobDataChangedV1Handler(handler func(context.Context, *P2JobDataChangedV1) error) *P2JobDataChangedV1Handler {
	h := &P2JobDataChangedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2JobDataChangedV1Handler) Event() interface{} {
	return &P2JobDataChangedV1{}
}

// 回调开发者注册的handle
func (h *P2JobDataChangedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2JobDataChangedV1))
}

// 消息处理器定义
type P2JobDataEmployedV1Handler struct {
	handler func(context.Context, *P2JobDataEmployedV1) error
}

func NewP2JobDataEmployedV1Handler(handler func(context.Context, *P2JobDataEmployedV1) error) *P2JobDataEmployedV1Handler {
	h := &P2JobDataEmployedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2JobDataEmployedV1Handler) Event() interface{} {
	return &P2JobDataEmployedV1{}
}

// 回调开发者注册的handle
func (h *P2JobDataEmployedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2JobDataEmployedV1))
}

// 消息处理器定义
type P2OffboardingUpdatedV1Handler struct {
	handler func(context.Context, *P2OffboardingUpdatedV1) error
}

func NewP2OffboardingUpdatedV1Handler(handler func(context.Context, *P2OffboardingUpdatedV1) error) *P2OffboardingUpdatedV1Handler {
	h := &P2OffboardingUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2OffboardingUpdatedV1Handler) Event() interface{} {
	return &P2OffboardingUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2OffboardingUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2OffboardingUpdatedV1))
}

// 消息处理器定义
type P2OrgRoleAuthorizationUpdatedV1Handler struct {
	handler func(context.Context, *P2OrgRoleAuthorizationUpdatedV1) error
}

func NewP2OrgRoleAuthorizationUpdatedV1Handler(handler func(context.Context, *P2OrgRoleAuthorizationUpdatedV1) error) *P2OrgRoleAuthorizationUpdatedV1Handler {
	h := &P2OrgRoleAuthorizationUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2OrgRoleAuthorizationUpdatedV1Handler) Event() interface{} {
	return &P2OrgRoleAuthorizationUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2OrgRoleAuthorizationUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2OrgRoleAuthorizationUpdatedV1))
}

// 消息处理器定义
type P2PersonCreatedV1Handler struct {
	handler func(context.Context, *P2PersonCreatedV1) error
}

func NewP2PersonCreatedV1Handler(handler func(context.Context, *P2PersonCreatedV1) error) *P2PersonCreatedV1Handler {
	h := &P2PersonCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2PersonCreatedV1Handler) Event() interface{} {
	return &P2PersonCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2PersonCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2PersonCreatedV1))
}

// 消息处理器定义
type P2PersonDeletedV1Handler struct {
	handler func(context.Context, *P2PersonDeletedV1) error
}

func NewP2PersonDeletedV1Handler(handler func(context.Context, *P2PersonDeletedV1) error) *P2PersonDeletedV1Handler {
	h := &P2PersonDeletedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2PersonDeletedV1Handler) Event() interface{} {
	return &P2PersonDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2PersonDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2PersonDeletedV1))
}

// 消息处理器定义
type P2PersonUpdatedV1Handler struct {
	handler func(context.Context, *P2PersonUpdatedV1) error
}

func NewP2PersonUpdatedV1Handler(handler func(context.Context, *P2PersonUpdatedV1) error) *P2PersonUpdatedV1Handler {
	h := &P2PersonUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2PersonUpdatedV1Handler) Event() interface{} {
	return &P2PersonUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2PersonUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2PersonUpdatedV1))
}

// 消息处理器定义
type P2PreHireUpdatedV1Handler struct {
	handler func(context.Context, *P2PreHireUpdatedV1) error
}

func NewP2PreHireUpdatedV1Handler(handler func(context.Context, *P2PreHireUpdatedV1) error) *P2PreHireUpdatedV1Handler {
	h := &P2PreHireUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2PreHireUpdatedV1Handler) Event() interface{} {
	return &P2PreHireUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2PreHireUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2PreHireUpdatedV1))
}
