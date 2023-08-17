// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	post "example/post"
)

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler FederationServiceErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
}

// FederationServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// FederationServiceDependencyServiceClient has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependencyServiceClient struct {
	Post_PostServiceClient post.PostServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// FederationServiceErrorHandler Federation Service often needs to convert errors received from downstream services.
// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
type FederationServiceErrorHandler func(ctx context.Context, methodName string, err error) error

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
)

// FederationServiceRecoveredError represents recovered error.
type FederationServiceRecoveredError struct {
	Message string
	Stack   []string
}

func (e *FederationServiceRecoveredError) Error() string {
	return fmt.Sprintf("recovered error: %s", e.Message)
}

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler FederationServiceErrorHandler
	client       *FederationServiceDependencyServiceClient
}

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument struct {
	Client *FederationServiceDependencyServiceClient
	Id     string
	Post   *Post
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument struct {
	Client *FederationServiceDependencyServiceClient
	Id     string
	Post   *post.Post
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if err := validateFederationServiceConfig(cfg); err != nil {
		return nil, err
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
		Name:    "post_service",
	})
	if err != nil {
		return nil, err
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		client: &FederationServiceDependencyServiceClient{
			Post_PostServiceClient: Post_PostServiceClient,
		},
	}, nil
}

func validateFederationServiceConfig(cfg FederationServiceConfig) error {
	if cfg.Client == nil {
		return fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	return nil
}

func withTimeoutFederationService[T any](ctx context.Context, method string, timeout time.Duration, fn func(context.Context) (*T, error)) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var (
		ret   *T
		errch = make(chan error)
	)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errch <- recoverErrorFederationService(r, debug.Stack())
			}
		}()

		res, err := fn(ctx)
		ret = res
		errch <- err
	}()
	select {
	case <-ctx.Done():
		status := grpcstatus.New(grpccodes.DeadlineExceeded, ctx.Err().Error())
		withDetails, err := status.WithDetails(&errdetails.ErrorInfo{
			Metadata: map[string]string{
				"method":  method,
				"timeout": timeout.String(),
			},
		})
		if err != nil {
			return nil, status.Err()
		}
		return nil, withDetails.Err()
	case err := <-errch:
		return ret, err
	}
}

func withRetryFederationService[T any](b backoff.BackOff, fn func() (*T, error)) (*T, error) {
	var res *T
	if err := backoff.Retry(func() (err error) {
		res, err = fn()
		return
	}, b); err != nil {
		return nil, err
	}
	return res, nil
}

func recoverErrorFederationService(v interface{}, rawStack []byte) *FederationServiceRecoveredError {
	msg := fmt.Sprint(v)
	lines := strings.Split(msg, "\n")
	if len(lines) <= 1 {
		lines := strings.Split(string(rawStack), "\n")
		stack := make([]string, 0, len(lines))
		for _, line := range lines {
			if line == "" {
				continue
			}
			stack = append(stack, strings.TrimPrefix(line, "\t"))
		}
		return &FederationServiceRecoveredError{
			Message: msg,
			Stack:   stack,
		}
	}
	// If panic occurs under singleflight, singleflight's recover catches the error and gives a stack trace.
	// Therefore, once the stack trace is removed.
	stack := make([]string, 0, len(lines))
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		stack = append(stack, strings.TrimPrefix(line, "\t"))
	}
	return &FederationServiceRecoveredError{
		Message: lines[0],
		Stack:   stack,
	}
}

func (s *FederationService) goWithRecover(eg *errgroup.Group, fn func() (interface{}, error)) {
	eg.Go(func() (e error) {
		defer func() {
			if r := recover(); r != nil {
				e = recoverErrorFederationService(r, debug.Stack())
			}
		}()
		_, err := fn()
		return err
	})
}

func (s *FederationService) outputErrorLog(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if status, ok := grpcstatus.FromError(err); ok {
		s.logger.ErrorContext(ctx, status.Message(),
			slog.Group("grpc_status",
				slog.String("code", status.Code().String()),
				slog.Any("details", status.Details()),
			),
		)
		return
	}
	var recoveredErr *FederationServiceRecoveredError
	if errors.As(err, &recoveredErr) {
		trace := make([]interface{}, 0, len(recoveredErr.Stack))
		for idx, stack := range recoveredErr.Stack {
			trace = append(trace, slog.String(fmt.Sprint(idx+1), stack))
		}
		s.logger.ErrorContext(ctx, recoveredErr.Message, slog.Group("stack_trace", trace...))
		return
	}
	s.logger.ErrorContext(ctx, err.Error())
}

// GetPost implements "federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = recoverErrorFederationService(r, debug.Stack())
			s.outputErrorLog(ctx, e)
		}
	}()
	res, err := withTimeoutFederationService[GetPostResponse](ctx, "federation.FederationService/GetPost", 1000000000 /* 1s */, func(ctx context.Context) (*GetPostResponse, error) {
		return s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument{
			Client: s.client,
			Id:     req.Id,
		})
	})
	if err != nil {
		s.outputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument) (*GetPostResponse, error) {
	s.logger.DebugContext(ctx, "resolve  federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *Post
	)

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "post"
	     message: "Post"
	     args { name: "id", by: "$.id" }
	   }
	*/
	resPostIface, err, _ := sg.Do("post_federation.Post", func() (interface{}, error) {
		valueMu.RLock()
		args := &Federation_PostArgument{
			Client: s.client,
			Id:     req.Id, // { name: "id", by: "$.id" }
		}
		valueMu.RUnlock()
		return s.resolve_Federation_Post(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resPost := resPostIface.(*Post)
	valueMu.Lock()
	valuePost = resPost // { name: "post", message: "Post" ... }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	ret.Post = valuePost // (grpc.federation.field).by = "post"

	s.logger.DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument) (*Post, error) {
	s.logger.DebugContext(ctx, "resolve  federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *post.Post
	)

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "post.PostService/GetPost"
	     request { field: "id", by: "$.id" }
	     response { name: "post", field: "post", autobind: true }
	   }
	*/
	resGetPostResponseIface, err, _ := sg.Do("post.PostService/GetPost", func() (interface{}, error) {
		valueMu.RLock()
		args := &post.GetPostRequest{
			Id: req.Id, // { field: "id", by: "$.id" }
		}
		valueMu.RUnlock()
		return s.client.Post_PostServiceClient.GetPost(ctx, args)
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
			return nil, err
		}
	}
	resGetPostResponse := resGetPostResponseIface.(*post.GetPostResponse)
	valueMu.Lock()
	valuePost = resGetPostResponse.GetPost() // { name: "post", field: "post", autobind: true }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = valuePost.GetId()           // { name: "post", autobind: true }
	ret.Title = valuePost.GetTitle()     // { name: "post", autobind: true }
	ret.Content = valuePost.GetContent() // { name: "post", autobind: true }

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Post(v *Post) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}
