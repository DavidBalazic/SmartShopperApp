package contextkeys

type ContextKey string

const (
    UserAgentKey ContextKey = "userAgent"
    IPKey        ContextKey = "ip"
)