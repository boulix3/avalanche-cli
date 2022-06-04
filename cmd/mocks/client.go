// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/ava-labs/avalanche-network-runner/client"

	mock "github.com/stretchr/testify/mock"

	rpcpb "github.com/ava-labs/avalanche-network-runner/rpcpb"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: ctx, name, execPath, opts
func (_m *Client) AddNode(ctx context.Context, name string, execPath string, opts ...client.OpOption) (*rpcpb.AddNodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, execPath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.AddNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...client.OpOption) *rpcpb.AddNodeResponse); ok {
		r0 = rf(ctx, name, execPath, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, name, execPath, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSubnets provides a mock function with given fields: ctx, numSubnets, opts
func (_m *Client) AddSubnets(ctx context.Context, numSubnets uint32, opts ...client.OpOption) (*rpcpb.AddSubnetsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, numSubnets)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.AddSubnetsResponse
	if rf, ok := ret.Get(0).(func(context.Context, uint32, ...client.OpOption) *rpcpb.AddSubnetsResponse); ok {
		r0 = rf(ctx, numSubnets, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AddSubnetsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, ...client.OpOption) error); ok {
		r1 = rf(ctx, numSubnets, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachPeer provides a mock function with given fields: ctx, nodeName
func (_m *Client) AttachPeer(ctx context.Context, nodeName string) (*rpcpb.AttachPeerResponse, error) {
	ret := _m.Called(ctx, nodeName)

	var r0 *rpcpb.AttachPeerResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.AttachPeerResponse); ok {
		r0 = rf(ctx, nodeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.AttachPeerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Client) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployBlockchains provides a mock function with given fields: ctx, opts
func (_m *Client) DeployBlockchains(ctx context.Context, opts ...client.OpOption) (*rpcpb.DeployBlockchainsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.DeployBlockchainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, ...client.OpOption) *rpcpb.DeployBlockchainsResponse); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.DeployBlockchainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...client.OpOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSnapshotNames provides a mock function with given fields: ctx
func (_m *Client) GetSnapshotNames(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx
func (_m *Client) Health(ctx context.Context) (*rpcpb.HealthResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.HealthResponse
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.HealthResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.HealthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshot provides a mock function with given fields: ctx, snapshotName, opts
func (_m *Client) LoadSnapshot(ctx context.Context, snapshotName string, opts ...client.OpOption) (*rpcpb.LoadSnapshotResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, snapshotName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.LoadSnapshotResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.LoadSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.LoadSnapshotResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, snapshotName, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *Client) Ping(ctx context.Context) (*rpcpb.PingResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.PingResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.PingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNode provides a mock function with given fields: ctx, name
func (_m *Client) RemoveNode(ctx context.Context, name string) (*rpcpb.RemoveNodeResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *rpcpb.RemoveNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.RemoveNodeResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RemoveNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveSnapshot provides a mock function with given fields: ctx, snapshotName
func (_m *Client) RemoveSnapshot(ctx context.Context, snapshotName string) (*rpcpb.RemoveSnapshotResponse, error) {
	ret := _m.Called(ctx, snapshotName)

	var r0 *rpcpb.RemoveSnapshotResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.RemoveSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RemoveSnapshotResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, snapshotName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartNode provides a mock function with given fields: ctx, name, opts
func (_m *Client) RestartNode(ctx context.Context, name string, opts ...client.OpOption) (*rpcpb.RestartNodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.RestartNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.RestartNodeResponse); ok {
		r0 = rf(ctx, name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.RestartNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, name, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSnapshot provides a mock function with given fields: ctx, snapshotName
func (_m *Client) SaveSnapshot(ctx context.Context, snapshotName string) (*rpcpb.SaveSnapshotResponse, error) {
	ret := _m.Called(ctx, snapshotName)

	var r0 *rpcpb.SaveSnapshotResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *rpcpb.SaveSnapshotResponse); ok {
		r0 = rf(ctx, snapshotName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.SaveSnapshotResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, snapshotName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendOutboundMessage provides a mock function with given fields: ctx, nodeName, peerID, op, msgBody
func (_m *Client) SendOutboundMessage(ctx context.Context, nodeName string, peerID string, op uint32, msgBody []byte) (*rpcpb.SendOutboundMessageResponse, error) {
	ret := _m.Called(ctx, nodeName, peerID, op, msgBody)

	var r0 *rpcpb.SendOutboundMessageResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint32, []byte) *rpcpb.SendOutboundMessageResponse); ok {
		r0 = rf(ctx, nodeName, peerID, op, msgBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.SendOutboundMessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, uint32, []byte) error); ok {
		r1 = rf(ctx, nodeName, peerID, op, msgBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: ctx, execPath, opts
func (_m *Client) Start(ctx context.Context, execPath string, opts ...client.OpOption) (*rpcpb.StartResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, execPath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rpcpb.StartResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, ...client.OpOption) *rpcpb.StartResponse); ok {
		r0 = rf(ctx, execPath, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StartResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...client.OpOption) error); ok {
		r1 = rf(ctx, execPath, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx
func (_m *Client) Status(ctx context.Context) (*rpcpb.StatusResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.StatusResponse
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.StatusResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: ctx
func (_m *Client) Stop(ctx context.Context) (*rpcpb.StopResponse, error) {
	ret := _m.Called(ctx)

	var r0 *rpcpb.StopResponse
	if rf, ok := ret.Get(0).(func(context.Context) *rpcpb.StopResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpcpb.StopResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamStatus provides a mock function with given fields: ctx, pushInterval
func (_m *Client) StreamStatus(ctx context.Context, pushInterval time.Duration) (<-chan *rpcpb.ClusterInfo, error) {
	ret := _m.Called(ctx, pushInterval)

	var r0 <-chan *rpcpb.ClusterInfo
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) <-chan *rpcpb.ClusterInfo); ok {
		r0 = rf(ctx, pushInterval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *rpcpb.ClusterInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Duration) error); ok {
		r1 = rf(ctx, pushInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URIs provides a mock function with given fields: ctx
func (_m *Client) URIs(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
