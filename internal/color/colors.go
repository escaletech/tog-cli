package color

import "github.com/fatih/color"

var URL = color.New(color.FgCyan).SprintFunc()
var Parameter = color.New(color.Bold, color.FgHiWhite).SprintFunc()
var Feedback = color.New(color.Faint).SprintFunc()
var Error = color.New(color.FgHiRed).SprintFunc()
