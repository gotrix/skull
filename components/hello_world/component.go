package component

import (
	"fmt"

	"github.com/gotrix/gotrix"
)

// Component struct. Implements gotix.Component.
type Component struct{}

// Include component func.
func (*Component) Include(params gotrix.ComponentParams) (string, error) {
	rows := ""
	for _, p := range params.Params() {
		rows += fmt.Sprintf("<tr><td>%s</td></tr>", p)
	}
	return fmt.Sprintf(
		`
		<div class="gtx-component hello-world">
			<h2>Hello %s!</h2>
			<p>This is hello_world.so component</p>
			<table>
				<thead>
				<tr>
					<th>Component parameter</th>
				<tr>
				</thead>
				<tbody>
				%s
				</tbody>
			</table>
		</div>`,
		params.Name(), rows), nil
}
