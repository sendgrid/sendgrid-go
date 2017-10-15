package plaintext

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFromHTML(t *testing.T) {
    test := "empty html"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html></html>")
        assert.NoError(t, err, test + " conversion error")
        assert.Equal(t, "", plain, test + " conversion failed")
	})

    test = "empty HTML"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<HTML></HTML>")
        assert.NoError(t, err, test + " conversion error")
        assert.Equal(t, "", plain, test + " conversion failed")
	})

    test = "minimal body"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html><body></body></html>")
        assert.NoError(t, err, test + " conversion error")
        assert.Equal(t, "", plain, test + " html conversion failed")
	})

    test = "simple body"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html><body background=\"#808080\">simple</body></html>")
        assert.NoError(t, err, test + " conversion error")
        assert.Equal(t, "simple", plain, test + " html conversion failed")
	})

    test = "line-broken body"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html><body>first<br>second<br/>third</body></html>")
        assert.NoError(t, err, test + " html conversion error")
        assert.Equal(t, "first\n\nsecond\n\nthird", plain, test + " html conversion failed")
	})

    test = "removes script contents"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html><body>non-<script>script</script>script</body></html>")
        assert.NoError(t, err, test + " html conversion error")
        assert.Equal(t, "non-script", plain, test + " html conversion failed")
	})

    test = "removes style contents"
    t.Run(test, func(t *testing.T) {
        plain, err := FromHTMLString("<html><body>non-<style>style</style>style</body></html>")
        assert.NoError(t, err, test + " html conversion error")
        assert.Equal(t, "non-style", plain, test + " html conversion failed")
	})

}
