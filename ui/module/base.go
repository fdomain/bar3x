package module

import (
	"image/draw"

	"github.com/shimmerglass/bar3x/ui"
	"github.com/shimmerglass/bar3x/ui/base"
)

type moduleBase struct {
	base.Base
	Root ui.Drawable
}

func newBase(p ui.ParentDrawable) moduleBase {
	return moduleBase{Base: base.NewBase(p)}
}

func (b *moduleBase) SetContext(ctx ui.Context) {
	b.Base.SetContext(ctx)
	if b.Root != nil {
		b.Root.SetContext(ctx)
	}
}

func (b *moduleBase) Add(ui.Drawable) {}

func (b *moduleBase) Children() []ui.Drawable {
	return []ui.Drawable{b.Root}
}

func (b *moduleBase) ChildContext(int) ui.Context {
	return b.Base.Context()
}

func (b *moduleBase) Width() int {
	return b.Root.Width()
}
func (b *moduleBase) OnWidthChange(c func(int)) {
	b.Root.OnWidthChange(c)
}
func (b *moduleBase) Height() int {
	return b.Root.Height()
}
func (b *moduleBase) OnHeightChange(c func(int)) {
	b.Root.OnHeightChange(c)
}

func (b *moduleBase) Draw(x, y int, im draw.Image) {
	if b.Root == nil {
		return
	}
	b.Root.Draw(x, y, im)
}

func (b *moduleBase) SendEvent(ev ui.Event) bool {
	if !b.Base.SendEvent(ev) {
		return false
	}

	return b.Root.SendEvent(ev)
}
