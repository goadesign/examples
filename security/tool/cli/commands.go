package cli

import (
	"github.com/goadesign/examples/security/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// AuthorizeOAuth2ProviderCommand is the command line data structure for the authorize action of OAuth2Provider
	AuthorizeOAuth2ProviderCommand struct {
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

	// GetTokenOAuth2ProviderCommand is the command line data structure for the get_token action of OAuth2Provider
	GetTokenOAuth2ProviderCommand struct {
		PrettyPrint bool
	}

	// SecuredAPIKeyCommand is the command line data structure for the secured action of api_key
	SecuredAPIKeyCommand struct {
		PrettyPrint bool
	}

	// UnsecuredAPIKeyCommand is the command line data structure for the unsecured action of api_key
	UnsecuredAPIKeyCommand struct {
		PrettyPrint bool
	}

	// SecuredBasicCommand is the command line data structure for the secured action of basic
	SecuredBasicCommand struct {
		PrettyPrint bool
	}

	// UnsecuredBasicCommand is the command line data structure for the unsecured action of basic
	UnsecuredBasicCommand struct {
		PrettyPrint bool
	}

	// SecuredJWTCommand is the command line data structure for the secured action of jwt
	SecuredJWTCommand struct {
		// Force auth failure via JWT validation middleware
		Fail        bool
		PrettyPrint bool
	}

	// SigninJWTCommand is the command line data structure for the signin action of jwt
	SigninJWTCommand struct {
		PrettyPrint bool
	}

	// UnsecuredJWTCommand is the command line data structure for the unsecured action of jwt
	UnsecuredJWTCommand struct {
		PrettyPrint bool
	}

	// SecuredOauth2Command is the command line data structure for the secured action of oauth2
	SecuredOauth2Command struct {
		PrettyPrint bool
	}

	// UnsecuredOauth2Command is the command line data structure for the unsecured action of oauth2
	UnsecuredOauth2Command struct {
		PrettyPrint bool
	}

	// WriteOauth2Command is the command line data structure for the write action of oauth2
	WriteOauth2Command struct {
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "authorize",
		Short: `Authorize OAuth2 client`,
	}
	tmp1 := new(AuthorizeOAuth2ProviderCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Provider [/oauth2/authorize]`,
		Short: `This resource implements the OAuth2 authorization code flow`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get_token",
		Short: `Get OAuth2 access token from authorization code or refresh token`,
	}
	tmp2 := new(GetTokenOAuth2ProviderCommand)
	sub = &cobra.Command{
		Use:   `OAuth2Provider [/oauth2/token]`,
		Short: `This resource implements the OAuth2 authorization code flow`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "secured",
		Short: `secured action`,
	}
	tmp3 := new(SecuredAPIKeyCommand)
	sub = &cobra.Command{
		Use:   `api_key [/api_key]`,
		Short: `This resource uses an API key to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(SecuredBasicCommand)
	sub = &cobra.Command{
		Use:   `basic [/basic]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(SecuredJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt [/jwt]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(SecuredOauth2Command)
	sub = &cobra.Command{
		Use:   `oauth2 [/oauth2/read]`,
		Short: `This resource uses OAuth2 to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `Creates a valid JWT`,
	}
	tmp7 := new(SigninJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt [/jwt/signin]`,
		Short: `This resource uses JWT to secure its endpoints`,
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
	tmp8 := new(UnsecuredAPIKeyCommand)
	sub = &cobra.Command{
		Use:   `api_key [/api_key/unsecured]`,
		Short: `This resource uses an API key to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(UnsecuredBasicCommand)
	sub = &cobra.Command{
		Use:   `basic [/basic/unsecured]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(UnsecuredJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt [/jwt/unsecured]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(UnsecuredOauth2Command)
	sub = &cobra.Command{
		Use:   `oauth2 [/oauth2/unsecured]`,
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
	tmp12 := new(WriteOauth2Command)
	sub = &cobra.Command{
		Use:   `oauth2 [/oauth2/write]`,
		Short: `This resource uses OAuth2 to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

// Run makes the HTTP request corresponding to the AuthorizeOAuth2ProviderCommand command.
func (cmd *AuthorizeOAuth2ProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/authorize"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AuthorizeOAuth2Provider(ctx, path, cmd.ClientID, cmd.ResponseType, &cmd.RedirectURI, &cmd.Scope, &cmd.State)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AuthorizeOAuth2ProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
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

// Run makes the HTTP request corresponding to the GetTokenOAuth2ProviderCommand command.
func (cmd *GetTokenOAuth2ProviderCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/token"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetTokenOAuth2Provider(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetTokenOAuth2ProviderCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredAPIKeyCommand command.
func (cmd *SecuredAPIKeyCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api_key"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredAPIKey(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredAPIKeyCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredAPIKeyCommand command.
func (cmd *UnsecuredAPIKeyCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api_key/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredAPIKey(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredAPIKeyCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredBasicCommand command.
func (cmd *SecuredBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredBasicCommand command.
func (cmd *UnsecuredBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/basic/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredJWTCommand command.
func (cmd *SecuredJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredJWT(ctx, path, &cmd.Fail)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var fail bool
	cc.Flags().BoolVar(&cmd.Fail, "fail", fail, `Force auth failure via JWT validation middleware`)
}

// Run makes the HTTP request corresponding to the SigninJWTCommand command.
func (cmd *SigninJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/signin"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninJWT(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredJWTCommand command.
func (cmd *UnsecuredJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jwt/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredJWT(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecuredOauth2Command command.
func (cmd *SecuredOauth2Command) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/read"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecuredOauth2(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecuredOauth2Command) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecuredOauth2Command command.
func (cmd *UnsecuredOauth2Command) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/unsecured"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecuredOauth2(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecuredOauth2Command) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the WriteOauth2Command command.
func (cmd *WriteOauth2Command) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/oauth2/write"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.WriteOauth2(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WriteOauth2Command) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
