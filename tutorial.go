package tutorial

import (
	"golang.org/x/image/colornames"
	"gomatcha.io/matcha/bridge"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
	"image/color"
)

func init() {
	// Registers a function with the objc bridge. This function returns
	// a view.View, which can be displayed in a MatchaViewController.
	bridge.RegisterFunc("github.com/overcyn/tutorial New", func() view.View {
		// Call the TutorialView initializer.
		v := NewTutorialView()
		v.TextColor = colornames.Red
		return v
	})
}

// Here is our root view.
type TutorialView struct {
	// All components must implement the view.View interface. A basic implementation
	// is provided by view.Embed.
	view.Embed
	TextColor color.Color
}

// This is our view's initializer.
func NewTutorialView() *TutorialView {
	return &TutorialView{}
}

// Similar to React's render function. Views specify their properties and
// children in Build().
func (v *TutorialView) Build(ctx view.Context) view.Model {
	l := &constraint.Layouter{}

	// Create a new textview.
	child := view.NewTextView()
	child.String = "Hello World"
	child.Style.SetTextColor(v.TextColor)
	child.Style.SetFont(text.DefaultBoldFont(50))
	child.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}

	// Layout is primarily done using constraints. More info can be
	// found in the matcha/layout/constraints docs.
	l.Add(child, func(s *constraint.Solver) {
		s.Top(20)
		s.Left(0)
	})

	// Returns the view's children, layout, and styling.
	return view.Model{
		Children: []view.View{child},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Lightgray},
	}
}
