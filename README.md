requires https://gopkg.in/yaml.v2

Below are the original requirements:

Write a command-line tool `danlister` in Go that will list files in a file system. 

It should accept the following CLI flags:

--help  <print help>
--path=<path to folder, required>
--recursive  (when set, list files recursively.  default is off)
--output=<json|yaml|text, default is text>

It should write its output to standard out, and be capable of displaying the output recursively.

Here are some examples of output:
TEXT
$ danlister --path=/foo/bar --output=text --recursive

/foo/bar
  hello
  afolder/
    file1
    file2
  asymlink* (/path/to/link)

JSON
$ danlister --path=/foo/bar --output=json --recursive

[
  {
    "ModifiedTime": "2015-09-18T10:35:49.1565989Z",
    "IsLink": false,
    "IsDir": false,
    "LinksTo": "",
    "Size": 10012,
    "Name": "hello",
    "Children": []
  },
  {
    "ModifiedTime": "2015-09-18T10:35:49.1565989Z",
    "IsLink": false,
    "IsDir": true,
    "LinksTo": "",
    "Size": 4096,
    "Name": "afolder",
    "Children": [
      {
        "ModifiedTime": "2015-09-18T10:35:49.1565989Z",
        "IsLink": false,
        "IsDir": false,
        "LinksTo": "",
        "Size": 4096,
        "Name": "file1",
        "Children": []
      },
      {
        "ModifiedTime": "2015-09-18T10:35:49.1565989Z",
        "IsLink": false,
        "IsDir": false,
        "LinksTo": "",
        "Size": 4096,
        "Name": "file2",
        "Children": []
      }
    ]
  },
  {
    "ModifiedTime": "2015-09-18T10:35:49.1565989Z",
    "IsLink": true,
    "IsDir": false,
    "LinksTo": "/path/to/link",
    "Size": 4096,
    "Name": "asymlink",
    "Children": []
  }
]
YAML
Same as the JSON output, but write YAML instead.

