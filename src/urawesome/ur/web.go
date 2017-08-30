package ur

import (
    "bytes"
    "html/template"
)

func WebView(g *Game) string {
    funcs := template.FuncMap{
        "isMod": func(a, b int) bool { return a % b == 0 },
        "loop": func(n uint) []struct{} { return make([]struct{}, n) },
    }

    t, err := template.New("ur.html").Funcs(funcs).ParseFiles("views/ur.html")
    if err != nil {
        panic(err)
    }

    var buf bytes.Buffer
    err = t.Execute(&buf, g)
    if err != nil {
        panic(err)
    }
    return buf.String()
}

