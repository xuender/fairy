package ui

import (
	"image/color"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
	"github.com/xuender/oils/logs"
)

type Theme struct {
	regular, bold, italic, boldItalic, monospace fyne.Resource
}

var _fonts = [...]string{"wqy-zenhei", "wqy", "simhei.ttf", "simkai.ttf", "simsun.ttf", "STHeiti", "ukai.ttc", "uming.ttc"}

func NewTheme() *Theme {
	fontPaths := findfont.List()

	for _, f := range _fonts {
		for _, path := range fontPaths {
			if strings.Contains(strings.ToLower(path), f) {
				os.Setenv("FYNE_FONT", path)
				// os.Setenv("FYNE_FONT_MONOSPACE", path)
				logs.Info("字体初始化成功", path)

				ret := &Theme{}
				ret.SetFonts(path, "")

				return ret
			}
		}
	}

	return &Theme{}
}

func (t *Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return t.monospace
	}

	if style.Bold {
		if style.Italic {
			return t.boldItalic
		}

		return t.bold
	}

	if style.Italic {
		return t.italic
	}

	return t.regular
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (t *Theme) SetFonts(regularFontPath string, monoFontPath string) {
	t.regular = theme.TextFont()
	t.bold = theme.TextBoldFont()
	t.italic = theme.TextItalicFont()
	t.boldItalic = theme.TextBoldItalicFont()
	t.monospace = theme.TextMonospaceFont()

	if regularFontPath != "" {
		t.regular = loadCustomFont(regularFontPath, "Regular", t.regular)
		t.bold = loadCustomFont(regularFontPath, "Bold", t.bold)
		t.italic = loadCustomFont(regularFontPath, "Italic", t.italic)
		t.boldItalic = loadCustomFont(regularFontPath, "BoldItalic", t.boldItalic)
	}

	if monoFontPath != "" {
		t.monospace = loadCustomFont(monoFontPath, "Regular", t.monospace)
	} else {
		t.monospace = t.regular
	}
}

func loadCustomFont(env, variant string, fallback fyne.Resource) fyne.Resource {
	variantPath := strings.ReplaceAll(env, "Regular", variant)

	res, err := fyne.LoadResourceFromPath(variantPath)
	if err != nil {
		logs.Error("Error loading specified font", err)

		return fallback
	}

	return res
}
