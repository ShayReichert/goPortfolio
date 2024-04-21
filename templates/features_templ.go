// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func features() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"competences\"><!--begin section-white--><div class=\"section-white small-padding-bottom\"><!--begin container--><div class=\"container\"><!--begin row--><div class=\"row margin-bottom-30\"><!--begin col-md-12--><div class=\"col-md-12 text-center\"><h2 class=\"section-title\">Mes Compétences</h2><div class=\"separator_wrapper\"><i class=\"icon icon-star-two red\"></i></div><p class=\"section-subtitle\">There are many variations of passages of Lorem Ipsum available, but the majority<br>have suffered alteration, by injected humour, or new randomised words.</p></div><!--end col-md-12--></div><!--end row--><!--begin row--><div class=\"row\"><!--begin col-sm-4--><div class=\"col-sm-4 features-item\"><div class=\"icon-wrapper\"><i class=\"icon icon-settings-streamline-2 features-icon\"></i></div><h3>Fully Customizable</h3><p>Curabitur quam etsum lacus etsumis nulat iaculis ets vitae etsum nisle varius sed feugiat ligula aliquam ets vitae dictis netsum.</p></div><!--end col-sm-4--><!--begin col-sm-4--><div class=\"col-sm-4 features-item\"><div class=\"icon-wrapper\"><i class=\"icon icon-streamline-umbrella-weather features-icon\"></i></div><h3>Responsive Design</h3><p>Curabitur quam etsum lacus etsumis nulat iaculis ets vitae etsum nisle varius sed feugiat ligula aliquam ets vitae dictis netsum.</p></div><!--end col-sm-4--><!--begin col-sm-4--><div class=\"col-sm-4 features-item\"><div class=\"icon-wrapper\"><i class=\"icon icon-speech-streamline-talk-user features-icon\"></i></div><h3>SEO Ready Code</h3><p>Curabitur quam etsum lacus etsumis nulat iaculis ets vitae etsum nisle varius sed feugiat ligula aliquam ets vitae dictis netsum.</p></div><!--end col-sm-4--></div><!--end row--></div><!--end container--></div><!--end section-white--></section><!--end features-->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
