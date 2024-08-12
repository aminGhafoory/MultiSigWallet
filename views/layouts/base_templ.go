// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

var JSonce = templ.NewOnceHandle()

func Base(title string, navbar templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layouts/base.templ`, Line: 9, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><link rel=\"icon\" type=\"image/x-icon\" href=\"/public/favicon.ico\"><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"/public/styles.css\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/public/apple-touch-icon.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/public/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/public/favicon-16x16.png\"><link rel=\"manifest\" href=\"/site.webmanifest\"><script src=\"https://cdn.jsdelivr.net/npm/uikit@3.21.5/dist/js/uikit.min.js\"></script><script src=\"https://cdn.jsdelivr.net/npm/uikit@3.21.5/dist/js/uikit-icons.min.js\"></script><script src=\"https://unpkg.com/htmx.org@1.9.9\" defer></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\t\t\t\tif (\n\t\t\t\t\tlocalStorage.getItem(\"color-theme\") === \"dark\" ||\n\t\t\t\t\t(!(\"color-theme\" in localStorage) &&\n\t\t\t\t\twindow.matchMedia(\"(prefers-color-scheme: dark)\").matches)\n\t\t\t\t) {\n\t\t\t\t\tdocument.documentElement.classList.add(\"dark\");\n\t\t\t\t} else {\n\t\t\t\t\tdocument.documentElement.classList.remove(\"dark\");\n\t\t\t\t}\n\t\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = JSonce.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"antialiased\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = navbar.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body><script>\n\t\t\tvar themeToggleBtn = document.getElementById(\"theme-toggle\");\n\t\t\t\n\t\t\tthemeToggleBtn.addEventListener(\"click\", function () {\n\t\t\t\t// if set via local storage previously\n\t\t\t\tif (localStorage.getItem(\"color-theme\")) {\n\t\t\t\tif (localStorage.getItem(\"color-theme\") === \"light\") {\n\t\t\t\t\tdocument.documentElement.classList.add(\"dark\");\n\t\t\t\t\tlocalStorage.setItem(\"color-theme\", \"dark\");\n\t\t\t\t} else {\n\t\t\t\t\tdocument.documentElement.classList.remove(\"dark\");\n\t\t\t\t\tlocalStorage.setItem(\"color-theme\", \"light\");\n\t\t\t\t}\n\t\t\t\n\t\t\t\t// if NOT set via local storage previously\n\t\t\t\t} else {\n\t\t\t\tif (document.documentElement.classList.contains(\"dark\")) {\n\t\t\t\t\tdocument.documentElement.classList.remove(\"dark\");\n\t\t\t\t\tlocalStorage.setItem(\"color-theme\", \"light\");\n\t\t\t\t} else {\n\t\t\t\t\tdocument.documentElement.classList.add(\"dark\");\n\t\t\t\t\tlocalStorage.setItem(\"color-theme\", \"dark\");\n\t\t\t\t}\n\t\t\t\t}\n\t\t\t});\n\t\t\t</script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
