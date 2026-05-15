curl \
-N \
-X POST \
-H 'content-type:application/json' \
-i \
-d '{
    "messages": [{
        "role": "user",
        "content": "先并行的读取十次调用readfile工具十次, 一次返回多个工具"
    }],
    "tools": [
        {
            "type": "function",
            "function": {
                "name": "readfile",
                "description": "读取文件内容",
                "parameters": {"type": "object", "properties": {"path": {"type": "string", "description":"文件绝对路径"}}, "required": ["path"]}
            }
        },
        {
            "type": "function",
            "function": {
                "name": "writefile",
                "description": "写出文件内容",
                "parameters": {"type": "object", "properties": {"path": {"type": "string", "description":"文件绝对路径"}, "content": {"type": "string", "description": "文件内容"}}, "required": ["path", "content"]}
            }
        }
    ]
}' \
http://localhost:6060/ollama_direct