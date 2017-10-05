package textprocess

import "testing"

func TestHtml2Text(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in: `<p class=MsoNormal><font size=2 face=Arial><span lang=EN-US style='font-size:10.0pt;font-family:Arial'>Thanks a lot for <span class=SpellE>YADA</span> yada <span class=SpellE>YADA</span> yada.<o:p></o:p></span></font></p> <p class=MsoNormal><font size=2 face=Arial><span lang=EN-US style='font-size:10.0pt;font-family:Arial'><o:p>&nbsp;</o:p></span></font></p> <p class=MsoNormal>
<font size=2 face=Arial><span lang=EN-US style='font-size: 10.0pt;font-family:Arial'>New Line.<o:p></o:p></span></font></p>`,
			out: `Thanks a lot for YADA yada YADA yada.
New Line.
`,
		},
		{
			in: `<table> <tr> <td align="center"><img src="logo.jpg"></td> </tr> <tr>
               <td>Content</td> </tr> <tr>

               <td align="center">Address</td> </tr> </table>`,
			out: `Content
Address
`,
		},
	}

	for i, tt := range tests {
		out := Html2Txt(tt.in)
		if tt.out != out {
			t.Errorf("Test %d failed: expected <%s> but got <%s>", i, tt.out, out)
		}
	}
}
