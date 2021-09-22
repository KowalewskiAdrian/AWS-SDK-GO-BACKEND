package awsbase

import (
	"context"
	"fmt"

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Builds the user-agent string for APN. Includes the required trailing comma.
func (apn APNInfo) BuildUserAgentString() string {
	builder := smithyhttp.NewUserAgentBuilder()
	builder.AddKeyValue("APN", "1.0")
	builder.AddKeyValue(apn.PartnerName, "1.0")
	for _, p := range apn.Products {
		builder.AddKeyValue(p.Name, p.Version)
	}
	return builder.Build() + ","
}

func apnUserAgentMiddleware(apn APNInfo) middleware.BuildMiddleware {
	return middleware.BuildMiddlewareFunc("APNUserAgent",
		func(ctx context.Context, in middleware.BuildInput, next middleware.BuildHandler) (middleware.BuildOutput, middleware.Metadata, error) {
			request, ok := in.Request.(*smithyhttp.Request)
			if !ok {
				return middleware.BuildOutput{}, middleware.Metadata{}, fmt.Errorf("unknown request type %T", in.Request)
			}

			prependUserAgentHeader(request, apn.BuildUserAgentString())

			return next.HandleBuild(ctx, in)
		},
	)
}

// Because the default User-Agent middleware prepends itself to the contents of the User-Agent header,
// we have to run after it and also prepend our custom User-Agent
func prependUserAgentHeader(request *smithyhttp.Request, value string) {
	current := request.Header.Get("User-Agent")
	if len(current) > 0 {
		current = value + " " + current
	} else {
		current = value
	}
	request.Header["User-Agent"] = append(request.Header["User-Agent"][:0], current)

}
