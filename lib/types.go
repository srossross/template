package lib

// Context struct is used to pass into a template as "."
type Context struct {
    Values map[string]interface{}
    Env map[string]string
}
