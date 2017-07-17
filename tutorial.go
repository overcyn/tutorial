package tutorial

import (
	"golang.org/x/image/colornames"
	"gomatcha.io/bridge"
	"gomatcha.io/matcha/layout/constraint"
	"gomatcha.io/matcha/paint"
	"gomatcha.io/matcha/text"
	"gomatcha.io/matcha/view"
	"gomatcha.io/matcha/view/textview"
)

// Here is our root view.
type TutorialView struct {
	// All components must implement the view.View interface. A basic implementation
	// is provided by view.Embed.
	*view.Embed
}

// This is our view's initializer.
func New(ctx *view.Context, key string) *TutorialView {
	// To prevent rebuilding the entire tree on every rerender, initializers will return
	// the previous view if it already exists. Most views will contain this bit
	// of boilerplate.
	if v, ok := ctx.Prev(key).(*TutorialView); ok {
		return v
	}
	// If there was no matching view, we create a new one.
	return &TutorialView{Embed: ctx.NewEmbed(key)}
}

// Similar to React's render function. Views specify their properties and
// children in Build().
func (v *TutorialView) Build(ctx *view.Context) *view.Model {
	l := constraint.New()

	// Get the textview for the given key (hellotext), either initializing it or fetching
	// the previous one.
	textv := textview.New(ctx, "hellotext")
	textv.String = "Hello World"
	textv.Style.SetTextColor(colornames.Red)
	textv.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Face:   "Bold",
		Size:   50,
	})
	textv.PaintStyle = &paint.Style{BackgroundColor: colornames.Blue}

	// Layout is primarily done using a constraints. More info can be
	// found in the matcha/layout/constraints docs.
	l.Add(textv, func(s *constraint.Solver) {
		s.Top(20)
		s.Left(0)
	})

	// Returns the view's children, layout, and styling.
	return &view.Model{
		Children: []view.View{textv},
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.Lightgray},
	}
}

func init() {
	// Registers a function with the objc bridge. This function returns
	// a view.Root, which can be display in MatchaViewController.
	bridge.RegisterFunc("github.com/overcyn/tutorial New", func() *view.Root {
		return view.NewRoot(view.ScreenFunc(func(ctx *view.Context) view.View {

			// Call the TutorialView initializer.
			return New(ctx, "")
		}))
	})
}
