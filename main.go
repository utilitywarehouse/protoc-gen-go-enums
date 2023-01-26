package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

var includeNested bool

func main() {
	protogen.Options{
		ParamFunc: func(name, value string) error {
			switch name {
			case "include_nested":
				switch value {
				case "true", "yes", "y", "1":
					includeNested = true
				}
			}
			return nil
		},
	}.Run(generate)
}

func generate(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		newFileEnumGenerator(f, gen).Generate()
	}
	return nil
}

type fileEnumGenerator struct {
	f   *protogen.File
	gen *protogen.Plugin

	allConsts map[string]struct{}
	gf        *protogen.GeneratedFile
}

func newFileEnumGenerator(f *protogen.File, gen *protogen.Plugin) *fileEnumGenerator {
	gen.SupportedFeatures |= uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	return &fileEnumGenerator{f: f, gen: gen, allConsts: make(map[string]struct{})}
}

func (eg *fileEnumGenerator) Generate() {
	eg.gf = eg.gen.NewGeneratedFile(eg.f.GeneratedFilenamePrefix+".pb.enums.go", eg.f.GoImportPath)
	eg.writeHeader()

	var seen int
	for _, enum := range eg.f.Enums {
		seen += eg.processEnum(enum)
	}

	if includeNested {
		for _, m := range eg.f.Messages {
			for _, enum := range m.Enums {
				seen += eg.processEnum(enum)
			}
		}
	}

	if seen == 0 {
		// No enums processed
		eg.gf.Skip()
	}
}

func (eg *fileEnumGenerator) writeHeader() {
	eg.gf.P("// Code generated by protoc-gen-go-enums. DO NOT EDIT.")
	eg.gf.P("// source: ", eg.f.Desc.Path())
	eg.gf.P()
	eg.gf.P("package ", eg.f.GoPackageName)
	eg.gf.P()
}

func (eg *fileEnumGenerator) processEnum(enum *protogen.Enum) (seen int) {
	if hasClash, clashing := eg.hasConstClash(enum); hasClash {
		println(fmt.Sprintf("skipped generating constants for enum %v as it would duplicate constant %s", enum.Desc.FullName(), clashing))
		return seen
	}

	eg.gf.P("const (")
	for _, v := range enum.Values {
		c := constNameFor(v)
		eg.allConsts[c] = struct{}{}
		if v.Desc.Options().(*descriptorpb.EnumValueOptions).GetDeprecated() {
			eg.gf.P("// Deprecated: Do not use.")
		}
		eg.gf.P(c, " = ", eg.golangValue(v))

		seen++
	}
	eg.gf.P(")")
	eg.gf.P()

	return seen
}

func (eg *fileEnumGenerator) hasConstClash(enum *protogen.Enum) (bool, string) {
	for _, v := range enum.Values {
		constName := constNameFor(v)
		if _, exists := eg.allConsts[constName]; exists {
			return true, constName
		}
	}
	return false, ""
}

func constNameFor(v *protogen.EnumValue) string {
	return string(v.Desc.Name())
}

func (eg *fileEnumGenerator) golangValue(e *protogen.EnumValue) string {
	pkg := string(eg.f.Desc.Package())
	typeName := strings.TrimPrefix(string(e.Parent.Desc.FullName()), ".")
	typeName = strings.TrimPrefix(typeName, pkg+".")

	parts := strings.Split(typeName, ".")
	if len(parts) > 1 {
		typeName = strings.Join(parts[0:len(parts)-1], "_")
	} else {
		typeName = parts[0]
	}

	return typeName + "_" + string(e.Desc.Name())
}
