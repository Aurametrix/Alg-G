package main

import (
  "github.com/matthewmueller/joy/dom/htmlbodyelement"
  "github.com/matthewmueller/joy/dom/window"
)

func main() {
  document := window.New().Document()

  a := document.CreateElement("a")
  println(a.NodeName())
  strong := document.CreateElement("strong")
  println(document.CreateElement("strong").OuterHTML())
  a.AppendChild(strong)

  strong.SetTextContent("hi world!")

  body := document.Body().(*htmlbodyelement.HTMLBodyElement)
  body.AppendChild(a)
  println(document.Body().OuterHTML())
}

            
 ;(function() {
  var pkg = {};
  pkg["52-basic-dom"] = (function() {
    function main () {
      var document = window.document;
      var a = document.createElement("a");
      console.log(a.nodeName);
      var strong = document.createElement("strong");
      console.log(document.createElement("strong").outerHTML);
      a.appendChild(strong);
      strong.textContent = "hi world!";
      var body = document.body;
      body.appendChild(a);
      console.log(document.body.outerHTML)
    };
    return {
      main: main
    };
  })();
  return pkg["52-basic-dom"].main();
})()
            
