package toolx

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/flitsinc/go-llms/tools"
	"github.com/metalim/jsonmap"
)

type readParam struct {
	DirPath  string `json:"dir_path,omitempty" description:"必须是绝对路径, 有值时不应该传 file_path"`
	FilePath string `json:"file_path,omitempty" description:"必须是绝对路径, 有值时不应该传 dir_path"`
}

func ReadTool() tools.Tool {
	return tools.Func("read file", "读取文件夹结构或者文件内容", "read", func(_ tools.Runner, p readParam) tools.Result {
		fmt.Printf("read tool called, param: %+v\n", p)
		if p.DirPath != "" {
			d, err := os.Open(p.DirPath)
			if err != nil {
				return tools.Error(err)
			}
			defer d.Close()
			s, err := readDirStructure(d)
			if err != nil {
				return tools.Errorf("%s, but meet error: %v", s, err)
			}
			return tools.Success(s)
		}
		if p.FilePath != "" {
			f, err := os.Open(p.DirPath)
			if err != nil {
				return tools.Error(err)
			}
			defer f.Close()
			b, err := io.ReadAll(f)
			if err != nil {
				return tools.Error(err)
			}
			return tools.Success(string(b))
		}
		return tools.Errorf("dir_path 和 file_path 都是空的，没法读文件夹和文件")
	})
}

func readDirStructure(dir *os.File) (string, error) {
	entries, err := dir.ReadDir(100)
	if err != nil {
		if entries != nil {
			return buildStructureByDirEntry(entries), err
		}
	}
	return buildStructureByDirEntry(entries), nil
}

func buildStructureByDirEntry(entries []os.DirEntry) string {
	var s strings.Builder
	for _, en := range entries {
		s.WriteString(en.Name())
		s.WriteRune('\n')
	}
	return s.String()
}

func ReadFileTool() tools.Tool {
	params := jsonmap.New()
	params.Push("path", map[string]any{
		"type":        "string",
		"description": "文件绝对路径",
	})
	tool := tools.External("read file", &tools.FunctionSchema{
		Name:        "readfile",
		Description: "读取文件内容",
		Parameters: tools.ValueSchema{
			Type:       "object",
			Properties: params,
		},
	}, func(r tools.Runner, params json.RawMessage) tools.Result {
		return tools.Success("pass to client execute")
	})
	return tool
}

func readFile(_ tools.Runner, params json.RawMessage) tools.Result {
	fmt.Printf("execute tool: readfile by param: %s\n", params)
	m := make(map[string]any)
	err := json.Unmarshal(params, &m)
	if err != nil {
		return tools.Error(err)
	}
	if path, ok := m["path"]; !ok {
		return tools.Errorf("no path param: %s\n", params)
	} else if p, ok := path.(string); !ok {
		return tools.Errorf("path type is wrong\n")
	} else {
		b, err := os.ReadFile(p)
		if err != nil {
			return tools.Errorf("cannot read file: %v\n", err)
		}
		fmt.Println("have read content")
		return tools.Successf("have read file, content: %s\n", b)
	}
}
