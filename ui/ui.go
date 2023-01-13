package ui

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
)

// nolint: gochecknoglobals
var _metaCN = map[string]string{
	"Unknown":     "未知",
	"Image":       "图片",
	"Video":       "视频",
	"Audio":       "音频",
	"Archive":     "归档",
	"Documents":   "文档",
	"Font":        "字体",
	"Application": "应用",
}

type UI struct {
	cfg      *pb.Config
	app      fyne.App
	main     fyne.Window
	tabs     *container.AppTabs
	commands []Command
}

func NewUI(cfg *pb.Config) *UI {
	ret := &UI{
		cfg: cfg,
		app: app.New(),
	}

	ret.main = ret.app.NewWindow("Fairy")
	ret.app.Settings().SetTheme(NewTheme())
	ret.createCommands()
	ret.tabs = ret.createTabs()

	return ret
}

func (p *UI) Run() {
	border := container.NewBorder(p.createToolbar(), nil, nil, nil, p.tabs)
	// nolint: gomnd
	p.main.Resize(fyne.NewSize(800, 600))
	p.main.SetContent(border)
	p.main.CenterOnScreen()
	p.main.Show()

	p.app.Run()
}

func (p *UI) createTabs() *container.AppTabs {
	items := make([]*container.TabItem, len(p.cfg.Group))

	for index, group := range p.cfg.Group {
		items[index] = p.group2TabItem(group)
	}

	tabs := container.NewAppTabs(items...)
	// tabs.SetTabLocation(container.TabLocationBottom)
	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}

func (p *UI) group2TabItem(group *pb.Group) *container.TabItem {
	items := make([]*widget.FormItem, len(pb.Meta_name))
	entries := make([]*widget.Entry, len(pb.Meta_name))

	var selected *widget.Entry

	for index, name := range pb.Meta_name {
		entry := widget.NewEntry()
		entry.SetPlaceHolder("目标目录")
		entry.OnCursorChanged = func() {
			if entry.Text != "" || selected == entry {
				return
			}

			selected = entry

			ShowFolderOpen(func(uri fyne.ListableURI, err error) {
				if err != nil || uri == nil {
					entry.SetText("")

					return
				}

				entry.SetText(uri.Path())
			}, p.main)
		}

		if path, has := group.Meta[name]; has {
			entry.SetText(path)
		}

		if chinese, has := _metaCN[name]; has {
			name = chinese
		}

		entries[index] = entry
		items[index] = &widget.FormItem{Text: name, Widget: entry}
	}

	form := &widget.Form{
		Items:      items,
		SubmitText: "执行",
		OnSubmit: func() {
			for index, entry := range entries {
				if entry.Text == "" {
					continue
				}

				group.Meta[pb.Meta_name[int32(index)]] = entry.Text
			}

			p.Save()
			// TODO 检查目录
			// TODO 执行文件迁移
		},
	}

	return container.NewTabItem(group.Watch, form)
}

func (p *UI) createToolbar() *widget.Toolbar {
	items := []widget.ToolbarItem{}

	for _, command := range p.commands {
		if command.Help == "Separator" {
			items = append(items, widget.NewToolbarSeparator())

			continue
		}

		if command.Help == "Spacer" {
			items = append(items, widget.NewToolbarSpacer())

			continue
		}

		if command.Icon == nil {
			continue
		}

		items = append(items, widget.NewToolbarAction(command.Icon, command.Call))
	}

	return widget.NewToolbar(items...)
}

func ShowFolderOpen(callback func(fyne.ListableURI, error), parent fyne.Window) {
	dialog := dialog.NewFolderOpen(callback, parent)

	dialog.SetConfirmText("选择目录")
	dialog.SetDismissText("取消")
	dialog.Show()
}

func (p *UI) createGroup() {
	ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, p.main)

			return
		}

		if uri == nil || uri.Path() == "" {
			return
		}

		for _, group := range p.cfg.Group {
			if group.Watch == uri.Path() {
				dialog.ShowInformation("路径重复", group.Watch, p.main)

				return
			}
			// TODO 子目录校验
		}

		logs.Info(uri.Path())

		group := &pb.Group{Meta: map[string]string{}, Watch: uri.Path()}
		p.cfg.Group = append(p.cfg.Group, group)
		p.Save()
		item := p.group2TabItem(group)
		p.tabs.Append(item)
		p.tabs.Select(item)
	}, p.main)
}

func (p *UI) Save() {
	logs.Debug("save")

	config := viper.ConfigFileUsed()

	if config == "" {
		home := base.Must1(homedir.Dir())
		config = filepath.Join(home, "fairy.toml")
	}

	file := base.Must1(os.Create(config))
	defer file.Close()

	encoder := toml.NewEncoder(file)
	_ = encoder.Encode(p.cfg)
}

func (p *UI) createCommands() {
	p.commands = []Command{
		// {Help: "First Page", Call: p.first, Icon: theme.MediaSkipPreviousIcon(), Key1: fyne.KeyHome},
		// {Help: "Previous Page", Call: p.previous, Icon: theme.NavigateBackIcon(), Key1: fyne.KeyPageUp},
		// {Help: "Next Page", Call: p.next, Icon: theme.NavigateNextIcon(), Key1: fyne.KeyPageDown, Key2: fyne.KeySpace},
		// {Help: "Last Page", Call: p.last, Icon: theme.MediaSkipNextIcon(), Key1: fyne.KeyEnd},
		{Help: "Create", Call: p.createGroup, Icon: theme.FolderNewIcon(), Key1: fyne.KeyQ},
		{Help: "Separator"},
		// {Help: "Full screen", Call: p.fullScreen, Icon: theme.ComputerIcon(), Key1: fyne.KeyF11},
		// {Help: "Actual Size", Call: p.modeActual, Icon: theme.ViewRestoreIcon(), Key1: fyne.KeyAsterisk, Key2: fyne.Key8},
		// {Help: "Fit to width", Call: p.modeWidth, Icon: theme.MoreHorizontalIcon(), Key1: fyne.KeyW},
		// {Help: "Fit to height", Call: p.modeHeight, Icon: theme.MoreVerticalIcon(), Key1: fyne.KeyH},
		// {Help: "Zoom In", Call: p.zoomIn, Icon: theme.ZoomInIcon(), Key1: fyne.KeyPlus, Key2: fyne.KeyEqual},
		// {Help: "Zoom Out", Call: p.zoomOut, Icon: theme.ZoomOutIcon(), Key1: fyne.KeyMinus},
		// {Help: "Rotate", Call: p.rotate, Icon: theme.ViewRefreshIcon(), Key1: fyne.KeyR},
		// {Help: "Close", Call: p.close, Key1: fyne.KeyEscape},
		// {Help: "To Top", Call: p.top, Key1: fyne.KeyUp},
		// {Help: "To Bottom", Call: p.bottom, Key1: fyne.KeyDown},
		{Help: "Quit", Call: p.app.Quit, Key1: fyne.KeyQ},
		{Help: "Spacer"},
		// {Help: "This Help", Call: p.showHelp, Icon: theme.HelpIcon(), Key1: fyne.KeyF1, Key2: fyne.KeyF10},
		{Help: "Quit", Call: p.app.Quit, Icon: theme.ContentRemoveIcon(), Key1: fyne.KeyQ},
	}
}
