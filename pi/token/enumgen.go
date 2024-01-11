// Code generated by "goki generate ./..."; DO NOT EDIT.

package token

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/goki/enums"
)

var _TokensValues = []Tokens{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176}

// TokensN is the highest valid value
// for type Tokens, plus one.
const TokensN Tokens = 177

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TokensNoOp() {
	var x [1]struct{}
	_ = x[None-(0)]
	_ = x[Error-(1)]
	_ = x[EOF-(2)]
	_ = x[EOL-(3)]
	_ = x[EOS-(4)]
	_ = x[Background-(5)]
	_ = x[Keyword-(6)]
	_ = x[KeywordConstant-(7)]
	_ = x[KeywordDeclaration-(8)]
	_ = x[KeywordNamespace-(9)]
	_ = x[KeywordPseudo-(10)]
	_ = x[KeywordReserved-(11)]
	_ = x[KeywordType-(12)]
	_ = x[Name-(13)]
	_ = x[NameBuiltin-(14)]
	_ = x[NameBuiltinPseudo-(15)]
	_ = x[NameOther-(16)]
	_ = x[NamePseudo-(17)]
	_ = x[NameType-(18)]
	_ = x[NameClass-(19)]
	_ = x[NameStruct-(20)]
	_ = x[NameField-(21)]
	_ = x[NameInterface-(22)]
	_ = x[NameConstant-(23)]
	_ = x[NameEnum-(24)]
	_ = x[NameEnumMember-(25)]
	_ = x[NameArray-(26)]
	_ = x[NameMap-(27)]
	_ = x[NameObject-(28)]
	_ = x[NameTypeParam-(29)]
	_ = x[NameFunction-(30)]
	_ = x[NameDecorator-(31)]
	_ = x[NameFunctionMagic-(32)]
	_ = x[NameMethod-(33)]
	_ = x[NameOperator-(34)]
	_ = x[NameConstructor-(35)]
	_ = x[NameException-(36)]
	_ = x[NameLabel-(37)]
	_ = x[NameEvent-(38)]
	_ = x[NameScope-(39)]
	_ = x[NameNamespace-(40)]
	_ = x[NameModule-(41)]
	_ = x[NamePackage-(42)]
	_ = x[NameLibrary-(43)]
	_ = x[NameVar-(44)]
	_ = x[NameVarAnonymous-(45)]
	_ = x[NameVarClass-(46)]
	_ = x[NameVarGlobal-(47)]
	_ = x[NameVarInstance-(48)]
	_ = x[NameVarMagic-(49)]
	_ = x[NameVarParam-(50)]
	_ = x[NameValue-(51)]
	_ = x[NameTag-(52)]
	_ = x[NameProperty-(53)]
	_ = x[NameAttribute-(54)]
	_ = x[NameEntity-(55)]
	_ = x[Literal-(56)]
	_ = x[LiteralDate-(57)]
	_ = x[LiteralOther-(58)]
	_ = x[LiteralBool-(59)]
	_ = x[LitStr-(60)]
	_ = x[LitStrAffix-(61)]
	_ = x[LitStrAtom-(62)]
	_ = x[LitStrBacktick-(63)]
	_ = x[LitStrBoolean-(64)]
	_ = x[LitStrChar-(65)]
	_ = x[LitStrDelimiter-(66)]
	_ = x[LitStrDoc-(67)]
	_ = x[LitStrDouble-(68)]
	_ = x[LitStrEscape-(69)]
	_ = x[LitStrHeredoc-(70)]
	_ = x[LitStrInterpol-(71)]
	_ = x[LitStrName-(72)]
	_ = x[LitStrOther-(73)]
	_ = x[LitStrRegex-(74)]
	_ = x[LitStrSingle-(75)]
	_ = x[LitStrSymbol-(76)]
	_ = x[LitStrFile-(77)]
	_ = x[LitNum-(78)]
	_ = x[LitNumBin-(79)]
	_ = x[LitNumFloat-(80)]
	_ = x[LitNumHex-(81)]
	_ = x[LitNumInteger-(82)]
	_ = x[LitNumIntegerLong-(83)]
	_ = x[LitNumOct-(84)]
	_ = x[LitNumImag-(85)]
	_ = x[Operator-(86)]
	_ = x[OperatorWord-(87)]
	_ = x[OpMath-(88)]
	_ = x[OpMathAdd-(89)]
	_ = x[OpMathSub-(90)]
	_ = x[OpMathMul-(91)]
	_ = x[OpMathDiv-(92)]
	_ = x[OpMathRem-(93)]
	_ = x[OpBit-(94)]
	_ = x[OpBitAnd-(95)]
	_ = x[OpBitOr-(96)]
	_ = x[OpBitNot-(97)]
	_ = x[OpBitXor-(98)]
	_ = x[OpBitShiftLeft-(99)]
	_ = x[OpBitShiftRight-(100)]
	_ = x[OpBitAndNot-(101)]
	_ = x[OpAsgn-(102)]
	_ = x[OpAsgnAssign-(103)]
	_ = x[OpAsgnInc-(104)]
	_ = x[OpAsgnDec-(105)]
	_ = x[OpAsgnArrow-(106)]
	_ = x[OpAsgnDefine-(107)]
	_ = x[OpMathAsgn-(108)]
	_ = x[OpMathAsgnAdd-(109)]
	_ = x[OpMathAsgnSub-(110)]
	_ = x[OpMathAsgnMul-(111)]
	_ = x[OpMathAsgnDiv-(112)]
	_ = x[OpMathAsgnRem-(113)]
	_ = x[OpBitAsgn-(114)]
	_ = x[OpBitAsgnAnd-(115)]
	_ = x[OpBitAsgnOr-(116)]
	_ = x[OpBitAsgnXor-(117)]
	_ = x[OpBitAsgnShiftLeft-(118)]
	_ = x[OpBitAsgnShiftRight-(119)]
	_ = x[OpBitAsgnAndNot-(120)]
	_ = x[OpLog-(121)]
	_ = x[OpLogAnd-(122)]
	_ = x[OpLogOr-(123)]
	_ = x[OpLogNot-(124)]
	_ = x[OpRel-(125)]
	_ = x[OpRelEqual-(126)]
	_ = x[OpRelNotEqual-(127)]
	_ = x[OpRelLess-(128)]
	_ = x[OpRelGreater-(129)]
	_ = x[OpRelLtEq-(130)]
	_ = x[OpRelGtEq-(131)]
	_ = x[OpList-(132)]
	_ = x[OpListEllipsis-(133)]
	_ = x[Punctuation-(134)]
	_ = x[PunctGp-(135)]
	_ = x[PunctGpLParen-(136)]
	_ = x[PunctGpRParen-(137)]
	_ = x[PunctGpLBrack-(138)]
	_ = x[PunctGpRBrack-(139)]
	_ = x[PunctGpLBrace-(140)]
	_ = x[PunctGpRBrace-(141)]
	_ = x[PunctSep-(142)]
	_ = x[PunctSepComma-(143)]
	_ = x[PunctSepPeriod-(144)]
	_ = x[PunctSepSemicolon-(145)]
	_ = x[PunctSepColon-(146)]
	_ = x[PunctStr-(147)]
	_ = x[PunctStrDblQuote-(148)]
	_ = x[PunctStrQuote-(149)]
	_ = x[PunctStrBacktick-(150)]
	_ = x[PunctStrEsc-(151)]
	_ = x[Comment-(152)]
	_ = x[CommentHashbang-(153)]
	_ = x[CommentMultiline-(154)]
	_ = x[CommentSingle-(155)]
	_ = x[CommentSpecial-(156)]
	_ = x[CommentPreproc-(157)]
	_ = x[CommentPreprocFile-(158)]
	_ = x[Text-(159)]
	_ = x[TextWhitespace-(160)]
	_ = x[TextSymbol-(161)]
	_ = x[TextPunctuation-(162)]
	_ = x[TextSpellErr-(163)]
	_ = x[TextStyle-(164)]
	_ = x[TextStyleDeleted-(165)]
	_ = x[TextStyleEmph-(166)]
	_ = x[TextStyleError-(167)]
	_ = x[TextStyleHeading-(168)]
	_ = x[TextStyleInserted-(169)]
	_ = x[TextStyleOutput-(170)]
	_ = x[TextStylePrompt-(171)]
	_ = x[TextStyleStrong-(172)]
	_ = x[TextStyleSubheading-(173)]
	_ = x[TextStyleTraceback-(174)]
	_ = x[TextStyleUnderline-(175)]
	_ = x[TextStyleLink-(176)]
}

var _TokensNameToValueMap = map[string]Tokens{
	`None`:                0,
	`none`:                0,
	`Error`:               1,
	`error`:               1,
	`EOF`:                 2,
	`eof`:                 2,
	`EOL`:                 3,
	`eol`:                 3,
	`EOS`:                 4,
	`eos`:                 4,
	`Background`:          5,
	`background`:          5,
	`Keyword`:             6,
	`keyword`:             6,
	`KeywordConstant`:     7,
	`keywordconstant`:     7,
	`KeywordDeclaration`:  8,
	`keyworddeclaration`:  8,
	`KeywordNamespace`:    9,
	`keywordnamespace`:    9,
	`KeywordPseudo`:       10,
	`keywordpseudo`:       10,
	`KeywordReserved`:     11,
	`keywordreserved`:     11,
	`KeywordType`:         12,
	`keywordtype`:         12,
	`Name`:                13,
	`name`:                13,
	`NameBuiltin`:         14,
	`namebuiltin`:         14,
	`NameBuiltinPseudo`:   15,
	`namebuiltinpseudo`:   15,
	`NameOther`:           16,
	`nameother`:           16,
	`NamePseudo`:          17,
	`namepseudo`:          17,
	`NameType`:            18,
	`nametype`:            18,
	`NameClass`:           19,
	`nameclass`:           19,
	`NameStruct`:          20,
	`namestruct`:          20,
	`NameField`:           21,
	`namefield`:           21,
	`NameInterface`:       22,
	`nameinterface`:       22,
	`NameConstant`:        23,
	`nameconstant`:        23,
	`NameEnum`:            24,
	`nameenum`:            24,
	`NameEnumMember`:      25,
	`nameenummember`:      25,
	`NameArray`:           26,
	`namearray`:           26,
	`NameMap`:             27,
	`namemap`:             27,
	`NameObject`:          28,
	`nameobject`:          28,
	`NameTypeParam`:       29,
	`nametypeparam`:       29,
	`NameFunction`:        30,
	`namefunction`:        30,
	`NameDecorator`:       31,
	`namedecorator`:       31,
	`NameFunctionMagic`:   32,
	`namefunctionmagic`:   32,
	`NameMethod`:          33,
	`namemethod`:          33,
	`NameOperator`:        34,
	`nameoperator`:        34,
	`NameConstructor`:     35,
	`nameconstructor`:     35,
	`NameException`:       36,
	`nameexception`:       36,
	`NameLabel`:           37,
	`namelabel`:           37,
	`NameEvent`:           38,
	`nameevent`:           38,
	`NameScope`:           39,
	`namescope`:           39,
	`NameNamespace`:       40,
	`namenamespace`:       40,
	`NameModule`:          41,
	`namemodule`:          41,
	`NamePackage`:         42,
	`namepackage`:         42,
	`NameLibrary`:         43,
	`namelibrary`:         43,
	`NameVar`:             44,
	`namevar`:             44,
	`NameVarAnonymous`:    45,
	`namevaranonymous`:    45,
	`NameVarClass`:        46,
	`namevarclass`:        46,
	`NameVarGlobal`:       47,
	`namevarglobal`:       47,
	`NameVarInstance`:     48,
	`namevarinstance`:     48,
	`NameVarMagic`:        49,
	`namevarmagic`:        49,
	`NameVarParam`:        50,
	`namevarparam`:        50,
	`NameValue`:           51,
	`namevalue`:           51,
	`NameTag`:             52,
	`nametag`:             52,
	`NameProperty`:        53,
	`nameproperty`:        53,
	`NameAttribute`:       54,
	`nameattribute`:       54,
	`NameEntity`:          55,
	`nameentity`:          55,
	`Literal`:             56,
	`literal`:             56,
	`LiteralDate`:         57,
	`literaldate`:         57,
	`LiteralOther`:        58,
	`literalother`:        58,
	`LiteralBool`:         59,
	`literalbool`:         59,
	`LitStr`:              60,
	`litstr`:              60,
	`LitStrAffix`:         61,
	`litstraffix`:         61,
	`LitStrAtom`:          62,
	`litstratom`:          62,
	`LitStrBacktick`:      63,
	`litstrbacktick`:      63,
	`LitStrBoolean`:       64,
	`litstrboolean`:       64,
	`LitStrChar`:          65,
	`litstrchar`:          65,
	`LitStrDelimiter`:     66,
	`litstrdelimiter`:     66,
	`LitStrDoc`:           67,
	`litstrdoc`:           67,
	`LitStrDouble`:        68,
	`litstrdouble`:        68,
	`LitStrEscape`:        69,
	`litstrescape`:        69,
	`LitStrHeredoc`:       70,
	`litstrheredoc`:       70,
	`LitStrInterpol`:      71,
	`litstrinterpol`:      71,
	`LitStrName`:          72,
	`litstrname`:          72,
	`LitStrOther`:         73,
	`litstrother`:         73,
	`LitStrRegex`:         74,
	`litstrregex`:         74,
	`LitStrSingle`:        75,
	`litstrsingle`:        75,
	`LitStrSymbol`:        76,
	`litstrsymbol`:        76,
	`LitStrFile`:          77,
	`litstrfile`:          77,
	`LitNum`:              78,
	`litnum`:              78,
	`LitNumBin`:           79,
	`litnumbin`:           79,
	`LitNumFloat`:         80,
	`litnumfloat`:         80,
	`LitNumHex`:           81,
	`litnumhex`:           81,
	`LitNumInteger`:       82,
	`litnuminteger`:       82,
	`LitNumIntegerLong`:   83,
	`litnumintegerlong`:   83,
	`LitNumOct`:           84,
	`litnumoct`:           84,
	`LitNumImag`:          85,
	`litnumimag`:          85,
	`Operator`:            86,
	`operator`:            86,
	`OperatorWord`:        87,
	`operatorword`:        87,
	`OpMath`:              88,
	`opmath`:              88,
	`OpMathAdd`:           89,
	`opmathadd`:           89,
	`OpMathSub`:           90,
	`opmathsub`:           90,
	`OpMathMul`:           91,
	`opmathmul`:           91,
	`OpMathDiv`:           92,
	`opmathdiv`:           92,
	`OpMathRem`:           93,
	`opmathrem`:           93,
	`OpBit`:               94,
	`opbit`:               94,
	`OpBitAnd`:            95,
	`opbitand`:            95,
	`OpBitOr`:             96,
	`opbitor`:             96,
	`OpBitNot`:            97,
	`opbitnot`:            97,
	`OpBitXor`:            98,
	`opbitxor`:            98,
	`OpBitShiftLeft`:      99,
	`opbitshiftleft`:      99,
	`OpBitShiftRight`:     100,
	`opbitshiftright`:     100,
	`OpBitAndNot`:         101,
	`opbitandnot`:         101,
	`OpAsgn`:              102,
	`opasgn`:              102,
	`OpAsgnAssign`:        103,
	`opasgnassign`:        103,
	`OpAsgnInc`:           104,
	`opasgninc`:           104,
	`OpAsgnDec`:           105,
	`opasgndec`:           105,
	`OpAsgnArrow`:         106,
	`opasgnarrow`:         106,
	`OpAsgnDefine`:        107,
	`opasgndefine`:        107,
	`OpMathAsgn`:          108,
	`opmathasgn`:          108,
	`OpMathAsgnAdd`:       109,
	`opmathasgnadd`:       109,
	`OpMathAsgnSub`:       110,
	`opmathasgnsub`:       110,
	`OpMathAsgnMul`:       111,
	`opmathasgnmul`:       111,
	`OpMathAsgnDiv`:       112,
	`opmathasgndiv`:       112,
	`OpMathAsgnRem`:       113,
	`opmathasgnrem`:       113,
	`OpBitAsgn`:           114,
	`opbitasgn`:           114,
	`OpBitAsgnAnd`:        115,
	`opbitasgnand`:        115,
	`OpBitAsgnOr`:         116,
	`opbitasgnor`:         116,
	`OpBitAsgnXor`:        117,
	`opbitasgnxor`:        117,
	`OpBitAsgnShiftLeft`:  118,
	`opbitasgnshiftleft`:  118,
	`OpBitAsgnShiftRight`: 119,
	`opbitasgnshiftright`: 119,
	`OpBitAsgnAndNot`:     120,
	`opbitasgnandnot`:     120,
	`OpLog`:               121,
	`oplog`:               121,
	`OpLogAnd`:            122,
	`oplogand`:            122,
	`OpLogOr`:             123,
	`oplogor`:             123,
	`OpLogNot`:            124,
	`oplognot`:            124,
	`OpRel`:               125,
	`oprel`:               125,
	`OpRelEqual`:          126,
	`oprelequal`:          126,
	`OpRelNotEqual`:       127,
	`oprelnotequal`:       127,
	`OpRelLess`:           128,
	`oprelless`:           128,
	`OpRelGreater`:        129,
	`oprelgreater`:        129,
	`OpRelLtEq`:           130,
	`oprellteq`:           130,
	`OpRelGtEq`:           131,
	`oprelgteq`:           131,
	`OpList`:              132,
	`oplist`:              132,
	`OpListEllipsis`:      133,
	`oplistellipsis`:      133,
	`Punctuation`:         134,
	`punctuation`:         134,
	`PunctGp`:             135,
	`punctgp`:             135,
	`PunctGpLParen`:       136,
	`punctgplparen`:       136,
	`PunctGpRParen`:       137,
	`punctgprparen`:       137,
	`PunctGpLBrack`:       138,
	`punctgplbrack`:       138,
	`PunctGpRBrack`:       139,
	`punctgprbrack`:       139,
	`PunctGpLBrace`:       140,
	`punctgplbrace`:       140,
	`PunctGpRBrace`:       141,
	`punctgprbrace`:       141,
	`PunctSep`:            142,
	`punctsep`:            142,
	`PunctSepComma`:       143,
	`punctsepcomma`:       143,
	`PunctSepPeriod`:      144,
	`punctsepperiod`:      144,
	`PunctSepSemicolon`:   145,
	`punctsepsemicolon`:   145,
	`PunctSepColon`:       146,
	`punctsepcolon`:       146,
	`PunctStr`:            147,
	`punctstr`:            147,
	`PunctStrDblQuote`:    148,
	`punctstrdblquote`:    148,
	`PunctStrQuote`:       149,
	`punctstrquote`:       149,
	`PunctStrBacktick`:    150,
	`punctstrbacktick`:    150,
	`PunctStrEsc`:         151,
	`punctstresc`:         151,
	`Comment`:             152,
	`comment`:             152,
	`CommentHashbang`:     153,
	`commenthashbang`:     153,
	`CommentMultiline`:    154,
	`commentmultiline`:    154,
	`CommentSingle`:       155,
	`commentsingle`:       155,
	`CommentSpecial`:      156,
	`commentspecial`:      156,
	`CommentPreproc`:      157,
	`commentpreproc`:      157,
	`CommentPreprocFile`:  158,
	`commentpreprocfile`:  158,
	`Text`:                159,
	`text`:                159,
	`TextWhitespace`:      160,
	`textwhitespace`:      160,
	`TextSymbol`:          161,
	`textsymbol`:          161,
	`TextPunctuation`:     162,
	`textpunctuation`:     162,
	`TextSpellErr`:        163,
	`textspellerr`:        163,
	`TextStyle`:           164,
	`textstyle`:           164,
	`TextStyleDeleted`:    165,
	`textstyledeleted`:    165,
	`TextStyleEmph`:       166,
	`textstyleemph`:       166,
	`TextStyleError`:      167,
	`textstyleerror`:      167,
	`TextStyleHeading`:    168,
	`textstyleheading`:    168,
	`TextStyleInserted`:   169,
	`textstyleinserted`:   169,
	`TextStyleOutput`:     170,
	`textstyleoutput`:     170,
	`TextStylePrompt`:     171,
	`textstyleprompt`:     171,
	`TextStyleStrong`:     172,
	`textstylestrong`:     172,
	`TextStyleSubheading`: 173,
	`textstylesubheading`: 173,
	`TextStyleTraceback`:  174,
	`textstyletraceback`:  174,
	`TextStyleUnderline`:  175,
	`textstyleunderline`:  175,
	`TextStyleLink`:       176,
	`textstylelink`:       176,
}

var _TokensDescMap = map[Tokens]string{
	0:   `None is the nil token value -- for non-terminal cases or TBD`,
	1:   `Error is an input that could not be tokenized due to syntax error etc`,
	2:   `EOF is end of file`,
	3:   `EOL is end of line (typically implicit -- used for rule matching)`,
	4:   `EOS is end of statement -- a key meta-token -- in C it is ;, in Go it is either ; or EOL`,
	5:   `Background is for syntax highlight styles based on these tokens`,
	6:   `Cat: Keywords (actual keyword is just the string)`,
	7:   ``,
	8:   ``,
	9:   ``,
	10:  ``,
	11:  ``,
	12:  ``,
	13:  `Cat: Names.`,
	14:  ``,
	15:  ``,
	16:  ``,
	17:  ``,
	18:  `SubCat: Type names`,
	19:  ``,
	20:  ``,
	21:  ``,
	22:  ``,
	23:  ``,
	24:  ``,
	25:  ``,
	26:  ``,
	27:  ``,
	28:  ``,
	29:  ``,
	30:  `SubCat: Function names`,
	31:  ``,
	32:  ``,
	33:  ``,
	34:  ``,
	35:  ``,
	36:  ``,
	37:  ``,
	38:  ``,
	39:  `SubCat: Scoping names`,
	40:  ``,
	41:  ``,
	42:  ``,
	43:  ``,
	44:  `SubCat: NameVar -- variable names`,
	45:  ``,
	46:  ``,
	47:  ``,
	48:  ``,
	49:  ``,
	50:  ``,
	51:  `SubCat: Value -- data-like elements`,
	52:  ``,
	53:  ``,
	54:  ``,
	55:  ``,
	56:  `Cat: Literals.`,
	57:  ``,
	58:  ``,
	59:  ``,
	60:  `SubCat: Literal Strings.`,
	61:  ``,
	62:  ``,
	63:  ``,
	64:  ``,
	65:  ``,
	66:  ``,
	67:  ``,
	68:  ``,
	69:  ``,
	70:  ``,
	71:  ``,
	72:  ``,
	73:  ``,
	74:  ``,
	75:  ``,
	76:  ``,
	77:  ``,
	78:  `SubCat: Literal Numbers.`,
	79:  ``,
	80:  ``,
	81:  ``,
	82:  ``,
	83:  ``,
	84:  ``,
	85:  ``,
	86:  `Cat: Operators.`,
	87:  ``,
	88:  `SubCat: Math operators`,
	89:  ``,
	90:  ``,
	91:  ``,
	92:  ``,
	93:  ``,
	94:  `SubCat: Bitwise operators`,
	95:  ``,
	96:  ``,
	97:  ``,
	98:  ``,
	99:  ``,
	100: ``,
	101: ``,
	102: `SubCat: Assign operators`,
	103: ``,
	104: ``,
	105: ``,
	106: ``,
	107: ``,
	108: `SubCat: Math Assign operators`,
	109: ``,
	110: ``,
	111: ``,
	112: ``,
	113: ``,
	114: `SubCat: Bitwise Assign operators`,
	115: ``,
	116: ``,
	117: ``,
	118: ``,
	119: ``,
	120: ``,
	121: `SubCat: Logical operators`,
	122: ``,
	123: ``,
	124: ``,
	125: `SubCat: Relational operators`,
	126: ``,
	127: ``,
	128: ``,
	129: ``,
	130: ``,
	131: ``,
	132: `SubCat: List operators`,
	133: ``,
	134: `Cat: Punctuation.`,
	135: `SubCat: Grouping punctuation`,
	136: ``,
	137: ``,
	138: ``,
	139: ``,
	140: ``,
	141: ``,
	142: `SubCat: Separator punctuation`,
	143: ``,
	144: ``,
	145: ``,
	146: ``,
	147: `SubCat: String punctuation`,
	148: ``,
	149: ``,
	150: ``,
	151: ``,
	152: `Cat: Comments.`,
	153: ``,
	154: ``,
	155: ``,
	156: ``,
	157: `SubCat: Preprocessor &#34;comments&#34;.`,
	158: ``,
	159: `Cat: Text.`,
	160: ``,
	161: ``,
	162: ``,
	163: ``,
	164: `SubCat: TextStyle (corresponds to Generic in chroma / pygments) todo: look in font deco for more`,
	165: ``,
	166: ``,
	167: ``,
	168: ``,
	169: ``,
	170: ``,
	171: ``,
	172: ``,
	173: ``,
	174: ``,
	175: ``,
	176: ``,
}

var _TokensMap = map[Tokens]string{
	0:   `None`,
	1:   `Error`,
	2:   `EOF`,
	3:   `EOL`,
	4:   `EOS`,
	5:   `Background`,
	6:   `Keyword`,
	7:   `KeywordConstant`,
	8:   `KeywordDeclaration`,
	9:   `KeywordNamespace`,
	10:  `KeywordPseudo`,
	11:  `KeywordReserved`,
	12:  `KeywordType`,
	13:  `Name`,
	14:  `NameBuiltin`,
	15:  `NameBuiltinPseudo`,
	16:  `NameOther`,
	17:  `NamePseudo`,
	18:  `NameType`,
	19:  `NameClass`,
	20:  `NameStruct`,
	21:  `NameField`,
	22:  `NameInterface`,
	23:  `NameConstant`,
	24:  `NameEnum`,
	25:  `NameEnumMember`,
	26:  `NameArray`,
	27:  `NameMap`,
	28:  `NameObject`,
	29:  `NameTypeParam`,
	30:  `NameFunction`,
	31:  `NameDecorator`,
	32:  `NameFunctionMagic`,
	33:  `NameMethod`,
	34:  `NameOperator`,
	35:  `NameConstructor`,
	36:  `NameException`,
	37:  `NameLabel`,
	38:  `NameEvent`,
	39:  `NameScope`,
	40:  `NameNamespace`,
	41:  `NameModule`,
	42:  `NamePackage`,
	43:  `NameLibrary`,
	44:  `NameVar`,
	45:  `NameVarAnonymous`,
	46:  `NameVarClass`,
	47:  `NameVarGlobal`,
	48:  `NameVarInstance`,
	49:  `NameVarMagic`,
	50:  `NameVarParam`,
	51:  `NameValue`,
	52:  `NameTag`,
	53:  `NameProperty`,
	54:  `NameAttribute`,
	55:  `NameEntity`,
	56:  `Literal`,
	57:  `LiteralDate`,
	58:  `LiteralOther`,
	59:  `LiteralBool`,
	60:  `LitStr`,
	61:  `LitStrAffix`,
	62:  `LitStrAtom`,
	63:  `LitStrBacktick`,
	64:  `LitStrBoolean`,
	65:  `LitStrChar`,
	66:  `LitStrDelimiter`,
	67:  `LitStrDoc`,
	68:  `LitStrDouble`,
	69:  `LitStrEscape`,
	70:  `LitStrHeredoc`,
	71:  `LitStrInterpol`,
	72:  `LitStrName`,
	73:  `LitStrOther`,
	74:  `LitStrRegex`,
	75:  `LitStrSingle`,
	76:  `LitStrSymbol`,
	77:  `LitStrFile`,
	78:  `LitNum`,
	79:  `LitNumBin`,
	80:  `LitNumFloat`,
	81:  `LitNumHex`,
	82:  `LitNumInteger`,
	83:  `LitNumIntegerLong`,
	84:  `LitNumOct`,
	85:  `LitNumImag`,
	86:  `Operator`,
	87:  `OperatorWord`,
	88:  `OpMath`,
	89:  `OpMathAdd`,
	90:  `OpMathSub`,
	91:  `OpMathMul`,
	92:  `OpMathDiv`,
	93:  `OpMathRem`,
	94:  `OpBit`,
	95:  `OpBitAnd`,
	96:  `OpBitOr`,
	97:  `OpBitNot`,
	98:  `OpBitXor`,
	99:  `OpBitShiftLeft`,
	100: `OpBitShiftRight`,
	101: `OpBitAndNot`,
	102: `OpAsgn`,
	103: `OpAsgnAssign`,
	104: `OpAsgnInc`,
	105: `OpAsgnDec`,
	106: `OpAsgnArrow`,
	107: `OpAsgnDefine`,
	108: `OpMathAsgn`,
	109: `OpMathAsgnAdd`,
	110: `OpMathAsgnSub`,
	111: `OpMathAsgnMul`,
	112: `OpMathAsgnDiv`,
	113: `OpMathAsgnRem`,
	114: `OpBitAsgn`,
	115: `OpBitAsgnAnd`,
	116: `OpBitAsgnOr`,
	117: `OpBitAsgnXor`,
	118: `OpBitAsgnShiftLeft`,
	119: `OpBitAsgnShiftRight`,
	120: `OpBitAsgnAndNot`,
	121: `OpLog`,
	122: `OpLogAnd`,
	123: `OpLogOr`,
	124: `OpLogNot`,
	125: `OpRel`,
	126: `OpRelEqual`,
	127: `OpRelNotEqual`,
	128: `OpRelLess`,
	129: `OpRelGreater`,
	130: `OpRelLtEq`,
	131: `OpRelGtEq`,
	132: `OpList`,
	133: `OpListEllipsis`,
	134: `Punctuation`,
	135: `PunctGp`,
	136: `PunctGpLParen`,
	137: `PunctGpRParen`,
	138: `PunctGpLBrack`,
	139: `PunctGpRBrack`,
	140: `PunctGpLBrace`,
	141: `PunctGpRBrace`,
	142: `PunctSep`,
	143: `PunctSepComma`,
	144: `PunctSepPeriod`,
	145: `PunctSepSemicolon`,
	146: `PunctSepColon`,
	147: `PunctStr`,
	148: `PunctStrDblQuote`,
	149: `PunctStrQuote`,
	150: `PunctStrBacktick`,
	151: `PunctStrEsc`,
	152: `Comment`,
	153: `CommentHashbang`,
	154: `CommentMultiline`,
	155: `CommentSingle`,
	156: `CommentSpecial`,
	157: `CommentPreproc`,
	158: `CommentPreprocFile`,
	159: `Text`,
	160: `TextWhitespace`,
	161: `TextSymbol`,
	162: `TextPunctuation`,
	163: `TextSpellErr`,
	164: `TextStyle`,
	165: `TextStyleDeleted`,
	166: `TextStyleEmph`,
	167: `TextStyleError`,
	168: `TextStyleHeading`,
	169: `TextStyleInserted`,
	170: `TextStyleOutput`,
	171: `TextStylePrompt`,
	172: `TextStyleStrong`,
	173: `TextStyleSubheading`,
	174: `TextStyleTraceback`,
	175: `TextStyleUnderline`,
	176: `TextStyleLink`,
}

// String returns the string representation
// of this Tokens value.
func (i Tokens) String() string {
	if str, ok := _TokensMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Tokens value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Tokens) SetString(s string) error {
	if val, ok := _TokensNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _TokensNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Tokens")
}

// Int64 returns the Tokens value as an int64.
func (i Tokens) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Tokens value from an int64.
func (i *Tokens) SetInt64(in int64) {
	*i = Tokens(in)
}

// Desc returns the description of the Tokens value.
func (i Tokens) Desc() string {
	if str, ok := _TokensDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TokensValues returns all possible values
// for the type Tokens.
func TokensValues() []Tokens {
	return _TokensValues
}

// Values returns all possible values
// for the type Tokens.
func (i Tokens) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TokensValues))
	for i, d := range _TokensValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Tokens.
func (i Tokens) IsValid() bool {
	_, ok := _TokensMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Tokens) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Tokens) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
