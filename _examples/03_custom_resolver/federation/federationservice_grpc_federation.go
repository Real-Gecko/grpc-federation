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
	user "example/user"
)

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver FederationServiceResolver // required
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
	// User_UserServiceClient create a gRPC Client to be used to call methods in user.UserService.
	User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
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
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
	// Resolve_Federation_ForNameless implements resolver for "federation.ForNameless".
	Resolve_Federation_ForNameless(context.Context, *Federation_ForNamelessArgument) (*ForNameless, error)
	// Resolve_Federation_Post_User implements resolver for "federation.Post.user".
	Resolve_Federation_Post_User(context.Context, *Federation_Post_UserArgument) (*User, error)
	// Resolve_Federation_Unused implements resolver for "federation.Unused".
	Resolve_Federation_Unused(context.Context, *Federation_UnusedArgument) (*Unused, error)
	// Resolve_Federation_User implements resolver for "federation.User".
	Resolve_Federation_User(context.Context, *Federation_UserArgument) (*User, error)
	// Resolve_Federation_User_Name implements resolver for "federation.User.name".
	Resolve_Federation_User_Name(context.Context, *Federation_User_NameArgument) (string, error)
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// Resolve_Federation_ForNameless resolve "federation.ForNameless".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_ForNameless(context.Context, *Federation_ForNamelessArgument) (ret *ForNameless, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_ForNameless not implemented")
	return
}

// Resolve_Federation_Post_User resolve "federation.Post.user".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_Post_User(context.Context, *Federation_Post_UserArgument) (ret *User, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_Post_User not implemented")
	return
}

// Resolve_Federation_Unused resolve "federation.Unused".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_Unused(context.Context, *Federation_UnusedArgument) (ret *Unused, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_Unused not implemented")
	return
}

// Resolve_Federation_User resolve "federation.User".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_User(context.Context, *Federation_UserArgument) (ret *User, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_User not implemented")
	return
}

// Resolve_Federation_User_Name resolve "federation.User.name".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_User_Name(context.Context, *Federation_User_NameArgument) (ret string, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_User_Name not implemented")
	return
}

// FederationServiceErrorHandler Federation Service often needs to convert errors received from downstream services.
// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
type FederationServiceErrorHandler func(ctx context.Context, methodName string, err error) error

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
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
	resolver     FederationServiceResolver
	client       *FederationServiceDependencyServiceClient
}

// Federation_ForNamelessArgument is argument for "federation.ForNameless" message.
type Federation_ForNamelessArgument struct {
	Client *FederationServiceDependencyServiceClient
	Bar    string
}

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument struct {
	Client *FederationServiceDependencyServiceClient
	Id     string
	Post   *Post
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument struct {
	Client                *FederationServiceDependencyServiceClient
	FederationForNameless *ForNameless
	Id                    string
	Post                  *post.Post
	Unused                *Unused
	User                  *User
}

// Federation_Post_UserArgument is custom resolver's argument for "user" field of "federation.Post" message.
type Federation_Post_UserArgument struct {
	Client                  *FederationServiceDependencyServiceClient
	Federation_PostArgument *Federation_PostArgument
}

// Federation_UnusedArgument is argument for "federation.Unused" message.
type Federation_UnusedArgument struct {
	Client *FederationServiceDependencyServiceClient
	Foo    string
}

// Federation_UserArgument is argument for "federation.User" message.
type Federation_UserArgument struct {
	Client  *FederationServiceDependencyServiceClient
	Content string
	Id      string
	Title   string
	U       *user.User
	UserId  string
}

// Federation_User_NameArgument is custom resolver's argument for "name" field of "federation.User" message.
type Federation_User_NameArgument struct {
	Client                  *FederationServiceDependencyServiceClient
	Federation_User         *User
	Federation_UserArgument *Federation_UserArgument
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
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
		Name:    "user_service",
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
		resolver:     cfg.Resolver,
		client: &FederationServiceDependencyServiceClient{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

func validateFederationServiceConfig(cfg FederationServiceConfig) error {
	if cfg.Client == nil {
		return fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	if cfg.Resolver == nil {
		return fmt.Errorf("Resolver field in FederationServiceConfig is not set. this field must be set")
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
	res, err := s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument{
		Client: s.client,
		Id:     req.Id,
	})
	if err != nil {
		s.outputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_ForNameless resolve "federation.ForNameless" message.
func (s *FederationService) resolve_Federation_ForNameless(ctx context.Context, req *Federation_ForNamelessArgument) (*ForNameless, error) {
	s.logger.DebugContext(ctx, "resolve  federation.ForNameless", slog.Any("message_args", s.logvalue_Federation_ForNamelessArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ret, err := s.resolver.Resolve_Federation_ForNameless(ctx, req)
	if err != nil {
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.ForNameless", slog.Any("federation.ForNameless", s.logvalue_Federation_ForNameless(ret)))
	return ret, nil
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
		sg                         singleflight.Group
		valueFederationForNameless *ForNameless
		valueMu                    sync.RWMutex
		valuePost                  *post.Post
		valueUnused                *Unused
		valueUser                  *User
	)
	// A tree view of message dependencies is shown below.
	/*
	            _federation_ForNameless ─┐
	                             unused ─┤
	   GetPost ─┐                        │
	                               user ─┤
	*/
	eg, ctx := errgroup.WithContext(ctx)

	s.goWithRecover(eg, func() (interface{}, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "_federation_ForNameless"
		     message: "ForNameless"
		     args { name: "bar", string: "bar" }
		   }
		*/
		resForNamelessIface, err, _ := sg.Do("_federation_ForNameless_federation.ForNameless", func() (interface{}, error) {
			valueMu.RLock()
			args := &Federation_ForNamelessArgument{
				Client: s.client,
				Bar:    "bar", // { name: "bar", string: "bar" }
			}
			valueMu.RUnlock()
			return s.resolve_Federation_ForNameless(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resForNameless := resForNamelessIface.(*ForNameless)
		valueMu.Lock()
		valueFederationForNameless = resForNameless // { name: "_federation_ForNameless", message: "ForNameless" ... }
		valueMu.Unlock()
		return nil, nil
	})

	s.goWithRecover(eg, func() (interface{}, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "unused"
		     message: "Unused"
		     args { name: "foo", string: "foo" }
		   }
		*/
		resUnusedIface, err, _ := sg.Do("unused_federation.Unused", func() (interface{}, error) {
			valueMu.RLock()
			args := &Federation_UnusedArgument{
				Client: s.client,
				Foo:    "foo", // { name: "foo", string: "foo" }
			}
			valueMu.RUnlock()
			return s.resolve_Federation_Unused(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resUnused := resUnusedIface.(*Unused)
		valueMu.Lock()
		valueUnused = resUnused // { name: "unused", message: "Unused" ... }
		valueMu.Unlock()
		return nil, nil
	})

	s.goWithRecover(eg, func() (interface{}, error) {

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

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "user"
		     message: "User"
		     args { inline: "post" }
		   }
		*/
		resUserIface, err, _ := sg.Do("user_federation.User", func() (interface{}, error) {
			valueMu.RLock()
			args := &Federation_UserArgument{
				Client:  s.client,
				Id:      valuePost.GetId(),      // { inline: "post" }
				Title:   valuePost.GetTitle(),   // { inline: "post" }
				Content: valuePost.GetContent(), // { inline: "post" }
				UserId:  valuePost.GetUserId(),  // { inline: "post" }
			}
			valueMu.RUnlock()
			return s.resolve_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resUser := resUserIface.(*User)
		valueMu.Lock()
		valueUser = resUser // { name: "user", message: "User" ... }
		valueMu.Unlock()
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.FederationForNameless = valueFederationForNameless
	req.Post = valuePost
	req.Unused = valueUnused
	req.User = valueUser

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = valuePost.GetId()           // { name: "post", autobind: true }
	ret.Title = valuePost.GetTitle()     // { name: "post", autobind: true }
	ret.Content = valuePost.GetContent() // { name: "post", autobind: true }
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.User, err = s.resolver.Resolve_Federation_Post_User(ctx, &Federation_Post_UserArgument{
			Client:                  s.client,
			Federation_PostArgument: req,
		})
		if err != nil {
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_Unused resolve "federation.Unused" message.
func (s *FederationService) resolve_Federation_Unused(ctx context.Context, req *Federation_UnusedArgument) (*Unused, error) {
	s.logger.DebugContext(ctx, "resolve  federation.Unused", slog.Any("message_args", s.logvalue_Federation_UnusedArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ret, err := s.resolver.Resolve_Federation_Unused(ctx, req)
	if err != nil {
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.Unused", slog.Any("federation.Unused", s.logvalue_Federation_Unused(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *FederationService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	s.logger.DebugContext(ctx, "resolve  federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
		valueU  *user.User
	)

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "user.UserService/GetUser"
	     request { field: "id", by: "$.user_id" }
	     response { name: "u", field: "user" }
	   }
	*/
	resGetUserResponseIface, err, _ := sg.Do("user.UserService/GetUser", func() (interface{}, error) {
		valueMu.RLock()
		args := &user.GetUserRequest{
			Id: req.UserId, // { field: "id", by: "$.user_id" }
		}
		valueMu.RUnlock()
		return s.client.User_UserServiceClient.GetUser(ctx, args)
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
			return nil, err
		}
	}
	resGetUserResponse := resGetUserResponseIface.(*user.GetUserResponse)
	valueMu.Lock()
	valueU = resGetUserResponse.GetUser() // { name: "u", field: "user" }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.U = valueU

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ret, err := s.resolver.Resolve_Federation_User(ctx, req)
	if err != nil {
		return nil, err
	}

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.Name, err = s.resolver.Resolve_Federation_User_Name(ctx, &Federation_User_NameArgument{
			Client:                  s.client,
			Federation_UserArgument: req,
			Federation_User:         ret,
		})
		if err != nil {
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Federation_ForNameless(v *ForNameless) slog.Value {
	return slog.GroupValue(
		slog.String("bar", v.GetBar()),
	)
}

func (s *FederationService) logvalue_Federation_ForNamelessArgument(v *Federation_ForNamelessArgument) slog.Value {
	return slog.GroupValue(
		slog.String("bar", v.Bar),
	)
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
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Unused(v *Unused) slog.Value {
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationService) logvalue_Federation_UnusedArgument(v *Federation_UnusedArgument) slog.Value {
	return slog.GroupValue(
		slog.String("foo", v.Foo),
	)
}

func (s *FederationService) logvalue_Federation_User(v *User) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}
