package rest

import (
	"github.com/imdario/mergo"
)

var (
	// DefaultCharset character encoding for rest rendering
	DefaultCharset = "UTF-8"
)

// Config is a struct for specifying configuration options for the rest.Render object.
type Config struct {
	// Appends the given character set to the Content-Type header. Default is "UTF-8".
	Charset string
	// Gzip enable it if you want to render with gzip compression. Default is false
	Gzip bool
	// Outputs human readable JSON.
	IndentJSON bool
	// Outputs human readable XML. Default is false.
	IndentXML bool
	// Prefixes the JSON output with the given bytes. Default is false.
	PrefixJSON []byte
	// Prefixes the XML output with the given bytes.
	PrefixXML []byte
	// Unescape HTML characters "&<>" to their original values. Default is false.
	UnEscapeHTML bool
	// Streams JSON responses instead of marshalling prior to sending. Default is false.
	StreamingJSON bool
	// Disables automatic rendering of http.StatusInternalServerError when an error occurs. Default is false.
	DisableHTTPErrorRendering bool
	// MarkdownSanitize sanitizes the markdown. Default is false.
	MarkdownSanitize bool
}

// DefaultConfig returns the default config for rest
func DefaultConfig() Config {
	return Config{
		Charset:                   DefaultCharset,
		IndentJSON:                false,
		IndentXML:                 false,
		PrefixJSON:                []byte(""),
		PrefixXML:                 []byte(""),
		UnEscapeHTML:              false,
		StreamingJSON:             false,
		DisableHTTPErrorRendering: false,
		MarkdownSanitize:          false,
	}
}

// Merge merges the default with the given config and returns the result
func (c Config) Merge(cfg []Config) (config Config) {

	if len(cfg) > 0 {
		config = cfg[0]
		mergo.Merge(&config, c)
	} else {
		_default := c
		config = _default
	}

	return
}

// MergeSingle merges the default with the given config and returns the result
func (c Config) MergeSingle(cfg Config) (config Config) {

	config = cfg
	mergo.Merge(&config, c)

	return
}
