package dslgen

var Fluent = []byte("package fluent\n\n// Matcher is a main API group-level entry point.\n// It's used to define and configure the group rules.\n// It also represents a map of all rule-local variables.\ntype Matcher map[string]Var\n\n// Import loads given package path into a rule group imports table.\n//\n// That table is used during the rules compilation.\n//\n// The table has the following effect on the rules:\n//\t* For type expressions, it's used to resolve the\n//\t  full package paths of qualified types, like `foo.Bar`.\n//\t  If Import(`a/b/foo`) is called, `foo.Bar` will match\n//\t  `a/b/foo.Bar` type during the pattern execution.\nfunc (m Matcher) Import(pkgPath string) {}\n\n// Match specifies a set of patterns that match a rule being defined.\n// Pattern matching succeeds if at least 1 pattern matches.\n//\n// If none of the given patterns matched, rule execution stops.\nfunc (m Matcher) Match(pattern string, alternatives ...string) Matcher {\n\treturn m\n}\n\n// Where applies additional constraint to a match.\n// If a given cond is not satisfied, a match is rejected and\n// rule execution stops.\nfunc (m Matcher) Where(cond bool) Matcher {\n\treturn m\n}\n\n// Report prints a message if associated rule match is successful.\n//\n// A message is a string that can contain interpolated expressions.\n// For every matched variable it's possible to interpolate\n// their printed representation into the message text with $<name>.\n// An entire match can be addressed with $$.\nfunc (m Matcher) Report(message string) Matcher {\n\treturn m\n}\n\n// Suggest assigns a quickfix suggestion for the matched code.\nfunc (m Matcher) Suggest(suggestion string) Matcher {\n\treturn m\n}\n\n// At binds the reported node to a named submatch.\n// If no explicit location is given, the outermost node ($$) is used.\nfunc (m Matcher) At(v Var) Matcher {\n\treturn m\n}\n\n// Var is a pattern variable that describes a named submatch.\ntype Var struct {\n\t// Pure reports whether expr matched by var is side-effect-free.\n\tPure bool\n\n\t// Const reports whether expr matched by var is a constant value.\n\tConst bool\n\n\t// Addressable reports whether the corresponding expression is addressable.\n\t// See https://golang.org/ref/spec#Address_operators.\n\tAddressable bool\n\n\t// Type is a type of a matched expr.\n\tType ExprType\n\n\t// Test is a captured node text as in the source code.\n\tText MatchedText\n}\n\n// ExprType describes a type of a matcher expr.\ntype ExprType struct {\n\t// Size represents expression type size in bytes.\n\tSize int\n}\n\n// AssignableTo reports whether a type is assign-compatible with a given type.\n// See https://golang.org/pkg/go/types/#AssignableTo.\nfunc (ExprType) AssignableTo(typ string) bool { return boolResult }\n\n// ConvertibleTo reports whether a type is conversible to a given type.\n// See https://golang.org/pkg/go/types/#ConvertibleTo.\nfunc (ExprType) ConvertibleTo(typ string) bool { return boolResult }\n\n// Implements reports whether a type implements a given interface.\n// See https://golang.org/pkg/go/types/#Implements.\nfunc (ExprType) Implements(typ string) bool { return boolResult }\n\n// Is reports whether a type is identical to a given type.\nfunc (ExprType) Is(typ string) bool { return boolResult }\n\n// MatchedText represents a source text associated with a matched node.\ntype MatchedText string\n\n// Matches reports whether the text matches the given regexp pattern.\nfunc (MatchedText) Matches(pattern string) bool { return boolResult }\n\n\n\nvar boolResult bool\n\n")
