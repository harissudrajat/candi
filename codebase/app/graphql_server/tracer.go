package graphqlserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	gqlerrors "github.com/golangid/graphql-go/errors"
	"github.com/golangid/graphql-go/introspection"
	"github.com/golangid/graphql-go/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/candishared"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/types"
	"pkg.agungdwiprasetyo.com/candi/config/env"
	"pkg.agungdwiprasetyo.com/candi/logger"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

const schemaRootInstropectionField = "__schema"

var gqlTypeNotShowLog = map[string]bool{
	"Query": true, "Mutation": true, "Subscription": true, "__Type": true, "__Schema": true,
}

// graphQLTracer struct
type graphQLTracer struct {
	midd types.GraphQLMiddlewareGroup
}

// newGraphQLTracer constructor
func newGraphQLTracer(midd types.GraphQLMiddlewareGroup) *graphQLTracer {
	return &graphQLTracer{
		midd: midd,
	}
}

// TraceQuery method
func (t *graphQLTracer) TraceQuery(ctx context.Context, queryString string, operationName string, variables map[string]interface{}, varTypes map[string]*introspection.Type) (context.Context, trace.TraceQueryFinishFunc) {

	globalTracer := opentracing.GlobalTracer()
	traceName := strings.TrimSuffix(fmt.Sprintf("GraphQL-Root:%s", operationName), ":")

	headers, _ := ctx.Value(candishared.ContextKeyHTTPHeader).(http.Header)
	var span opentracing.Span
	if spanCtx, err := globalTracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(headers)); err != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, traceName)
		ext.SpanKindRPCServer.Set(span)
	} else {
		span = globalTracer.StartSpan(traceName, ext.RPCServerOption((spanCtx)))
		ctx = opentracing.ContextWithSpan(ctx, span)
		ext.SpanKindRPCClient.Set(span)
	}

	if len(headers) > 0 {
		span.SetTag("http.headers", string(candihelper.ToBytes(headers)))
	}
	span.SetTag("graphql.query", queryString)
	span.SetTag("graphql.operationName", operationName)
	if len(variables) != 0 {
		span.SetTag("graphql.variables", variables)
	}

	return ctx, func(data []byte, errs []*gqlerrors.QueryError) {
		defer span.Finish()

		span.LogKV("data", string(data))
		if len(errs) > 0 {
			span.LogKV("errors", string(candihelper.ToBytes(errs)))
			ext.Error.Set(span, true)
		}
		logger.LogGreen("graphql " + tracer.GetTraceURL(ctx))
	}
}

// TraceField method
func (t *graphQLTracer) TraceField(ctx context.Context, label, typeName, fieldName string, trivial bool, args map[string]interface{}) (context.Context, trace.TraceFieldFinishFunc) {
	start := time.Now()
	if middFunc, ok := t.midd[fmt.Sprintf("%s.%s", typeName, fieldName)]; ok {
		ctx = middFunc(ctx)
	}
	return ctx, func(data []byte, err *gqlerrors.QueryError) {
		end := time.Now()
		if env.BaseEnv().DebugMode && !trivial && !gqlTypeNotShowLog[typeName] && fieldName != schemaRootInstropectionField {
			statusColor := candihelper.Green
			status := " OK  "
			if err != nil {
				statusColor = candihelper.Red
				status = "ERROR"
			}

			arg, _ := json.Marshal(args)
			fmt.Fprintf(os.Stdout, "%s[GRAPHQL]%s => %s %10s %s | %v | %s %s %s | %13v | %s %s %s | %s\n",
				candihelper.White, candihelper.Reset,
				candihelper.Blue, typeName, candihelper.Reset,
				end.Format("2006/01/02 - 15:04:05"),
				statusColor, status, candihelper.Reset,
				end.Sub(start),
				candihelper.Magenta, label, candihelper.Reset,
				arg,
			)
		}
	}
}
