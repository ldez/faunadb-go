package faunadb

type fnApply struct {
	Expr   `json:"-"`
	string `json:"-"`
}

type abortFn struct {
	fnApply
	Abort Expr `json:"abort"`
}

type doFn struct {
	fnApply
	Do Expr `json:"do"`
}

type ifFn struct {
	fnApply
	Else Expr `json:"else"`
	If   Expr `json:"if"`
	Then Expr `json:"then"`
}

type lambdaFn struct {
	fnApply
	Expression Expr `json:"expr"`
	Lambda     Expr `json:"lambda"`
}

type atFn struct {
	fnApply
	At         Expr `json:"at"`
	Expression Expr `json:"expr"`
}

type letFn struct {
	fnApply
	In  Expr `json:"in"`
	Let Expr `json:"let"`
}

type varFn struct {
	fnApply
	Var Expr `json:"var"`
}

type callFn struct {
	fnApply
	Call   Expr `json:"call"`
	Params Expr `json:"arguments"`
}

type queryFn struct {
	fnApply
	Query Expr `json:"query"`
}

type mapFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Map        Expr `json:"map"`
}

type foreachFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Foreach    Expr `json:"foreach"`
}

type filterFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Filter     Expr `json:"filter"`
}

type takeFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Take       Expr `json:"take"`
}

type dropFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Drop       Expr `json:"drop"`
}

type prependFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Prepend    Expr `json:"prepend"`
}

type appendFn struct {
	fnApply
	Append     Expr `json:"append"`
	Collection Expr `json:"collection"`
}

type isEmptyFn struct {
	fnApply
	IsEmpty Expr `json:"is_empty"`
}

type isNonEmptyFn struct {
	fnApply
	IsNonEmpty Expr `json:"is_nonempty"`
}

type getFn struct {
	fnApply
	Get Expr `json:"get"`
	TS  Expr `json:"ts,omitempty"`
}

type legacyRefFn struct {
	fnApply
	Ref Expr `json:"@ref"`
}

type refFn struct {
	fnApply
	ID  Expr `json:"id,omitempty"`
	Ref Expr `json:"ref"`
}

type keyFromSecretFn struct {
	fnApply
	KeyFromSecret Expr `json:"key_from_secret"`
}

type existsFn struct {
	fnApply
	Exists Expr `json:"exists"`
	TS     Expr `json:"ts,omitempty"`
}

type paginateFn struct {
	fnApply
	After    Expr `json:"after,omitempty"`
	Before   Expr `json:"before,omitempty"`
	Events   Expr `json:"events,omitempty"`
	Paginate Expr `json:"paginate"`
	Size     Expr `json:"size,omitempty"`
	Sources  Expr `json:"sources,omitempty"`
	TS       Expr `json:"ts,omitempty"`
}

type createFn struct {
	fnApply
	Create Expr `json:"create"`
	Params Expr `json:"params"`
}

type createClassFn struct {
	fnApply
	CreateClass Expr `json:"create_class"`
}

type createCollectionFn struct {
	fnApply
	CreateCollection Expr `json:"create_collection"`
}

type createDatabaseFn struct {
	fnApply
	CreateDatabase Expr `json:"create_database"`
}

type createIndexFn struct {
	fnApply
	CreateIndex Expr `json:"create_index"`
}

type createKeyFn struct {
	fnApply
	CreateKey Expr `json:"create_key"`
}

type createFunctionFn struct {
	fnApply
	CreateFunction Expr `json:"create_function"`
}

type createRoleFn struct {
	fnApply
	CreateRole Expr `json:"create_role"`
}

type moveDatabaseFn struct {
	fnApply
	MoveDatabase Expr `json:"move_database"`
	To           Expr `json:"to"`
}

type updateFn struct {
	fnApply
	Params Expr `json:"params"`
	Update Expr `json:"update"`
}

type replaceFn struct {
	fnApply
	Params  Expr `json:"params"`
	Replace Expr `json:"replace"`
}

type deleteFn struct {
	fnApply
	Delete Expr `json:"delete"`
}

type insertFn struct {
	fnApply
	Action Expr `json:"action"`
	Insert Expr `json:"insert"`
	Params Expr `json:"params"`
	Ts     Expr `json:"ts"`
}

type removeFn struct {
	fnApply
	Action Expr `json:"action"`
	Remove Expr `json:"remove"`
	Ts     Expr `json:"ts"`
}

type formatFn struct {
	fnApply
	Format Expr `json:"format"`
	Values Expr `json:"values"`
}

type concatFn struct {
	fnApply
	Concat    Expr `json:"concat"`
	Separator Expr `json:"separator,omitempty"`
}

type casefoldFn struct {
	fnApply
	Casefold   Expr `json:"casefold"`
	Normalizer Expr `json:"normalizer,omitempty"`
}

type startsWithFn struct {
	fnApply
	Search     Expr `json:"search"`
	StartsWith Expr `json:"startswith"`
}

type endsWithFn struct {
	fnApply
	EndsWith Expr `json:"endswith"`
	Search   Expr `json:"search"`
}

type containsStrFn struct {
	fnApply
	ContainsStr Expr `json:"containsstr"`
	Search      Expr `json:"search"`
}

type containsStrRegexFn struct {
	fnApply
	ContainsStrRegex Expr `json:"containsstrregex"`
	Pattern          Expr `json:"pattern"`
}

type regexEscapeFn struct {
	fnApply
	RegexEscape Expr `json:"regexescape"`
}

type findStrFn struct {
	fnApply
	Find    Expr `json:"find"`
	FindStr Expr `json:"findstr"`
	Start   Expr `json:"start,omitempty"`
}

type findStrRegexFn struct {
	fnApply
	FindStrRegex Expr `json:"findstrregex"`
	Pattern      Expr `json:"pattern"`
}

type lengthFn struct {
	fnApply
	Length Expr `json:"length"`
}

type lowercaseFn struct {
	fnApply
	Lowercase Expr `json:"lowercase"`
}

type lTrimFn struct {
	fnApply
	LTrim Expr `json:"ltrim"`
}

type repeatFn struct {
	fnApply
	Number Expr `json:"number"`
	Repeat Expr `json:"repeat"`
}

type replaceStrFn struct {
	fnApply
	Find       Expr `json:"find"`
	Replace    Expr `json:"replace"`
	ReplaceStr Expr `json:"replacestr"`
}

type replaceStrRegexFn struct {
	fnApply
	Pattern         Expr `json:"pattern"`
	Replace         Expr `json:"replace"`
	ReplaceStrRegex Expr `json:"replacestrregex"`
	First           Expr `json:"first,omitempty"`
}

type rTrimFn struct {
	fnApply
	RTrim Expr `json:"rtrim"`
}

type spaceFn struct {
	fnApply
	Space Expr `json:"space"`
}

type subStringFn struct {
	fnApply
	Length    Expr `json:"length,omitempty"`
	Start     Expr `json:"start"`
	SubString Expr `json:"substring"`
}

type titleCaseFn struct {
	fnApply
	Titlecase Expr `json:"titlecase"`
}

type trimFn struct {
	fnApply
	Trim Expr `json:"trim"`
}

type upperCaseFn struct {
	fnApply
	UpperCase Expr `json:"uppercase"`
}

type timeFn struct {
	fnApply
	Time Expr `json:"time"`
}

type timeAddFn struct {
	fnApply
	Offset  Expr `json:"offset"`
	TimeAdd Expr `json:"time_add"`
	Unit    Expr `json:"unit"`
}

type timeSubtractFn struct {
	fnApply
	Offset       Expr `json:"offset"`
	TimeSubtract Expr `json:"time_subtract"`
	Unit         Expr `json:"unit"`
}

type timeDiffFn struct {
	fnApply
	Other    Expr `json:"other"`
	TimeDiff Expr `json:"time_diff"`
	Unit     Expr `json:"unit"`
}

type dateFn struct {
	fnApply
	Date Expr `json:"date"`
}

type epochFn struct {
	fnApply
	Epoch Expr `json:"epoch"`
	Unit  Expr `json:"unit"`
}

type nowFn struct {
	fnApply
	Now Expr `json:"now"`
}

type singletonFn struct {
	fnApply
	Singleton Expr `json:"singleton"`
}

type eventsFn struct {
	fnApply
	Events Expr `json:"events"`
}

type matchFn struct {
	fnApply
	Match Expr `json:"match"`
	Terms Expr `json:"terms,omitempty"`
}

type unionFn struct {
	fnApply
	Union Expr `json:"union"`
}

type mergeFn struct {
	fnApply
	Lambda Expr `json:"lambda,omitempty"`
	Merge  Expr `json:"merge"`
	With   Expr `json:"with"`
}

type reduceFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Initial    Expr `json:"initial"`
	Reduce     Expr `json:"reduce"`
}

type intersectionFn struct {
	fnApply
	Intersection Expr `json:"intersection"`
}

type differenceFn struct {
	fnApply
	Difference Expr `json:"difference"`
}

type distinctFn struct {
	fnApply
	Distinct Expr `json:"distinct"`
}

type joinFn struct {
	fnApply
	Join Expr `json:"join"`
	With Expr `json:"with"`
}

type rangeFn struct {
	fnApply
	From  Expr `json:"from"`
	Range Expr `json:"range"`
	To    Expr `json:"to"`
}

type loginFn struct {
	fnApply
	Login  Expr `json:"login"`
	Params Expr `json:"params"`
}

type logoutFn struct {
	fnApply
	Logout Expr `json:"logout"`
}

type identifyFn struct {
	fnApply
	Identify Expr `json:"identify"`
	Password Expr `json:"password"`
}

type identityFn struct {
	fnApply
	Identity Expr `json:"identity"`
}

type hasIdentityFn struct {
	fnApply
	HasIdentity Expr `json:"has_identity"`
}

type nextIdFn struct {
	fnApply
	NextID Expr `json:"next_id"`
}

type newIdFn struct {
	fnApply
	NewID Expr `json:"new_id"`
}

type databaseFn struct {
	fnApply
	Database Expr `json:"database"`
	Scope    Expr `json:"scope,omitempty"`
}

type indexFn struct {
	fnApply
	Index Expr `json:"index"`
	Scope Expr `json:"scope,omitempty"`
}

type classFn struct {
	fnApply
	Class Expr `json:"class"`
	Scope Expr `json:"scope,omitempty"`
}

type collectionFn struct {
	fnApply
	Collection Expr `json:"collection"`
	Scope      Expr `json:"scope,omitempty"`
}

type documentsFn struct {
	fnApply
	Documents Expr `json:"documents"`
}

type functionFn struct {
	fnApply
	Function Expr `json:"function"`
	Scope    Expr `json:"scope,omitempty"`
}

type roleFn struct {
	fnApply
	Role  Expr `json:"role"`
	Scope Expr `json:"scope,omitempty"`
}

type collectionsFn struct {
	fnApply
	Collections Expr `json:"collections"`
}

type classesFn struct {
	fnApply
	Classes Expr `json:"classes"`
}

type indexesFn struct {
	fnApply
	Indexes Expr `json:"indexes"`
}

type databasesFn struct {
	fnApply
	Databases Expr `json:"databases"`
}

type functionsFn struct {
	fnApply
	Functions Expr `json:"functions"`
}

type rolesFn struct {
	fnApply
	Roles Expr `json:"roles"`
}

type keysFn struct {
	fnApply
	Keys Expr `json:"keys"`
}

type tokensFn struct {
	fnApply
	Tokens Expr `json:"tokens"`
}

type credentialsFn struct {
	fnApply
	Credentials Expr `json:"credentials"`
}

type equalsFn struct {
	fnApply
	Equals Expr `json:"equals"`
}

type containsFn struct {
	fnApply
	Contains Expr `json:"contains"`
	Value    Expr `json:"in"`
}

type absFn struct {
	fnApply
	Abs Expr `json:"abs"`
}

type acosFn struct {
	fnApply
	Acos Expr `json:"acos"`
}

type asinFn struct {
	fnApply
	Asin Expr `json:"asin"`
}

type atanFn struct {
	fnApply
	Atan Expr `json:"atan"`
}

type addFn struct {
	fnApply
	Add Expr `json:"add"`
}

type bitAndFn struct {
	fnApply
	BitAnd Expr `json:"bitand"`
}

type bitNotFn struct {
	fnApply
	BitNot Expr `json:"bitnot"`
}

type bitOrFn struct {
	fnApply
	BitOr Expr `json:"bitor"`
}

type bitXorFn struct {
	fnApply
	BitXor Expr `json:"bitxor"`
}

type ceilFn struct {
	fnApply
	Ceil Expr `json:"ceil"`
}

type cosFn struct {
	fnApply
	Cos Expr `json:"cos"`
}

type coshFn struct {
	fnApply
	Cosh Expr `json:"cosh"`
}

type degreesFn struct {
	fnApply
	Degrees Expr `json:"degrees"`
}

type divideFn struct {
	fnApply
	Divide Expr `json:"divide"`
}

type expFn struct {
	fnApply
	Exp Expr `json:"exp"`
}

type floorFn struct {
	fnApply
	Floor Expr `json:"floor"`
}

type hypotFn struct {
	fnApply
	B     Expr `json:"b"`
	Hypot Expr `json:"hypot"`
}

type lnFn struct {
	fnApply
	Ln Expr `json:"ln"`
}

type logFn struct {
	fnApply
	Log Expr `json:"log"`
}

type maxFn struct {
	fnApply
	Max Expr `json:"max"`
}

type minFn struct {
	fnApply
	Min Expr `json:"min"`
}

type moduloFn struct {
	fnApply
	Modulo Expr `json:"modulo"`
}

type multiplyFn struct {
	fnApply
	Multiply Expr `json:"multiply"`
}

type powFn struct {
	fnApply
	Exp Expr `json:"exp"`
	Pow Expr `json:"pow"`
}

type radiansFn struct {
	fnApply
	Radians Expr `json:"radians"`
}

type roundFn struct {
	fnApply
	Round     Expr `json:"round"`
	Precision Expr `json:"precision,omitempty"`
}

type signFn struct {
	fnApply
	Sign Expr `json:"sign"`
}

type sinFn struct {
	fnApply
	Sin Expr `json:"sin"`
}

type sinhFn struct {
	fnApply
	Sinh Expr `json:"sinh"`
}

type sqrtFn struct {
	fnApply
	Sqrt Expr `json:"sqrt"`
}

type subtractFn struct {
	fnApply
	Subtract Expr `json:"subtract"`
}

type tanFn struct {
	fnApply
	Tan Expr `json:"tan"`
}

type tanhFn struct {
	fnApply
	Tanh Expr `json:"tanh"`
}

type truncFn struct {
	fnApply
	Trunc     Expr `json:"trunc"`
	Precision Expr `json:"precision,omitempty"`
}

type anyFn struct {
	fnApply
	Any Expr `json:"any"`
}

type allFn struct {
	fnApply
	All Expr `json:"all"`
}

type countFn struct {
	fnApply
	Count Expr `json:"count"`
}

type sumFn struct {
	fnApply
	Sum Expr `json:"sum"`
}

type meanFn struct {
	fnApply
	Mean Expr `json:"mean"`
}

type ltFn struct {
	fnApply
	LT Expr `json:"lt"`
}

type lteFn struct {
	fnApply
	LTE Expr `json:"lte"`
}

type gtFn struct {
	fnApply
	GT Expr `json:"gt"`
}

type gteFn struct {
	fnApply
	GTE Expr `json:"gte"`
}

type andFn struct {
	fnApply
	And Expr `json:"and"`
}

type orFn struct {
	fnApply
	Or Expr `json:"or"`
}

type notFn struct {
	fnApply
	Not Expr `json:"not"`
}

type selectFn struct {
	fnApply
	Default Expr `json:"default,omitempty"`
	From    Expr `json:"from"`
	Select  Expr `json:"select"`
}

type selectAllFn struct {
	fnApply
	Default   Expr `json:"default,omitempty"`
	From      Expr `json:"from"`
	SelectAll Expr `json:"select_all"`
}

type toStringFn struct {
	fnApply
	ToString Expr `json:"to_string"`
}

type toNumberFn struct {
	fnApply
	ToNumber Expr `json:"to_number"`
}

type toTimeFn struct {
	fnApply
	ToTime Expr `json:"to_time"`
}

type toSecondsFn struct {
	fnApply
	ToSeconds Expr `json:"to_seconds"`
}

type toMillisFn struct {
	fnApply
	ToMillis Expr `json:"to_millis"`
}

type toMicrosFn struct {
	fnApply
	ToMicros Expr `json:"to_micros"`
}

type yearFn struct {
	fnApply
	Year Expr `json:"year"`
}

type monthFn struct {
	fnApply
	Month Expr `json:"month"`
}

type hourFn struct {
	fnApply
	Hour Expr `json:"hour"`
}

type minuteFn struct {
	fnApply
	Minute Expr `json:"minute"`
}

type secondFn struct {
	fnApply
	Second Expr `json:"second"`
}

type dayOfMonthFn struct {
	fnApply
	DayOfMonth Expr `json:"day_of_month"`
}

type dayOfWeekFn struct {
	fnApply
	DayOfWeek Expr `json:"day_of_week"`
}

type dayOfYearFn struct {
	fnApply
	DayOfYear Expr `json:"day_of_year"`
}

type toDateFn struct {
	fnApply
	ToDate Expr `json:"to_date"`
}

type isNumberFn struct {
	fnApply
	IsNumber Expr `json:"is_number"`
}

type isDoubleFn struct {
	fnApply
	IsDouble Expr `json:"is_double"`
}

type isIntegerFn struct {
	fnApply
	IsInteger Expr `json:"is_integer"`
}

type isBooleanFn struct {
	fnApply
	IsBoolean Expr `json:"is_boolean"`
}

type isTimestampFn struct {
	fnApply
	IsTimestamp Expr `json:"is_timestamp"`
}

type isBytesFn struct {
	fnApply
	IsBytes Expr `json:"is_bytes"`
}

type isNullFn struct {
	fnApply
	IsNull Expr `json:"is_null"`
}

type isDateFn struct {
	fnApply
	IsDate Expr `json:"is_date"`
}

type isStringFn struct {
	fnApply
	IsString Expr `json:"is_string"`
}

type isArrayFn struct {
	fnApply
	IsArray Expr `json:"is_array"`
}

type isObjectFn struct {
	fnApply
	IsObject Expr `json:"is_object"`
}

type isRefFn struct {
	fnApply
	IsRef Expr `json:"is_ref"`
}

type isSetFn struct {
	fnApply
	IsSet Expr `json:"is_set"`
}

type isDocFn struct {
	fnApply
	IsDoc Expr `json:"is_doc"`
}

type isLambdaFn struct {
	fnApply
	IsLambda Expr `json:"is_lambda"`
}

type isCollectionFn struct {
	fnApply
	IsCollection Expr `json:"is_collection"`
}

type isDatabaseFn struct {
	fnApply
	IsDatabase Expr `json:"is_database"`
}

type isFunctionFn struct {
	fnApply
	IsFunction Expr `json:"is_function"`
}

type isKeyFn struct {
	fnApply
	IsKey Expr `json:"is_key"`
}

type isTokenFn struct {
	fnApply
	IsToken Expr `json:"is_token"`
}

type isCredentialsFn struct {
	fnApply
	IsCredentials Expr `json:"is_credentials"`
}

type isRoleFn struct {
	fnApply
	IsRole Expr `json:"is_role"`
}

type isIndexFn struct {
	fnApply
	IsIndex Expr `json:"is_index"`
}
