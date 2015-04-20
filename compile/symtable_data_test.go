// Test data generated by make_symtable_test.py - do not edit

package compile

import (
	"github.com/ncw/gpython/py"
)

var symtableTestData = []struct {
	in            string
	mode          string // exec, eval or single
	out           *SymTable
	exceptionType *py.Type
	errString     string
}{
	{"1", "eval", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols:           Symbols{},
		Children:          map[string]*SymTable{},
	}, nil, ""},
	{"a*b*c", "eval", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"a": Symbol{
				Flags: defUse,
				Scope: scopeGlobalImplicit,
			},
			"b": Symbol{
				Flags: defUse,
				Scope: scopeGlobalImplicit,
			},
			"c": Symbol{
				Flags: defUse,
				Scope: scopeGlobalImplicit,
			},
		},
		Children: map[string]*SymTable{},
	}, nil, ""},
	{"def fn(): pass", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         false,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{},
				Symbols:           Symbols{},
				Children:          map[string]*SymTable{},
			},
		},
	}, nil, ""},
	{"def fn(a,b):\n e=1\n return a*b*c*d*e", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         false,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      true,
				NeedsClassClosure: false,
				Varnames:          []string{"a", "b"},
				Symbols: Symbols{
					"a": Symbol{
						Flags: defParam | defUse,
						Scope: scopeLocal,
					},
					"b": Symbol{
						Flags: defParam | defUse,
						Scope: scopeLocal,
					},
					"c": Symbol{
						Flags: defUse,
						Scope: scopeGlobalImplicit,
					},
					"d": Symbol{
						Flags: defUse,
						Scope: scopeGlobalImplicit,
					},
					"e": Symbol{
						Flags: defLocal | defUse,
						Scope: scopeLocal,
					},
				},
				Children: map[string]*SymTable{},
			},
		},
	}, nil, ""},
	{"def fn(a,b):\n def nested(c,d):\n  return a*b*c*d*e", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         true,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         true,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{"a", "b"},
				Symbols: Symbols{
					"a": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
					"b": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
					"nested": Symbol{
						Flags: defLocal,
						Scope: scopeLocal,
					},
				},
				Children: map[string]*SymTable{
					"nested": &SymTable{
						Type:              FunctionBlock,
						Name:              "nested",
						Lineno:            2,
						Unoptimized:       0,
						Nested:            true,
						Free:              true,
						ChildFree:         false,
						Generator:         false,
						Varargs:           false,
						Varkeywords:       false,
						ReturnsValue:      true,
						NeedsClassClosure: false,
						Varnames:          []string{"c", "d"},
						Symbols: Symbols{
							"a": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
							"b": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
							"c": Symbol{
								Flags: defParam | defUse,
								Scope: scopeLocal,
							},
							"d": Symbol{
								Flags: defParam | defUse,
								Scope: scopeLocal,
							},
							"e": Symbol{
								Flags: defUse,
								Scope: scopeGlobalImplicit,
							},
						},
						Children: map[string]*SymTable{},
					},
				},
			},
		},
	}, nil, ""},
	{"def fn(a:\"a\",*arg:\"arg\",b:\"b\"=1,c:\"c\"=2,**kwargs:\"kw\") -> \"ret\":\n    def fn(A,b):\n        e=1\n        return a*arg*b*c*kwargs*A*e*glob", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         true,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         true,
				Generator:         false,
				Varargs:           true,
				Varkeywords:       true,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{"a", "b", "c", "arg", "kwargs"},
				Symbols: Symbols{
					"a": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
					"arg": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
					"b": Symbol{
						Flags: defParam,
						Scope: scopeLocal,
					},
					"c": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
					"fn": Symbol{
						Flags: defLocal,
						Scope: scopeLocal,
					},
					"kwargs": Symbol{
						Flags: defParam,
						Scope: scopeCell,
					},
				},
				Children: map[string]*SymTable{
					"fn": &SymTable{
						Type:              FunctionBlock,
						Name:              "fn",
						Lineno:            2,
						Unoptimized:       0,
						Nested:            true,
						Free:              true,
						ChildFree:         false,
						Generator:         false,
						Varargs:           false,
						Varkeywords:       false,
						ReturnsValue:      true,
						NeedsClassClosure: false,
						Varnames:          []string{"A", "b"},
						Symbols: Symbols{
							"A": Symbol{
								Flags: defParam | defUse,
								Scope: scopeLocal,
							},
							"a": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
							"arg": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
							"b": Symbol{
								Flags: defParam | defUse,
								Scope: scopeLocal,
							},
							"c": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
							"e": Symbol{
								Flags: defLocal | defUse,
								Scope: scopeLocal,
							},
							"glob": Symbol{
								Flags: defUse,
								Scope: scopeGlobalImplicit,
							},
							"kwargs": Symbol{
								Flags: defUse,
								Scope: scopeFree,
							},
						},
						Children: map[string]*SymTable{},
					},
				},
			},
		},
	}, nil, ""},
	{"def fn(a):\n    global b\n    b = a", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"b": Symbol{
				Flags: defGlobal,
				Scope: scopeGlobalExplicit,
			},
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         false,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{"a"},
				Symbols: Symbols{
					"a": Symbol{
						Flags: defParam | defUse,
						Scope: scopeLocal,
					},
					"b": Symbol{
						Flags: defGlobal | defLocal,
						Scope: scopeGlobalExplicit,
					},
				},
				Children: map[string]*SymTable{},
			},
		},
	}, nil, ""},
	{"def fn(a):\n    b = 6\n    global b\n    b = a", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         false,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"b": Symbol{
				Flags: defGlobal,
				Scope: scopeGlobalExplicit,
			},
			"fn": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"fn": &SymTable{
				Type:              FunctionBlock,
				Name:              "fn",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         false,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{"a"},
				Symbols: Symbols{
					"a": Symbol{
						Flags: defParam | defUse,
						Scope: scopeLocal,
					},
					"b": Symbol{
						Flags: defGlobal | defLocal,
						Scope: scopeGlobalExplicit,
					},
				},
				Children: map[string]*SymTable{},
			},
		},
	}, nil, ""},
	{"def outer():\n   x = 1\n   def inner():\n       nonlocal x\n       x = 2", "exec", &SymTable{
		Type:              ModuleBlock,
		Name:              "top",
		Lineno:            0,
		Unoptimized:       optTopLevel,
		Nested:            false,
		Free:              false,
		ChildFree:         true,
		Generator:         false,
		Varargs:           false,
		Varkeywords:       false,
		ReturnsValue:      false,
		NeedsClassClosure: false,
		Varnames:          []string{},
		Symbols: Symbols{
			"outer": Symbol{
				Flags: defLocal,
				Scope: scopeLocal,
			},
		},
		Children: map[string]*SymTable{
			"outer": &SymTable{
				Type:              FunctionBlock,
				Name:              "outer",
				Lineno:            1,
				Unoptimized:       0,
				Nested:            false,
				Free:              false,
				ChildFree:         true,
				Generator:         false,
				Varargs:           false,
				Varkeywords:       false,
				ReturnsValue:      false,
				NeedsClassClosure: false,
				Varnames:          []string{},
				Symbols: Symbols{
					"inner": Symbol{
						Flags: defLocal,
						Scope: scopeLocal,
					},
					"x": Symbol{
						Flags: defLocal,
						Scope: scopeCell,
					},
				},
				Children: map[string]*SymTable{
					"inner": &SymTable{
						Type:              FunctionBlock,
						Name:              "inner",
						Lineno:            3,
						Unoptimized:       0,
						Nested:            true,
						Free:              true,
						ChildFree:         false,
						Generator:         false,
						Varargs:           false,
						Varkeywords:       false,
						ReturnsValue:      false,
						NeedsClassClosure: false,
						Varnames:          []string{},
						Symbols: Symbols{
							"x": Symbol{
								Flags: defLocal | defNonlocal,
								Scope: scopeFree,
							},
						},
						Children: map[string]*SymTable{},
					},
				},
			},
		},
	}, nil, ""},
	{"def outer():\n   def inner():\n       nonlocal x\n       x = 2", "exec", nil, py.SyntaxError, "no binding for nonlocal 'x' found"},
}