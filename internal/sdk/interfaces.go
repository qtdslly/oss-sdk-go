package sdk

// Invalidator provides access to a type's invalidate method to make it
// invalidate it cache.
//
// e.g oss.SafeCredentialsProvider's Invalidate method.
type Invalidator interface {
	Invalidate()
}
