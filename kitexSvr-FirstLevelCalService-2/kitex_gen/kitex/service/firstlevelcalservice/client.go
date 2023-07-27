// Code generated by Kitex v0.6.1. DO NOT EDIT.

package firstlevelcalservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	service "kitexSvr-FirstLevelCalService/kitex_gen/kitex/service"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Add(ctx context.Context, request *service.Request, callOptions ...callopt.Option) (r *service.Response, err error)
	Sub(ctx context.Context, request *service.Request, callOptions ...callopt.Option) (r *service.Response, err error)
}

// NewClient creates a client for the idlManager defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFirstLevelCalServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the idlManager defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFirstLevelCalServiceClient struct {
	*kClient
}

func (p *kFirstLevelCalServiceClient) Add(ctx context.Context, request *service.Request, callOptions ...callopt.Option) (r *service.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Add(ctx, request)
}

func (p *kFirstLevelCalServiceClient) Sub(ctx context.Context, request *service.Request, callOptions ...callopt.Option) (r *service.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Sub(ctx, request)
}
