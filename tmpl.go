package schemax

const lDAPSyntaxTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:=" DESC " -}}
{{- $numOID:=.OID -}}
{{- $desc:=.Desc -}}
{{- $open -}}{{- $numOID -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const matchingRuleTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $name:=.Name -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:=" DESC " -}}
{{- $namel:=" NAME " -}}
{{- $stxl:=" SYNTAX " -}}
{{- $numOID:=.OID -}}
{{- $desc:=.Desc -}}
{{- $sytx:=.Syntax.OID -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $sytx -}}
{{- $stxl -}} {{- $sytx -}}
{{- end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const attributeTypeTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $namel:=" NAME " -}}
{{- $name:=.Name -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:=" DESC " -}}
{{- $eqll:=" EQUALITY " -}}
{{- $subl:=" SUBSTR " -}}
{{- $ordl:=" ORDERING " -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $obs:=(IsObsolete) -}}
{{- $usagel:=" USAGE " -}}
{{- $usage:=Usage -}}
{{- $desc:=.Desc -}}
{{- $stxl:=" SYNTAX " -}}
{{- $sytx:=.Syntax.OID -}}
{{- $eql:=.Equality.OID -}}
{{- $ord:=.Ordering.OID -}}
{{- $sub:=.Substring.OID -}}
{{- $supl:=" SUP " -}}
{{- $sup:=.SuperType.OID -}}
{{- $numOID:=.OID -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $obsl -}}
{{end -}}
{{if $sup -}}
{{- $supl -}} {{- $sup -}}
{{- end -}}
{{if $sytx -}}
{{- $stxl -}} {{- $sytx -}}
{{- end -}}
{{if $eql -}}
{{- $eqll -}} {{- $eql -}}
{{- end -}}
{{if $sub -}}
{{- $subl -}} {{- $sub -}}
{{- end -}}
{{if $ord -}}
{{- $ordl -}} {{- $ord -}}
{{- end -}}
{{if $usage -}}
{{- $usagel -}} {{- $usage -}}
{{- end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const matchingRuleUseTmpl = `{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $name:=.Name -}}
{{- $extn:=(ExtensionSet) -}}
{{- $descl:=" DESC " -}}
{{- $namel:=" NAME " -}}
{{- $appl:=" APPLIES " -}}
{{- $numOID:=.OID -}}
{{- $desc:=.Desc -}}
{{- $apps:=.Applies -}}
{{- $open -}}{{- $numOID -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $apps -}}
{{- $appl -}} {{- $apps -}}
{{- end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const objectClassTmpl = `{{- $name:=.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $obs:=(IsObsolete) -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $namel:=" NAME " -}}
{{- $descl:=" DESC " -}}
{{- $mustl:=" MUST " -}}
{{- $mayl:=" MAY " -}}
{{- $supl:=" SUP " -}}
{{- $suplen:=SuperLen -}}
{{- $mustlen:=MustLen -}}
{{- $maylen:=MayLen -}}
{{- $id:=.OID -}}
{{- $desc:=.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $obsl -}}
{{end -}}
{{if gt $suplen 0 -}}
{{- $supl -}}{{- .SuperClasses -}}
{{end -}}
{{if gt $mustlen 0 -}}
{{- $mustl -}}{{- .Must -}}
{{end -}}
{{if gt $maylen 0 -}}
{{- $mayl -}}{{- .May -}}
{{end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const dITContentRuleTmpl = `{{- $name:=.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $obs:=(IsObsolete) -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $namel:=" NAME " -}}
{{- $descl:=" DESC " -}}
{{- $mustl:=" MUST " -}}
{{- $mayl:=" MAY " -}}
{{- $notl:=" NOT " -}}
{{- $auxl:=" AUX " -}}
{{- $auxlen:=AuxLen -}}
{{- $mustlen:=MustLen -}}
{{- $maylen:=MayLen -}}
{{- $notlen:=NotLen -}}
{{- $id:=(StructuralOID) -}}
{{- $desc:=.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $obsl -}}
{{end -}}
{{if gt $auxlen 0 -}}
{{- $auxl -}}{{- .Aux -}}
{{end -}}
{{if gt $mustlen 0 -}}
{{- $mustl -}}{{- .Must -}}
{{end -}}
{{if gt $maylen 0 -}}
{{- $mayl -}}{{- .May -}}
{{end -}}
{{if gt $notlen 0 -}}
{{- $notl -}}{{- .Not -}}
{{end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const nameFormTmpl = `{{- $name:=.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $extn:=(ExtensionSet) -}}
{{- $obs:=(IsObsolete) -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $namel:=" NAME " -}}
{{- $descl:=" DESC " -}}
{{- $mustl:=" MUST " -}}
{{- $mayl:=" MAY " -}}
{{- $ocl:=" OC " -}}
{{- $oc:=.Structural.OID -}}
{{- $maylen:=MayLen -}}
{{- $id:=.OID -}}
{{- $desc:=.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $obsl -}}
{{end -}}
{{- $ocl -}}{{- $oc -}}
{{- $mustl -}}{{- .Must -}}
{{if gt $maylen 0 -}}
{{- $mayl -}}{{- .May -}}
{{end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`

const dITStructureRuleTmpl = `{{- $name:=.Name -}}
{{- $open:="( " -}}
{{- $close:=" )" -}}
{{- $obs:=(IsObsolete) -}}
{{- $obsl:=" OBSOLETE " -}}
{{- $extn:=(ExtensionSet) -}}
{{- $namel:=" NAME " -}}
{{- $descl:=" DESC " -}}
{{- $forml:=" FORM " -}}
{{- $supl:=" SUP " -}}
{{- $suplen:=SuperLen -}}
{{- $id:=.ID -}}
{{- $form:=.Form.OID -}}
{{- $desc:=.Desc -}}
{{- $open -}}{{- $id -}}
{{if $name -}}
{{- $namel -}}{{- $name -}}
{{- end -}}
{{if $desc -}}
{{- $descl -}}'{{- $desc -}}'
{{- end -}}
{{if $obs -}}
{{- $obsl -}}
{{end -}}
{{- $forml -}}{{- $form -}}
{{if gt $suplen 0 -}}
{{- $supl -}}{{- .SuperRules -}}
{{end -}}
{{if $extn -}} {{$extn}}{{- end -}}
{{$close -}}`
