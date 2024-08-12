// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package newContract

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/aminghafoory/multisig2/views/components"
	"github.com/aminghafoory/multisig2/views/layouts"
	"net/http"
)

func NewContractPage(title string, r *http.Request) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"uk-child-width-expand@s uk-text-center \"><div class=\"uk-flex-center-auto py-10 px-10\"><div class=\"uk-card uk-card-body uk-card-default uk-margin-left uk-margin-right\"><form action=\"/contracts\" method=\"post\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.CSRFfield(r).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<fieldset class=\"uk-fieldset\"><legend class=\"uk-legend uk-h1 \">NewContract</legend><div class=\"uk-margin\"><input class=\"uk-input\" type=\"text\" id=\"contract-name\" name=\"contract-name\" placeholder=\"contract name\" aria-label=\"Input\" required></div><div class=\"uk-margin\"><input class=\"uk-input\" type=\"text\" id=\"nconfirm\" name=\"nconfirm\" placeholder=\"number of confirms\" aria-label=\"Input\" required></div><div class=\"uk-margin\"><input class=\"uk-input\" type=\"text\" id=\"owners\" name=\"owners\" placeholder=\"owners\" aria-label=\"Input\" required></div><div class=\"uk-margin\"><input class=\"uk-input\" type=\"text\" id=\"wallet_id\" name=\"wallet_id\" placeholder=\"wallet_id\" aria-label=\"Input\" required></div><div class=\"uk-margin\"><input class=\"uk-input\" type=\"password\" name=\"password\" id=\"password\" placeholder=\"password\" aria-label=\"Input\" required></div><button class=\"uk-button uk-button-primary \" type=\"submit\">submit</button></fieldset></form></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base(title, components.Navbar(r)).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

//
//
// password                  string
// 		ContractName              string
// 		Owners                    []common.Address
// 		NumOfConfirmationRequired *big.Int
