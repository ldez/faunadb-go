package faunadb

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	noArgs    = 1
	varArgs   = 2
	optArgs   = 3
	basicArgs = 4
)

func printFn(name string, args ...Expr) string {
	return _fnCall(name, basicArgs, nil, args...)
}

func printScopedFn(name string, arg1 Expr, scope Expr) string {
	if scope != nil {
		return printFn("Scoped"+name, arg1, scope)
	}
	return printFn(name, arg1)
}

func printKeyScopedFn(name string, expr Expr) string {
	if expr != Null() {
		return printFn("Scoped"+name, expr)
	}
	return printNoArgsFn(name)

}

func printNoArgsFn(name string) string {
	return _fnCall(name, noArgs, nil, nil)
}

func _fnCall(name string, argType int, mapping unescapedObj, args ...Expr) string {
	pattern := "%s(%s)"
	res := ""
	switch argType {
	case noArgs:
		res = fmt.Sprintf(pattern, name, "")
	case basicArgs:
		strArgs := make([]string, len(args))
		for idx, v := range args {
			strArgs[idx] = v.String()
		}
		res = fmt.Sprintf(pattern, name, strings.Join(strArgs, ", "))
	case varArgs:
		if len(args) == 1 && reflect.ValueOf(args).Type().Implements(arrType) {
			nestedArgs := reflect.ValueOf(args).Interface().(unescapedArr)
			strArgs := make([]string, len(args))
			for idx, v := range nestedArgs {
				strArgs[idx] = v.String()
			}
			res = fmt.Sprintf(pattern, name, strings.Join(strArgs, ", "))
		} else {
			res = _fnCall(name, basicArgs, mapping, args...)
		}
	case optArgs:
		if mapping == nil || len(mapping) == 0 {
			panic(fmt.Sprintf("%s.String() accepts optional parameters, should have a non-zero mapping", name))
		}
		strArgs := []string{}
		for _, v := range args {
			strArgs = append(strArgs, v.String())
		}
		for k, v := range mapping {
			if v != nil {
				strArgs = append(strArgs, fmt.Sprintf(k, v))
			}
		}
		res = fmt.Sprintf(pattern, name, strings.Join(strArgs, ", "))
	}
	return res
}

func printVarArgFn(name string, args ...Expr) string {
	return _fnCall(name, varArgs, nil, args...)
}

func printFnWithOpts(name string, mapping unescapedObj, args ...Expr) string {
	return _fnCall(name, optArgs, mapping, args...)
}

func (fn legacyRefFn) String() string {
	return printFn("Ref", fn.Ref)
}

func (fn refFn) String() string {
	return printFn("RefCollection", fn.Ref, fn.ID)
}

func (fn abortFn) String() string {
	return printFn("Abort", fn.Abort)
}

func (fn doFn) String() string {
	return printFn("Do", fn.Do)
}

func (fn ifFn) String() string {
	return printFn("If", fn.If, fn.Then, fn.Else)
}

func (fn lambdaFn) String() string {
	return printFn("Lambda", fn.Lambda, fn.Expression)
}

func (fn atFn) String() string {
	return printFn("At", fn.At, fn.Expression)
}

func (fn letFn) String() string {
	bindings := []string{}
	tpe := reflect.TypeOf(fn.Let)
	if tpe == reflect.TypeOf(unescapedArr{}) {
		arr := fn.Let.(unescapedArr)
		for _, elem := range arr {
			m := elem.(unescapedObj)
			for k, v := range m {
				bindings = append(bindings, fmt.Sprintf("Bind(%s, %s)", k, v))
			}
		}
	} else {
		panic(fmt.Sprintf("Let bindings should be an unescapedArr, %s instead", tpe))
	}
	return fmt.Sprintf("Let().%s.In(%s)", strings.Join(bindings, "."), fn.In)
}

func (fn varFn) String() string {
	return printFn("Var", fn.Var)
}

func (fn callFn) String() string {
	return printFn("Call", fn.Call, fn.Params)
}

func (fn queryFn) String() string {
	return printFn("Query", fn.Query)
}

func (fn mapFn) String() string {
	return printFn("Map", fn.Map)
}

func (fn foreachFn) String() string {
	return printFn("Foreach", fn.Foreach, fn.Collection)
}

func (fn filterFn) String() string {
	return printFn("Filter", fn.Filter, fn.Collection)
}

func (fn takeFn) String() string {
	return printFn("Take", fn.Take)
}

func (fn dropFn) String() string {
	return printFn("Drop", fn.Drop)
}

func (fn prependFn) String() string {
	return printFn("Prepend", fn.Prepend, fn.Collection)

}

func (fn appendFn) String() string {
	return printFn("Append", fn.Append, fn.Collection)

}

func (fn isEmptyFn) String() string {
	return printFn("IsEmpty", fn.IsEmpty)
}

func (fn isNonEmptyFn) String() string {
	return printFn("IsNonEmpty", fn.IsNonEmpty)
}

func (fn getFn) String() string {
	return printFnWithOpts("Get", unescapedObj{
		"TS(%s)": fn.TS,
	}, fn.Get)
}

func (fn keyFromSecretFn) String() string {
	return printFn("KeyFromSecret", fn.KeyFromSecret)
}

func (fn existsFn) String() string {
	return printFnWithOpts("Exists", unescapedObj{
		"TS(%s)": fn.TS}, fn.Exists)
}

func (fn paginateFn) String() string {
	return printFnWithOpts("Paginate", unescapedObj{
		"After(%s)":   (fn.After),
		"Before(%s)":  (fn.Before),
		"Events(%s)":  (fn.Events),
		"Size(%s)":    (fn.Size),
		"Sources(%s)": (fn.Sources),
		"TS(%s)":      (fn.TS),
	}, fn.Paginate)
}

func (fn createFn) String() string {
	return printFn("Create", fn.Create)
}

func (fn createClassFn) String() string {
	return printFn("CreateClass", fn.CreateClass)
}

func (fn createCollectionFn) String() string {
	return printFn("CreateCollection", fn.CreateCollection)
}

func (fn createDatabaseFn) String() string {
	return printFn("CreateDatabase", fn.CreateDatabase)
}

func (fn createIndexFn) String() string {
	return printFn("CreateIndex", fn.CreateIndex)
}

func (fn createKeyFn) String() string {
	return printFn("CreateKey", fn.CreateKey)
}

func (fn createFunctionFn) String() string {
	return printFn("CreateFunction", fn.CreateFunction)
}

func (fn createRoleFn) String() string {
	return printFn("CreateRole", fn.CreateRole)
}

func (fn moveDatabaseFn) String() string {
	return printFn("MoveDatabase", fn.MoveDatabase, fn.To)
}

func (fn updateFn) String() string {
	return printFn("Update", fn.Update)
}

func (fn replaceFn) String() string {
	return printFn("Replace", fn.Replace)
}

func (fn deleteFn) String() string {
	return printFn("Delete", fn.Delete)
}

func (fn insertFn) String() string {
	return printFn("Insert", fn.Insert, fn.Ts, fn.Action, fn.Params)
}

func (fn removeFn) String() string {
	return printFn("Remove", fn.Remove, fn.Ts, fn.Action)
}

func (fn formatFn) String() string {
	return printFn("Format", fn.Format, fn.Values)
}

func (fn concatFn) String() string {
	return printFnWithOpts("Concat", unescapedObj{
		"Separator(%s)": fn.Separator,
	}, fn.Concat)
}

func (fn casefoldFn) String() string {
	return printFnWithOpts("Casefold", unescapedObj{
		"%s": fn.Normalizer,
	}, fn.Casefold)
}

func (fn startsWithFn) String() string {
	return printFn("StartsWith", fn.StartsWith, fn.Search)
}

func (fn endsWithFn) String() string {
	return printFn("EndsWith", fn.EndsWith, fn.Search)
}

func (fn containsStrFn) String() string {
	return printFn("ContainsStr", fn.ContainsStr, fn.Search)
}

func (fn containsStrRegexFn) String() string {
	return printFn("ContainsStrRegex", fn.ContainsStrRegex, fn.Pattern)

}

func (fn regexEscapeFn) String() string {
	return printFn("RegexEscape", fn.RegexEscape)
}

func (fn findStrFn) String() string {
	return printFnWithOpts("FindStr", unescapedObj{
		"Start": fn.Start,
	}, fn.FindStr, fn.Find)
}

func (fn findStrRegexFn) String() string {
	return printFn("FindStrRegex", fn.FindStrRegex, fn.Pattern)

}

func (fn lengthFn) String() string {
	return printFn("Length", fn.Length)
}

func (fn lowercaseFn) String() string {
	return printFn("LowerCase", fn.Lowercase)
}

func (fn lTrimFn) String() string {
	return printFn("LTrim", fn.LTrim)
}

func (fn repeatFn) String() string {
	return printFn("Repeat", fn.Repeat)
}

func (fn replaceStrFn) String() string {
	return printFn("ReplaceStr", fn.ReplaceStr, fn.Find, fn.Replace)
}

func (fn replaceStrRegexFn) String() string {
	return printFnWithOpts("ReplaceStrRegex", unescapedObj{
		"OnlyFirst()": fn.First,
	}, fn.ReplaceStrRegex, fn.Pattern, fn.Replace)
}

func (fn rTrimFn) String() string {
	return printFn("RTrim", fn.RTrim)
}

func (fn spaceFn) String() string {
	return printFn("Space", fn.Space)
}

func (fn subStringFn) String() string {
	return printFnWithOpts("SubString", unescapedObj{
		"Length(%s)": fn.Length,
	}, fn.SubString, fn.Start)
}

func (fn titleCaseFn) String() string {
	return printFn("TitleCase", fn.Titlecase)
}

func (fn trimFn) String() string {
	return printFn("Trim", fn.Trim)
}

func (fn upperCaseFn) String() string {
	return printFn("UpperCase", fn.UpperCase)
}
func (fn timeFn) String() string {
	return printFn("Time", fn.Time)
}

func (fn timeAddFn) String() string {
	return printFn("TimeAdd", fn.TimeAdd, fn.Offset, fn.Unit)

}

func (fn timeSubtractFn) String() string {
	return printFn("TimeSubtract", fn.TimeSubtract, fn.Offset, fn.Unit)

}

func (fn timeDiffFn) String() string {
	return printFn("TimeDiff", fn.TimeDiff, fn.Other, fn.Unit)

}

func (fn dateFn) String() string {
	return printFn("Date", fn.Date)
}

func (fn epochFn) String() string {
	return printFn("Epoch", fn.Epoch)
}

func (fn nowFn) String() string {
	return printNoArgsFn("Now")
}

func (fn singletonFn) String() string {
	return printFn("Singleton", fn.Singleton)
}

func (fn eventsFn) String() string {
	return printFn("Events", fn.Events)
}

func (fn matchFn) String() string {
	if fn.Terms != nil {
		return printFn("MatchTerms", fn.Match, fn.Terms)
	}
	return printFn("Match", fn.Match)
}

func (fn unionFn) String() string {
	return printFn("Union", fn.Union)
}

func (fn mergeFn) String() string {
	return printFnWithOpts("Merge", unescapedObj{
		"ConflictResolver(%s)": fn.Lambda,
	}, fn.Merge, fn.With)
}

func (fn reduceFn) String() string {
	return printFn("Reduce", fn.Reduce, fn.Initial, fn.Collection)
}

func (fn intersectionFn) String() string {
	return printFn("Intersection", fn.Intersection)
}

func (fn differenceFn) String() string {
	return printFn("Difference", fn.Difference)
}

func (fn distinctFn) String() string {
	return printFn("Distinct", fn.Distinct)
}

func (fn joinFn) String() string {
	return printFn("Join", fn.Join)
}

func (fn rangeFn) String() string {
	return printFn("Range", fn.Range, fn.From, fn.To)

}

func (fn loginFn) String() string {
	return printFn("Login", fn.Login, fn.Params)

}

func (fn logoutFn) String() string {
	return printFn("Logout", fn.Logout)
}

func (fn identifyFn) String() string {
	return printFn("Identify", fn.Identify, fn.Password)

}

func (fn identityFn) String() string {
	return printNoArgsFn("Identity")
}

func (fn hasIdentityFn) String() string {
	return printNoArgsFn("HasIdentity")
}

func (fn nextIdFn) String() string {
	return printNoArgsFn("NextID")
}

func (fn newIdFn) String() string {
	return printNoArgsFn("NewId")
}

func (fn databaseFn) String() string {
	return printScopedFn("Database", fn.Database, fn.Scope)
}

func (fn indexFn) String() string {
	return printScopedFn("Index", fn.Index, fn.Scope)
}

func (fn classFn) String() string {
	return printScopedFn("Class", fn.Class, fn.Scope)
}

func (fn collectionFn) String() string {
	return printScopedFn("Collection", fn.Collection, fn.Scope)
}

func (fn documentsFn) String() string {
	return printFn("Documents", fn.Documents)
}

func (fn functionFn) String() string {
	return printScopedFn("Function", fn.Function, fn.Scope)
}

func (fn roleFn) String() string {
	return printScopedFn("Role", fn.Role, fn.Scope)
}

func (fn classesFn) String() string {
	return printKeyScopedFn("Classes", fn.Classes)
}

func (fn collectionsFn) String() string {
	return printKeyScopedFn("Collections", fn.Collections)
}

func (fn indexesFn) String() string {
	return printKeyScopedFn("Indexes", fn.Indexes)
}

func (fn databasesFn) String() string {
	return printKeyScopedFn("Databases", fn.Databases)
}

func (fn functionsFn) String() string {
	return printKeyScopedFn("Functions", fn.Functions)
}

func (fn rolesFn) String() string {
	return printKeyScopedFn("Roles", fn.Roles)
}

func (fn keysFn) String() string {
	return printKeyScopedFn("Keys", fn.Keys)
}

func (fn tokensFn) String() string {
	return printKeyScopedFn("Tokens", fn.Tokens)
}

func (fn credentialsFn) String() string {
	return printKeyScopedFn("Credentials", fn.Credentials)
}

func (fn equalsFn) String() string {
	return printFn("Equals", fn.Equals)
}

func (fn containsFn) String() string {
	return printFn("Contains", fn.Contains, fn.Value)

}

func (fn absFn) String() string {
	return printFn("Abs", fn.Abs)
}

func (fn acosFn) String() string {
	return printFn("Acos", fn.Acos)
}

func (fn asinFn) String() string {
	return printFn("Asin", fn.Asin)
}

func (fn atanFn) String() string {
	return printFn("Atan", fn.Atan)
}

func (fn addFn) String() string {
	return printFn("Add", fn.Add)
}

func (fn bitAndFn) String() string {
	return printFn("BitAnd", fn.BitAnd)
}

func (fn bitNotFn) String() string {
	return printFn("BitNot", fn.BitNot)
}

func (fn bitOrFn) String() string {
	return printFn("BitOr", fn.BitOr)
}

func (fn bitXorFn) String() string {
	return printFn("BitXor", fn.BitXor)
}

func (fn ceilFn) String() string {
	return printFn("Ceil", fn.Ceil)
}

func (fn cosFn) String() string {
	return printFn("Cos", fn.Cos)
}

func (fn coshFn) String() string {
	return printFn("Cosh", fn.Cosh)
}

func (fn degreesFn) String() string {
	return printFn("Degrees", fn.Degrees)
}

func (fn divideFn) String() string {
	return printFn("Divide", fn.Divide)
}

func (fn expFn) String() string {
	return printFn("Exp", fn.Exp)
}

func (fn floorFn) String() string {
	return printFn("Floor", fn.Floor)
}

func (fn hypotFn) String() string {
	return printFn("Hypot", fn.Hypot)
}

func (fn lnFn) String() string {
	return printFn("Ln", fn.Ln)
}

func (fn logFn) String() string {
	return printFn("Log", fn.Log)
}

func (fn maxFn) String() string {
	return printFn("Max", fn.Max)
}

func (fn minFn) String() string {
	return printFn("Min", fn.Min)
}

func (fn moduloFn) String() string {
	return printFn("Modulo", fn.Modulo)
}

func (fn multiplyFn) String() string {
	return printFn("Multiply", fn.Multiply)
}

func (fn powFn) String() string {
	return printFn("Pow", fn.Pow)
}

func (fn radiansFn) String() string {
	return printFn("Radians", fn.Radians)
}

func (fn roundFn) String() string {
	return printFnWithOpts("Round", unescapedObj{
		"Precision(%s)": fn.Precision,
	}, fn.Round)
}

func (fn signFn) String() string {
	return printFn("Sign", fn.Sign)
}

func (fn sinFn) String() string {
	return printFn("Sin", fn.Sin)
}

func (fn sinhFn) String() string {
	return printFn("Sinh", fn.Sinh)
}

func (fn sqrtFn) String() string {
	return printFn("Sqrt", fn.Sqrt)
}

func (fn subtractFn) String() string {
	return printFn("Subtract", fn.Subtract)
}

func (fn tanFn) String() string {
	return printFn("Tan", fn.Tan)
}

func (fn tanhFn) String() string {
	return printFn("Tanh", fn.Tanh)
}

func (fn truncFn) String() string {
	return printFnWithOpts("Trunc", unescapedObj{
		"Precision(%s)": fn.Precision,
	}, fn.Trunc)
}

func (fn anyFn) String() string {
	return printFn("Any", fn.Any)
}

func (fn allFn) String() string {
	return printFn("All", fn.All)
}

func (fn countFn) String() string {
	return printFn("Count", fn.Count)
}

func (fn sumFn) String() string {
	return printFn("Sum", fn.Sum)
}

func (fn meanFn) String() string {
	return printFn("Mean", fn.Mean)
}

func (fn ltFn) String() string {
	return printFn("LT", fn.LT)
}

func (fn lteFn) String() string {
	return printFn("LTE", fn.LTE)
}

func (fn gtFn) String() string {
	return printFn("GT", fn.GT)
}

func (fn gteFn) String() string {
	return printFn("GTE", fn.GTE)
}

func (fn andFn) String() string {
	return printFn("And", fn.And)
}

func (fn orFn) String() string {
	return printFn("Or", fn.Or)
}

func (fn notFn) String() string {
	return printFn("Not", fn.Not)
}

func (fn selectFn) String() string {
	return printFnWithOpts("Select", unescapedObj{
		"Default(%s)": fn.Default,
	}, fn.Select, fn.From)
}

func (fn selectAllFn) String() string {
	return printFnWithOpts("SelectAll", unescapedObj{
		"SelectAll(%s)": fn.Default,
	}, fn.SelectAll, fn.From)
}

func (fn toStringFn) String() string {
	return printFn("ToString", fn.ToString)
}

func (fn toNumberFn) String() string {
	return printFn("ToNumber", fn.ToNumber)
}

func (fn toTimeFn) String() string {
	return printFn("ToTime", fn.ToTime)
}

func (fn toSecondsFn) String() string {
	return printFn("ToSeconds", fn.ToSeconds)
}

func (fn toMillisFn) String() string {
	return printFn("ToMillis", fn.ToMillis)
}

func (fn toMicrosFn) String() string {
	return printFn("ToMicros", fn.ToMicros)
}

func (fn yearFn) String() string {
	return printFn("Year", fn.Year)
}

func (fn monthFn) String() string {
	return printFn("Month", fn.Month)
}

func (fn hourFn) String() string {
	return printFn("Hour", fn.Hour)
}

func (fn minuteFn) String() string {
	return printFn("Minute", fn.Minute)
}

func (fn secondFn) String() string {
	return printFn("Second", fn.Second)
}

func (fn dayOfMonthFn) String() string {
	return printFn("DayOfMonth", fn.DayOfMonth)
}

func (fn dayOfWeekFn) String() string {
	return printFn("DayOfWeek", fn.DayOfWeek)
}

func (fn dayOfYearFn) String() string {
	return printFn("DayOfYear", fn.DayOfYear)
}

func (fn toDateFn) String() string {
	return printFn("ToDate", fn.ToDate)
}

func (fn isNumberFn) String() string {
	return printFn("IsNumber", fn.IsNumber)
}

func (fn isDoubleFn) String() string {
	return printFn("IsDouble", fn.IsDouble)
}

func (fn isIntegerFn) String() string {
	return printFn("IsInteger", fn.IsInteger)
}

func (fn isBooleanFn) String() string {
	return printFn("IsBoolean", fn.IsBoolean)
}

func (fn isNullFn) String() string {
	return printFn("IsNull", fn.IsNull)
}

func (fn isBytesFn) String() string {
	return printFn("IsBytes", fn.IsBytes)
}

func (fn isTimestampFn) String() string {
	return printFn("IsTimestamp", fn.IsTimestamp)
}

func (fn isDateFn) String() string {
	return printFn("IsDate", fn.IsDate)
}

func (fn isStringFn) String() string {
	return printFn("IsString", fn.IsString)
}

func (fn isArrayFn) String() string {
	return printFn("IsArray", fn.IsArray)
}

func (fn isObjectFn) String() string {
	return printFn("IsObject", fn.IsObject)
}

func (fn isRefFn) String() string {
	return printFn("IsRef", fn.IsRef)
}

func (fn isSetFn) String() string {
	return printFn("IsSet", fn.IsSet)
}

func (fn isDocFn) String() string {
	return printFn("IsDoc", fn.IsDoc)
}

func (fn isLambdaFn) String() string {
	return printFn("IsLambda", fn.IsLambda)
}

func (fn isCollectionFn) String() string {
	return printFn("IsCollection", fn.IsCollection)
}

func (fn isDatabaseFn) String() string {
	return printFn("IsDatabase", fn.IsDatabase)
}

func (fn isIndexFn) String() string {
	return printFn("IsIndex", fn.IsIndex)
}

func (fn isFunctionFn) String() string {
	return printFn("IsFunction", fn.IsFunction)
}

func (fn isKeyFn) String() string {
	return printFn("IsKey", fn.IsKey)
}

func (fn isTokenFn) String() string {
	return printFn("IsToken", fn.IsToken)
}

func (fn isCredentialsFn) String() string {
	return printFn("IsCredentials", fn.IsCredentials)
}

func (fn isRoleFn) String() string {
	return printFn("IsRole", fn.IsRole)
}
