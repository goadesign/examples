package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/examples/security/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// SecuredAPIKeyEndpointsCommand is the command line data structure for the secured action of APIKeyEndpoints
	SecuredAPIKeyEndpointsCommand struct {
		PrettyPrint bool
	}

	// UnsecuredAPIKeyEndpointsCommand is the command line data structure for the unsecured action of APIKeyEndpoints
	UnsecuredAPIKeyEndpointsCommand struct {
		PrettyPrint bool
	}

	// SecuredBasicAuthEndpointsCommand is the command line data structure for the secured action of BasicAuthEndpoints
	SecuredBasicAuthEndpointsCommand struct {
		PrettyPrint bool
	}

	// UnsecuredBasicAuthEndpointsCommand is the command line data structure for the unsecured action of BasicAuthEndpoints
	UnsecuredBasicAuthEndpointsCommand struct {
		PrettyPrint bool
	}

	// SecuredJWTEndpointsCommand is the command line data structure for the secured action of JWTEndpoints
	SecuredJWTEndpointsCommand struct {
		// Force auth failure via JWT validation middleware
		Fail        bool
		PrettyPrint bool
	}

	// UnsecuredJWTEndpointsCommand is the command line data structure for the unsecured action of JWTEndpoints
	UnsecuredJWTEndpointsCommand struct {
		PrettyPrint bool
	}

	// SecuredOAuth2EndpointsCommand is the command line data structure for the secured action of OAuth2Endpoints
	SecuredOAuth2EndpointsCommand struct {
		PrettyPrint bool
	}

	// UnsecuredOAuth2EndpointsCommand is the command line data structure for the unsecured action of OAuth2Endpoints
	UnsecuredOAuth2EndpointsCommand struct {
		PrettyPrint bool
	}

	// WriteOAuth2EndpointsCommand is the command line data structure for the write action of OAuth2Endpoints
	WriteOAuth2EndpointsCommand struct {
		PrettyPrint bool
	}

	// ExchangeTokenOAuth2ProviderCommand is the command line data structure for the exchange_token action of OAuth2Provider
	ExchangeTokenOAuth2ProviderCommand struct {
		PrettyPrint bool
	}

	// RefreshTokenOAuth2ProviderCommand is the command line data structure for the refresh_token action of OAuth2Provider
	RefreshTokenOAuth2ProviderCommand struct {
		Payload     string
		PrettyPrint bool
	}

	// RequestAuthOAuth2ProviderCommand is the command line data structure for the request_auth action of OAuth2Provider
	RequestAuthOAuth2ProviderCommand struct {
		// The client identifier
		ClientID string
		// Redirection endpoint
		RedirectURI string
		// Value MUST be set to "code"
		ResponseType string
		// The scope of the access request
		Scope string
		// An opaque value used by the client to maintain state between the request and callback
		State       string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "exchange_token",
		Short: `Request coming from client to retrieve access token from authorization code`,
	}
	tmp1 := new(ExchangeTokenOAuth2ProviderCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Provider [/oauth2/token]`,
		Short: `This resource implements the OAuth2 authorization code flow`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "refresh_token",
		Short: `Request coming from client to refresh expired access token`,
	}
	tmp2 := new(RefreshTokenOAuth2ProviderCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Provider [/oauth2/token/refresh]`,
		Short: `This resource implements the OAuth2 authorization code flow`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "request_auth",
		Short: `Request coming from client requesting authorization code`,
	}
	tmp3 := new(RequestAuthOAuth2ProviderCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Provider [/oauth2/auth]`,
		Short: `This resource implements the OAuth2 authorization code flow`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "secured",
		Short: `secured action`,
	}
	tmp4 := new(SecuredAPIKeyEndpointsCommand)
	sub = &cobra.Command{
		Use:   `APIKeyEndpoints [/api_key]`,
		Short: `This resource uses an API key to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(SecuredBasicAuthEndpointsCommand)
	sub = &cobra.Command{
		Use:   `BasicAuthEndpoints [/basic]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(SecuredJWTEndpointsCommand)
	sub = &cobra.Command{
		Use:   `JWTEndpoints [/jwt]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(SecuredOAuth2EndpointsCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Endpoints [/oauth2/read]`,
		Short: `This resource uses OAuth2 to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "unsecured",
		Short: `unsecured action`,
	}
	tmp8 := new(UnsecuredAPIKeyEndpointsCommand)
	sub = &cobra.Command{
		Use:   `APIKeyEndpoints [/api_key/unsecured]`,
		Short: `This resource uses an API key to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(UnsecuredBasicAuthEndpointsCommand)
	sub = &cobra.Command{
		Use:   `BasicAuthEndpoints [/basic/unsecured]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(UnsecuredJWTEndpointsCommand)
	sub = &cobra.Command{
		Use:   `JWTEndpoints [/jwt/unsecured]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(UnsecuredOAuth2EndpointsCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Endpoints [/oauth2/unsecured]`,
		Short: `This resource uses OAuth2 to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "write",
		Short: `This action requires an additional scope on top of "api:read"`,
	}
	tmp12 := new(WriteOAuth2EndpointsCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Endpoints [/oauth2/write]`,
		Short: `This resource uses OAuth2 to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

// Run makes the HTTP request corresponding to the SecuredAPIKeyEndpointsCommand command.
func (cmd *SecuredAPIKeyEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api_key"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredAPIKeyEndpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredAPIKeyEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredAPIKeyEndpointsCommand command.
func (cmd *UnsecuredAPIKeyEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api_key/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredAPIKeyEndpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredAPIKeyEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredBasicAuthEndpointsCommand command.
func (cmd *SecuredBasicAuthEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredBasicAuthEndpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredBasicAuthEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredBasicAuthEndpointsCommand command.
func (cmd *UnsecuredBasicAuthEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredBasicAuthEndpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredBasicAuthEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredJWTEndpointsCommand command.
func (cmd *SecuredJWTEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredJWTEndpoints(ctx, path, &cmd.Fail)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredJWTEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var fail bool
	cc.Flags().BoolVar(&cmd.Fail, "fail", fail, `Force auth failure via JWT validation middleware`)
}

// Run makes the HTTP request corresponding to the UnsecuredJWTEndpointsCommand command.
func (cmd *UnsecuredJWTEndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredJWTEndpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredJWTEndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredOAuth2EndpointsCommand command.
func (cmd *SecuredOAuth2EndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/read"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredOAuth2Endpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredOAuth2EndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredOAuth2EndpointsCommand command.
func (cmd *UnsecuredOAuth2EndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredOAuth2Endpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredOAuth2EndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the WriteOAuth2EndpointsCommand command.
func (cmd *WriteOAuth2EndpointsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/write"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.WriteOAuth2Endpoints(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WriteOAuth2EndpointsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ExchangeTokenOAuth2ProviderCommand command.
func (cmd *ExchangeTokenOAuth2ProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/token"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ExchangeTokenOAuth2Provider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ExchangeTokenOAuth2ProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the RefreshTokenOAuth2ProviderCommand command.
func (cmd *RefreshTokenOAuth2ProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/token/refresh"
	}
	var payload client.RefreshTokenOAuth2ProviderPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RefreshTokenOAuth2Provider(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RefreshTokenOAuth2ProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}

// Run makes the HTTP request corresponding to the RequestAuthOAuth2ProviderCommand command.
func (cmd *RequestAuthOAuth2ProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/auth"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RequestAuthOAuth2Provider(ctx, path, cmd.ClientID, cmd.ResponseType, &cmd.RedirectURI, &cmd.Scope, &cmd.State)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RequestAuthOAuth2ProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var clientID string
	cc.Flags().StringVar(&cmd.ClientID, "client_id", clientID, `The client identifier`)
	var redirectURI string
	cc.Flags().StringVar(&cmd.RedirectURI, "redirect_uri", redirectURI, `Redirection endpoint`)
	var responseType string
	cc.Flags().StringVar(&cmd.ResponseType, "response_type", responseType, `Value MUST be set to "code"`)
	var scope string
	cc.Flags().StringVar(&cmd.Scope, "scope", scope, `The scope of the access request`)
	var state string
	cc.Flags().StringVar(&cmd.State, "state", state, `An opaque value used by the client to maintain state between the request and callback`)
}
