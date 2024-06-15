package schemax

import (
	"bytes"
	"text/template"
)

func newTemplate(name string) *template.Template {
	return template.New(name)
}

func newBuf() *bytes.Buffer {
	return new(bytes.Buffer)
}

func funcMap(fm map[string]any) template.FuncMap {
	return template.FuncMap(fm)
}

const lDAPSyntaxTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $hindent:=.HIndent -}}
{{- $descl:="DESC " -}}
{{- $numOID:=.Definition.OID -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $numOID -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $extn -}}{{- $extn -}}{{- end -}}
{{$close -}}`

const matchingRuleTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $name:=.Definition.Name -}}
{{- $hindent:=.HIndent -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:="DESC " -}}
{{- $namel:="NAME " -}}
{{- $stxl:="SYNTAX " -}}
{{- $numOID:=.Definition.OID -}}
{{- $desc:=.Definition.Desc -}}
{{- $obs:=(Obsolete) -}}
{{- $obsl:="OBSOLETE" -}}
{{- $sytx:=(Syntax) -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{if $sytx -}}
{{- $hindent -}}{{- $stxl -}} {{- $sytx -}}
{{- end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`

const attributeTypeTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $namel:="NAME " -}}
{{- $name:=.Definition.Name -}}
{{- $extn:=(ExtensionSet) -}}
{{- $hindent:=.HIndent -}}
{{- $descl:="DESC " -}}
{{- $eqll:="EQUALITY " -}}
{{- $subl:="SUBSTR " -}}
{{- $ordl:="ORDERING " -}}
{{- $obs:=(Obsolete) -}}
{{- $collective:=(Collective) -}}
{{- $single:=(IsSingleVal) -}}
{{- $nousermod:=(IsNoUserMod) -}}
{{- $usagel:="USAGE " -}}
{{- $usage:=Usage -}}
{{- $desc:=.Definition.Desc -}}
{{- $stxl:="SYNTAX " -}}
{{- $sytx:=(Syntax) -}}
{{- $eql:=(Equality) -}}
{{- $ord:=(Ordering) -}}
{{- $sub:=(Substring) -}}
{{- $supl:="SUP " -}}
{{- $obsl:="OBSOLETE" -}}
{{- $sv:="SINGLE-VALUE" -}}
{{- $coll:="COLLECTIVE" -}}
{{- $nomod:="NO-USER-MODIFICATIONS" -}}
{{- $sup:=(SuperType) -}}
{{- $numOID:=.Definition.OID -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{if $sup -}}
{{- $hindent -}}{{- $supl -}} {{- $sup -}}
{{- end -}}
{{if $eql -}}
{{- $hindent -}}{{- $eqll -}} {{- $eql -}}
{{- end -}}
{{if $sub -}}
{{- $hindent -}}{{- $subl -}} {{- $sub -}}
{{- end -}}
{{if $ord -}}
{{- $hindent -}}{{- $ordl -}} {{- $ord -}}
{{- end -}}
{{if $sytx -}}
{{- $hindent -}}{{- $stxl -}} {{- $sytx -}}
{{- end -}}
{{if $single -}}
{{- $hindent -}}{{- $sv -}}
{{- else if $collective -}}
{{- $hindent -}}{{- $coll -}}
{{- end -}}
{{if $nousermod -}}
{{- $hindent -}}{{- $nomod -}}
{{- end -}}
{{if $usage -}}
{{- $hindent -}}{{- $usagel -}} {{- $usage -}}
{{- end -}}
{{if $extn -}}{{- $extn}}{{- end -}}
{{$close -}}`

const matchingRuleUseTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $name:=.Definition.Name -}}
{{- $hindent:=.HIndent -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:="DESC " -}}
{{- $namel:="NAME " -}}
{{- $appl:="APPLIES " -}}
{{- $numOID:=(MatchingRuleOID) -}}
{{- $applied:=Applied -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $applied -}}
{{- $hindent -}}{{- $appl -}}{{- $applied -}}
{{end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`

const objectClassTmpl = `{{- $name:=.Definition.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $obs:=(Obsolete) -}}
{{- $hindent:=.HIndent -}}
{{- $obsl:="OBSOLETE " -}}
{{- $kind:=(Kind) -}}
{{- $namel:="NAME " -}}
{{- $descl:="DESC " -}}
{{- $mustl:="MUST " -}}
{{- $mayl:="MAY " -}}
{{- $supl:="SUP " -}}
{{- $suplen:=SuperLen -}}
{{- $mustlen:=MustLen -}}
{{- $maylen:=MayLen -}}
{{- $id:=.Definition.OID -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent  -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{if gt $suplen 0 -}}
{{- $hindent -}}{{- $supl -}}{{- .Definition.SuperClasses -}}
{{end -}}
{{- $hindent -}}{{- $kind -}}
{{if gt $mustlen 0 -}}
{{- $hindent -}}{{- $mustl -}}{{- .Definition.Must -}}
{{end -}}
{{if gt $maylen 0 -}}
{{- $hindent -}}{{- $mayl -}}{{- .Definition.May -}}
{{end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`

const dITContentRuleTmpl = `{{- $name:=.Definition.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $hindent:=.HIndent -}}
{{- $obs:=(Obsolete) -}}
{{- $obsl:="OBSOLETE " -}}
{{- $namel:="NAME " -}}
{{- $descl:="DESC " -}}
{{- $mustl:="MUST " -}}
{{- $mayl:="MAY " -}}
{{- $notl:="NOT " -}}
{{- $auxl:="AUX " -}}
{{- $auxlen:=AuxLen -}}
{{- $mustlen:=MustLen -}}
{{- $maylen:=MayLen -}}
{{- $notlen:=NotLen -}}
{{- $id:=(StructuralOID) -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{if gt $auxlen 0 -}}
{{- $hindent -}}{{- $auxl -}}{{- .Definition.Aux -}}
{{end -}}
{{if gt $mustlen 0 -}}
{{- $hindent -}}{{- $mustl -}}{{- .Definition.Must -}}
{{end -}}
{{if gt $maylen 0 -}}
{{- $hindent -}}{{- $mayl -}}{{- .Definition.May -}}
{{end -}}
{{if gt $notlen 0 -}}
{{- $hindent -}}{{- $notl -}}{{- .Definition.Not -}}
{{end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`

const nameFormTmpl = `{{- $name:=.Definition.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $hindent:=.HIndent -}}
{{- $obs:=(Obsolete) -}}
{{- $obsl:="OBSOLETE " -}}
{{- $namel:="NAME " -}}
{{- $descl:="DESC " -}}
{{- $mustl:="MUST " -}}
{{- $mayl:="MAY " -}}
{{- $ocl:="OC " -}}
{{- $oc:=.Definition.Structural.OID -}}
{{- $maylen:=MayLen -}}
{{- $id:=.Definition.OID -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{- $hindent -}}{{- $ocl -}}{{- $oc -}}
{{- $hindent -}}{{- $mustl -}}{{- .Definition.Must -}}
{{if gt $maylen 0 -}}
{{- $hindent -}}{{- $mayl -}}{{- .Definition.May -}}
{{end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`

const dITStructureRuleTmpl = `{{- $name:=.Definition.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $obs:=(Obsolete) -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $extn:=(ExtensionSet) -}}
{{- $hindent:=.HIndent -}}
{{- $namel:="NAME " -}}
{{- $descl:="DESC " -}}
{{- $forml:="FORM " -}}
{{- $supl:="SUP " -}}
{{- $suplen:=SuperLen -}}
{{- $id:=.Definition.ID -}}
{{- $form:=.Definition.Form.OID -}}
{{- $desc:=.Definition.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $hindent -}}{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $hindent -}}{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $hindent -}}{{- $obsl -}}
{{end -}}
{{- $hindent -}}{{- $forml -}}{{- $form -}}
{{if gt $suplen 0 -}}
{{- $hindent -}}{{- $supl -}}{{- .Definition.SuperRules -}}
{{end -}}
{{if $extn -}}{{$extn}}{{- end -}}
{{$close -}}`
