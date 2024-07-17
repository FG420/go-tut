module example.com/m

go 1.22.5

replace example.com/wiki => ../gowiki

require (
	example.com/template v1.0.0
	example.com/wiki v1.0.0
)

replace example.com/template => ../tmpl
