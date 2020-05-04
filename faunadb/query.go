package faunadb

// Event's action types. Usually used as a parameter for Insert or Remove functions.
//
// See: https://app.fauna.com/documentation/reference/queryapi#simple-type-events
const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionAdd    = "add"
	ActionRemove = "remove"
)

// Time unit. Usually used as a parameter for Time functions.
//
// See: https://app.fauna.com/documentation/reference/queryapi#epochnum-unit
const (
	TimeUnitDay         = "day"
	TimeUnitHalfDay     = "half day"
	TimeUnitHour        = "hour"
	TimeUnitMinute      = "minute"
	TimeUnitSecond      = "second"
	TimeUnitMillisecond = "millisecond"
	TimeUnitMicrosecond = "microsecond"
	TimeUnitNanosecond  = "nanosecond"
)

// Normalizers for Casefold
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
const (
	NormalizerNFKCCaseFold = "NFKCCaseFold"
	NormalizerNFC          = "NFC"
	NormalizerNFD          = "NFD"
	NormalizerNFKC         = "NFKC"
	NormalizerNFKD         = "NFKD"
)

// Helper functions

func varargs(expr ...interface{}) interface{} {
	if len(expr) == 1 {
		return expr[0]
	}

	return expr
}

// Optional parameters

// EventsOpt is an boolean optional parameter that describes if the query should include historical events.
// For more information about events, check https://app.fauna.com/documentation/reference/queryapi#simple-type-events.
//
// Functions that accept this optional parameter are: Paginate.
//
// Deprecated: The Events function was renamed to EventsOpt to support the new history API.
// EventsOpt is provided here for backwards compatibility. Instead of using Paginate with the EventsOpt parameter,
// you should use the new Events function.
func EventsOpt(events interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "events", wrap(events))
	}
}

// TS is a timestamp optional parameter that specifies in which timestamp a query should be executed.
//
// Functions that accept this optional parameter are: Get, Insert, Remove, Exists, and Paginate.
func TS(timestamp interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "ts", wrap(timestamp))
	}
}

// After is an optional parameter used when cursoring that refers to the specified cursor's the next page, inclusive.
// For more information about pages, check https://app.fauna.com/documentation/reference/queryapi#simple-type-pages.
//
// Functions that accept this optional parameter are: Paginate.
func After(ref interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "after", wrap(ref))
	}
}

// Before is an optional parameter used when cursoring that refers to the specified cursor's previous page, exclusive.
// For more information about pages, check https://app.fauna.com/documentation/reference/queryapi#simple-type-pages.
//
// Functions that accept this optional parameter are: Paginate.
func Before(ref interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "before", wrap(ref))
	}
}

// Size is a numeric optional parameter that specifies the size of a pagination cursor.
//
// Functions that accept this optional parameter are: Paginate.
func Size(size interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "size", wrap(size))
	}
}

// Start is a numeric optional parameter that specifies the start of where to search.
//
// Functions that accept this optional parameter are: FindStr and FindStrRegex.
func Start(start interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "start", wrap(start))
	}
}

// StrLength is a numeric optional parameter that specifies the amount to copy.
//
// Functions that accept this optional parameter are: FindStr and FindStrRegex.
func StrLength(length interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "length", wrap(length))
	}
}

func updateField(m optionalKeysMapping, field string, nv Expr) {
	*(m[field]) = nv
}

// OnlyFirst is a boolean optional parameter that only replace the first string
//
// Functions that accept this optional parameter are: ReplaceStrRegex
func OnlyFirst() OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "first", BooleanV(true))
	}
}

// Sources is a boolean optional parameter that specifies if a pagination cursor should include
// the source sets along with each element.
//
// Functions that accept this optional parameter are: Paginate.
func Sources(sources interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "sources", wrap(sources))
	}
}

// Default is an optional parameter that specifies the default value for a select operation when
// the desired value path is absent.
//
// Functions that accept this optional parameter are: Select.
func Default(value interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "default", wrap(value))
	}
}

// Separator is a string optional parameter that specifies the separator for a concat operation.
//
// Functions that accept this optional parameter are: Concat.
func Separator(sep interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "separator", wrap(sep))
	}
}

// Precision is an optional parameter that specifies the precision for a Trunc and Round operations.
//
// Functions that accept this optional parameter are: Round and Trunc.
func Precision(precision interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "precision", wrap(precision))
	}
}

// ConflictResolver is an optional parameter that specifies the lambda for resolving Merge conflicts
//
// Functions that accept this optional parameter are: Merge
func ConflictResolver(lambda interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "lambda", wrap(lambda))
	}
}

// Normalizer is a string optional parameter that specifies the normalization function for casefold operation.
//
// Functions that accept this optional parameter are: Casefold.
func Normalizer(norm interface{}) OptionalParameter {
	return func(m optionalKeysMapping) {
		updateField(m, "normalizer", wrap(norm))
	}
}

// LetBuilder builds Let expressions
type LetBuilder struct {
	bindings unescapedArr
}

// Bind binds a variable name to a value and returns a LetBuilder
func (lb *LetBuilder) Bind(key string, in interface{}) *LetBuilder {
	binding := make(unescapedObj, 1)
	binding[key] = wrap(in)
	lb.bindings = append(lb.bindings, binding)
	return lb
}

// In sets the expression to be evaluated and returns the prepared Let.
func (lb *LetBuilder) In(in Expr) Expr {
	return letFn{
		Let: wrap(lb.bindings),
		In:  in,
	}
}

// Values

// Ref creates a new RefV value with the provided ID.
//
// Parameters:
//  id string - A string representation of a reference type.
//
// Returns:
//  Ref - A new reference type.
//
// See: https://app.fauna.com/documentation/reference/queryapi#special-type
func Ref(id string) Expr { return legacyRefFn{Ref: wrap(id)} }

// RefClass creates a new Ref based on the provided class and ID.
//
// Parameters:
//  classRef Ref - A class reference.
//  id string|int64 - The document ID.
//
// Deprecated: Use RefCollection instead, RefClass is kept for backwards compatibility
//
// Returns:
//  Ref - A new reference type.
//
// See: https://app.fauna.com/documentation/reference/queryapi#special-type
func RefClass(classRef, id interface{}) Expr { return refFn{Ref: wrap(classRef), ID: wrap(id)} }

// RefCollection creates a new Ref based on the provided collection and ID.
//
// Parameters:
//  collectionRef Ref - A collection reference.
//  id string|int64 - The document ID.
//
// Returns:
//  Ref - A new reference type.
//
// See: https://app.fauna.com/documentation/reference/queryapi#special-type
func RefCollection(collectionRef, id interface{}) Expr {
	return refFn{Ref: wrap(collectionRef), ID: wrap(id)}
}

// Null creates a NullV value.
//
// Returns:
//  Value - A null value.
//
// See: https://app.fauna.com/documentation/reference/queryapi#simple-type
func Null() Expr { return NullV{} }

// Basic forms

// Abort aborts the execution of the query
//
// Parameters:
//  msg string - An error message.
//
// Returns:
//  Error
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Abort(msg interface{}) Expr { return abortFn{Abort: wrap(msg)} }

// Do sequentially evaluates its arguments, and returns the last expression.
// If no expressions are provided, do returns an error.
//
// Parameters:
//  exprs []Expr - A variable number of expressions to be evaluated.
//
// Returns:
//  Value - The result of the last expression in the list.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Do(exprs ...interface{}) Expr { return doFn{Do: wrap(varargs(exprs))} }

// If evaluates and returns then or elze depending on the value of cond.
// If cond evaluates to anything other than a boolean, if returns an “invalid argument” error
//
// Parameters:
//  cond bool - A boolean expression.
//  then Expr - The expression to run if condition is true.
//  elze Expr - The expression to run if condition is false.
//
// Returns:
//  Value - The result of either then or elze expression.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func If(cond, then, elze interface{}) Expr {
	return ifFn{
		If:   wrap(cond),
		Then: wrap(then),
		Else: wrap(elze),
	}
}

// Lambda creates an anonymous function. Mostly used with Collection functions.
//
// Parameters:
//  varName string|[]string - A string or an array of strings of arguments name to be bound in the body of the lambda.
//  expr Expr - An expression used as the body of the lambda.
//
// Returns:
//  Value - The result of the body expression.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Lambda(varName, expr interface{}) Expr {
	return lambdaFn{
		Lambda:     wrap(varName),
		Expression: wrap(expr),
	}
}

// At execute an expression at a given timestamp.
//
// Parameters:
//  timestamp time - The timestamp in which the expression will be evaluated.
//  expr Expr - An expression to be evaluated.
//
// Returns:
//  Value - The result of the given expression.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func At(timestamp, expr interface{}) Expr {
	return atFn{
		At:         wrap(timestamp),
		Expression: wrap(expr),
	}
}

// Let binds values to one or more variables.
//
// Returns:
//  *LetBuilder - Returns a LetBuilder.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Let() *LetBuilder { return &LetBuilder{nil} }

// Var refers to a value of a variable on the current lexical scope.
//
// Parameters:
//  name string - The variable name.
//
// Returns:
//  Value - The variable value.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Var(name string) Expr { return varFn{Var: wrap(name)} }

// Call invokes the specified function passing in a variable number of arguments
//
// Parameters:
//  ref Ref - The reference to the user defined functions to call.
//  args []Value - A series of values to pass as arguments to the user defined function.
//
// Returns:
//  Value - The return value of the user defined function.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Call(ref interface{}, args ...interface{}) Expr {
	return callFn{Call: wrap(ref), Params: wrap(varargs(args...))}
}

// Query creates an instance of the @query type with the specified lambda
//
// Parameters:
//  lambda Lambda - A lambda representation. See Lambda() function.
//
// Returns:
//  Query - The lambda wrapped in a @query type.
//
// See: https://app.fauna.com/documentation/reference/queryapi#basic-forms
func Query(lambda interface{}) Expr { return queryFn{Query: wrap(lambda)} }

// Collections

// Map applies the lambda expression on each element of a collection or Page.
// It returns the result of each application on a collection of the same type.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. See Lambda() function.
//
// Returns:
//  []Value - A new collection with elements transformed by the lambda function.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Map(coll, lambda interface{}) Expr { return mapFn{Map: wrap(lambda), Collection: wrap(coll)} }

// Foreach applies the lambda expression on each element of a collection or Page.
// The original collection is returned.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. See Lambda() function.
//
// Returns:
//  []Value - The original collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Foreach(coll, lambda interface{}) Expr {
	return foreachFn{Foreach: wrap(lambda), Collection: wrap(coll)}
}

// Filter applies the lambda expression on each element of a collection or Page.
// It returns a new collection of the same type containing only the elements in which the
// function application returned true.
//
// Parameters:
//  coll []Value - The collection of elements to iterate.
//  lambda Lambda - A lambda function to be applied to each element of the collection. The lambda function must return a boolean value. See Lambda() function.
//
// Returns:
//  []Value - A new collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Filter(coll, lambda interface{}) Expr {
	return filterFn{Filter: wrap(lambda), Collection: wrap(coll)}
}

// Take returns a new collection containing num elements from the head of the original collection.
//
// Parameters:
//  num int64 - The number of elements to take from the collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Take(num, coll interface{}) Expr { return takeFn{Take: wrap(num), Collection: wrap(coll)} }

// Drop returns a new collection containing the remaining elements from the original collection
// after num elements have been removed.
//
// Parameters:
//  num int64 - The number of elements to drop from the collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Drop(num, coll interface{}) Expr { return dropFn{Drop: wrap(num), Collection: wrap(coll)} }

// Prepend returns a new collection that is the result of prepending elems to coll.
//
// Parameters:
//  elems []Value - Elements to add to the beginning of the other collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Prepend(elems, coll interface{}) Expr {
	return prependFn{Prepend: wrap(elems), Collection: wrap(coll)}
}

// Append returns a new collection that is the result of appending elems to coll.
//
// Parameters:
//  elems []Value - Elements to add to the end of the other collection.
//  coll []Value - The collection of elements.
//
// Returns:
//  []Value - A new collection.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func Append(elems, coll interface{}) Expr {
	return appendFn{Append: wrap(elems), Collection: wrap(coll)}
}

// IsEmpty returns true if the collection is the empty set, else false.
//
// Parameters:
//  coll []Value - The collection of elements.
//
// Returns:
//   bool - True if the collection is empty, else false.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func IsEmpty(coll interface{}) Expr { return isEmptyFn{IsEmpty: wrap(coll)} }

// IsNonEmpty returns false if the collection is the empty set, else true
//
// Parameters:
//  coll []Value - The collection of elements.
//
// Returns:
//   bool - True if the collection is not empty, else false.
//
// See: https://app.fauna.com/documentation/reference/queryapi#collections
func IsNonEmpty(coll interface{}) Expr { return isNonEmptyFn{IsNonEmpty: wrap(coll)} }

// Read

// Get retrieves the document identified by the provided ref. Optional parameters: TS.
//
// Parameters:
//  ref Ref|SetRef - The reference to the object or a set reference.
//
// Optional parameters:
//  ts time - The snapshot time at which to get the document. See TS() function.
//
// Returns:
//  Object - The object requested.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func Get(ref interface{}, options ...OptionalParameter) Expr {
	fn := getFn{Get: wrap(ref)}
	mappings := optionalKeysMapping{
		"ts": &(fn.TS),
	}
	applyOptionals(mappings, options)
	return fn
}

// KeyFromSecret retrieves the key object from the given secret.
//
// Parameters:
//  secret string - The token secret.
//
// Returns:
//  Key - The key object related to the token secret.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func KeyFromSecret(secret interface{}) Expr { return keyFromSecretFn{KeyFromSecret: wrap(secret)} }

// Exists returns boolean true if the provided ref exists (in the case of an document),
// or is non-empty (in the case of a set), and false otherwise. Optional parameters: TS.
//
// Parameters:
//  ref Ref - The reference to the object. It could be a document reference of a object reference like a collection.
//
// Optional parameters:
//  ts time - The snapshot time at which to check for the document's existence. See TS() function.
//
// Returns:
//  bool - true if the reference exists, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func Exists(ref interface{}, options ...OptionalParameter) Expr {
	fn := existsFn{Exists: wrap(ref)}
	mappings := optionalKeysMapping{
		"ts": &(fn.TS),
	}
	applyOptionals(mappings, options)
	return fn

}

// Paginate retrieves a page from the provided set.
//
// Parameters:
//  set SetRef - A set reference to paginate over. See Match() or MatchTerm() functions.
//
// Optional parameters:
//  after Cursor - Return the next page of results after this cursor (inclusive). See After() function.
//  before Cursor - Return the previous page of results before this cursor (exclusive). See Before() function.
//  sources bool - If true, include the source sets along with each element. See Sources() function.
//  ts time - The snapshot time at which to get the document. See TS() function.
//
// Returns:
//  Page - A page of elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func Paginate(set interface{}, options ...OptionalParameter) Expr {
	fn := paginateFn{Paginate: wrap(set)}
	mappings := optionalKeysMapping{
		"after":   &(fn.After),
		"before":  &(fn.Before),
		"events":  &(fn.Events),
		"size":    &(fn.Size),
		"sources": &(fn.Sources),
		"ts":      &(fn.TS),
	}
	applyOptionals(mappings, options)
	return fn
}

// Write

// Create creates an document of the specified collection.
//
// Parameters:
//  ref Ref - A collection reference.
//  params Object - An object with attributes of the document created.
//
// Returns:
//  Object - A new document of the collection referenced.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Create(ref, params interface{}) Expr { return createFn{Create: wrap(ref), Params: wrap(params)} }

// CreateClass creates a new class.
//
// Parameters:
//  params Object - An object with attributes of the class.
//
// Deprecated: Use CreateCollection instead, CreateClass is kept for backwards compatibility
//
// Returns:
//  Object - The new created class object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateClass(params interface{}) Expr { return createClassFn{CreateClass: wrap(params)} }

// CreateCollection creates a new collection.
//
// Parameters:
//  params Object - An object with attributes of the collection.
//
// Returns:
//  Object - The new created collection object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateCollection(params interface{}) Expr {
	return createCollectionFn{CreateCollection: wrap(params)}
}

// CreateDatabase creates an new database.
//
// Parameters:
//  params Object - An object with attributes of the database.
//
// Returns:
//  Object - The new created database object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateDatabase(params interface{}) Expr { return createDatabaseFn{CreateDatabase: wrap(params)} }

// CreateIndex creates a new index.
//
// Parameters:
//  params Object - An object with attributes of the index.
//
// Returns:
//  Object - The new created index object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateIndex(params interface{}) Expr { return createIndexFn{CreateIndex: wrap(params)} }

// CreateKey creates a new key.
//
// Parameters:
//  params Object - An object with attributes of the key.
//
// Returns:
//  Object - The new created key object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateKey(params interface{}) Expr { return createKeyFn{CreateKey: wrap(params)} }

// CreateFunction creates a new function.
//
// Parameters:
//  params Object - An object with attributes of the function.
//
// Returns:
//  Object - The new created function object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateFunction(params interface{}) Expr { return createFunctionFn{CreateFunction: wrap(params)} }

// CreateRole creates a new role.
//
// Parameters:
//  params Object - An object with attributes of the role.
//
// Returns:
//  Object - The new created role object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func CreateRole(params interface{}) Expr { return createRoleFn{CreateRole: wrap(params)} }

// MoveDatabase moves a database to a new hierachy.
//
// Parameters:
//  from Object - Source reference to be moved.
//  to Object   - New parent database reference.
//
// Returns:
//  Object - instance.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func MoveDatabase(from interface{}, to interface{}) Expr {
	return moveDatabaseFn{MoveDatabase: wrap(from), To: wrap(to)}
}

// Update updates the provided document.
//
// Parameters:
//  ref Ref - The reference to update.
//  params Object - An object representing the parameters of the document.
//
// Returns:
//  Object - The updated object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Update(ref, params interface{}) Expr { return updateFn{Update: wrap(ref), Params: wrap(params)} }

// Replace replaces the provided document.
//
// Parameters:
//  ref Ref - The reference to replace.
//  params Object - An object representing the parameters of the document.
//
// Returns:
//  Object - The replaced object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Replace(ref, params interface{}) Expr { return replaceFn{Replace: wrap(ref), Params: wrap(params)} }

// Delete deletes the provided document.
//
// Parameters:
//  ref Ref - The reference to delete.
//
// Returns:
//  Object - The deleted object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Delete(ref interface{}) Expr { return deleteFn{Delete: wrap(ref)} }

// Insert adds an event to the provided document's history.
//
// Parameters:
//  ref Ref - The reference to insert against.
//  ts time - The valid time of the inserted event.
//  action string - Whether the event shoulde be a ActionCreate, ActionUpdate or ActionDelete.
//
// Returns:
//  Object - The deleted object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Insert(ref, ts, action, params interface{}) Expr {
	return insertFn{Insert: wrap(ref), Ts: wrap(ts), Action: wrap(action), Params: wrap(params)}
}

// Remove deletes an event from the provided document's history.
//
// Parameters:
//  ref Ref - The reference of the document whose event should be removed.
//  ts time - The valid time of the inserted event.
//  action string - The event action (ActionCreate, ActionUpdate or ActionDelete) that should be removed.
//
// Returns:
//  Object - The deleted object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#write-functions
func Remove(ref, ts, action interface{}) Expr {
	return removeFn{Remove: wrap(ref), Ts: wrap(ts), Action: wrap(action)}
}

// String

// Format formats values into a string.
//
// Parameters:
//  format string - format a string with format specifiers.
//
// Optional parameters:
//  values []string - list of values to format into string.
//
// Returns:
//  string - A string.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Format(format interface{}, values ...interface{}) Expr {
	return formatFn{Format: wrap(format), Values: wrap(varargs(values...))}
}

// Concat concatenates a list of strings into a single string.
//
// Parameters:
//  terms []string - A list of strings to concatenate.
//
// Optional parameters:
//  separator string - The separator to use between each string. See Separator() function.
//
// Returns:
//  string - A string with all terms concatenated.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Concat(terms interface{}, options ...OptionalParameter) Expr {
	fn := concatFn{Concat: wrap(terms)}
	mappings := optionalKeysMapping{
		"separator": &(fn.Separator),
	}
	applyOptionals(mappings, options)
	return fn
}

// Casefold normalizes strings according to the Unicode Standard section 5.18 "Case Mappings".
//
// Parameters:
//  str string - The string to casefold.
//
// Optional parameters:
//  normalizer string - The algorithm to use. One of: NormalizerNFKCCaseFold, NormalizerNFC, NormalizerNFD, NormalizerNFKC, NormalizerNFKD.
//
// Returns:
//  string - The normalized string.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Casefold(str interface{}, options ...OptionalParameter) Expr {
	fn := casefoldFn{Casefold: wrap(str)}
	mappings := optionalKeysMapping{
		"normalizer": &(fn.Normalizer),
	}
	applyOptionals(mappings, options)
	return fn
}

// StartsWith returns true if the string starts with the given prefix value, or false if otherwise
//
// Parameters:
//
//  value  string -  the string to evaluate
//  search string -  the prefix to search for
//
// Returns:
//   boolean       - does `value` start with `search
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/startswith
func StartsWith(value interface{}, search interface{}) Expr {
	return startsWithFn{StartsWith: wrap(value), Search: wrap(search)}
}

// EndsWith returns true if the string ends with the given suffix value, or false if otherwise
//
// Parameters:
//
// value  string  -  the string to evaluate
// search  string -  the suffix to search for
//
// Returns:
// boolean       - does `value` end with `search`
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/endswith
func EndsWith(value interface{}, search interface{}) Expr {
	return endsWithFn{EndsWith: wrap(value), Search: wrap(search)}
}

// ContainsStr returns true if the string contains the given substring, or false if otherwise
//
// Parameters:
//
// value string  -  the string to evaluate
// search string -  the substring to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/containsstr
func ContainsStr(value interface{}, search interface{}) Expr {
	return containsStrFn{ContainsStr: wrap(value), Search: wrap(search)}
}

// ContainsStrRegex returns true if the string contains the given pattern, or false if otherwise
//
// Parameters:
//
// value   string      -  the string to evaluate
// pattern string      -  the pattern to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/containsstrregex
func ContainsStrRegex(value interface{}, pattern interface{}) Expr {
	return containsStrRegexFn{ContainsStrRegex: wrap(value), Pattern: wrap(pattern)}
}

// RegexEscape It takes a string and returns a regex which matches the input string verbatim.
//
// Parameters:
//
// value  string     - the string to analyze
// pattern       -  the pattern to search for
//
// Returns:
// boolean      - was the search result found
//
// See https://docs.fauna.com/fauna/current/api/fql/functions/regexescape
func RegexEscape(value interface{}) Expr {
	return regexEscapeFn{RegexEscape: wrap(value)}
}

// FindStr locates a substring in a source string.  Optional parameters: Start
//
// Parameters:
//  str string  - The source string
//  find string - The string to locate
//
// Optional parameters:
//  start int - a position to start the search. See Start() function.
//
// Returns:
//  string - The offset of where the substring starts or -1 if not found
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func FindStr(str, find interface{}, options ...OptionalParameter) Expr {
	fn := findStrFn{FindStr: wrap(str), Find: wrap(find)}
	mappings := optionalKeysMapping{
		"start": &(fn.Start),
	}
	applyOptionals(mappings, options)
	return fn
}

// FindStrRegex locates a java regex pattern in a source string.  Optional parameters: Start
//
// Parameters:
//  str string      - The sourcestring
//  pattern string  - The pattern to locate.
//
// Optional parameters:
//  start long - a position to start the search.  See Start() function.
//
// Returns:
//  string - The offset of where the substring starts or -1 if not found
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func FindStrRegex(str, pattern interface{}, options ...OptionalParameter) Expr {
	return findStrRegexFn{FindStrRegex: wrap(str), Pattern: wrap(pattern)}
}

// Length finds the length of a string in codepoints
//
// Parameters:
//  str string - A string to find the length in codepoints
//
// Returns:
//  int - A length of a string.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Length(str interface{}) Expr { return lengthFn{Length: wrap(str)} }

// LowerCase changes all characters in the string to lowercase
//
// Parameters:
//  str string - A string to convert to lowercase
//
// Returns:
//  string - A string in lowercase.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func LowerCase(str interface{}) Expr { return lowercaseFn{Lowercase: wrap(str)} }

// LTrim returns a string wtih leading white space removed.
//
// Parameters:
//  str string - A string to remove leading white space
//
// Returns:
//  string - A string with all leading white space removed
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func LTrim(str interface{}) Expr { return lTrimFn{LTrim: wrap(str)} }

// Repeat returns a string wtih repeated n times
//
// Parameters:
//  str string - A string to repeat
//  number int - The number of times to repeat the string
//
// Returns:
//  string - A string concatendanted the specified number of times
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Repeat(str, number interface{}) Expr { return repeatFn{Repeat: wrap(str), Number: wrap(number)} }

// ReplaceStr returns a string with every occurence of the "find" string changed to "replace" string
//
// Parameters:
//  str string     - A source string
//  find string    - The substring to locate in in the source string
//  replace string - The string to replaice the "find" string when located
//
// Returns:
//  string - returns a string with every occurence of the "find" string changed to "replace"
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func ReplaceStr(str, find, replace interface{}) Expr {
	return replaceStrFn{
		ReplaceStr: wrap(str),
		Find:       wrap(find),
		Replace:    wrap(replace),
	}
}

// ReplaceStrRegex returns a string with occurence(s) of the java regular expression "pattern" changed to "replace" string.   Optional parameters: OnlyFirst
//
// Parameters:
//  value string   - The source string
//  pattern string - A java regular expression to locate
//  replace string - The string to replace the pattern when located
//
// Optional parameters:
//  OnlyFirst - Only replace the first found pattern.  See OnlyFirst() function.
//
// Returns:
//  string - A string with occurence(s) of the java regular expression "pattern" changed to "replace" string
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func ReplaceStrRegex(value, pattern, replace interface{}, options ...OptionalParameter) Expr {
	fn := replaceStrRegexFn{
		ReplaceStrRegex: wrap(value),
		Pattern:         wrap(pattern),
		Replace:         wrap(replace),
	}
	mappings := optionalKeysMapping{
		"first": &(fn.First),
	}
	applyOptionals(mappings, options)
	return fn
}

// RTrim returns a string wtih trailing white space removed.
//
// Parameters:
//  str string - A string to remove trailing white space
//
// Returns:
//  string - A string with all trailing white space removed
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func RTrim(str interface{}) Expr { return rTrimFn{RTrim: wrap(str)} }

// Space function returns "N" number of spaces
//
// Parameters:
//  value int - the number of spaces
//
// Returns:
//  string - function returns string with n spaces
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Space(value interface{}) Expr { return spaceFn{Space: wrap(value)} }

// SubString returns a subset of the source string.   Optional parameters: StrLength
//
// Parameters:
//  str string - A source string
//  start int  - The position in the source string where SubString starts extracting characters
//
// Optional parameters:
//  StrLength int - A value for the length of the extracted substring. See StrLength() function.
//
// Returns:
//  string - function returns a subset of the source string
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func SubString(str, start interface{}, options ...OptionalParameter) Expr {
	fn := subStringFn{SubString: wrap(str), Start: wrap(start)}
	mappings := optionalKeysMapping{
		"length": &(fn.Length),
	}
	applyOptionals(mappings, options)
	return fn
}

// TitleCase changes all characters in the string to TitleCase
//
// Parameters:
//  str string - A string to convert to TitleCase
//
// Returns:
//  string - A string in TitleCase.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func TitleCase(str interface{}) Expr { return titleCaseFn{Titlecase: wrap(str)} }

// Trim returns a string wtih trailing white space removed.
//
// Parameters:
//  str string - A string to remove trailing white space
//
// Returns:
//  string - A string with all trailing white space removed
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func Trim(str interface{}) Expr { return trimFn{Trim: wrap(str)} }

// UpperCase changes all characters in the string to uppercase
//
// Parameters:
//  string - A string to convert to uppercase
//
// Returns:
//  string - A string in uppercase.
//
// See: https://app.fauna.com/documentation/reference/queryapi#string-functions
func UpperCase(str interface{}) Expr { return upperCaseFn{UpperCase: wrap(str)} }

// Time and Date

// Time constructs a time from a ISO 8601 offset date/time string.
//
// Parameters:
//  str string - A string to convert to a time object.
//
// Returns:
//  time - A time object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#time-and-date
func Time(str interface{}) Expr { return timeFn{Time: wrap(str)} }

// TimeAdd returns a new time or date with the offset in terms of the unit
// added.
//
// Parameters:
// base        -  the base time or data
// offset      -  the number of units
// unit        -  the unit type
//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timeadd
func TimeAdd(base interface{}, offset interface{}, unit interface{}) Expr {
	return timeAddFn{TimeAdd: wrap(base), Offset: wrap(offset), Unit: wrap(unit)}
}

// TimeSubtract returns a new time or date with the offset in terms of the unit
// subtracted.
//
// Parameters:
// base        -  the base time or data
// offset      -  the number of units
// unit        -  the unit type
//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timesubtract
func TimeSubtract(base interface{}, offset interface{}, unit interface{}) Expr {
	return timeSubtractFn{TimeSubtract: wrap(base), Offset: wrap(offset), Unit: wrap(unit)}
}

// TimeDiff returns the number of intervals in terms of the unit between
// two times or dates. Both start and finish must be of the same
// type.
//
// Parameters:
//   start the starting time or date, inclusive
//   finish the ending time or date, exclusive
//   unit the unit type//
// Returns:
// Expr
//
//See: https://docs.fauna.com/fauna/current/api/fql/functions/timediff
func TimeDiff(start interface{}, finish interface{}, unit interface{}) Expr {
	return timeDiffFn{TimeDiff: wrap(start), Other: wrap(finish), Unit: wrap(unit)}
}

// Date constructs a date from a ISO 8601 offset date/time string.
//
// Parameters:
//  str string - A string to convert to a date object.
//
// Returns:
//  date - A date object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#time-and-date
func Date(str interface{}) Expr { return dateFn{Date: wrap(str)} }

// Epoch constructs a time relative to the epoch "1970-01-01T00:00:00Z".
//
// Parameters:
//  num int64 - The number of units from Epoch.
//  unit string - The unit of number. One of TimeUnitSecond, TimeUnitMillisecond, TimeUnitMicrosecond, TimeUnitNanosecond.
//
// Returns:
//  time - A time object.
//
// See: https://app.fauna.com/documentation/reference/queryapi#time-and-date
func Epoch(num, unit interface{}) Expr { return epochFn{Epoch: wrap(num), Unit: wrap(unit)} }

// Now returns the current snapshot time.
//
// Returns:
// Expr
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/now
func Now() Expr {
	return nowFn{Now: NullV{}}
}

// Set

// Singleton returns the history of the document's presence of the provided ref.
//
// Parameters:
//  ref Ref - The reference of the document for which to retrieve the singleton set.
//
// Returns:
//  SetRef - The singleton SetRef.
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Singleton(ref interface{}) Expr { return singletonFn{Singleton: wrap(ref)} }

// Events returns the history of document's data of the provided ref.
//
// Parameters:
//  refSet Ref|SetRef - A reference or set reference to retrieve an event set from.
//
// Returns:
//  SetRef - The events SetRef.
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Events(refSet interface{}) Expr { return eventsFn{Events: wrap(refSet)} }

// Match returns the set of documents for the specified index.
//
// Parameters:
//  ref Ref - The reference of the index to match against.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Match(ref interface{}) Expr { return matchFn{Match: wrap(ref), Terms: nil} }

// MatchTerm returns th set of documents that match the terms in an index.
//
// Parameters:
//  ref Ref - The reference of the index to match against.
//  terms []Value - A list of terms used in the match.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func MatchTerm(ref, terms interface{}) Expr { return matchFn{Match: wrap(ref), Terms: wrap(terms)} }

// Union returns the set of documents that are present in at least one of the specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to union together.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Union(sets ...interface{}) Expr { return unionFn{Union: wrap(varargs(sets...))} }

// Merge two or more objects..
//
// Parameters:
//   merge merge the first object.
//   with the second object or a list of objects
//   lambda a lambda to resolve possible conflicts
//
// Returns:
// merged object
//
func Merge(merge interface{}, with interface{}, lambda ...OptionalParameter) Expr {
	fn := mergeFn{Merge: wrap(merge), With: wrap(with)}
	mappings := optionalKeysMapping{
		"lambda": &(fn.Lambda),
	}
	applyOptionals(mappings, lambda)
	return fn
}

// Reduce function applies a reducer Lambda function serially to each member of the collection to produce a single value.
//
// Parameters:
// lambda     Expr  - The accumulator function
// initial    Expr  - The initial value
// collection Expr  - The collection to be reduced
//
// Returns:
// Expr
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/reduce
func Reduce(lambda, initial interface{}, collection interface{}) Expr {
	return reduceFn{
		Reduce:     wrap(lambda),
		Initial:    wrap(initial),
		Collection: wrap(collection),
	}
}

// Intersection returns the set of documents that are present in all of the specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to intersect.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Intersection(sets ...interface{}) Expr {
	return intersectionFn{Intersection: wrap(varargs(sets...))}
}

// Difference returns the set of documents that are present in the first set but not in
// any of the other specified sets.
//
// Parameters:
//  sets []SetRef - A list of SetRef to diff.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Difference(sets ...interface{}) Expr { return differenceFn{Difference: wrap(varargs(sets...))} }

// Distinct returns the set of documents with duplicates removed.
//
// Parameters:
//  set []SetRef - A list of SetRef to remove duplicates from.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Distinct(set interface{}) Expr { return distinctFn{Distinct: wrap(set)} }

// Join derives a set of resources by applying each document in the source set to the target set.
//
// Parameters:
//  source SetRef - A SetRef of the source set.
//  target Lambda - A Lambda that will accept each element of the source Set and return a Set.
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Join(source, target interface{}) Expr { return joinFn{Join: wrap(source), With: wrap(target)} }

// Range filters the set based on the lower/upper bounds (inclusive).
//
// Parameters:
//  set SetRef - Set to be filtered.
//  from - lower bound.
//  to - upper bound
//
// Returns:
//  SetRef
//
// See: https://app.fauna.com/documentation/reference/queryapi#sets
func Range(set interface{}, from interface{}, to interface{}) Expr {
	return rangeFn{Range: wrap(set), From: wrap(from), To: wrap(to)}
}

// Authentication

// Login creates a token for the provided ref.
//
// Parameters:
//  ref Ref - A reference with credentials to authenticate against.
//  params Object - An object of parameters to pass to the login function
//    - password: The password used to login
//
// Returns:
//  Key - a key with the secret to login.
//
// See: https://app.fauna.com/documentation/reference/queryapi#authentication
func Login(ref, params interface{}) Expr {
	return loginFn{Login: wrap(ref), Params: wrap(params)}
}

// Logout deletes the current session token. If invalidateAll is true, logout will delete all tokens associated with the current session.
//
// Parameters:
//  invalidateAll bool - If true, log out all tokens associated with the current session.
//
// See: https://app.fauna.com/documentation/reference/queryapi#authentication
func Logout(invalidateAll interface{}) Expr { return logoutFn{Logout: wrap(invalidateAll)} }

// Identify checks the given password against the provided ref's credentials.
//
// Parameters:
//  ref Ref - The reference to check the password against.
//  password string - The credentials password to check.
//
// Returns:
//  bool - true if the password is correct, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#authentication
func Identify(ref, password interface{}) Expr {
	return identifyFn{Identify: wrap(ref), Password: wrap(password)}
}

// Identity returns the document reference associated with the current key.
//
// For example, the current key token created using:
//	Create(Tokens(), Obj{"document": someRef})
// or via:
//	Login(someRef, Obj{"password":"sekrit"})
// will return "someRef" as the result of this function.
//
// Returns:
//  Ref - The reference associated with the current key.
//
// See: https://app.fauna.com/documentation/reference/queryapi#authentication
func Identity() Expr { return identityFn{Identity: NullV{}} }

// HasIdentity checks if the current key has an identity associated to it.
//
// Returns:
//  bool - true if the current key has an identity, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#authentication
func HasIdentity() Expr { return hasIdentityFn{HasIdentity: NullV{}} }

// Miscellaneous

// NextID produces a new identifier suitable for use when constructing refs.
//
// Deprecated: Use NewId instead
//
// Returns:
//  string - The new ID.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func NextID() Expr { return nextIdFn{NextID: NullV{}} }

// NewId produces a new identifier suitable for use when constructing refs.
//
// Returns:
//  string - The new ID.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func NewId() Expr { return newIdFn{NewID: NullV{}} }

// Database creates a new database ref.
//
// Parameters:
//  name string - The name of the database.
//
// Returns:
//  Ref - The database reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Database(name interface{}) Expr { return databaseFn{Database: wrap(name), Scope: nil} }

// ScopedDatabase creates a new database ref inside a database.
//
// Parameters:
//  name string - The name of the database.
//  scope Ref - The reference of the database's scope.
//
// Returns:
//  Ref - The database reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedDatabase(name interface{}, scope interface{}) Expr {
	return databaseFn{
		Database: wrap(name),
		Scope:    wrap(scope),
	}
}

// Index creates a new index ref.
//
// Parameters:
//  name string - The name of the index.
//
// Returns:
//  Ref - The index reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Index(name interface{}) Expr { return indexFn{Index: wrap(name)} }

// ScopedIndex creates a new index ref inside a database.
//
// Parameters:
//  name string - The name of the index.
//  scope Ref - The reference of the index's scope.
//
// Returns:
//  Ref - The index reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedIndex(name interface{}, scope interface{}) Expr {
	return indexFn{
		Index: wrap(name),
		Scope: wrap(scope),
	}
}

// Class creates a new class ref.
//
// Parameters:
//  name string - The name of the class.
//
// Deprecated: Use Collection instead, Class is kept for backwards compatibility
//
// Returns:
//  Ref - The class reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Class(name interface{}) Expr { return classFn{Class: wrap(name)} }

// Collection creates a new collection ref.
//
// Parameters:
//  name string - The name of the collection.
//
// Returns:
//  Ref - The collection reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Collection(name interface{}) Expr { return collectionFn{Collection: wrap(name)} }

//Documents returns a set of all documents in the given collection.
// A set must be paginated in order to retrieve its values.
//
// Parameters:
// collection  ref  - a reference to the collection
//
// Returns:
// Expr  - A new Expr instance
//
// See:  https://docs.fauna.com/fauna/current/api/fql/functions/Documents
func Documents(collection interface{}) Expr {
	return documentsFn{Documents: wrap(collection)}
}

// ScopedClass creates a new class ref inside a database.
//
// Parameters:
//  name string - The name of the class.
//  scope Ref - The reference of the class's scope.
//
// Deprecated: Use ScopedCollection instead, ScopedClass is kept for backwards compatibility
//
// Returns:
//  Ref - The collection reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedClass(name interface{}, scope interface{}) Expr {
	return classFn{Class: wrap(name), Scope: wrap(scope)}
}

// ScopedCollection creates a new collection ref inside a database.
//
// Parameters:
//  name string - The name of the collection.
//  scope Ref - The reference of the collection's scope.
//
// Returns:
//  Ref - The collection reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedCollection(name interface{}, scope interface{}) Expr {
	return collectionFn{Collection: wrap(name), Scope: wrap(scope)}
}

// Function create a new function ref.
//
// Parameters:
//  name string - The name of the functions.
//
// Returns:
//  Ref - The function reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Function(name interface{}) Expr { return functionFn{Function: wrap(name)} }

// ScopedFunction creates a new function ref inside a database.
//
// Parameters:
//  name string - The name of the function.
//  scope Ref - The reference of the function's scope.
//
// Returns:
//  Ref - The function reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedFunction(name interface{}, scope interface{}) Expr {
	return functionFn{Function: wrap(name), Scope: wrap(scope)}
}

// Role create a new role ref.
//
// Parameters:
//  name string - The name of the role.
//
// Returns:
//  Ref - The role reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Role(name interface{}) Expr { return roleFn{Role: wrap(name)} }

// ScopedRole create a new role ref.
//
// Parameters:
//  name string - The name of the role.
//  scope Ref - The reference of the role's scope.
//
// Returns:
//  Ref - The role reference.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedRole(name, scope interface{}) Expr { return roleFn{Role: wrap(name), Scope: wrap(scope)} }

// Classes creates a native ref for classes.
//
// Deprecated: Use Collections instead, Classes is kept for backwards compatibility
//
// Returns:
//  Ref - The reference of the class set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Classes() Expr { return classesFn{Classes: NullV{}} }

// Collections creates a native ref for collections.
//
// Returns:
//  Ref - The reference of the collections set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Collections() Expr { return collectionsFn{Collections: NullV{}} }

// ScopedClasses creates a native ref for classes inside a database.
//
// Parameters:
//  scope Ref - The reference of the class set's scope.
//
// Deprecated: Use ScopedCollections instead, ScopedClasses is kept for backwards compatibility
//
// Returns:
//  Ref - The reference of the class set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedClasses(scope interface{}) Expr { return classesFn{Classes: wrap(scope)} }

// ScopedCollections creates a native ref for collections inside a database.
//
// Parameters:
//  scope Ref - The reference of the collections set's scope.
//
// Returns:
//  Ref - The reference of the collections set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedCollections(scope interface{}) Expr {
	return collectionsFn{Collections: wrap(scope)}
}

// Indexes creates a native ref for indexes.
//
// Returns:
//  Ref - The reference of the index set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Indexes() Expr { return indexesFn{Indexes: NullV{}} }

// ScopedIndexes creates a native ref for indexes inside a database.
//
// Parameters:
//  scope Ref - The reference of the index set's scope.
//
// Returns:
//  Ref - The reference of the index set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedIndexes(scope interface{}) Expr { return indexesFn{Indexes: wrap(scope)} }

// Databases creates a native ref for databases.
//
// Returns:
//  Ref - The reference of the datbase set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Databases() Expr { return databasesFn{Databases: NullV{}} }

// ScopedDatabases creates a native ref for databases inside a database.
//
// Parameters:
//  scope Ref - The reference of the database set's scope.
//
// Returns:
//  Ref - The reference of the database set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedDatabases(scope interface{}) Expr { return databasesFn{Databases: wrap(scope)} }

// Functions creates a native ref for functions.
//
// Returns:
//  Ref - The reference of the function set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Functions() Expr { return functionsFn{Functions: NullV{}} }

// ScopedFunctions creates a native ref for functions inside a database.
//
// Parameters:
//  scope Ref - The reference of the function set's scope.
//
// Returns:
//  Ref - The reference of the function set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedFunctions(scope interface{}) Expr { return functionsFn{Functions: wrap(scope)} }

// Roles creates a native ref for roles.
//
// Returns:
//  Ref - The reference of the roles set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Roles() Expr { return rolesFn{Roles: NullV{}} }

// ScopedRoles creates a native ref for roles inside a database.
//
// Parameters:
//  scope Ref - The reference of the role set's scope.
//
// Returns:
//  Ref - The reference of the role set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedRoles(scope interface{}) Expr { return rolesFn{Roles: wrap(scope)} }

// Keys creates a native ref for keys.
//
// Returns:
//  Ref - The reference of the key set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Keys() Expr { return keysFn{Keys: NullV{}} }

// ScopedKeys creates a native ref for keys inside a database.
//
// Parameters:
//  scope Ref - The reference of the key set's scope.
//
// Returns:
//  Ref - The reference of the key set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedKeys(scope interface{}) Expr { return keysFn{Keys: wrap(scope)} }

// Tokens creates a native ref for tokens.
//
// Returns:
//  Ref - The reference of the token set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Tokens() Expr { return tokensFn{Tokens: NullV{}} }

// ScopedTokens creates a native ref for tokens inside a database.
//
// Parameters:
//  scope Ref - The reference of the token set's scope.
//
// Returns:
//  Ref - The reference of the token set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedTokens(scope interface{}) Expr { return tokensFn{Tokens: wrap(scope)} }

// Credentials creates a native ref for credentials.
//
// Returns:
//  Ref - The reference of the credential set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Credentials() Expr { return credentialsFn{Credentials: NullV{}} }

// ScopedCredentials creates a native ref for credentials inside a database.
//
// Parameters:
//  scope Ref - The reference of the credential set's scope.
//
// Returns:
//  Ref - The reference of the credential set.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func ScopedCredentials(scope interface{}) Expr {
	return credentialsFn{Credentials: wrap(scope)}
}

// Equals checks if all args are equivalents.
//
// Parameters:
//  args []Value - A collection of expressions to check for equivalence.
//
// Returns:
//  bool - true if all elements are equals, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Equals(args ...interface{}) Expr { return equalsFn{Equals: wrap(varargs(args...))} }

// Contains checks if the provided value contains the path specified.
//
// Parameters:
//  path Path - An array representing a path to check for the existence of. Path can be either strings or ints.
//  value Object - An object to search against.
//
// Returns:
//  bool - true if the path contains any value, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Contains(path, value interface{}) Expr {
	return containsFn{Contains: wrap(path), Value: wrap(value)}
}

// Abs computes the absolute value of a number.
//
// Parameters:
//  value number - The number to take the absolute value of
//
// Returns:
//  number - The abosulte value of a number
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Abs(value interface{}) Expr { return absFn{Abs: wrap(value)} }

// Acos computes the arccosine of a number.
//
// Parameters:
//  value number - The number to take the arccosine of
//
// Returns:
//  number - The arccosine of a number
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Acos(value interface{}) Expr { return acosFn{Acos: wrap(value)} }

// Asin computes the arcsine of a number.
//
// Parameters:
//  value number - The number to take the arcsine of
//
// Returns:
//  number - The arcsine of a number
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Asin(value interface{}) Expr { return asinFn{Asin: wrap(value)} }

// Atan computes the arctan of a number.
//
// Parameters:
//  value number - The number to take the arctan of
//
// Returns:
//  number - The arctan of a number
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Atan(value interface{}) Expr { return atanFn{Atan: wrap(value)} }

// Add computes the sum of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to sum together.
//
// Returns:
//  number - The sum of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Add(args ...interface{}) Expr { return addFn{Add: wrap(varargs(args...))} }

// BitAnd computes the and of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to and together.
//
// Returns:
//  number - The and of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func BitAnd(args ...interface{}) Expr { return bitAndFn{BitAnd: wrap(varargs(args...))} }

// BitNot computes the 2's complement of a number
//
// Parameters:
//  value number - A numbers to not
//
// Returns:
//  number - The not of an element
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func BitNot(value interface{}) Expr { return bitNotFn{BitNot: wrap(value)} }

// BitOr computes the OR of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to OR together.
//
// Returns:
//  number - The OR of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func BitOr(args ...interface{}) Expr { return bitOrFn{BitOr: wrap(varargs(args...))} }

// BitXor computes the XOR of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to XOR together.
//
// Returns:
//  number - The XOR of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func BitXor(args ...interface{}) Expr { return bitXorFn{BitXor: wrap(varargs(args...))} }

// Ceil computes the largest integer greater than or equal to
//
// Parameters:
//  value number - A numbers to compute the ceil of
//
// Returns:
//  number - The ceil of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Ceil(value interface{}) Expr { return ceilFn{Ceil: wrap(value)} }

// Cos computes the Cosine of a number
//
// Parameters:
//  value number - A number to compute the cosine of
//
// Returns:
//  number - The cosine of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Cos(value interface{}) Expr { return cosFn{Cos: wrap(value)} }

// Cosh computes the Hyperbolic Cosine of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic cosine of
//
// Returns:
//  number - The Hyperbolic cosine of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Cosh(value interface{}) Expr { return coshFn{Cosh: wrap(value)} }

// Degrees computes the degress of a number
//
// Parameters:
//  value number - A number to compute the degress of
//
// Returns:
//  number - The degrees of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Degrees(value interface{}) Expr { return degreesFn{Degrees: wrap(value)} }

// Divide computes the quotient of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the quotient of.
//
// Returns:
//  number - The quotient of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Divide(args ...interface{}) Expr { return divideFn{Divide: wrap(varargs(args...))} }

// Exp computes the Exp of a number
//
// Parameters:
//  value number - A number to compute the exp of
//
// Returns:
//  number - The exp of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Exp(value interface{}) Expr { return expFn{Exp: wrap(value)} }

// Floor computes the Floor of a number
//
// Parameters:
//  value number - A number to compute the Floor of
//
// Returns:
//  number - The Floor of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Floor(value interface{}) Expr { return floorFn{Floor: wrap(value)} }

// Hypot computes the Hypotenuse of two numbers
//
// Parameters:
//  a number - A side of a right triangle
//  b number - A side of a right triangle
//
// Returns:
//  number - The hypotenuse of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Hypot(a, b interface{}) Expr { return hypotFn{Hypot: wrap(a), B: wrap(b)} }

// Ln computes the natural log of a number
//
// Parameters:
//  value number - A number to compute the natural log of
//
// Returns:
//  number - The ln of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Ln(value interface{}) Expr { return lnFn{Ln: wrap(value)} }

// Log computes the Log of a number
//
// Parameters:
//  value number - A number to compute the Log of
//
// Returns:
//  number - The Log of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Log(value interface{}) Expr { return logFn{Log: wrap(value)} }

// Max computes the max of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to find the max of.
//
// Returns:
//  number - The max of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Max(args ...interface{}) Expr { return maxFn{Max: wrap(varargs(args...))} }

// Min computes the Min of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to find the min of.
//
// Returns:
//  number - The min of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Min(args ...interface{}) Expr { return minFn{Min: wrap(varargs(args...))} }

// Modulo computes the reminder after the division of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the quotient of. The remainder will be returned.
//
// Returns:
//  number - The remainder of the quotient of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Modulo(args ...interface{}) Expr { return moduloFn{Modulo: wrap(varargs(args...))} }

// Multiply computes the product of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to multiply together.
//
// Returns:
//  number - The multiplication of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Multiply(args ...interface{}) Expr { return multiplyFn{Multiply: wrap(varargs(args...))} }

// Pow computes the Power of a number
//
// Parameters:
//  base number - A number which is the base
//  exp number  - A number which is the exponent
//
// Returns:
//  number - The Pow of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Pow(base, exp interface{}) Expr { return powFn{Pow: wrap(base), Exp: wrap(exp)} }

// Radians computes the Radians of a number
//
// Parameters:
//  value number - A number which is convert to radians
//
// Returns:
//  number - The Radians of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Radians(value interface{}) Expr { return radiansFn{Radians: wrap(value)} }

// Round a number at the given percission
//
// Parameters:
//  value number - The number to truncate
//  precision number - precision where to truncate, defaults is 2
//
// Returns:
//  number - The Rounded value.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Round(value interface{}, options ...OptionalParameter) Expr {
	fn := roundFn{Round: wrap(value)}
	mappings := optionalKeysMapping{
		"precision": &(fn.Precision),
	}
	applyOptionals(mappings, options)
	return fn
}

// Sign computes the Sign of a number
//
// Parameters:
//  value number - A number to compute the Sign of
//
// Returns:
//  number - The Sign of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Sign(value interface{}) Expr { return signFn{Sign: wrap(value)} }

// Sin computes the Sine of a number
//
// Parameters:
//  value number - A number to compute the Sine of
//
// Returns:
//  number - The Sine of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Sin(value interface{}) Expr { return sinFn{Sin: wrap(value)} }

// Sinh computes the Hyperbolic Sine of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic Sine of
//
// Returns:
//  number - The Hyperbolic Sine of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Sinh(value interface{}) Expr { return sinhFn{Sinh: wrap(value)} }

// Sqrt computes the square root of a number
//
// Parameters:
//  value number - A number to compute the square root of
//
// Returns:
//  number - The square root of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Sqrt(value interface{}) Expr { return sqrtFn{Sqrt: wrap(value)} }

// Subtract computes the difference of a list of numbers.
//
// Parameters:
//  args []number - A collection of numbers to compute the difference of.
//
// Returns:
//  number - The difference of all elements.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Subtract(args ...interface{}) Expr { return subtractFn{Subtract: wrap(varargs(args...))} }

// Tan computes the Tangent of a number
//
// Parameters:
//  value number - A number to compute the Tangent of
//
// Returns:
//  number - The Tangent of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Tan(value interface{}) Expr { return tanFn{Tan: wrap(value)} }

// Tanh computes the Hyperbolic Tangent of a number
//
// Parameters:
//  value number - A number to compute the Hyperbolic Tangent of
//
// Returns:
//  number - The Hyperbolic Tangent of value
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Tanh(value interface{}) Expr { return tanhFn{Tanh: wrap(value)} }

// Trunc truncates a number at the given percission
//
// Parameters:
//  value number - The number to truncate
//  precision number - precision where to truncate, defaults is 2
//
// Returns:
//  number - The truncated value.
//
// See: https://app.fauna.com/documentation/reference/queryapi#mathematical-functions
func Trunc(value interface{}, options ...OptionalParameter) Expr {
	fn := truncFn{Trunc: wrap(value)}
	mappings := optionalKeysMapping{
		"precision": &(fn.Precision),
	}
	applyOptionals(mappings, options)
	return fn
}

// Any evaluates to true if any element of the collection is true.
//
// Parameters:
// collection  - the collection
//
// Returns:
// Expr
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/any
func Any(collection interface{}) Expr {
	return anyFn{Any: wrap(collection)}
}

// All evaluates to true if all elements of the collection are true.
//
// Parameters:
// collection - the collection
//
// Returns:
// Expr
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/all
func All(collection interface{}) Expr {
	return allFn{All: wrap(collection)}
}

// Count returns the number of elements in the collection.
//
// Parameters:
// collection Expr - the collection
//
// Returns:
// a new Expr instance
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/count
func Count(collection interface{}) Expr {
	return countFn{Count: wrap(collection)}
}

// Sum sums the elements in the collection.
//
// Parameters:
// collection Expr - the collection
//
// Returns:
// a new Expr instance
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/sum
func Sum(collection interface{}) Expr {
	return sumFn{Sum: wrap(collection)}
}

// Mean returns the mean of all elements in the collection.
//
// Parameters:
//
// collection Expr - the collection
//
// Returns:
// a new Expr instance
//
// See: https://docs.fauna.com/fauna/current/api/fql/functions/mean
func Mean(collection interface{}) Expr {
	return meanFn{Mean: wrap(collection)}
}

// LT returns true if each specified value is less than all the subsequent values. Otherwise LT returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are less than each other from left to right.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func LT(args ...interface{}) Expr { return ltFn{LT: wrap(varargs(args...))} }

// LTE returns true if each specified value is less than or equal to all subsequent values. Otherwise LTE returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are less than of equals each other from left to right.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func LTE(args ...interface{}) Expr { return lteFn{LTE: wrap(varargs(args...))} }

// GT returns true if each specified value is greater than all subsequent values. Otherwise GT returns false.
// and false otherwise.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are greather than to each other from left to right.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func GT(args ...interface{}) Expr { return gtFn{GT: wrap(varargs(args...))} }

// GTE returns true if each specified value is greater than or equal to all subsequent values. Otherwise GTE returns false.
//
// Parameters:
//  args []number - A collection of terms to compare.
//
// Returns:
//  bool - true if all elements are greather than or equals to each other from left to right.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func GTE(args ...interface{}) Expr { return gteFn{GTE: wrap(varargs(args...))} }

// And returns the conjunction of a list of boolean values.
//
// Parameters:
//  args []bool - A collection to compute the conjunction of.
//
// Returns:
//  bool - true if all elements are true, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func And(args ...interface{}) Expr { return andFn{And: wrap(varargs(args...))} }

// Or returns the disjunction of a list of boolean values.
//
// Parameters:
//  args []bool - A collection to compute the disjunction of.
//
// Returns:
//  bool - true if at least one element is true, false otherwise.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Or(args ...interface{}) Expr { return orFn{Or: wrap(varargs(args...))} }

// Not returns the negation of a boolean value.
//
// Parameters:
//  boolean bool - A boolean to produce the negation of.
//
// Returns:
//  bool - The value negated.
//
// See: https://app.fauna.com/documentation/reference/queryapi#miscellaneous-functions
func Not(boolean interface{}) Expr { return notFn{Not: wrap(boolean)} }

// Select traverses into the provided value, returning the value at the given path.
//
// Parameters:
//  path []Path - An array representing a path to pull from an object. Path can be either strings or numbers.
//  value Object - The object to select from.
//
// Optional parameters:
//  default Value - A default value if the path does not exist. See Default() function.
//
// Returns:
//  Value - The value at the given path location.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func Select(path, value interface{}, options ...OptionalParameter) Expr {
	fn := selectFn{Select: wrap(path), From: wrap(value)}
	mappings := optionalKeysMapping{
		"default": &fn.Default,
	}
	applyOptionals(mappings, options)
	return fn
}

// SelectAll traverses into the provided value flattening all values under the desired path.
//
// Parameters:
//  path []Path - An array representing a path to pull from an object. Path can be either strings or numbers.
//  value Object - The object to select from.
//
// Returns:
//  Value - The value at the given path location.
//
// See: https://app.fauna.com/documentation/reference/queryapi#read-functions
func SelectAll(path, value interface{}) Expr {
	return selectAllFn{SelectAll: wrap(path), From: wrap(value)}
}

// ToString attempts to convert an expression to a string literal.
//
// Parameters:
//   value Object - The expression to convert.
//
// Returns:
//   string - A string literal.
func ToString(value interface{}) Expr {
	return toStringFn{ToString: wrap(value)}
}

// ToNumber attempts to convert an expression to a numeric literal -
// either an int64 or float64.
//
// Parameters:
//   value Object - The expression to convert.
//
// Returns:
//   number - A numeric literal.
func ToNumber(value interface{}) Expr {
	return toNumberFn{ToNumber: wrap(value)}
}

// ToTime attempts to convert an expression to a time literal.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
func ToTime(value interface{}) Expr {
	return toTimeFn{ToTime: wrap(value)}
}

// ToSeconds converts a time expression to seconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
func ToSeconds(value interface{}) Expr {
	return toSecondsFn{ToSeconds: wrap(value)}
}

// ToMillis converts a time expression to milliseconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
func ToMillis(value interface{}) Expr {
	return toMillisFn{ToMillis: wrap(value)}
}

// ToMicros converts a time expression to microseconds since the UNIX epoch.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - A time literal.
func ToMicros(value interface{}) Expr {
	return toMicrosFn{ToMicros: wrap(value)}
}

// Year returns the time expression's year, following the ISO-8601 standard.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
func Year(value interface{}) Expr {
	return yearFn{Year: wrap(value)}
}

// Month returns a time expression's month of the year, from 1 to 12.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - Month.
func Month(value interface{}) Expr {
	return monthFn{Month: wrap(value)}
}

// Hour returns a time expression's hour of the day, from 0 to 23.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
func Hour(value interface{}) Expr {
	return hourFn{Hour: wrap(value)}
}

// Minute returns a time expression's minute of the hour, from 0 to 59.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
func Minute(value interface{}) Expr {
	return minuteFn{Minute: wrap(value)}
}

// Second returns a time expression's second of the minute, from 0 to 59.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - year.
func Second(value interface{}) Expr {
	return secondFn{Second: wrap(value)}
}

// DayOfMonth returns a time expression's day of the month, from 1 to 31.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - day of month.
func DayOfMonth(value interface{}) Expr {
	return dayOfMonthFn{DayOfMonth: wrap(value)}
}

// DayOfWeek returns a time expression's day of the week following ISO-8601 convention, from 1 (Monday) to 7 (Sunday).
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - day of week.
func DayOfWeek(value interface{}) Expr {
	return dayOfWeekFn{DayOfWeek: wrap(value)}
}

// DayOfYear eturns a time expression's day of the year, from 1 to 365, or 366 in a leap year.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   time - Day of the year.
func DayOfYear(value interface{}) Expr {
	return dayOfYearFn{DayOfYear: wrap(value)}
}

// ToDate attempts to convert an expression to a date literal.
//
// Parameters:
//    value Object - The expression to convert.
//
// Returns:
//   date - A date literal.
func ToDate(value interface{}) Expr {
	return toDateFn{ToDate: wrap(value)}
}

// IsNumber checks if the expression is a number
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool      -  returns true if the expression is a number
func IsNumber(expr interface{}) Expr {
	return isNumberFn{IsNumber: wrap(expr)}
}

// IsDouble checks if the expression is a double
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a double
func IsDouble(expr interface{}) Expr {
	return isDoubleFn{IsDouble: wrap(expr)}
}

// IsInteger checks if the expression is an integer
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is an integer
func IsInteger(expr interface{}) Expr {
	return isIntegerFn{IsInteger: wrap(expr)}
}

// IsBoolean checks if the expression is a boolean
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a boolean
func IsBoolean(expr interface{}) Expr {
	return isBooleanFn{IsBoolean: wrap(expr)}
}

// IsNull checks if the expression is null
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is null
func IsNull(expr interface{}) Expr {
	return isNullFn{IsNull: wrap(expr)}
}

// IsBytes checks if the expression are bytes
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression are bytes
func IsBytes(expr interface{}) Expr {
	return isBytesFn{IsBytes: wrap(expr)}
}

// IsTimestamp checks if the expression is a timestamp
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a timestamp
func IsTimestamp(expr interface{}) Expr {
	return isTimestampFn{IsTimestamp: wrap(expr)}
}

// IsDate checks if the expression is a date
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a date
func IsDate(expr interface{}) Expr {
	return isDateFn{IsDate: wrap(expr)}
}

// IsString checks if the expression is a string
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a string
func IsString(expr interface{}) Expr {
	return isStringFn{IsString: wrap(expr)}
}

// IsArray checks if the expression is an array
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is an array
func IsArray(expr interface{}) Expr {
	return isArrayFn{IsArray: wrap(expr)}
}

// IsObject checks if the expression is an object
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is an object
func IsObject(expr interface{}) Expr {
	return isObjectFn{IsObject: wrap(expr)}
}

// IsRef checks if the expression is a ref
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a ref
func IsRef(expr interface{}) Expr {
	return isRefFn{IsRef: wrap(expr)}
}

// IsSet checks if the expression is a set
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a set
func IsSet(expr interface{}) Expr {
	return isSetFn{IsSet: wrap(expr)}
}

// IsDoc checks if the expression is a document
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a document
func IsDoc(expr interface{}) Expr {
	return isDocFn{IsDoc: wrap(expr)}
}

// IsLambda checks if the expression is a Lambda
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a Lambda
func IsLambda(expr interface{}) Expr {
	return isLambdaFn{IsLambda: wrap(expr)}
}

// IsCollection checks if the expression is a collection
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a collection
func IsCollection(expr interface{}) Expr {
	return isCollectionFn{IsCollection: wrap(expr)}
}

// IsDatabase checks if the expression is a database
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a database
func IsDatabase(expr interface{}) Expr {
	return isDatabaseFn{IsDatabase: wrap(expr)}
}

// IsIndex checks if the expression is an index
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is an index
func IsIndex(expr interface{}) Expr {
	return isIndexFn{IsIndex: wrap(expr)}
}

// IsFunction checks if the expression is a function
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a function
func IsFunction(expr interface{}) Expr {
	return isFunctionFn{IsFunction: wrap(expr)}
}

// IsKey checks if the expression is a key
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a key
func IsKey(expr interface{}) Expr {
	return isKeyFn{IsKey: wrap(expr)}
}

// IsToken checks if the expression is a token
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a token
func IsToken(expr interface{}) Expr {
	return isTokenFn{IsToken: wrap(expr)}
}

// IsCredentials checks if the expression is a credentials
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a credential
func IsCredentials(expr interface{}) Expr {
	return isCredentialsFn{IsCredentials: wrap(expr)}
}

// IsRole checks if the expression is a role
//
// Parameters:
//  expr Expr - The expression to check.
//
// Returns:
//  bool         -  returns true if the expression is a role
func IsRole(expr interface{}) Expr {
	return isRoleFn{IsRole: wrap(expr)}
}
