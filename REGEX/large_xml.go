decoder := xml.NewDecoder(xmlFile) 

for { 
    // Read tokens from the XML document in a stream. 
    t, _ := decoder.Token() 
    if t == nil { 
        break 
    } 
    // Inspect the type of the token just read. 
    switch se := t.(type) { 
    case xml.StartElement: 
        // If we just read a StartElement token 
        // ...and its name is "page" 
        if se.Name.Local == "page" { 
            var p Page 
            // decode a whole chunk of following XML into the
            // variable p which is a Page (se above) 
            decoder.DecodeElement(&p, &se) 
            // Do some stuff with the page. 
            p.Title = CanonicalizeTitle(p.Title)
            ...
        } 
...
