// Code generated by "stringer -type=Tokens"; DO NOT EDIT.

package token

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[None-0]
	_ = x[Error-1]
	_ = x[EOF-2]
	_ = x[EOL-3]
	_ = x[EOS-4]
	_ = x[Background-5]
	_ = x[Keyword-6]
	_ = x[KeywordConstant-7]
	_ = x[KeywordDeclaration-8]
	_ = x[KeywordNamespace-9]
	_ = x[KeywordPseudo-10]
	_ = x[KeywordReserved-11]
	_ = x[KeywordType-12]
	_ = x[Name-13]
	_ = x[NameBuiltin-14]
	_ = x[NameBuiltinPseudo-15]
	_ = x[NameOther-16]
	_ = x[NamePseudo-17]
	_ = x[NameType-18]
	_ = x[NameClass-19]
	_ = x[NameStruct-20]
	_ = x[NameField-21]
	_ = x[NameInterface-22]
	_ = x[NameConstant-23]
	_ = x[NameEnum-24]
	_ = x[NameEnumMember-25]
	_ = x[NameArray-26]
	_ = x[NameMap-27]
	_ = x[NameObject-28]
	_ = x[NameTypeParam-29]
	_ = x[NameFunction-30]
	_ = x[NameDecorator-31]
	_ = x[NameFunctionMagic-32]
	_ = x[NameMethod-33]
	_ = x[NameOperator-34]
	_ = x[NameConstructor-35]
	_ = x[NameException-36]
	_ = x[NameLabel-37]
	_ = x[NameEvent-38]
	_ = x[NameScope-39]
	_ = x[NameNamespace-40]
	_ = x[NameModule-41]
	_ = x[NamePackage-42]
	_ = x[NameLibrary-43]
	_ = x[NameVar-44]
	_ = x[NameVarAnonymous-45]
	_ = x[NameVarClass-46]
	_ = x[NameVarGlobal-47]
	_ = x[NameVarInstance-48]
	_ = x[NameVarMagic-49]
	_ = x[NameVarParam-50]
	_ = x[NameValue-51]
	_ = x[NameTag-52]
	_ = x[NameProperty-53]
	_ = x[NameAttribute-54]
	_ = x[NameEntity-55]
	_ = x[Literal-56]
	_ = x[LiteralDate-57]
	_ = x[LiteralOther-58]
	_ = x[LiteralBool-59]
	_ = x[LitStr-60]
	_ = x[LitStrAffix-61]
	_ = x[LitStrAtom-62]
	_ = x[LitStrBacktick-63]
	_ = x[LitStrBoolean-64]
	_ = x[LitStrChar-65]
	_ = x[LitStrDelimiter-66]
	_ = x[LitStrDoc-67]
	_ = x[LitStrDouble-68]
	_ = x[LitStrEscape-69]
	_ = x[LitStrHeredoc-70]
	_ = x[LitStrInterpol-71]
	_ = x[LitStrName-72]
	_ = x[LitStrOther-73]
	_ = x[LitStrRegex-74]
	_ = x[LitStrSingle-75]
	_ = x[LitStrSymbol-76]
	_ = x[LitStrFile-77]
	_ = x[LitNum-78]
	_ = x[LitNumBin-79]
	_ = x[LitNumFloat-80]
	_ = x[LitNumHex-81]
	_ = x[LitNumInteger-82]
	_ = x[LitNumIntegerLong-83]
	_ = x[LitNumOct-84]
	_ = x[LitNumImag-85]
	_ = x[Operator-86]
	_ = x[OperatorWord-87]
	_ = x[OpMath-88]
	_ = x[OpMathAdd-89]
	_ = x[OpMathSub-90]
	_ = x[OpMathMul-91]
	_ = x[OpMathDiv-92]
	_ = x[OpMathRem-93]
	_ = x[OpBit-94]
	_ = x[OpBitAnd-95]
	_ = x[OpBitOr-96]
	_ = x[OpBitNot-97]
	_ = x[OpBitXor-98]
	_ = x[OpBitShiftLeft-99]
	_ = x[OpBitShiftRight-100]
	_ = x[OpBitAndNot-101]
	_ = x[OpAsgn-102]
	_ = x[OpAsgnAssign-103]
	_ = x[OpAsgnInc-104]
	_ = x[OpAsgnDec-105]
	_ = x[OpAsgnArrow-106]
	_ = x[OpAsgnDefine-107]
	_ = x[OpMathAsgn-108]
	_ = x[OpMathAsgnAdd-109]
	_ = x[OpMathAsgnSub-110]
	_ = x[OpMathAsgnMul-111]
	_ = x[OpMathAsgnDiv-112]
	_ = x[OpMathAsgnRem-113]
	_ = x[OpBitAsgn-114]
	_ = x[OpBitAsgnAnd-115]
	_ = x[OpBitAsgnOr-116]
	_ = x[OpBitAsgnXor-117]
	_ = x[OpBitAsgnShiftLeft-118]
	_ = x[OpBitAsgnShiftRight-119]
	_ = x[OpBitAsgnAndNot-120]
	_ = x[OpLog-121]
	_ = x[OpLogAnd-122]
	_ = x[OpLogOr-123]
	_ = x[OpLogNot-124]
	_ = x[OpRel-125]
	_ = x[OpRelEqual-126]
	_ = x[OpRelNotEqual-127]
	_ = x[OpRelLess-128]
	_ = x[OpRelGreater-129]
	_ = x[OpRelLtEq-130]
	_ = x[OpRelGtEq-131]
	_ = x[OpList-132]
	_ = x[OpListEllipsis-133]
	_ = x[Punctuation-134]
	_ = x[PunctGp-135]
	_ = x[PunctGpLParen-136]
	_ = x[PunctGpRParen-137]
	_ = x[PunctGpLBrack-138]
	_ = x[PunctGpRBrack-139]
	_ = x[PunctGpLBrace-140]
	_ = x[PunctGpRBrace-141]
	_ = x[PunctSep-142]
	_ = x[PunctSepComma-143]
	_ = x[PunctSepPeriod-144]
	_ = x[PunctSepSemicolon-145]
	_ = x[PunctSepColon-146]
	_ = x[PunctStr-147]
	_ = x[PunctStrDblQuote-148]
	_ = x[PunctStrQuote-149]
	_ = x[PunctStrBacktick-150]
	_ = x[PunctStrEsc-151]
	_ = x[Comment-152]
	_ = x[CommentHashbang-153]
	_ = x[CommentMultiline-154]
	_ = x[CommentSingle-155]
	_ = x[CommentSpecial-156]
	_ = x[CommentPreproc-157]
	_ = x[CommentPreprocFile-158]
	_ = x[Text-159]
	_ = x[TextWhitespace-160]
	_ = x[TextSymbol-161]
	_ = x[TextPunctuation-162]
	_ = x[TextSpellErr-163]
	_ = x[TextStyle-164]
	_ = x[TextStyleDeleted-165]
	_ = x[TextStyleEmph-166]
	_ = x[TextStyleError-167]
	_ = x[TextStyleHeading-168]
	_ = x[TextStyleInserted-169]
	_ = x[TextStyleOutput-170]
	_ = x[TextStylePrompt-171]
	_ = x[TextStyleStrong-172]
	_ = x[TextStyleSubheading-173]
	_ = x[TextStyleTraceback-174]
	_ = x[TextStyleUnderline-175]
	_ = x[TextStyleLink-176]
	_ = x[TokensN-177]
}

const _Tokens_name = "NoneErrorEOFEOLEOSBackgroundKeywordKeywordConstantKeywordDeclarationKeywordNamespaceKeywordPseudoKeywordReservedKeywordTypeNameNameBuiltinNameBuiltinPseudoNameOtherNamePseudoNameTypeNameClassNameStructNameFieldNameInterfaceNameConstantNameEnumNameEnumMemberNameArrayNameMapNameObjectNameTypeParamNameFunctionNameDecoratorNameFunctionMagicNameMethodNameOperatorNameConstructorNameExceptionNameLabelNameEventNameScopeNameNamespaceNameModuleNamePackageNameLibraryNameVarNameVarAnonymousNameVarClassNameVarGlobalNameVarInstanceNameVarMagicNameVarParamNameValueNameTagNamePropertyNameAttributeNameEntityLiteralLiteralDateLiteralOtherLiteralBoolLitStrLitStrAffixLitStrAtomLitStrBacktickLitStrBooleanLitStrCharLitStrDelimiterLitStrDocLitStrDoubleLitStrEscapeLitStrHeredocLitStrInterpolLitStrNameLitStrOtherLitStrRegexLitStrSingleLitStrSymbolLitStrFileLitNumLitNumBinLitNumFloatLitNumHexLitNumIntegerLitNumIntegerLongLitNumOctLitNumImagOperatorOperatorWordOpMathOpMathAddOpMathSubOpMathMulOpMathDivOpMathRemOpBitOpBitAndOpBitOrOpBitNotOpBitXorOpBitShiftLeftOpBitShiftRightOpBitAndNotOpAsgnOpAsgnAssignOpAsgnIncOpAsgnDecOpAsgnArrowOpAsgnDefineOpMathAsgnOpMathAsgnAddOpMathAsgnSubOpMathAsgnMulOpMathAsgnDivOpMathAsgnRemOpBitAsgnOpBitAsgnAndOpBitAsgnOrOpBitAsgnXorOpBitAsgnShiftLeftOpBitAsgnShiftRightOpBitAsgnAndNotOpLogOpLogAndOpLogOrOpLogNotOpRelOpRelEqualOpRelNotEqualOpRelLessOpRelGreaterOpRelLtEqOpRelGtEqOpListOpListEllipsisPunctuationPunctGpPunctGpLParenPunctGpRParenPunctGpLBrackPunctGpRBrackPunctGpLBracePunctGpRBracePunctSepPunctSepCommaPunctSepPeriodPunctSepSemicolonPunctSepColonPunctStrPunctStrDblQuotePunctStrQuotePunctStrBacktickPunctStrEscCommentCommentHashbangCommentMultilineCommentSingleCommentSpecialCommentPreprocCommentPreprocFileTextTextWhitespaceTextSymbolTextPunctuationTextSpellErrTextStyleTextStyleDeletedTextStyleEmphTextStyleErrorTextStyleHeadingTextStyleInsertedTextStyleOutputTextStylePromptTextStyleStrongTextStyleSubheadingTextStyleTracebackTextStyleUnderlineTextStyleLinkTokensN"

var _Tokens_index = [...]uint16{0, 4, 9, 12, 15, 18, 28, 35, 50, 68, 84, 97, 112, 123, 127, 138, 155, 164, 174, 182, 191, 201, 210, 223, 235, 243, 257, 266, 273, 283, 296, 308, 321, 338, 348, 360, 375, 388, 397, 406, 415, 428, 438, 449, 460, 467, 483, 495, 508, 523, 535, 547, 556, 563, 575, 588, 598, 605, 616, 628, 639, 645, 656, 666, 680, 693, 703, 718, 727, 739, 751, 764, 778, 788, 799, 810, 822, 834, 844, 850, 859, 870, 879, 892, 909, 918, 928, 936, 948, 954, 963, 972, 981, 990, 999, 1004, 1012, 1019, 1027, 1035, 1049, 1064, 1075, 1081, 1093, 1102, 1111, 1122, 1134, 1144, 1157, 1170, 1183, 1196, 1209, 1218, 1230, 1241, 1253, 1271, 1290, 1305, 1310, 1318, 1325, 1333, 1338, 1348, 1361, 1370, 1382, 1391, 1400, 1406, 1420, 1431, 1438, 1451, 1464, 1477, 1490, 1503, 1516, 1524, 1537, 1551, 1568, 1581, 1589, 1605, 1618, 1634, 1645, 1652, 1667, 1683, 1696, 1710, 1724, 1742, 1746, 1760, 1770, 1785, 1797, 1806, 1822, 1835, 1849, 1865, 1882, 1897, 1912, 1927, 1946, 1964, 1982, 1995, 2002}

func (i Tokens) String() string {
	if i < 0 || i >= Tokens(len(_Tokens_index)-1) {
		return "Tokens(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Tokens_name[_Tokens_index[i]:_Tokens_index[i+1]]
}

func (i *Tokens) FromString(s string) error {
	for j := 0; j < len(_Tokens_index)-1; j++ {
		if s == _Tokens_name[_Tokens_index[j]:_Tokens_index[j+1]] {
			*i = Tokens(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Tokens")
}
