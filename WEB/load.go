// Page struct, which will contain template
type Page struct {
    Title    string
    Body     template.HTML
    UserData template.HTML
    Bar      template.HTML
}
 
// Loads a page for use
func loadPage(title string, r *http.Request) (*Page, error) {
 
    filename, option, usr, bar := user.LoadUserInfo(title, r)
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
 
    return &Page{Title: title, Body: template.HTML(body), UserData: (template.HTML(usr) + template.HTML(option)), Bar: template.HTML(bar)}, nil
}
