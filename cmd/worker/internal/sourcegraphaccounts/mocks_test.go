// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package sourcegraphaccounts

import (
	"context"
	"sync"

	v1 "github.com/sourcegraph/sourcegraph-accounts-sdk-go/clients/v1"
	database "github.com/sourcegraph/sourcegraph/internal/database"
	extsvc "github.com/sourcegraph/sourcegraph/internal/extsvc"
)

// MockNotificationsSubscriberStore is a mock implementation of the
// notificationsSubscriberStore interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/worker/internal/sourcegraphaccounts)
// used for unit testing.
type MockNotificationsSubscriberStore struct {
	// GetSAMSUserByIDFunc is an instance of a mock function object
	// controlling the behavior of the method GetSAMSUserByID.
	GetSAMSUserByIDFunc *NotificationsSubscriberStoreGetSAMSUserByIDFunc
	// HardDeleteUsersFunc is an instance of a mock function object
	// controlling the behavior of the method HardDeleteUsers.
	HardDeleteUsersFunc *NotificationsSubscriberStoreHardDeleteUsersFunc
	// ListUserExternalAccountsFunc is an instance of a mock function object
	// controlling the behavior of the method ListUserExternalAccounts.
	ListUserExternalAccountsFunc *NotificationsSubscriberStoreListUserExternalAccountsFunc
}

// NewMockNotificationsSubscriberStore creates a new mock of the
// notificationsSubscriberStore interface. All methods return zero values
// for all results, unless overwritten.
func NewMockNotificationsSubscriberStore() *MockNotificationsSubscriberStore {
	return &MockNotificationsSubscriberStore{
		GetSAMSUserByIDFunc: &NotificationsSubscriberStoreGetSAMSUserByIDFunc{
			defaultHook: func(context.Context, string) (r0 *v1.User, r1 error) {
				return
			},
		},
		HardDeleteUsersFunc: &NotificationsSubscriberStoreHardDeleteUsersFunc{
			defaultHook: func(context.Context, []int32) (r0 error) {
				return
			},
		},
		ListUserExternalAccountsFunc: &NotificationsSubscriberStoreListUserExternalAccountsFunc{
			defaultHook: func(context.Context, database.ExternalAccountsListOptions) (r0 []*extsvc.Account, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockNotificationsSubscriberStore creates a new mock of the
// notificationsSubscriberStore interface. All methods panic on invocation,
// unless overwritten.
func NewStrictMockNotificationsSubscriberStore() *MockNotificationsSubscriberStore {
	return &MockNotificationsSubscriberStore{
		GetSAMSUserByIDFunc: &NotificationsSubscriberStoreGetSAMSUserByIDFunc{
			defaultHook: func(context.Context, string) (*v1.User, error) {
				panic("unexpected invocation of MockNotificationsSubscriberStore.GetSAMSUserByID")
			},
		},
		HardDeleteUsersFunc: &NotificationsSubscriberStoreHardDeleteUsersFunc{
			defaultHook: func(context.Context, []int32) error {
				panic("unexpected invocation of MockNotificationsSubscriberStore.HardDeleteUsers")
			},
		},
		ListUserExternalAccountsFunc: &NotificationsSubscriberStoreListUserExternalAccountsFunc{
			defaultHook: func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error) {
				panic("unexpected invocation of MockNotificationsSubscriberStore.ListUserExternalAccounts")
			},
		},
	}
}

// surrogateMockNotificationsSubscriberStore is a copy of the
// notificationsSubscriberStore interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/worker/internal/sourcegraphaccounts).
// It is redefined here as it is unexported in the source package.
type surrogateMockNotificationsSubscriberStore interface {
	GetSAMSUserByID(context.Context, string) (*v1.User, error)
	HardDeleteUsers(context.Context, []int32) error
	ListUserExternalAccounts(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error)
}

// NewMockNotificationsSubscriberStoreFrom creates a new mock of the
// MockNotificationsSubscriberStore interface. All methods delegate to the
// given implementation, unless overwritten.
func NewMockNotificationsSubscriberStoreFrom(i surrogateMockNotificationsSubscriberStore) *MockNotificationsSubscriberStore {
	return &MockNotificationsSubscriberStore{
		GetSAMSUserByIDFunc: &NotificationsSubscriberStoreGetSAMSUserByIDFunc{
			defaultHook: i.GetSAMSUserByID,
		},
		HardDeleteUsersFunc: &NotificationsSubscriberStoreHardDeleteUsersFunc{
			defaultHook: i.HardDeleteUsers,
		},
		ListUserExternalAccountsFunc: &NotificationsSubscriberStoreListUserExternalAccountsFunc{
			defaultHook: i.ListUserExternalAccounts,
		},
	}
}

// NotificationsSubscriberStoreGetSAMSUserByIDFunc describes the behavior
// when the GetSAMSUserByID method of the parent
// MockNotificationsSubscriberStore instance is invoked.
type NotificationsSubscriberStoreGetSAMSUserByIDFunc struct {
	defaultHook func(context.Context, string) (*v1.User, error)
	hooks       []func(context.Context, string) (*v1.User, error)
	history     []NotificationsSubscriberStoreGetSAMSUserByIDFuncCall
	mutex       sync.Mutex
}

// GetSAMSUserByID delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockNotificationsSubscriberStore) GetSAMSUserByID(v0 context.Context, v1 string) (*v1.User, error) {
	r0, r1 := m.GetSAMSUserByIDFunc.nextHook()(v0, v1)
	m.GetSAMSUserByIDFunc.appendCall(NotificationsSubscriberStoreGetSAMSUserByIDFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetSAMSUserByID
// method of the parent MockNotificationsSubscriberStore instance is invoked
// and the hook queue is empty.
func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) SetDefaultHook(hook func(context.Context, string) (*v1.User, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetSAMSUserByID method of the parent MockNotificationsSubscriberStore
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) PushHook(hook func(context.Context, string) (*v1.User, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) SetDefaultReturn(r0 *v1.User, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*v1.User, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) PushReturn(r0 *v1.User, r1 error) {
	f.PushHook(func(context.Context, string) (*v1.User, error) {
		return r0, r1
	})
}

func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) nextHook() func(context.Context, string) (*v1.User, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) appendCall(r0 NotificationsSubscriberStoreGetSAMSUserByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// NotificationsSubscriberStoreGetSAMSUserByIDFuncCall objects describing
// the invocations of this function.
func (f *NotificationsSubscriberStoreGetSAMSUserByIDFunc) History() []NotificationsSubscriberStoreGetSAMSUserByIDFuncCall {
	f.mutex.Lock()
	history := make([]NotificationsSubscriberStoreGetSAMSUserByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// NotificationsSubscriberStoreGetSAMSUserByIDFuncCall is an object that
// describes an invocation of method GetSAMSUserByID on an instance of
// MockNotificationsSubscriberStore.
type NotificationsSubscriberStoreGetSAMSUserByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *v1.User
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c NotificationsSubscriberStoreGetSAMSUserByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c NotificationsSubscriberStoreGetSAMSUserByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// NotificationsSubscriberStoreHardDeleteUsersFunc describes the behavior
// when the HardDeleteUsers method of the parent
// MockNotificationsSubscriberStore instance is invoked.
type NotificationsSubscriberStoreHardDeleteUsersFunc struct {
	defaultHook func(context.Context, []int32) error
	hooks       []func(context.Context, []int32) error
	history     []NotificationsSubscriberStoreHardDeleteUsersFuncCall
	mutex       sync.Mutex
}

// HardDeleteUsers delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockNotificationsSubscriberStore) HardDeleteUsers(v0 context.Context, v1 []int32) error {
	r0 := m.HardDeleteUsersFunc.nextHook()(v0, v1)
	m.HardDeleteUsersFunc.appendCall(NotificationsSubscriberStoreHardDeleteUsersFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the HardDeleteUsers
// method of the parent MockNotificationsSubscriberStore instance is invoked
// and the hook queue is empty.
func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) SetDefaultHook(hook func(context.Context, []int32) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// HardDeleteUsers method of the parent MockNotificationsSubscriberStore
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) PushHook(hook func(context.Context, []int32) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, []int32) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, []int32) error {
		return r0
	})
}

func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) nextHook() func(context.Context, []int32) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) appendCall(r0 NotificationsSubscriberStoreHardDeleteUsersFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// NotificationsSubscriberStoreHardDeleteUsersFuncCall objects describing
// the invocations of this function.
func (f *NotificationsSubscriberStoreHardDeleteUsersFunc) History() []NotificationsSubscriberStoreHardDeleteUsersFuncCall {
	f.mutex.Lock()
	history := make([]NotificationsSubscriberStoreHardDeleteUsersFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// NotificationsSubscriberStoreHardDeleteUsersFuncCall is an object that
// describes an invocation of method HardDeleteUsers on an instance of
// MockNotificationsSubscriberStore.
type NotificationsSubscriberStoreHardDeleteUsersFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 []int32
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c NotificationsSubscriberStoreHardDeleteUsersFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c NotificationsSubscriberStoreHardDeleteUsersFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// NotificationsSubscriberStoreListUserExternalAccountsFunc describes the
// behavior when the ListUserExternalAccounts method of the parent
// MockNotificationsSubscriberStore instance is invoked.
type NotificationsSubscriberStoreListUserExternalAccountsFunc struct {
	defaultHook func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error)
	hooks       []func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error)
	history     []NotificationsSubscriberStoreListUserExternalAccountsFuncCall
	mutex       sync.Mutex
}

// ListUserExternalAccounts delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockNotificationsSubscriberStore) ListUserExternalAccounts(v0 context.Context, v1 database.ExternalAccountsListOptions) ([]*extsvc.Account, error) {
	r0, r1 := m.ListUserExternalAccountsFunc.nextHook()(v0, v1)
	m.ListUserExternalAccountsFunc.appendCall(NotificationsSubscriberStoreListUserExternalAccountsFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// ListUserExternalAccounts method of the parent
// MockNotificationsSubscriberStore instance is invoked and the hook queue
// is empty.
func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) SetDefaultHook(hook func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListUserExternalAccounts method of the parent
// MockNotificationsSubscriberStore instance invokes the hook at the front
// of the queue and discards it. After the queue is empty, the default hook
// function is invoked for any future action.
func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) PushHook(hook func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) SetDefaultReturn(r0 []*extsvc.Account, r1 error) {
	f.SetDefaultHook(func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) PushReturn(r0 []*extsvc.Account, r1 error) {
	f.PushHook(func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error) {
		return r0, r1
	})
}

func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) nextHook() func(context.Context, database.ExternalAccountsListOptions) ([]*extsvc.Account, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) appendCall(r0 NotificationsSubscriberStoreListUserExternalAccountsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// NotificationsSubscriberStoreListUserExternalAccountsFuncCall objects
// describing the invocations of this function.
func (f *NotificationsSubscriberStoreListUserExternalAccountsFunc) History() []NotificationsSubscriberStoreListUserExternalAccountsFuncCall {
	f.mutex.Lock()
	history := make([]NotificationsSubscriberStoreListUserExternalAccountsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// NotificationsSubscriberStoreListUserExternalAccountsFuncCall is an object
// that describes an invocation of method ListUserExternalAccounts on an
// instance of MockNotificationsSubscriberStore.
type NotificationsSubscriberStoreListUserExternalAccountsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 database.ExternalAccountsListOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*extsvc.Account
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c NotificationsSubscriberStoreListUserExternalAccountsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c NotificationsSubscriberStoreListUserExternalAccountsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
