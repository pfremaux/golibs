package web

import (
	"fmt"
	"net/http"
)

type ServerExposureConfig struct {
	Port              int
	CorsAllowOrigin   string
	CorsAllowOMethods string
	AuthTokenOverride string
	TokenValidator    func(token string) error
	mux               *http.ServeMux
}

type WebServerConfig struct {
	Host            string
	DebugMode       bool
	PublicExposure  ServerExposureConfig
	PrivateExposure ServerExposureConfig
}

var defaultConfig = WebServerConfig{
	Host: "",
	PublicExposure: ServerExposureConfig{
		Port:              8080,
		CorsAllowOrigin:   "*",
		CorsAllowOMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AuthTokenOverride: "",
		TokenValidator: func(token string) error {
			// Default implementation: accept any token
			return nil
		},
	},
	PrivateExposure: ServerExposureConfig{
		Port:              8081,
		CorsAllowOrigin:   "*",
		CorsAllowOMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AuthTokenOverride: "",
		TokenValidator: func(token string) error {
			// Default implementation: accept any token
			return nil
		},
	},

	DebugMode: true,
}

type WebServerOperations interface {
	RegisterPublicEndpoint(path string, handler http.HandlerFunc)
	RegisterPrivateEndpoint(path string, handler http.HandlerFunc)
	Listen()
}

func (wsc *WebServerConfig) RegisterPublicEndpoint(path string, handler http.HandlerFunc) {
	wsc.PublicExposure.mux.HandleFunc(path, corsHandler(authHandler(handler, wsc), wsc))
}

func (wsc *WebServerConfig) RegisterPrivateEndpoint(path string, handler http.HandlerFunc) {
	wsc.PrivateExposure.mux.HandleFunc(path, corsHandler(authHandler(handler, wsc), wsc))
}

func (wsc *WebServerConfig) Listen() {
	/*address := fmt.Sprintf("%s:%d", wsc.Host, wsc.PublicExposure.Port)
	if wsc.DebugMode {
		fmt.Printf("Starting web server at %s\n", address)
	}
	if err := http.ListenAndServe(address, wsc.PublicExposure.mux); err != nil {
		panic(err)
	}*/
	go func() {
		address := fmt.Sprintf("%s:%d", wsc.Host, wsc.PublicExposure.Port)
		if wsc.DebugMode {
			fmt.Printf("Starting web server at %s\n", address)
		}
		if err := http.ListenAndServe(address, wsc.PublicExposure.mux); err != nil {
			panic(err)
		}
	}()

	// Not a go routine to keep app running
	address := fmt.Sprintf("%s:%d", wsc.Host, wsc.PrivateExposure.Port)
	if wsc.DebugMode {
		fmt.Printf("Starting web server at %s\n", address)
	}
	if err := http.ListenAndServe(address, wsc.PrivateExposure.mux); err != nil {
		panic(err)
	}
}

func NewWebServer(host string, port int) *WebServerConfig {
	return &WebServerConfig{
		Host: host,
		PublicExposure: ServerExposureConfig{
			Port:              port,
			CorsAllowOrigin:   defaultConfig.PublicExposure.CorsAllowOrigin,
			CorsAllowOMethods: defaultConfig.PublicExposure.CorsAllowOMethods,
			AuthTokenOverride: defaultConfig.PublicExposure.AuthTokenOverride,
			TokenValidator:    defaultConfig.PublicExposure.TokenValidator,
			mux:               http.NewServeMux(),
		},
		PrivateExposure: ServerExposureConfig{
			Port:              port + 1,
			CorsAllowOrigin:   defaultConfig.PrivateExposure.CorsAllowOrigin,
			CorsAllowOMethods: defaultConfig.PrivateExposure.CorsAllowOMethods,
			AuthTokenOverride: defaultConfig.PrivateExposure.AuthTokenOverride,
			TokenValidator:    defaultConfig.PrivateExposure.TokenValidator,
			mux:               http.NewServeMux(),
		},
		DebugMode: defaultConfig.DebugMode,
	}
}

func corsHandler(h http.HandlerFunc, config *WebServerConfig) http.HandlerFunc {
	// todo pfr validate inputs
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("cors 1")
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", config.PublicExposure.CorsAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", config.PublicExposure.CorsAllowOMethods)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Println("cors 2")
		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			fmt.Printf("OPTIONS request received\n")
			w.WriteHeader(http.StatusOK)
			fmt.Println("cors 3")
			return
		}

		// Call the original handler
		fmt.Println("cors 4")
		h.ServeHTTP(w, r)
	}
}

func authHandler(h http.HandlerFunc, config *WebServerConfig) http.HandlerFunc {
	// todo pfr validate inputs
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth 1")
		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			fmt.Printf("OPTIONS request received\n")
			w.WriteHeader(http.StatusOK)
			fmt.Println("auth 2")
			return
		}

		token := r.Header.Get("Authorization")

		fmt.Println("auth 3")
		err := authenticate(token, config)
		if err != nil {

			fmt.Println("auth 4")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("auth 5")
		// Call the original handler
		h.ServeHTTP(w, r)
	}
}

func authenticate(token string, config *WebServerConfig) error { // TODO PFR differentiate pub priv
	if token == "" && config.PublicExposure.AuthTokenOverride == "" {
		fmt.Println("auth skipped")
		return nil
	}
	if config.PublicExposure.AuthTokenOverride != "" && token != config.PublicExposure.AuthTokenOverride {
		return fmt.Errorf("unauthorized: %v", token)
	}
	return config.PublicExposure.TokenValidator(token)
}

func JsonResponse(data string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, data)
	}
}
