# _,/??:=range

import os,times,system,strutils,nre

proc readTemplate(filename:string,firstNest:int = 0) : string =
  let f = open(filename,FileMode.fmRead)
  let html = f.readAll().replace("\n"," ").replace(re"\s\s+"," ").replace("> <","><")
  f.close()
  result = ""
  let spls = html.split(re"({{|}})")
  var isStringMode = spls[0] != "{{"
  var nest = firstNest
  let indent = 4
  for spl in spls:
    if spl == "{{" :
      isStringMode = false
      continue
    if spl == "}}" :
      isStringMode = true
      continue
    let normalized = spl.strip()
    if normalized == "" : continue
    result &= " ".repeat(nest * indent)
    if isStringMode:
      result &= "w.Write([]byte(`" & spl & "`))\n"
      continue
    let tag = normalized.strip().replace(".","").strip()
    if tag == "end" :
      nest -= 1
      result = result[0..^(indent+1)] & "}\n"
      continue
    if tag.startsWith("if") :
      nest += 1
      result &= tag.replace(".","") & " != \"\" {\n"
      continue
    if tag.startsWith("else"):
      if "if" in tag :
        result = result[0..^(indent+1)] &  "} " & tag.replace(".","") & " != \"\" {\n"
      else:
        result = result[0..^(indent+1)] & "} else {\n"
      continue
    if tag.startsWith("template"):
      let ext = ".tmpl"
      let nextFileName = tag.replace(".","").replace("\"","").replace("template","").strip() & ext
      result &= nextFileName.readTemplate(nest)[(indent * nest)..^1] & "\n"
      continue
    if tag.startsWith("range"):
      nest += 1
      result &= "for _ , ??? := " & tag.replace(".","") & "  {\n"
      continue
    result &= "w.Write([]byte(" & tag & "))\n"

proc toGolang(filename:string) : string =
  result = ("""package main
  import (
    "net/http"
  )
  func """ & filename.replace("/","").replace(".","") & """(rw http.ResponseWriter) {
    w, ok := rw.(http.ResponseWriter);
    if !ok { return }
    w.WriteHeader(http.StatusOK)
    w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
  """).replace("\n\t","\n").replace("\n  ","\n").replace("  ","    ")
  result &= filename.readTemplate(1)
  result &= "}\n"


let args = commandLineParams()
if args.len() == 0 : quit()
let goCode = toGolang(args[0])
echo goCode