// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"
	"sync"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/singleflight"

	post "example/post"
)

// Org_Federation_CreatePostArgument is argument for "org.federation.CreatePost" message.
type Org_Federation_CreatePostArgument[T any] struct {
	Content string
	Title   string
	UserId  string
	Client  T
}

// Org_Federation_CreatePostResponseArgument is argument for "org.federation.CreatePostResponse" message.
type Org_Federation_CreatePostResponseArgument[T any] struct {
	Content string
	Cp      *CreatePost
	P       *post.Post
	Res     *post.CreatePostResponse
	Title   string
	UserId  string
	Client  T
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// Org_Post_PostServiceClient create a gRPC Client to be used to call methods in org.post.PostService.
	Org_Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
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

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
	Org_Post_PostServiceClient post.PostServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_Org_Post_PostService_CreatePost = "/org.post.PostService/CreatePost"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	Org_Post_PostServiceClient, err := cfg.Client.Org_Post_PostServiceClient(FederationServiceClientConfig{
		Service: "org.post.PostService",
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
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*celtypes.FieldType{
		"grpc.federation.private.CreatePostArgument": {
			"title":   grpcfed.NewCELFieldType(celtypes.StringType, "Title"),
			"content": grpcfed.NewCELFieldType(celtypes.StringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
		"grpc.federation.private.CreatePostResponseArgument": {
			"title":   grpcfed.NewCELFieldType(celtypes.StringType, "Title"),
			"content": grpcfed.NewCELFieldType(celtypes.StringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
		cel.CustomTypeAdapter(celHelper.TypeAdapter()),
		cel.CustomTypeProvider(celHelper.TypeProvider()),
	)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Org_Post_PostServiceClient: Org_Post_PostServiceClient,
		},
	}, nil
}

// CreatePost implements "org.federation.FederationService/CreatePost" method.
func (s *FederationService) CreatePost(ctx context.Context, req *CreatePostRequest) (res *CreatePostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/CreatePost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_CreatePostResponse(ctx, &Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]{
		Client:  s.client,
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_CreatePost resolve "org.federation.CreatePost" message.
func (s *FederationService) resolve_Org_Federation_CreatePost(ctx context.Context, req *Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]) (*CreatePost, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CreatePost")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.CreatePost", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostArgument(req)))
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.CreatePostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// create a message value to be returned.
	ret := &CreatePost{}

	// field binding section.
	// (grpc.federation.field).by = "$.title"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.title", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Title = value.(string)
	}
	// (grpc.federation.field).by = "$.content"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.content", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Content = value.(string)
	}
	// (grpc.federation.field).by = "$.user_id"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.UserId = value.(string)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePost", slog.Any("org.federation.CreatePost", s.logvalue_Org_Federation_CreatePost(ret)))
	return ret, nil
}

// resolve_Org_Federation_CreatePostResponse resolve "org.federation.CreatePostResponse" message.
func (s *FederationService) resolve_Org_Federation_CreatePostResponse(ctx context.Context, req *Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]) (*CreatePostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CreatePostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.CreatePostResponse", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostResponseArgument(req)))
	var (
		sg       singleflight.Group
		valueCp  *CreatePost
		valueMu  sync.RWMutex
		valueP   *post.Post
		valueRes *post.CreatePostResponse
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.CreatePostResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "cp"
	     message {
	       name: "CreatePost"
	       args: [
	         { name: "title", by: "$.title" },
	         { name: "content", by: "$.content" },
	         { name: "user_id", by: "$.user_id" }
	       ]
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("cp", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "title", by: "$.title" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.title", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Title = value.(string)
			}
			// { name: "content", by: "$.content" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.content", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Content = value.(string)
			}
			// { name: "user_id", by: "$.user_id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.UserId = value.(string)
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_CreatePost(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*CreatePost)
		valueMu.Lock()
		valueCp = value // { name: "cp", message: "CreatePost" ... }
		envOpts = append(envOpts, cel.Variable("cp", cel.ObjectType("org.federation.CreatePost")))
		evalValues["cp"] = valueCp
		valueMu.Unlock()
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "org.post.PostService/CreatePost"
	       request { field: "post", by: "cp" }
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("res", func() (any, error) {
			valueMu.RLock()
			args := &post.CreatePostRequest{}
			// { field: "post", by: "cp" }
			{
				value, err := grpcfed.EvalCEL(s.env, "cp", envOpts, evalValues, reflect.TypeOf((*CreatePost)(nil)))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Post = s.cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(value.(*CreatePost))
			}
			valueMu.RUnlock()
			return s.client.Org_Post_PostServiceClient.CreatePost(ctx, args)
		})
		if err != nil {
			if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_Post_PostService_CreatePost, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
		}
		value := valueIface.(*post.CreatePostResponse)
		valueMu.Lock()
		valueRes = value
		envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("org.post.CreatePostResponse")))
		evalValues["res"] = valueRes
		valueMu.Unlock()
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "p"
	     by: "res.post"
	   }
	*/
	{
		valueIface, err, _ := sg.Do("p", func() (any, error) {
			valueMu.RLock()
			valueMu.RUnlock()
			return grpcfed.EvalCEL(s.env, "res.post", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*post.Post)
		valueMu.Lock()
		valueP = value
		envOpts = append(envOpts, cel.Variable("p", cel.ObjectType("org.post.Post")))
		evalValues["p"] = valueP
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Cp = valueCp
	req.P = valueP
	req.Res = valueRes

	// create a message value to be returned.
	ret := &CreatePostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "p"
	{
		value, err := grpcfed.EvalCEL(s.env, "p", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Post = s.cast_Org_Post_Post__to__Org_Federation_Post(value.(*post.Post))
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePostResponse", slog.Any("org.federation.CreatePostResponse", s.logvalue_Org_Federation_CreatePostResponse(ret)))
	return ret, nil
}

// cast_Org_Federation_CreatePost__to__Org_Post_CreatePost cast from "org.federation.CreatePost" to "org.post.CreatePost".
func (s *FederationService) cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(from *CreatePost) *post.CreatePost {
	if from == nil {
		return nil
	}

	return &post.CreatePost{
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

// cast_Org_Post_Post__to__Org_Federation_Post cast from "org.post.Post" to "org.federation.Post".
func (s *FederationService) cast_Org_Post_Post__to__Org_Federation_Post(from *post.Post) *Post {
	if from == nil {
		return nil
	}

	return &Post{
		Id:      from.GetId(),
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

func (s *FederationService) logvalue_Org_Federation_CreatePost(v *CreatePost) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.String("user_id", v.GetUserId()),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostArgument(v *Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponse(v *CreatePostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Org_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponseArgument(v *Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Org_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.String("user_id", v.GetUserId()),
	)
}
